package main

import (
	"log"

	"github.com/bohdan-kozlo/todo-app"
	"github.com/bohdan-kozlo/todo-app/internal/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)

	err := srv.Run("8000", handlers.InitRoutes())
	if err != nil {
		log.Fatal("error while running server: ", err)
	}
}
