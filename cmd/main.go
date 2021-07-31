package main

import (
	todo "github.com/GolangLev/Goland"
	"github.com/GolangLev/Goland/pkg/handler"
	"github.com/GolangLev/Goland/pkg/repository"
	"github.com/GolangLev/Goland/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	// Отображение зависимостей /*Репозиторий, сервисы общаются с репозиторием,
	//а хэндлеры общаются с сервисами. Бизнес логика приложения*/
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s ", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
