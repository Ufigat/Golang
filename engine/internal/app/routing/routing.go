package routing

import (
	"encoding/json"
	"engine/internal/app/delivery"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/engines", delivery.PostEngineUserCars)
	e.GET("/engine", delivery.GetEngine)

	showRotes(e)
}

func showRotes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		logrus.Fatal("fatal error parsing routes")
	}

	logrus.Infoln(string(data))
}
