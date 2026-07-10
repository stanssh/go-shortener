package handlers

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stanssh/go-shortener/internal/service"
	"github.com/stanssh/go-shortener/internal/storage"
	"github.com/stretchr/testify/assert"
)

func Test_Save(t *testing.T) {
	myStorage := storage.NewMem()
	service.SetStorage(myStorage)

	tests := []struct {
		name   string
		method string
		path   string
		link   string
		header map[string]string
		want   int
	}{
		{
			name:   "normal(plaintext)",
			method: http.MethodGet,
			header: map[string]string{"x-forwarded-proto": "http", "content-type": "text/plain"},
			path:   "/",
			link:   "https://google.com",
			want:   201,
		},
		{
			name:   "no-x-forwarded",
			method: http.MethodGet,
			header: map[string]string{"x-forwarded-proto": "", "content-type": "text/plain"},
			path:   "/",
			link:   "https://hz.com",
			want:   201,
		},
		{
			name:   "TLS",
			method: http.MethodGet,
			header: map[string]string{"x-forwarded-proto": "https", "content-type": "text/plain"},
			path:   "/",
			link:   "https://hz2.com",
			want:   201,
		},
		{
			name:   "bad content-type",
			method: http.MethodGet,
			header: map[string]string{"x-forwarded-proto": "http", "content-type": "text-plain"},
			path:   "/",
			link:   "https://google.com",
			want:   400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(
				http.MethodGet,
				tt.path,
				strings.NewReader(tt.link),
			)

			if val, ok := tt.header["x-forwarded-proto"]; ok {
				if val == "https" {
					request.TLS = &tls.ConnectionState{}
				}
			}

			for k, v := range tt.header {
				request.Header.Set(k, v)
			}

			myRecorder := httptest.NewRecorder()
			Save(myRecorder, request)

			result := myRecorder.Result()

			assert.Equal(t, tt.want, result.StatusCode)

		})
	}
}
