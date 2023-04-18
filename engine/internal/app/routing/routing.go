package routing

import (
	"engine/internal/app/delivery"
	"engine/pkg/rabbitmq"
)

func Init(c *rabbitmq.Connect) {
	createConsumers(c)

	e := &delivery.Engine{Conn: c}
	go e.GetEngines()
	go e.GetEngine()
}

func createConsumers(c *rabbitmq.Connect) {
	c.ConsumeMessage("GetEngines", "GetEngines")
	c.ConsumeMessage("GetEngine", "GetEngine")
}
