package routing

import (
	"encoding/json"
	"fmt"
	"gateway/internal/app/delivery"
	"log"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	us := e.Group("/users")
	us.GET("/user-cars", delivery.GetUserCars)
	us.GET("/user-engines", delivery.GetUserEngines)

	ca := e.Group("/cars")
	ca.GET("/car-engine", delivery.GetCarEngine)
	ca.GET("/cars-engines-brand", delivery.GetCarEnginesByBrand)

	showRotes(e)
}

func showRotes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Println("fatal error parsing routes")
	}

	fmt.Println(string(data))
}
