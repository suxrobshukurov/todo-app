package main

import (
	"log"

	"github.com/suxrobshukurov/todo-app"
	"github.com/suxrobshukurov/todo-app/pkg/handler"
	"github.com/suxrobshukurov/todo-app/pkg/repository"
	"github.com/suxrobshukurov/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http sever: %s", err.Error())
	}
}
