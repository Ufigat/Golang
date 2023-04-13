package usecase

import (
	"encoding/json"

	"gateway/pkg/rabbitmq"
	"gateway/pkg/response/car"
	"gateway/pkg/response/engine"
	"gateway/pkg/response/fault"
	"gateway/pkg/util"
	"gateway/pkg/websocket"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type Usacase struct {
	Conn        *rabbitmq.Connect
	SendCar     <-chan amqp.Delivery
	SendEngines <-chan amqp.Delivery
}

func (u *Usacase) GetCarEnginesByBrand(client *websocket.Client, brand string) {
	err := u.Conn.ProduceGetCar(brand)
	if err != nil {
		log.Errorln("GetCarsEngineByBrand #1 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		// client.WritePumpClose <- true
		return
	}

	msgSendCar := <-u.SendCar

	var carsByBrandResp car.DataResponse

	err = json.Unmarshal(msgSendCar.Body, &carsByBrandResp)
	if err != nil {
		log.Errorln("GetCarsEngineByBrand #2 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		// client.WritePumpClose <- true
		return
	}

	if carsByBrandResp.Error != nil {
		log.Errorln("GetCarsEngineByBrand #3 ", carsByBrandResp.Error.Message)

		client.Send <- &util.Response{Error: carsByBrandResp.Error}
		// client.WritePumpClose <- true
		return
	}

	var engineIDs []int

	for i := range carsByBrandResp.Data {
		engineIDs = append(engineIDs, carsByBrandResp.Data[i].EngineID)
	}

	err = u.Conn.ProduceGetEngines(engineIDs)
	if err != nil {
		log.Errorln("GetCarsEngineByBrand #4 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		// client.WritePumpClose <- true
		return
	}

	msgSendEngines := <-u.SendEngines

	var engineDataResp engine.DataResponse

	err = json.Unmarshal(msgSendEngines.Body, &engineDataResp)
	if err != nil {
		log.Errorln("GetCarsEngineByBrand #5 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		// client.WritePumpClose <- true
		return
	}

	if engineDataResp.Error != nil {
		log.Errorln("GetCarsEngineByBrand #6 ", engineDataResp.Error.Error())

		client.Send <- &util.Response{Error: engineDataResp.Error}
		// client.WritePumpClose <- true
		return
	}

	carEngineByBrandResp := &engine.ByBrandResponse{
		Brand:   brand,
		Engines: engineDataResp.Data,
	}

	client.Send <- &util.Response{Data: carEngineByBrandResp}
	// client.WritePumpClose <- true
}
