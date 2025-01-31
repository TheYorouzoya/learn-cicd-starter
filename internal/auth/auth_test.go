package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  http.Header
		want    string
		wantErr bool
	}{
		{
			name:    "valid header",
			header:  http.Header{"Authorization": []string{"ApiKey testApiKey"}},
			want:    "testApiKey",
			wantErr: false,
		},
		{
			name:    "missing token header",
			header:  http.Header{"Authorization": []string{}},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Header contains more than one API Key",
			header:  http.Header{"Authorization": []string{"ApiKey testApiKey anotherApiKey anotherAnotherAPIKey"}},
			want:    "testApiKey",
			wantErr: false,
		},
		{
			name:    "Header contains multiple tokens",
			header:  http.Header{"Authorization": []string{"Bearer testBearerToken", "ApiKey testApiKey"}},
			want:    "",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			key, err := GetAPIKey(test.header)
			if (err != nil) != test.wantErr {
				t.Errorf("GetApiKey() returned error: %v with output: %v, expected error: %v", err, key, test.wantErr)
				return
			}
			if test.want != key {
				t.Errorf("APIKey mismatch, wanted: %s, got: %s", test.want, key)
			}
		})
	}
}
