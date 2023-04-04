package routing

import (
	"encoding/json"
	"fmt"
	"log"
	"user/internal/app/delivery"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/user-cars", delivery.GetUserCars)
	e.GET("/user-cars-engine", delivery.GetUserCarEngines)
	showRotes(e)
}

func showRotes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Println("fatal error parsing routes")
	}

	fmt.Println(string(data))
}
