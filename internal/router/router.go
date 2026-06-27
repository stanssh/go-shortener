package router

import (
	"net/http"

	"github.com/stanssh/go-shortener/internal/handler"
)

type Router struct {
	Mux *http.ServeMux
}

func New() *Router {
	return &Router{
		Mux: http.NewServeMux(),
	}
}

func (r *Router) Run() error {
	r.Mux.HandleFunc("POST /", handler.Store)
	return http.ListenAndServe(":8080", r.Mux)
}
