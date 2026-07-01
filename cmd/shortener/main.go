package main

import (
	"log"

	"github.com/stanssh/go-shortener/internal/app"
)

func main() {

	log.Print("Starting app..")

	// app := app.New()

	log.Fatal(app.Run())

}
