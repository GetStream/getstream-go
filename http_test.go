package getstream

import (
	"net/url"
	"reflect"
	"testing"
)

func TestBuildPath(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		pathParams map[string]string
		want       string
	}{
		{
			name:       "No parameters",
			path:       "/api/resource",
			pathParams: nil,
			want:       "/api/resource",
		},
		{
			name: "With parameters",
			path: "/api/{resource}/{id}",
			pathParams: map[string]string{
				"resource": "user",
				"id":       "123",
			},
			want: "/api/user/123",
		},
		{
			name: "Escaped characters",
			path: "/api/{query}",
			pathParams: map[string]string{
				"query": "special char/=&%",
			},
			want: "/api/special+char%2F%3D%26%25",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildPath(tt.path, tt.pathParams); got != tt.want {
				t.Errorf("buildPath(%q, %v) = %q, want %q", tt.path, tt.pathParams, got, tt.want)
			}
		})
	}
}

func TestExtractQueryParams(t *testing.T) {
	t.Run("Extract query params from GetCallRequest", func(t *testing.T) {
		request := &GetCallRequest{
			MembersLimit: PtrTo(10),
			Notify:       PtrTo(true),
			Ring:         PtrTo(false),
			Video:        PtrTo(true),
		}

		expected := url.Values{
			"members_limit": []string{"10"},
			"notify":        []string{"true"},
			"ring":          []string{"false"},
			"video":         []string{"true"},
		}

		result := extractQueryParams(request)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("extractQueryParams() = %v, want %v", result, expected)
		}
	})

	t.Run("Extract query params from nil request", func(t *testing.T) {
		result := extractQueryParams(nil)

		if len(result) != 0 {
			t.Errorf("extractQueryParams(nil) = %v, want empty url.Values", result)
		}
	})

}
