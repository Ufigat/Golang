package routing

import (
	"car/internal/app/delivery"
	"encoding/json"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/cars", delivery.GetCars)

	ca := e.Group("/car/")
	ca.GET("user-engines", delivery.GetUserCarsWithEngines)

	ca.GET("engines-brand", delivery.GetCarsWithEnginesByBrand)
	ca.GET("engines", delivery.GetCarEngine)

	showRotes(e)
}

func showRotes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Fatal("fatal error parsing routes")
	}

	log.Infoln(string(data))
}
