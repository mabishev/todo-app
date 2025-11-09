package main

import (
	"log"

	"github.com/mabishev/todo-app.git"
	"github.com/mabishev/todo-app.git/pkg/handler"
	"github.com/mabishev/todo-app.git/pkg/repository"
	"github.com/mabishev/todo-app.git/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
