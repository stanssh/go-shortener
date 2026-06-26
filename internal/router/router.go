package router

import (
	"net/http"

	"github.com/stanssh/go-shortener/internal/handler"
)

type Router struct {
	Mux *http.ServeMux
}

func New() *Router {
	return &Router{}
}

func (r *Router) Run() {
	r.Mux = http.NewServeMux()
	r.Mux.HandleFunc("POST /", handler.Root)

}
