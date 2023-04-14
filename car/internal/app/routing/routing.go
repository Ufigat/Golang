package routing

import (
	"car/internal/app/delivery"
	database "car/pkg/postgres"
	"car/pkg/rabbitmq"
	"log"
)

func Init(c *rabbitmq.Connect) {
	// ca := e.Group("/cars")
	// ca.POST("", delivery.GetCars)
	// ca.POST("/engines", delivery.GetCarEngines)
	// //ca.GET("/:brand/engines-brand", delivery.GetCarsByBrand)
	// ca.GET("/:id/engine", delivery.GetCarEngine)

	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("fatal DB connect error: %s", err.Error())
	}

	err = rabbitmq.ConnRabbit()
	if err != nil {
		log.Fatalf("fatal rabbitmq connect error: %s", err.Error())
	}

	c.ConsumeMessage("SendCar", "SendCar")
	go delivery.GetCarsByBrand()

	// showRoutes(e)
}

// func showRoutes(e *echo.Echo) {
// 	data, err := json.MarshalIndent(e.Routes(), "", "  ")
// 	if err != nil {
// 		log.Fatal("fatal error parsing routes")
// 	}

// 	log.Infoln(string(data))
// }
