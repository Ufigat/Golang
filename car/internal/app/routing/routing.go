package routing

import (
	"car/internal/app/delivery"
	"car/pkg/rabbitmq"
)

func Init(c *rabbitmq.Connect) {
	createConsumers(c)

	d := &delivery.Car{Conn: c}
	go d.GetCarsByBrand()
	go d.GetCarEngine()
	go d.GetCars()
	go d.GetCarEngines()
}

func createConsumers(c *rabbitmq.Connect) {
	c.ConsumeMessage("GetCar", "GetCar")
	c.ConsumeMessage("GetCars", "GetCars")
	c.ConsumeMessage("GetCarEngine", "GetCarEngine")
	c.ConsumeMessage("GetCarEngines", "GetCarEngines")
}
