package routing

import (
	"user/internal/app/delivery"
	"user/pkg/rabbitmq"
)

func Init(c *rabbitmq.Connect) {
	createConsumers(c)

	u := &delivery.User{Conn: c}
	go u.GetUserCars()
}

func createConsumers(c *rabbitmq.Connect) {
	c.ConsumeMessage("GetUserCars", "GetUserCars", "", true, false, false, false, nil)
}
