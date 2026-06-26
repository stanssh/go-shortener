package main

import (
	"log"
	"net/http"

	"github.com/stanssh/go-shortener/internal/router"
)

func main() {

	log.Print("Starting app..")
	myRouter := router.New()
	myRouter.Run()

	log.Print("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", myRouter.Mux))
}
