package rabbitmq

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Connect struct {
	Conn *amqp.Connection

	QueueMap map[string]amqp.Queue

	QueueChannel map[string]*Consumer
}

func NewConnect() *Connect {
	return &Connect{
		Conn:         &amqp.Connection{},
		QueueMap:     map[string]amqp.Queue{},
		QueueChannel: map[string]*Consumer{},
	}
}

type Consumer struct {
	Channel      *amqp.Channel
	DeliveryChan <-chan amqp.Delivery
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

	err = c.create()
	if err != nil {
		log.Errorln("ConnRabbit #1 ", err.Error())

		return err
	}

	return nil
}

func (c *Connect) create() error {
	channels := viper.GetStringSlice("channels")
	queues := viper.GetStringSlice("queue")

	for cs := range channels {
		ch, err := c.Conn.Channel()
		if err != nil {
			log.Errorln("create #1 ", err.Error())

			return err
		}

		c.QueueChannel[channels[cs]] = &Consumer{Channel: ch, DeliveryChan: make(chan amqp.Delivery)}

		c.QueueMap[queues[cs]], err = ch.QueueDeclare(
			channels[cs], // name
			false,        // durable
			false,        // delete when unused
			false,        // exclusive
			false,        // no-wait
			nil,          // arguments
		)
		if err != nil {
			log.Errorln("create #2 ", err.Error())

			return err
		}
	}

	return nil
}

func (c *Connect) ProduceMessage(message []byte, channelName string, queueName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := c.QueueChannel[channelName].Channel.PublishWithContext(ctx,
		"",                         // exchange
		c.QueueMap[queueName].Name, // routing key
		false,                      // mandatory
		false,                      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Errorln(fmt.Sprintln("ProduceMessage error:  ", channelName, queueName), err.Error())

		return err
	}

	return nil
}

func (c *Connect) ConsumeMessage(channelName string, queueName string) {
	messages, err := c.QueueChannel[channelName].Channel.Consume(
		c.QueueMap[queueName].Name, // queue name
		"",                         // consumer
		true,                       // auto-ack
		false,                      // exclusive
		false,                      // no local
		false,                      // no wait
		nil,                        // arguments
	)
	if err != nil {
		log.Errorln(fmt.Sprintln("ConsumeMessage error:  ", channelName, queueName), err.Error())

		return
	}

	c.QueueChannel[channelName].DeliveryChan = messages
}
