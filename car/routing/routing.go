package routing

import (
	// "gateway/internal/delivery"

	"car/internal/app/delivery"
	"encoding/json"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	//ca := e.Group("/cars")
	e.GET("/cars", delivery.GetUserCars)
	e.GET("/car-engines", delivery.GetUserCarsWithEngines)
	e.GET("/car-engines-brand", delivery.GetCarsWithEnginesByBrand)

	// us := e.Group("/users")
	// us.GET("/user-cars", delivery.GetUserCars)
	//us.GET("/user-engines", delivery.GetUserEngines)
	showRotes(e)
}

func showRotes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Println("fatal error parsing routes")
	}

	fmt.Println(string(data))
}
