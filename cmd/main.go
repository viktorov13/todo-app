package main

import (
	"log"

	"github.com/spf13/viper"
	todo_app "github.com/viktorov13/todo-app"
	"github.com/viktorov13/todo-app/pkg/handler"
	"github.com/viktorov13/todo-app/pkg/repository"
	"github.com/viktorov13/todo-app/pkg/service"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo_app.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
