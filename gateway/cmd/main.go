package main

import (
	"gateway/internal/app/routing"
	"gateway/pkg/rabbitmq"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
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
	conn := rabbitmq.NewConnect()

	err := rabbitmq.ConnRabbit(conn)
	if err != nil {
		log.Fatalf("fatal rabbitmq connect error: %s", err.Error())
	}

	e := echo.New()
	routing.InitRoutes(e, conn)

	e.Logger.Fatal(e.Start(viper.GetString("app.port")))
}
