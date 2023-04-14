package usecase

import (
	"encoding/json"

	"gateway/pkg/rabbitmq"
	"gateway/pkg/response/car"
	"gateway/pkg/response/engine"
	"gateway/pkg/response/fault"
	"gateway/pkg/util"
	"gateway/pkg/websocket"

	log "github.com/sirupsen/logrus"
)

type Usacase struct {
	Room *websocket.Room
	Conn *rabbitmq.Connect
}

func (u *Usacase) GetCarEnginesByBrand(userID int, brand string) {
	client := u.Room.Clients[userID]

	err := u.Conn.ProduceMessage([]byte(brand), "GetCar", "GetCar")
	if err != nil {
		log.Errorln("GetCarsEngineByBrand #1 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		return
	}

	msgSendCar := <-u.Conn.QueueChannel["SendCar"].DeliveryChan

	var carsByBrandResp car.DataResponse

	err = json.Unmarshal(msgSendCar.Body, &carsByBrandResp)
	if err != nil {
		log.Errorln("GetCarsEngineByBrand #2 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		return
	}

	if carsByBrandResp.Error != nil {
		log.Errorln("GetCarsEngineByBrand #3 ", carsByBrandResp.Error.Message)

		client.Send <- &util.Response{Error: carsByBrandResp.Error}
		return
	}

	var engineIDs []int

	for i := range carsByBrandResp.Data {
		engineIDs = append(engineIDs, carsByBrandResp.Data[i].EngineID)
	}

	value, err := json.Marshal(engineIDs)
	if err != nil {
		log.Errorln("CarEngines #1 ", err.Error())

		client.Send <- &util.Response{Error: carsByBrandResp.Error}
		return
	}

	err = u.Conn.ProduceMessage(value, "GetEngines", "GetEngines")
	if err != nil {
		log.Errorln("GetCarsEngineByBrand #4 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		return
	}

	msgSendEngines := <-u.Conn.QueueChannel["SendEngines"].DeliveryChan

	var engineDataResp engine.DataResponse

	err = json.Unmarshal(msgSendEngines.Body, &engineDataResp)
	if err != nil {
		log.Errorln("GetCarsEngineByBrand #5 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		return
	}

	if engineDataResp.Error != nil {
		log.Errorln("GetCarsEngineByBrand #6 ", engineDataResp.Error.Error())

		client.Send <- &util.Response{Error: engineDataResp.Error}
		return
	}

	carEngineByBrandResp := &engine.ByBrandResponse{
		Brand:   brand,
		Engines: engineDataResp.Data,
	}

	client.Send <- &util.Response{Data: carEngineByBrandResp}
}

func (u *Usacase) GetCarEngine(userID int, brand string) {
	// client := u.Room.Clients[userID]

	// err := u.Conn.ProduceMessage([]byte(brand), "GetCar", "GetCar")
	// if err != nil {
	// 	log.Errorln("GetCarsEngineByBrand #1 ", err.Error())

	// 	client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
	// 	return
	// }

	// msgSendCar := <-u.SendCar

	// var carsByBrandResp car.DataResponse

	// err = json.Unmarshal(msgSendCar.Body, &carsByBrandResp)
	// if err != nil {
	// 	log.Errorln("GetCarsEngineByBrand #2 ", err.Error())

	// 	client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
	// 	return
	// }

	// if carsByBrandResp.Error != nil {
	// 	log.Errorln("GetCarsEngineByBrand #3 ", carsByBrandResp.Error.Message)

	// 	client.Send <- &util.Response{Error: carsByBrandResp.Error}
	// 	return
	// }

	// var engineIDs []int

	// for i := range carsByBrandResp.Data {
	// 	engineIDs = append(engineIDs, carsByBrandResp.Data[i].EngineID)
	// }

	// value, err := json.Marshal(engineIDs)
	// if err != nil {
	// 	log.Errorln("CarEngines #1 ", err.Error())

	// 	client.Send <- &util.Response{Error: carsByBrandResp.Error}
	// 	return
	// }

	// err = u.Conn.ProduceMessage(value, "GetEnginesChan", "GetEnginesChan")
	// if err != nil {
	// 	log.Errorln("GetCarsEngineByBrand #4 ", err.Error())

	// 	client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
	// 	return
	// }

	// msgSendEngines := <-u.SendEngines

	// var engineDataResp engine.DataResponse

	// err = json.Unmarshal(msgSendEngines.Body, &engineDataResp)
	// if err != nil {
	// 	log.Errorln("GetCarsEngineByBrand #5 ", err.Error())

	// 	client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
	// 	return
	// }

	// if engineDataResp.Error != nil {
	// 	log.Errorln("GetCarsEngineByBrand #6 ", engineDataResp.Error.Error())

	// 	client.Send <- &util.Response{Error: engineDataResp.Error}
	// 	return
	// }

	// carEngineByBrandResp := &engine.ByBrandResponse{
	// 	Brand:   brand,
	// 	Engines: engineDataResp.Data,
	// }

	// client.Send <- &util.Response{Data: carEngineByBrandResp}
}
