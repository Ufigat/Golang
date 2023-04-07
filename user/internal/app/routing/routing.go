package routing

import (
	"encoding/json"
	"user/internal/app/delivery"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func InitRoutes(e *echo.Echo) {
	us := e.Group("/user")
	us.GET("/:id/cars", delivery.GetUserCars)
	//us.GET("/cars-engine", delivery.GetUserCarEngines)

	showRotes(e)
}

func showRotes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Fatal("fatal error parsing routes")
	}

	log.Infoln(string(data))
}
