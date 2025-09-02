package main

import (
	"github.com/bohdan-kozlo/todo-app"
	"github.com/bohdan-kozlo/todo-app/internal/database"
	"github.com/bohdan-kozlo/todo-app/internal/handler"
	"github.com/bohdan-kozlo/todo-app/internal/repository"
	"github.com/bohdan-kozlo/todo-app/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("error loading env variables: ", err)
	}

	if err := initConfig(); err != nil {
		logrus.Fatal("error while initializing configs: ", err)
	}

	db := database.InitDb()

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	err = srv.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		logrus.Fatal("error while running server: ", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
