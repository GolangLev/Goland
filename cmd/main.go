package main

import (
	todo "github.com/GolangLev/Goland"
	"github.com/GolangLev/Goland/pkg/handler"
	"github.com/GolangLev/Goland/pkg/repository"
	"github.com/GolangLev/Goland/pkg/service"
	"log"
)

func main() {
	// Отображение зависимостей /*Репозиторий, сервисы общаются с репозиторием,
	//а хэндлеры общаются с сервисами. Бизнес логика приложения*/
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s ", err.Error())
	}
}
