package delivery

import (
	"encoding/json"
	"engine/internal/app/infrastructure/repository"
	"engine/pkg/rabbitmq"
	"engine/pkg/request/engine"
	"engine/pkg/response/fault"
	"engine/pkg/util"
	"fmt"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

func GetEngines(c *rabbitmq.Connect, delivery <-chan amqp.Delivery) {
	for d := range delivery {
		var engineIDsReq []int

		err := json.Unmarshal(d.Body, &engineIDsReq)
		if err != nil {
			log.Errorln("GetEngines #1 ", err.Error())

			c.ProduceSendEngines(&util.Response{Error: fault.NewResponse(err.Error())})
			continue
		}

		fmt.Println(engineIDsReq)
		err = engine.ValidationIDs(engineIDsReq)
		if err != nil {
			log.Infoln("GetEngine #2 ", err.Error())

			c.ProduceSendEngines(&util.Response{Error: fault.NewResponse(err.Error())})
			continue
		}

		response, err := repository.GetEngines(engineIDsReq)
		if err != nil {
			log.Errorln("GetEngines #3 ", err.Error())

			c.ProduceSendEngines(&util.Response{Error: fault.NewResponse(err.Error())})
			continue
		}

		c.ProduceSendEngines(&util.Response{Data: response})
	}
}

func GetEngine(c *rabbitmq.Connect, delivery <-chan amqp.Delivery) {
	for d := range delivery {
		engineID, err := strconv.Atoi(string(d.Body))
		if err != nil {
			log.Errorln("GetEngine #1 ", err.Error())

			c.ProduceSendEngines(&util.Response{Error: fault.NewResponse(err.Error())})
			continue
		}

		err = engine.ValidationID(engineID)
		if err != nil {
			log.Infoln("GetEngine #2 ", err.Error())

			c.ProduceSendEngines(&util.Response{Error: fault.NewResponse(err.Error())})
			continue
		}

		response, err := repository.GetEngine(engineID)
		if err != nil {
			log.Errorln("GetEngine #3 ", err.Error())

			c.ProduceSendEngines(&util.Response{Error: fault.NewResponse(err.Error())})
			continue
		}

		c.ProduceSendEngines(&util.Response{Data: response})
	}
}
