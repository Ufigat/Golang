package main

import (
	"user/internal/app/routing"
	"user/pkg/postgres"
	"user/pkg/rabbitmq"

	log "github.com/sirupsen/logrus"

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
	err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("fatal DB connect error: %s", err.Error())
	}

	conn := rabbitmq.NewConnect()

	err = rabbitmq.ConnRabbit(conn)
	if err != nil {
		log.Fatalf("fatal rabbitmq connect error: %s", err.Error())
	}

	e := echo.New()
	routing.Init(conn)

	e.Logger.Fatal(e.Start(viper.GetString("app.port")))
}
