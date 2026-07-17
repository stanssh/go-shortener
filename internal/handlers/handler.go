package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/stanssh/go-shortener/internal/config"
	"github.com/stanssh/go-shortener/internal/service"
)

// type Server struct {
// 	srv service.Service
// }

// var srv Server

func getScheme(r *http.Request) string {
	if r.TLS != nil {
		return "https"
	}
	if h := r.Header.Get("X-Forwarded-proto"); h != "" {
		return h
	}
	return "http"
}

var BaseURL string = ""

func SetBaseURL(c *config.Config) {
	BaseURL = c.BaseURL
}

func Save(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "text/plain" {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	myLinkToStore, err := io.ReadAll(r.Body)

	hash, err := service.StoreURL(string(myLinkToStore))

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	prefix := fmt.Sprintf("%s://%s", getScheme(r), r.Host)

	if BaseURL != "" {
		prefix = BaseURL
	}

	fmt.Fprintf(w, "%s/%s", prefix, hash)
}

func Get(w http.ResponseWriter, r *http.Request) {

	val, err := service.RetrieveURL(string(r.PathValue("id")))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
	w.Header().Add("Location", val)
	w.WriteHeader(http.StatusTemporaryRedirect)
	// fmt.Fprintf(w, "your string is: %s\n", hash)
}
