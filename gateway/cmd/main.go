package main

import (
	"gateway/internal/app/routing"
	"gateway/pkg/rabbitmq"
	"gateway/pkg/websocket"

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

	ws := &websocket.Room{Clients: make(map[int]*websocket.Client),
		Register:   make(chan *websocket.Client),
		Unregister: make(chan *websocket.Client),
	}

	go ws.Work()

	defer func() {
		close(ws.Register)
		close(ws.Unregister)
	}()

	e := echo.New()
	routing.InitRoutes(e, conn, ws)

	e.Logger.Fatal(e.Start(viper.GetString("app.port")))
}
