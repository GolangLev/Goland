package main

import (
	todo "github.com/GolangLev/Goland"
	"github.com/GolangLev/Goland/pkg/handler"
	"github.com/GolangLev/Goland/pkg/repository"
	"github.com/GolangLev/Goland/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	//.Env /*С помощью библиотеки godotenv считываем пароль для базы данных*/
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("error loading env variabless: %s", err.Error())
	}

	//.yml /*С файла .yml считываем конфигурацию базы данных*/
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	// Отображение зависимостей /*Репозиторий, сервисы общаются с репозиторием,
	//а хэндлеры общаются с сервисами. Бизнес логика приложения*/
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	//Server /*Запускаем сервер*/
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running server: %s ", err.Error())
	}
}

//InitConfig /*Инициализация файла .yml - файл конфигурации*/
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
