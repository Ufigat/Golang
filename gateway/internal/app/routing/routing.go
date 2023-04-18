package routing

import (
	"encoding/json"
	"gateway/internal/app/delivery"
	"gateway/pkg/rabbitmq"
	"gateway/pkg/websocket"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func InitRoutes(e *echo.Echo, r *rabbitmq.Connect, wr *websocket.Room) {
	createConsumers(r)
	initWs(e, wr)
	initHttp(e, r, wr)

	showRoutes(e)
}

func createConsumers(c *rabbitmq.Connect) {
	c.ConsumeMessage("SendCar", "SendCar")
	c.ConsumeMessage("SendEngines", "SendEngines")
	c.ConsumeMessage("SendCarEngine", "SendCarEngine")
	c.ConsumeMessage("SendEngine", "SendEngine")
}

func initWs(e *echo.Echo, wr *websocket.Room) {
	w := &delivery.Ws{Room: wr}

	ws := e.Group("/ws")
	ws.GET("/connect/:id", w.WsConnect)
}

func initHttp(e *echo.Echo, r *rabbitmq.Connect, wr *websocket.Room) {
	d := &delivery.Delivery{Conn: r, Room: wr}

	us := e.Group("/user/:client/:id")
	us.GET("/cars", delivery.GetUserCars)
	us.GET("/engines", delivery.GetUserEngines)

	ca := e.Group("/cars/:client")
	ca.GET("/:brand/engines-brand", d.GetCarEnginesByBrand)
	ca.GET("/:car/engine", d.GetCarEngine)
}

func showRoutes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Fatal("fatal error parsing routes")
	}

	log.Infoln(string(data))
}
