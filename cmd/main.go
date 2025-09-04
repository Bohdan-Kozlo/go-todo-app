package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

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

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatal("error while running http server: ", err)
		}
	}()

	logrus.Println("Todo app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Println("Todo app shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatal("error while shutting down http server: ", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
