// Dependency Root
package app

import (
	"github.com/stanssh/go-shortener/internal/router"
	"github.com/stanssh/go-shortener/internal/service"
	"github.com/stanssh/go-shortener/internal/storage"
)

type App struct {
	Service *service.Service
	Router  *router.Router
	Storage *storage.InMEM
}

func New() *App {
	myStorage := storage.NewMem()

	service := &service.Service{Storage: myStorage}
	router := router.New()

	return &App{
		Service: service,
		Router:  router,
		Storage: myStorage,
	}

}

func (a *App) Run() error {
	return a.Router.Run()
}
