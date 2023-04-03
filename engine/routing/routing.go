package routing

import (
	// "gateway/internal/delivery"

	"engine/internal/app/delivery"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {

	e.POST("/engines", delivery.PostEngineUserCars)
	//ca := e.Group("/cars")
	//ca.GET("/car-engine", delivery.GetCarEngine)
	//ca.GET("/cars-engines-brand", delivery.GetCarEngineByBrand)

	// us := e.Group("/users")
	// us.GET("/user-cars", delivery.GetUserCars)
	//us.GET("/user-engines", delivery.GetUserEngines)
}
