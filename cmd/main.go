package main

import (
	restgo "github.com/go-rest-api"
	"github.com/go-rest-api/pkg/handler"
	"github.com/go-rest-api/pkg/repository"
	"github.com/go-rest-api/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init config: $s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to init db: %s", err.Error())
	}

	repo, _ := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	server := new(restgo.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRouter()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
