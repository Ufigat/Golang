package rabbitmq

import (
	"context"
	"encoding/json"
	"engine/pkg/util"
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
}

type Queue struct {
	EnginesQueue amqp.Queue

	SendEnginesQueue amqp.Queue

	EngineQueue amqp.Queue

	SendEngineQueue amqp.Queue
}

type Channel struct {
	EnginesChan *amqp.Channel

	SendEnginesChan *amqp.Channel

	EngineChan *amqp.Channel

	SendEngineChan *amqp.Channel
}

func NewConnect() *Connect {
	return &Connect{
		Conn: &amqp.Connection{},
		Queue: &Queue{
			EnginesQueue:     amqp.Queue{},
			SendEnginesQueue: amqp.Queue{},
			EngineQueue:      amqp.Queue{},
			SendEngineQueue:  amqp.Queue{},
		},
		Channel: &Channel{
			EnginesChan:     &amqp.Channel{},
			SendEnginesChan: &amqp.Channel{},
			EngineChan:      &amqp.Channel{},
			SendEngineChan:  &amqp.Channel{},
		},
	}
}

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

	c.Channel.EnginesChan, err = c.Conn.Channel()
	if err != nil {
		log.Errorln("Create #1 ", err.Error())

		return err
	}

	c.Queue.EnginesQueue, err = c.Channel.EnginesChan.QueueDeclare(
		"GetEngines", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Errorln("Create #2 ", err.Error())

		return err
	}

	c.Channel.SendEnginesChan, err = c.Conn.Channel()
	if err != nil {
		log.Errorln("Create #3 ", err.Error())

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
		log.Errorln("Create #4 ", err.Error())

		return err
	}

	return nil
}

func (c *Connect) ProduceSendEngines(resp *util.Response) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)

	}

	err = c.Channel.SendEnginesChan.PublishWithContext(ctx,
		"",                            // exchange
		c.Queue.SendEnginesQueue.Name, // routing key
		false,                         // mandatory
		false,                         // immediate
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

func (c *Connect) ConsumeEnginesChan() <-chan amqp.Delivery {
	messages, err := c.Channel.EnginesChan.Consume(
		c.Queue.EnginesQueue.Name, // queue name
		"",                        // consumer
		true,                      // auto-ack
		false,                     // exclusive
		false,                     // no local
		false,                     // no wait
		nil,                       // arguments
	)
	if err != nil {
		log.Errorln("ConsumeEnginesChan #1 ", err.Error())
	}

	return messages
}
