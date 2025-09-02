package main

import (
	"log"

	"github.com/bohdan-kozlo/todo-app"
	"github.com/bohdan-kozlo/todo-app/internal/database"
	"github.com/bohdan-kozlo/todo-app/internal/handler"
	"github.com/bohdan-kozlo/todo-app/internal/repository"
	"github.com/bohdan-kozlo/todo-app/internal/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env variables: ", err)
	}

	if err := initConfig(); err != nil {
		log.Fatal("error while initializing configs: ", err)
	}

	db := database.InitDb()

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	err = srv.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		log.Fatal("error while running server: ", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
