package routing

import (
	"encoding/json"
	"engine/internal/app/delivery"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/engines", delivery.PostEngineUserCars)
	e.GET("/engine", delivery.GetEngine)

	showRotes(e)
}

func showRotes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Println("fatal error parsing routes")
	}

	fmt.Println(string(data))
}
