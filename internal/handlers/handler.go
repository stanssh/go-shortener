package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/stanssh/go-shortener/internal/service"
)

// type Server struct {
// 	srv service.Service
// }

// var srv Server

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
	fmt.Fprintf(w, "your string is: %s\n", hash)
}

func Get(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "text/plain" {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// myHash, err := io.ReadAll(r.Body)
	// r.URL.w

	val, err := service.RetrieveURL(string(r.PathValue("id")))

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Location", val)
	w.WriteHeader(http.StatusTemporaryRedirect)
	// fmt.Fprintf(w, "your string is: %s\n", hash)
}
