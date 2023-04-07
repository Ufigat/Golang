package routing

import (
	"encoding/json"
	"gateway/internal/app/delivery"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func InitRoutes(e *echo.Echo) {
	us := e.Group("/user/:id")
	us.GET("/cars", delivery.GetUserCars)
	us.GET("/engines", delivery.GetUserEngines)

	ca := e.Group("/cars")
	ca.GET("/:id/engine", delivery.GetCarEngine)
	ca.GET("/:brand/engines-brand", delivery.GetCarEnginesByBrand)

	showRoutes(e)
}

func showRoutes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Fatal("fatal error parsing routes")
	}

	log.Infoln(string(data))
}
