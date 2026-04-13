package main

import (
	"log"

	todo_app "github.com/viktorov13/todo-app"
	"github.com/viktorov13/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo_app.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
