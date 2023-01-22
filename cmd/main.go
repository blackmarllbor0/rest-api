package main

import (
	restgo "github.com/go-rest-api"
	"github.com/go-rest-api/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	server := new(restgo.Server)
	if err := server.Run("8080", handlers.InitRouter()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
