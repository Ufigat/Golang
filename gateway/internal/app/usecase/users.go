package usecase

import (
	"encoding/json"

	"gateway/pkg/response/car"
	"gateway/pkg/response/engine"
	"gateway/pkg/response/fault"
	"gateway/pkg/util"

	log "github.com/sirupsen/logrus"
)

func (u *Usacase) GetUserCars(clientID int, brand string) {
	client := u.Room.Clients[clientID]

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

	var engineResp engine.DataResponse

	err = json.Unmarshal(msgSendEngines.Body, &engineResp)
	if err != nil {
		log.Errorln("GetCarsEngineByBrand #5 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		return
	}

	if engineResp.Error != nil {
		log.Errorln("GetCarsEngineByBrand #6 ", engineResp.Error.Error())

		client.Send <- &util.Response{Error: engineResp.Error}
		return
	}

	carEngineByBrandResp := &engine.ByBrandResponse{
		Brand:   brand,
		Engines: engineResp.Data,
	}

	client.Send <- &util.Response{Data: carEngineByBrandResp}
}

func (u *Usacase) GetUserEngines(clientID int, carID string) {
	client := u.Room.Clients[clientID]

	err := u.Conn.ProduceMessage([]byte(carID), "GetUserEngines", "GetUserEngines")
	if err != nil {
		log.Errorln("GetUserEngines #1 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		return
	}

	msgSendCar := <-u.Conn.QueueChannel["SendCarEngine"].DeliveryChan

	var carResp car.CarEngineResponse

	err = json.Unmarshal(msgSendCar.Body, &carResp)
	if err != nil {
		log.Errorln("GetUserEngines #2 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		return
	}

	if carResp.Error != nil {
		log.Errorln("GetUserEngines #3 ", carResp.Error.Message)

		client.Send <- &util.Response{Error: carResp.Error}
		return
	}

	value, err := json.Marshal(carResp.Data.EngineID)
	if err != nil {
		log.Errorln("GetUserEngines #4 ", err.Error())

		client.Send <- &util.Response{Error: carResp.Error}
		return
	}

	err = u.Conn.ProduceMessage(value, "GetEngine", "GetEngine")
	if err != nil {
		log.Errorln("GetUserEngines #5 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		return
	}

	msgSendEngines := <-u.Conn.QueueChannel["SendEngine"].DeliveryChan

	var engineResp engine.EnigneResponse

	err = json.Unmarshal(msgSendEngines.Body, &engineResp)
	if err != nil {
		log.Errorln("GetUserEngines #6 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		return
	}

	if engineResp.Error != nil {
		log.Errorln("GetUserEngines #7 ", engineResp.Error.Error())

		client.Send <- &util.Response{Error: engineResp.Error}
		return
	}

	carEngineResp := &engine.ForCar{
		ID:     carID,
		Engine: engineResp.Data,
	}

	client.Send <- &util.Response{Data: carEngineResp}
}
