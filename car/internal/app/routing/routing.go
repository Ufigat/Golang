package routing

import (
	"car/internal/app/delivery"
	"encoding/json"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/cars", delivery.GetCars)
	e.GET("/car-user-engines", delivery.GetUserCarsWithEngines)

	e.GET("/car-engines-brand", delivery.GetCarsWithEnginesByBrand)
	e.GET("/car-engines", delivery.GetCarEngine)

	showRotes(e)
}

func showRotes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Println("fatal error parsing routes")
	}

	fmt.Println(string(data))
}
