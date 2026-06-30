package getstream

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type getBodyTestRequest struct {
	Foo string `json:"foo"`
}

type getBodyTestResponse struct{}

const getBodyTestJSON = `{"foo":"bar"}`

// A JSON body must carry GetBody and an accurate ContentLength.
func TestNewRequest_JSONBody_SetsGetBodyAndContentLength(t *testing.T) {
	client, _ := newClient("k", "s", WithBaseUrl("https://api.example.com"))

	req, err := newRequest(client, context.Background(), http.MethodPost, "/v1/x", nil, &getBodyTestRequest{Foo: "bar"}, nil)
	require.NoError(t, err)
	require.NotNil(t, req.GetBody, "JSON body must be replayable")
	require.Equal(t, int64(len(getBodyTestJSON)), req.ContentLength)

	b, err := io.ReadAll(req.Body)
	require.NoError(t, err)
	require.Equal(t, getBodyTestJSON, string(b))

	// GetBody must be usable repeatedly and always yield the full payload.
	for i := 0; i < 3; i++ {
		rc, err := req.GetBody()
		require.NoError(t, err)
		gb, err := io.ReadAll(rc)
		require.NoError(t, err)
		require.NoError(t, rc.Close())
		require.Equal(t, getBodyTestJSON, string(gb))
	}
}

// After the live body is drained (as the first attempt writes it out), GetBody
// must still recover identical bytes for the retry.
func TestNewRequest_GetBody_RewindableAfterDrain(t *testing.T) {
	client, _ := newClient("k", "s", WithBaseUrl("https://api.example.com"))

	req, err := newRequest(client, context.Background(), http.MethodPost, "/v1/x", nil, &getBodyTestRequest{Foo: "bar"}, nil)
	require.NoError(t, err)

	drained, err := io.ReadAll(req.Body)
	require.NoError(t, err)
	require.Equal(t, getBodyTestJSON, string(drained))

	rc, err := req.GetBody()
	require.NoError(t, err)
	replay, err := io.ReadAll(rc)
	require.NoError(t, err)
	require.Equal(t, getBodyTestJSON, string(replay))
}

// Backward compat: the wire bytes must stay exactly json.Marshal(data).
func TestNewRequest_JSONBody_WireBytesUnchanged(t *testing.T) {
	client, _ := newClient("k", "s", WithBaseUrl("https://api.example.com"))
	data := map[string]any{"field1": "value1", "field2": 2}

	req, err := newRequest(client, context.Background(), http.MethodPost, "/v1/x", nil, data, nil)
	require.NoError(t, err)

	want, err := json.Marshal(data)
	require.NoError(t, err)
	got, err := io.ReadAll(req.Body)
	require.NoError(t, err)
	require.Equal(t, string(want), string(got))
	require.Equal(t, int64(len(want)), req.ContentLength)
}

// Multipart uploads must also be replayable via GetBody.
func TestNewRequest_Multipart_SetsRetryableBody(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "f.txt")
	require.NoError(t, os.WriteFile(path, []byte("hello-file-content"), 0o600))

	client, _ := newClient("k", "s", WithBaseUrl("https://api.example.com"))
	req, err := newRequest(client, context.Background(), http.MethodPost, "/upload", nil, &UploadFileRequest{File: PtrTo(path)}, nil)
	require.NoError(t, err)
	require.NotNil(t, req.GetBody, "multipart upload must be replayable")
	require.True(t, strings.HasPrefix(req.Header.Get("Content-Type"), "multipart/form-data"))

	first, err := io.ReadAll(req.Body)
	require.NoError(t, err)
	require.Equal(t, int64(len(first)), req.ContentLength)
	require.Contains(t, string(first), "hello-file-content")

	rc, err := req.GetBody()
	require.NoError(t, err)
	second, err := io.ReadAll(rc)
	require.NoError(t, err)
	require.Equal(t, first, second, "GetBody must reproduce identical multipart bytes")
}

// TestNewRequest_GET_NoBody confirms GET requests are unchanged (no body, no GetBody).
func TestNewRequest_GET_NoBody(t *testing.T) {
	client, _ := newClient("k", "s", WithBaseUrl("https://api.example.com"))
	var data any

	req, err := newRequest(client, context.Background(), http.MethodGet, "/v1/x", nil, data, nil)
	require.NoError(t, err)
	require.Nil(t, req.Body)
	require.Nil(t, req.GetBody)
}

// Arbitrary streaming readers stay non-rewindable (no GetBody) by design.
func TestNewRequest_StreamingReader_NoGetBody(t *testing.T) {
	client, _ := newClient("k", "s", WithBaseUrl("https://api.example.com"))

	req, err := newRequest(client, context.Background(), http.MethodPut, "/v1/x", nil, strings.NewReader("raw"), nil)
	require.NoError(t, err)
	require.NotNil(t, req.Body)
	require.Nil(t, req.GetBody, "arbitrary streaming readers cannot be made rewindable")

	b, err := io.ReadAll(req.Body)
	require.NoError(t, err)
	require.Equal(t, "raw", string(b))
}

// logRequest must not consume the live body or drop GetBody.
func TestLogRequest_PreservesBodyAndGetBody(t *testing.T) {
	client, _ := newClient("k", "s", WithBaseUrl("https://api.example.com"))
	req, err := newRequest(client, context.Background(), http.MethodPost, "/v1/x", nil, &getBodyTestRequest{Foo: "bar"}, nil)
	require.NoError(t, err)

	client.logRequest(req)

	body, err := io.ReadAll(req.Body)
	require.NoError(t, err)
	require.Equal(t, getBodyTestJSON, string(body), "logRequest must not consume the live body")

	rc, err := req.GetBody()
	require.NoError(t, err)
	gb, err := io.ReadAll(rc)
	require.NoError(t, err)
	require.Equal(t, getBodyTestJSON, string(gb))
}

// refusedStreamOnce models net/http's HTTP/2 retry contract: REFUSED_STREAM is
// retried only if the body can be rewound via GetBody. It consumes the body on
// the first attempt, then rewinds via GetBody (or errors if GetBody is nil).
type refusedStreamOnce struct {
	attempts int
	bodies   [][]byte
}

func (f *refusedStreamOnce) Do(req *http.Request) (*http.Response, error) {
	f.attempts++

	var b []byte
	if req.Body != nil {
		b, _ = io.ReadAll(req.Body)
		_ = req.Body.Close()
	}
	f.bodies = append(f.bodies, b)

	if f.attempts == 1 {
		if req.GetBody == nil {
			return nil, errors.New("http2: Transport: cannot retry err [stream error: stream ID 1; REFUSED_STREAM; received from peer] after Request.Body was written; define Request.GetBody to avoid this error")
		}
		nb, err := req.GetBody()
		if err != nil {
			return nil, err
		}
		req.Body = nb
		return f.Do(req)
	}

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader("{}")),
		Header:     make(http.Header),
	}, nil
}

// A POST whose stream is refused once now succeeds, with an identical replayed body.
func TestMakeRequest_RetriesOnRefusedStream(t *testing.T) {
	fake := &refusedStreamOnce{}
	client, err := newClient("k", "s", WithBaseUrl("https://api.example.com"), WithHTTPClient(fake))
	require.NoError(t, err)

	_, err = MakeRequest(client, context.Background(), http.MethodPost, "/v1/x", nil, &getBodyTestRequest{Foo: "bar"}, &getBodyTestResponse{}, nil)
	require.NoError(t, err)
	require.Equal(t, 2, fake.attempts, "transport should have retried the refused stream once")
	require.Len(t, fake.bodies, 2)
	require.Equal(t, getBodyTestJSON, string(fake.bodies[0]))
	require.Equal(t, fake.bodies[0], fake.bodies[1], "replayed body must be identical to the first attempt")
}

// A bodied request without GetBody (the pre-fix shape) still fails, so the
// positive test above isn't tautological.
func TestRefusedStreamTransport_FailsWithoutGetBody(t *testing.T) {
	fake := &refusedStreamOnce{}
	req, err := http.NewRequest(http.MethodPost, "https://api.example.com", io.NopCloser(bytes.NewReader([]byte(getBodyTestJSON))))
	require.NoError(t, err)
	require.Nil(t, req.GetBody, "guard: this request intentionally has no GetBody")

	_, err = fake.Do(req)
	require.Error(t, err)
	require.Contains(t, err.Error(), "define Request.GetBody")
	require.Equal(t, 1, fake.attempts)
}
