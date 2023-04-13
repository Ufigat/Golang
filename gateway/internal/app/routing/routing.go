package routing

import (
	"encoding/json"
	"gateway/internal/app/delivery"
	"gateway/pkg/rabbitmq"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func InitRoutes(e *echo.Echo, c *rabbitmq.Connect) {

	// us := e.Group("/user/:id")
	// us.GET("/cars", delivery.GetUserCars)
	// us.GET("/engines", delivery.GetUserEngines)

	// ca := e.Group("/cars")
	// ca.GET("/:id/engine", delivery.GetCarEngine)
	// ca.GET("/:brand/engines-brand", delivery.GetCarEnginesByBrand)

	// sendCar := c.ConsumeSendCar()
	// sendEngines := c.ConsumeSendEngines()
	//d := &delivery.Delivery{Conn: c, SendCar: sendCar, SendEngines: sendEngines}

	// ca := e.Group("/cars")
	// ca.GET("/:id/engine", d.GetCarEngine)
	// ca.GET("/:brand", d.GetCarEnginesByBrand)

	ws := e.Group("/ws")
	ws.GET("/connect/:id", delivery.WsConnect)

	showRoutes(e)
}

func showRoutes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Fatal("fatal error parsing routes")
	}

	log.Infoln(string(data))
}
