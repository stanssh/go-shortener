package handlers

import (
	"net/http"
)

type Router struct {
	Mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		Mux: http.NewServeMux(),
	}
}

func (r *Router) Run() error {
	r.Mux.HandleFunc("POST /", Save)
	r.Mux.HandleFunc("GET /{id}", Get)
	return http.ListenAndServe(":8080", r.Mux)
}
