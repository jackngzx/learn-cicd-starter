package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		input       http.Header
		want        string
		expectedErr bool
	}{
		{
			name:        "common",
			input:       http.Header{"Authorization": []string{"ApiKey 12345"}},
			want:        "1234",
			expectedErr: false,
		},
		{
			name:        "no authorization header",
			input:       http.Header{"Accept": []string{"application/json"}},
			want:        "",
			expectedErr: true,
		},
		{
			name:        "malformed authorization header",
			input:       http.Header{"Authorization": []string{"token123"}},
			want:        "",
			expectedErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.input)
			if (err != nil) != tt.expectedErr {
				t.Errorf("GetAPIKey() error = %v, expectedErr = %v", err, tt.expectedErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
