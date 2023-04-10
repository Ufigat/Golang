package routing

import (
	"car/internal/app/delivery"
	"encoding/json"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func InitRoutes(e *echo.Echo) {
	ca := e.Group("/cars")
	ca.POST("/", delivery.PostCars)
	ca.POST("/engines", delivery.PostCarEngines)
	ca.GET("/:brand/engines-brand", delivery.GetCarsByBrand)
	ca.GET("/:id/engine", delivery.GetCarEngine)

	showRoutes(e)
}

func showRoutes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Fatal("fatal error parsing routes")
	}

	log.Infoln(string(data))
}
