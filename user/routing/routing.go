package routing

import (
	"user/internal/delivery"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	//ca := e.Group("/cars")
	e.GET("/user-cars", delivery.GetUserCars)
	//ca.GET("/cars-engines-brand", delivery.GetCarEngineByBrand)

	// us := e.Group("/users")
	// us.GET("/user-cars", delivery.GetUserCars)
	//us.GET("/user-engines", delivery.GetUserEngines)
}
