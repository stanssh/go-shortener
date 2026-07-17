package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	Mux *chi.Mux
}

func NewRouter() *Router {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)

	return &Router{
		Mux: mux,
	}
}

func (r *Router) Run() error {
	r.Mux.HandleFunc("POST /", Save)
	r.Mux.HandleFunc("GET /{id}", Get)
	return http.ListenAndServe(":8080", r.Mux)
}
