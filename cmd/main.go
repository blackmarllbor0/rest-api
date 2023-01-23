package main

import (
	restgo "github.com/go-rest-api"
	"github.com/go-rest-api/pkg/handler"
	"github.com/go-rest-api/pkg/repository"
	"github.com/go-rest-api/pkg/service"
	"log"
)

func main() {
	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	server := new(restgo.Server)
	if err := server.Run("8080", handlers.InitRouter()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
