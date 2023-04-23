package delivery

import (
	"encoding/json"
	"engine/internal/app/infrastructure/repository"
	"engine/pkg/rabbitmq"
	"engine/pkg/request/engine"
	"engine/pkg/response/fault"
	"engine/pkg/util"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Engine struct {
	Conn *rabbitmq.Connect
}

func (e *Engine) GetEngines() {
	for d := range e.Conn.QueueChannel["GetEngines"].DeliveryChan {
		var engineIDsReq []int

		err := json.Unmarshal(d.Body, &engineIDsReq)
		if err != nil {
			log.Errorln("GetEngines #1 ", err.Error())

			err = e.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendEngines",
				"SendEngines",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetEngines #2 ", err.Error())

				break
			}
			continue
		}

		err = engine.ValidationIDs(engineIDsReq)
		if err != nil {
			log.Infoln("GetEngines #3 ", err.Error())

			err = e.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendEngines",
				"SendEngines",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetEngines #4 ", err.Error())

				break
			}
			continue
		}

		resp, err := repository.GetEngines(engineIDsReq)
		if err != nil {
			log.Errorln("GetEngines #5 ", err.Error())

			err = e.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendEngines",
				"SendEngines",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetEngines #6 ", err.Error())

				break
			}
			continue
		}

		err = e.Conn.ProduceMessage(&util.Response{Data: resp},
			"SendEngines",
			"SendEngines",
			"",
			false,
			false,
			"text/plain",
		)
		if err != nil {
			log.Fatalf("GetEngines #7 ", err.Error())

			break
		}
	}
}

func (e *Engine) GetEngine() {
	for d := range e.Conn.QueueChannel["GetEngine"].DeliveryChan {
		engineID, err := strconv.Atoi(string(d.Body))
		if err != nil {
			log.Errorln("GetEngine #1 ", err.Error())

			err = e.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendEngine",
				"SendEngine",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetEngine #2 ", err.Error())

				break
			}
			continue
		}

		err = engine.ValidationID(engineID)
		if err != nil {
			log.Infoln("GetEngine #3 ", err.Error())

			err = e.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendEngine",
				"SendEngine",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetEngine #4 ", err.Error())

				break
			}
			continue
		}

		response, err := repository.GetEngine(engineID)
		if err != nil {
			log.Errorln("GetEngine #5 ", err.Error())

			err = e.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendEngine",
				"SendEngine",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetEngine #6 ", err.Error())

				break
			}
			continue
		}

		err = e.Conn.ProduceMessage(&util.Response{Data: response},
			"SendEngine",
			"SendEngine",
			"",
			false,
			false,
			"text/plain",
		)
		if err != nil {
			log.Fatalf("GetEngine #7 ", err.Error())

			break
		}
	}
}
