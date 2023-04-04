package main

import (
	"log"
	"user/internal/app/routing"
	database "user/pkg/postgres"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal config connect error: %s", err.Error())
	}
}

func main() {
	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("fatal DB connect error: %s", err.Error())
	}

	e := echo.New()
	routing.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8083"))
}
