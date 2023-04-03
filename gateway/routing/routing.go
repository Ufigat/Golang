package routing

import (
	"encoding/json"
	"fmt"
	"gateway/internal/delivery"
	"log"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	ca := e.Group("/cars")
	//ca.GET("/car-engine", delivery.GetCarEngine)
	ca.GET("/cars-engines-brand", delivery.GetCarEngineByBrand)

	us := e.Group("/users")
	us.GET("/user-cars", delivery.GetUserCars)
	us.GET("/user-engines", delivery.GetUserEngines)

	showRotes(e)
}

func showRotes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Println("fatal error parsing routes")
	}

	fmt.Println(string(data))
}
