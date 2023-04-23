package delivery

import (
	"car/internal/app/infrastructure/repository"
	"car/pkg/rabbitmq"
	"car/pkg/response/fault"
	"car/pkg/util"
	"encoding/json"
	"strconv"

	carReq "car/pkg/request/car"

	log "github.com/sirupsen/logrus"
)

type Car struct {
	Conn *rabbitmq.Connect
}

func (c *Car) GetCars() {
	for dc := range c.Conn.QueueChannel["GetCars"].DeliveryChan {

		var carsIDs carReq.IDsRequest

		err := json.Unmarshal(dc.Body, &carsIDs)
		if err != nil {
			err = c.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendCars",
				"SendCars",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetCars #1 ", err.Error())

				break
			}
			continue
		}

		err = carsIDs.ValidationIDs()
		if err != nil {
			log.Errorln("GetCars #2 ", err.Error())

			err = c.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendCars",
				"SendCars",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetCars #3 ", err.Error())

				break
			}
			continue
		}

		resp, err := repository.GetCarForUser(carsIDs.CarsIDs)
		if err != nil {
			log.Errorln("GetCars #4 ", err.Error())

			err = c.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendCars",
				"SendCars",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetCars #5 ", err.Error())

				break
			}
			continue
		}

		err = c.Conn.ProduceMessage(&util.Response{Data: resp},
			"SendCars",
			"SendCars",
			"",
			false,
			false,
			"text/plain",
		)
		if err != nil {
			log.Fatalf("GetCars #6 ", err.Error())

			break
		}
	}
}

func (c *Car) GetCarEngines() {
	for dc := range c.Conn.QueueChannel["GetCarEngines"].DeliveryChan {

		var carsIDs carReq.IDsRequest

		err := json.Unmarshal(dc.Body, &carsIDs)
		if err != nil {
			err = c.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendCarEngines",
				"SendCarEngines",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetCarEngines #1 ", err.Error())

				break
			}
			continue
		}

		err = carsIDs.ValidationIDs()
		if err != nil {
			log.Errorln("GetCarEngines #2 ", err.Error())

			err = c.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendCarEngines",
				"SendCarEngines",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetCarEngines #3 ", err.Error())

				break
			}
			continue
		}

		resp, err := repository.GetCarEngineForUser(carsIDs.CarsIDs)
		if err != nil {
			log.Errorln("GetCarEngines #4 ", err.Error())

			err = c.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendCarEngines",
				"SendCarEngines",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetCarEngines #5 ", err.Error())

				break
			}
			continue
		}

		err = c.Conn.ProduceMessage(&util.Response{Data: resp},
			"SendCarEngines",
			"SendCarEngines",
			"",
			false,
			false,
			"text/plain",
		)
		if err != nil {
			log.Fatalf("GetCarEngines #6 ", err.Error())

			break
		}
	}
}

func (c *Car) GetCarsByBrand() {
	for dc := range c.Conn.QueueChannel["GetCar"].DeliveryChan {

		req := &carReq.Request{
			Brand: string(dc.Body),
		}

		err := req.ValidationBrand()
		if err != nil {
			log.Errorln("GetCarsByBrand #1 ", err.Error())

			err = c.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendCar",
				"SendCar",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetCarsByBrand #2 ", err.Error())

				break
			}
			continue
		}

		resp, err := repository.GetCarEngineByBrand(req)
		if err != nil {
			log.Errorln("GetCarsByBrand #3 ", err.Error())

			err = c.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendCar",
				"SendCar",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetCarsByBrand #4 ", err.Error())

				break
			}
			continue
		}

		err = c.Conn.ProduceMessage(&util.Response{Data: resp},
			"SendCar",
			"SendCar",
			"",
			false,
			false,
			"text/plain",
		)
		if err != nil {
			log.Fatalf("GetCarsByBrand #5 ", err.Error())

			break
		}
	}
}

func (c *Car) GetCarEngine() {
	for dc := range c.Conn.QueueChannel["GetCarEngine"].DeliveryChan {

		carID, err := strconv.Atoi(string(dc.Body))
		if err != nil {
			log.Errorln("GetCarEngine #1 ", err.Error())

			err = c.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendCarEngine",
				"SendCarEngine",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetCarEngine #2 ", err.Error())

				break
			}
			continue
		}

		req := &carReq.Request{
			ID: carID,
		}

		err = req.ValidationID()
		if err != nil {
			log.Errorln("GetCarEngine #3 ", err.Error())

			err = c.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendCarEngine",
				"SendCarEngine",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetCarEngine #4 ", err.Error())

				break
			}
			continue
		}

		resp, err := repository.GetCarEngine(req)
		if err != nil {
			log.Errorln("GetCarEngine #5 ", err.Error())

			err = c.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendCarEngine",
				"SendCarEngine",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetCarEngine #6 ", err.Error())

				break
			}
			continue
		}

		err = c.Conn.ProduceMessage(&util.Response{Data: resp},
			"SendCarEngine",
			"SendCarEngine",
			"",
			false,
			false,
			"text/plain",
		)
		if err != nil {
			log.Fatalf("GetCarEngine #7 ", err.Error())

			break
		}
	}
}
