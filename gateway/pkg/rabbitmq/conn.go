package rabbitmq

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Connect struct {
	Conn *amqp.Connection

	Queue *Queue

	Channel *Channel
}

type Queue struct {
	CarQueue amqp.Queue

	SendCarQueue amqp.Queue
}

type Channel struct {
	CarChan *amqp.Channel

	SendCarChan *amqp.Channel
}

func NewConnect() *Connect {
	return &Connect{
		Conn: &amqp.Connection{},
		Queue: &Queue{
			CarQueue:     amqp.Queue{},
			SendCarQueue: amqp.Queue{},
		},
		Channel: &Channel{
			CarChan:     &amqp.Channel{},
			SendCarChan: &amqp.Channel{},
		},
	}
}

func ConnRabbit(c *Connect) error {
	var err error
	c.Conn, err = amqp.Dial("amqp://rmuser:rmpassword@localhost:5672/")

	if err != nil {
		log.Errorln("ConnRabbit #1 ", err.Error())

		return err
	}

	err = c.Create()
	if err != nil {
		log.Errorln("ConnRabbit #1 ", err.Error())

		return err
	}

	return nil
}

func (c *Connect) Create() error {
	var err error
	c.Channel.CarChan, err = c.Conn.Channel()
	if err != nil {
		log.Errorln("Create #1 ", err.Error())

		return err
	}

	fmt.Println("after after", c.Channel.CarChan)

	c.Queue.CarQueue, err = c.Channel.CarChan.QueueDeclare(
		"GetCar", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Errorln("Create #2 ", err.Error())

		return err
	}

	c.Channel.SendCarChan, err = c.Conn.Channel()
	if err != nil {
		log.Errorln("Create #3 ", err.Error())

		return err
	}

	c.Queue.SendCarQueue, err = c.Channel.SendCarChan.QueueDeclare(
		"SendCar", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Errorln("Create #4 ", err.Error())

		return err
	}

	return nil
}

func (c *Connect) ProduceGetCar(message string) error {
	fmt.Println("GetCarMessage GetCarMessage")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := c.Channel.CarChan.PublishWithContext(ctx,
		"",                    // exchange
		c.Queue.CarQueue.Name, // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	if err != nil {
		log.Errorln("GetCarMessage #1 ", err.Error())

		return err
	}

	return nil
}

func (c *Connect) ConsumeSendCar() <-chan amqp.Delivery {
	messages, err := c.Channel.SendCarChan.Consume(
		c.Queue.SendCarQueue.Name, // queue name
		"",                        // consumer
		true,                      // auto-ack
		false,                     // exclusive
		false,                     // no local
		false,                     // no wait
		nil,                       // arguments
	)
	if err != nil {
		log.Errorln("ConsumeSendCar #1 ", err.Error())
	}

	return messages
}
