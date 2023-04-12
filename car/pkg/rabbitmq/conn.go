package rabbitmq

import (
	"car/pkg/util"
	"context"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Conn *amqp.Connection

var CarQueue amqp.Queue

var CarChan *amqp.Channel

var SendCarQueue amqp.Queue

var SendCarChan *amqp.Channel

var Delivery chan amqp.Delivery

func ConnRabbit() error {
	var err error
	Conn, err = amqp.Dial("amqp://rmuser:rmpassword@localhost:5672/")

	if err != nil {
		log.Errorln("ConnRabbit #1 ", err.Error())

		return err
	}

	err = Create()
	if err != nil {
		log.Errorln("ConnRabbit #1 ", err.Error())

		return err
	}

	return nil
}

func Create() error {
	var err error
	CarChan, err = Conn.Channel()
	if err != nil {
		log.Errorln("Create #1 ", err.Error())

		return err
	}

	CarQueue, err = CarChan.QueueDeclare(
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

	SendCarChan, err = Conn.Channel()
	if err != nil {
		log.Errorln("Create #3 ", err.Error())

		return err
	}

	SendCarQueue, err = SendCarChan.QueueDeclare(
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

func SendCarMessage(resp *util.Response) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println("Body for SendCarMessage")

	err = SendCarChan.PublishWithContext(ctx,
		"",                // exchange
		SendCarQueue.Name, // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Errorln("GetCarMessage #1 ", err.Error())

		return err
	}

	return nil
}

func ConsumeGetCarMessage() <-chan amqp.Delivery {
	messages, err := CarChan.Consume(
		CarQueue.Name, // queue name
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no local
		false,         // no wait
		nil,           // arguments
	)
	if err != nil {
		log.Errorln("ConsumeSendCar #1 ", err.Error())
	}

	return messages
}
