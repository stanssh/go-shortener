// Dependency Root
package app

import (
	"log"

	"github.com/stanssh/go-shortener/internal/config"
	"github.com/stanssh/go-shortener/internal/handlers"
	"github.com/stanssh/go-shortener/internal/service"
	"github.com/stanssh/go-shortener/internal/storage"
)

// type App struct {
// 	Service *service.Service
// 	Storage *storage.InMEM
// }

// func New() *App {
// 	myStorage := storage.NewMem()

// 	service := &service.Service{Storage: myStorage}
// 	router := router.New()

// 	return &App{
// 		Service: service,
// 		Router:  router,
// 		Storage: myStorage,
// 	}

// }

// func (a *App) Run() error {
// 	return a.Router.Run()
// }

func Run() error {
	cfg := config.Load()

	myStorage := storage.NewMem()

	service.SetStorage(myStorage)
	handlers.SetBaseURL(cfg)

	log.Print("Starting app..")

	router := handlers.NewRouter()
	return router.Run(cfg)

}
