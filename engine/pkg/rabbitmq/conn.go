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

type ListQueue struct {
	Queues []Queue `mapstructure:"queues"`
}

type Queue struct {
	Name      string     `mapstructure:"name"`
	Durable   bool       `mapstructure:"durable"`
	Delete    bool       `mapstructure:"delete"`
	Exclusive bool       `mapstructure:"exclusive"`
	NoWait    bool       `mapstructure:"nowait"`
	Args      amqp.Table `mapstructure:"args"`
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
	var ql ListQueue

	err := viper.Unmarshal(&ql)
	if err != nil {
		log.Errorln("create #1 ", err.Error())

		return err
	}

	for cs := range channels {
		ch, err := c.Conn.Channel()
		if err != nil {
			log.Errorln("create #2 ", err.Error())

			return err
		}

		c.QueueChannel[channels[cs]] = &Consumer{Channel: ch, DeliveryChan: make(chan amqp.Delivery)}

		c.QueueMap[ql.Queues[cs].Name], err = ch.QueueDeclare(
			ql.Queues[cs].Name,      // name
			ql.Queues[cs].Durable,   // durable
			ql.Queues[cs].Delete,    // delete when unused
			ql.Queues[cs].Exclusive, // exclusive
			ql.Queues[cs].NoWait,    // no-wait
			ql.Queues[cs].Args,      // arguments
		)
		if err != nil {
			log.Errorln("create #3 ", err.Error())

			return err
		}

	}

	return nil
}

func (c *Connect) ProduceMessage(resp *util.Response,
	channelName string,
	queueName string,
	exchange string,
	mandatory bool,
	immediate bool,
	cType string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	value, err := json.Marshal(resp)
	if err != nil {
		log.Errorln("ProduceMessage serialization error: ", err.Error())

		return nil
	}

	err = c.QueueChannel[channelName].Channel.PublishWithContext(ctx,
		exchange,                   // exchange
		c.QueueMap[queueName].Name, // routing key
		mandatory,                  // mandatory
		immediate,                  // immediate
		amqp.Publishing{
			ContentType: cType,
			Body:        []byte(value),
		})
	if err != nil {
		log.Errorln(fmt.Sprintln("ProduceMessage error:  ", channelName, queueName), err.Error())

		return err
	}

	return nil
}

func (c *Connect) ConsumeMessage(channelName string,
	queueName string,
	consumer string,
	ack bool,
	exclusive bool,
	local bool,
	wait bool,
	args amqp.Table,
) {

	messages, err := c.QueueChannel[channelName].Channel.Consume(
		c.QueueMap[queueName].Name, // queue name
		consumer,                   // consumer
		ack,                        // auto-ack
		exclusive,                  // exclusive
		local,                      // no local
		wait,                       // no wait
		args,                       // arguments
	)
	if err != nil {
		log.Errorln(fmt.Sprintln("ConsumeMessage error:  ", channelName, queueName), err.Error())

		return
	}

	c.QueueChannel[channelName].DeliveryChan = messages
}
