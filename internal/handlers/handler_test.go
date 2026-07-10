package handlers

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stanssh/go-shortener/internal/service"
	"github.com/stanssh/go-shortener/internal/storage"
	"github.com/stretchr/testify/assert"
)

// var myStorage service.Repository

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
			method: http.MethodPost,
			header: map[string]string{"x-forwarded-proto": "http", "content-type": "text/plain"},
			path:   "/",
			link:   "https://google.com",
			want:   201,
		},
		{
			name:   "no-x-forwarded",
			method: http.MethodPost,
			header: map[string]string{"x-forwarded-proto": "", "content-type": "text/plain"},
			path:   "/",
			link:   "https://hz.com",
			want:   201,
		},
		{
			name:   "TLS",
			method: http.MethodPost,
			header: map[string]string{"x-forwarded-proto": "https", "content-type": "text/plain"},
			path:   "/",
			link:   "https://hz2.com",
			want:   201,
		},
		{
			name:   "bad content-type",
			method: http.MethodPost,
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

			if !assert.Equal(t, tt.want, result.StatusCode) {
				t.Errorf("expected: %v, got %v", tt.want, result.StatusCode)
			}

		})
	}
}

func Test_Get(t *testing.T) {
	myStorage := storage.NewMem()
	service.SetStorage(myStorage)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{id}", Get) // Регистрируем роут с шаблоном
	mux.HandleFunc("POST /", Save)   // Регистрируем роут с шаблоном

	myRecorder := httptest.NewRecorder()

	// Вместо прямого вызова Get(myRecorder, getRequest):
	// mux.ServeHTTP(myRecorder, getRequest)

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

			path: "/",
			link: "https://googles.com",
			want: 307,
		},
		// {
		// name:   "no-x-forwarded",
		// method: http.MethodGet,
		// header: map[string]string{"x-forwarded-proto": "", "content-type": "text/plain"},
		// path:   "/",
		// link:   "https://hz.com",
		// want:   201,
		// },
		// {
		// name:   "TLS",
		// method: http.MethodGet,
		// header: map[string]string{"x-forwarded-proto": "https", "content-type": "text/plain"},
		// path:   "/",
		// link:   "https://hz2.com",
		// want:   201,
		// },
		// {
		// name:   "bad content-type",
		// method: http.MethodGet,
		// header: map[string]string{"x-forwarded-proto": "http", "content-type": "text-plain"},
		// path:   "/",
		// link:   "https://google.com",
		// want:   400,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saveRequest := httptest.NewRequest(
				http.MethodPost,
				tt.path,
				strings.NewReader(tt.link),
			)

			for k, v := range tt.header {
				saveRequest.Header.Set(k, v)
			}

			mux.ServeHTTP(myRecorder, saveRequest)

			// Save(myRecorder, saveRequest)

			assert.Equal(t, myRecorder.Result().StatusCode, 201)

			saveResult, _ := io.ReadAll(myRecorder.Body)
			suffixArray := strings.Split(string(saveResult), "/")

			getRequest := httptest.NewRequest(
				http.MethodGet,
				"/"+string(suffixArray[len(suffixArray)-1]),
				nil,
			)

			for k, v := range tt.header {
				getRequest.Header.Set(k, v)
			}

			getRecorder := httptest.NewRecorder()

			mux.ServeHTTP(getRecorder, getRequest)

			assert.Equal(t, tt.want, getRecorder.Result().StatusCode)

		})
	}
}
