package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Println("Starting broker service...")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
