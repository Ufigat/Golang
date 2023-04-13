package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Connect struct {
	Conn *amqp.Connection

	Queue *Queue

	Channel *Channel

	QueueMap map[string]*Queue

	QueueChannel map[string]*Queue
}

type Queue struct {
	CarQueue amqp.Queue

	SendCarQueue amqp.Queue

	EnginesQueue amqp.Queue

	SendEnginesQueue amqp.Queue
}

type Channel struct {
	CarChan *amqp.Channel

	SendCarChan *amqp.Channel

	EnginesChan *amqp.Channel

	SendEnginesChan *amqp.Channel
}

func NewConnect() *Connect {
	return &Connect{
		Conn: &amqp.Connection{},
		Queue: &Queue{
			CarQueue:         amqp.Queue{},
			SendCarQueue:     amqp.Queue{},
			EnginesQueue:     amqp.Queue{},
			SendEnginesQueue: amqp.Queue{},
		},
		Channel: &Channel{
			CarChan:         &amqp.Channel{},
			SendCarChan:     &amqp.Channel{},
			EnginesChan:     &amqp.Channel{},
			SendEnginesChan: &amqp.Channel{},
		},
		QueueMap:     map[string]*Queue{},
		QueueChannel: map[string]*Queue{},
	}
}

var ChannelSlice []string

var QueueChannel []string

func ConnRabbit(c *Connect) error {
	var err error
	rabbitInfo := fmt.Sprint("amqp://", viper.GetString("rabbit.user"),
		viper.GetString("rabbit.password"), viper.GetString("rabbit.host"), viper.GetString("rabbit.port"))

	c.Conn, err = amqp.Dial(rabbitInfo)
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

	c.Channel.EnginesChan, err = c.Conn.Channel()
	if err != nil {
		log.Errorln("Create #5 ", err.Error())

		return err
	}

	c.Queue.EnginesQueue, err = c.Channel.EnginesChan.QueueDeclare(
		"GetEnginesChan", // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Errorln("Create #6 ", err.Error())

		return err
	}

	c.Channel.SendEnginesChan, err = c.Conn.Channel()
	if err != nil {
		log.Errorln("Create #7 ", err.Error())

		return err
	}

	c.Queue.SendEnginesQueue, err = c.Channel.SendEnginesChan.QueueDeclare(
		"SendEngines", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Errorln("Create #8 ", err.Error())

		return err
	}

	return nil
}

func (c *Connect) ProduceGetCar(message string) error {
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

func (c *Connect) ProduceGetEngines(req []int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	value, err := json.Marshal(req)
	if err != nil {
		log.Errorln("CarEngines #1 ", err.Error())

		return err
	}

	err = c.Channel.CarChan.PublishWithContext(ctx,
		"",                        // exchange
		c.Queue.EnginesQueue.Name, // routing key
		false,                     // mandatory
		false,                     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(value),
		})

	if err != nil {
		log.Errorln("ProduceGetEngines #1 ", err.Error())

		return err
	}

	return nil
}

func (c *Connect) ConsumeSendEngines() <-chan amqp.Delivery {
	messages, err := c.Channel.SendEnginesChan.Consume(
		c.Queue.SendEnginesQueue.Name, // queue name
		"",                            // consumer
		true,                          // auto-ack
		false,                         // exclusive
		false,                         // no local
		false,                         // no wait
		nil,                           // arguments
	)
	if err != nil {
		log.Errorln("ConsumeSendEngines #1 ", err.Error())
	}

	return messages
}
