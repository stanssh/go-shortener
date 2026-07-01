package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// type Server struct {
// 	srv service.Service
// }

// var srv Server

func Store(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "text/plain" {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	myLinkToStore, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "your string is: %s\n", myLinkToStore)
}
