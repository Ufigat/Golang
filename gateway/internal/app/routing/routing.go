package routing

import (
	"encoding/json"
	"gateway/internal/app/delivery"
	"gateway/pkg/rabbitmq"
	"gateway/pkg/websocket"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func InitRoutes(e *echo.Echo, c *rabbitmq.Connect, r *websocket.Room) {

	// us := e.Group("/user/:id")
	// us.GET("/cars", delivery.GetUserCars)
	// us.GET("/engines", delivery.GetUserEngines)

	// ca := e.Group("/cars")
	// ca.GET("/:id/engine", delivery.GetCarEngine)
	// ca.GET("/:brand/engines-brand", delivery.GetCarEnginesByBrand)

	c.ConsumeMessage("SendCar", "SendCar")
	c.ConsumeMessage("SendEngines", "SendEngines")
	d := &delivery.Delivery{Conn: c, Room: r}

	ca := e.Group("/cars/:id")
	ca.GET("/:brand/engines-brand", d.GetCarEnginesByBrand)
	// ca.GET("/:id/engine", d.GetCarEngine)

	w := &delivery.Ws{Room: r}

	ws := e.Group("/ws")
	ws.GET("/connect/:id", w.WsConnect)

	showRoutes(e)
}

func showRoutes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Fatal("fatal error parsing routes")
	}

	log.Infoln(string(data))
}
