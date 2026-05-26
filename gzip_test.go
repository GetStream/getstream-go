package getstream

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// gzipTestResponse is a minimal response shape used by the gzip tests below.
// It avoids depending on any generated model so the tests stay focused on
// transport-level behaviour (gzip request advertisement + gzip response
// decoding) rather than on a particular endpoint's schema.
type gzipTestResponse struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}

// TestGzipRequestAdvertisesAcceptEncoding verifies the CHA-2964 invariant:
// when the SDK issues a request through its default *http.Client, net/http
// auto-adds an "Accept-Encoding: gzip" header. The SDK itself does NOT set
// the header (which would otherwise silently disable auto-decoding).
func TestGzipRequestAdvertisesAcceptEncoding(t *testing.T) {
	var capturedAcceptEncoding string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedAcceptEncoding = r.Header.Get("Accept-Encoding")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","count":1}`))
	}))
	defer server.Close()

	client, err := newClient("testKey", "testSecret", WithBaseUrl(server.URL))
	require.NoError(t, err)

	var result gzipTestResponse
	_, err = MakeRequest[any, gzipTestResponse](
		client, context.Background(), http.MethodGet, "/api/v2/app", nil, nil, &result, nil,
	)
	require.NoError(t, err)

	// net/http.Transport auto-adds Accept-Encoding: gzip when:
	//   (a) the header is not set manually anywhere in the request chain
	//   (b) Transport.DisableCompression is false
	// Both hold for the SDK's default client, so the server must observe gzip.
	assert.True(t,
		strings.Contains(capturedAcceptEncoding, "gzip"),
		"expected Accept-Encoding header to contain \"gzip\", got %q", capturedAcceptEncoding,
	)
}

// TestGzipResponseAutoDecoded verifies that a gzip-encoded response body is
// transparently decoded by net/http before the SDK unmarshals it. The mock
// server returns Content-Encoding: gzip with a gzipped JSON payload; the SDK
// must produce the uncompressed struct.
func TestGzipResponseAutoDecoded(t *testing.T) {
	expected := gzipTestResponse{Message: "hello-from-gzip", Count: 42}

	payload, err := json.Marshal(expected)
	require.NoError(t, err)

	var gzBuf bytes.Buffer
	gzWriter := gzip.NewWriter(&gzBuf)
	_, err = gzWriter.Write(payload)
	require.NoError(t, err)
	require.NoError(t, gzWriter.Close())
	gzippedPayload := gzBuf.Bytes()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Sanity: only emit gzip if the client advertised support. This mirrors
		// the production server's behaviour and keeps the test honest.
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			http.Error(w, "client did not advertise gzip support", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		_, _ = w.Write(gzippedPayload)
	}))
	defer server.Close()

	client, err := newClient("testKey", "testSecret", WithBaseUrl(server.URL))
	require.NoError(t, err)

	var result gzipTestResponse
	_, err = MakeRequest[any, gzipTestResponse](
		client, context.Background(), http.MethodGet, "/api/v2/app", nil, nil, &result, nil,
	)
	require.NoError(t, err)

	assert.Equal(t, expected.Message, result.Message, "expected gzipped response to be auto-decoded into struct")
	assert.Equal(t, expected.Count, result.Count, "expected gzipped response to be auto-decoded into struct")
}
