package usecase

import (
	"encoding/json"

	carReq "gateway/pkg/request/car"
	"gateway/pkg/response/car"
	"gateway/pkg/response/engine"
	"gateway/pkg/response/fault"
	"gateway/pkg/response/user"
	"gateway/pkg/util"

	log "github.com/sirupsen/logrus"
)

func (u *Usacase) GetUserCars(clientID int, userID string) {
	client := u.Room.Clients[clientID]

	err := u.Conn.ProduceMessage([]byte(userID), "GetUserCars", "GetUserCars", "", false, false, "text/plain")
	if err != nil {
		log.Errorln("GetUserCars #1 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		return
	}

	msgSendUserCars := <-u.Conn.QueueChannel["SendUserCars"].DeliveryChan

	var userResp user.DataResponse

	err = json.Unmarshal(msgSendUserCars.Body, &userResp)
	if err != nil {
		log.Errorln("GetUserCars #2 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		return
	}

	if userResp.Error != nil {
		log.Errorln("GetUserCars #3 ", userResp.Error.Message)

		client.Send <- &util.Response{Error: userResp.Error}
		return
	}

	carReq := &carReq.CarsRequest{CarsIDs: userResp.Data.CarIDs}

	value, err := json.Marshal(carReq)
	if err != nil {
		log.Errorln("GetUserCars #1 ", err.Error())

		client.Send <- &util.Response{Error: userResp.Error}
		return
	}

	err = u.Conn.ProduceMessage(value, "GetCars", "GetCars", "", false, false, "text/plain")
	if err != nil {
		log.Errorln("GetUserCars #4 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		return
	}

	msgSendCars := <-u.Conn.QueueChannel["SendCars"].DeliveryChan

	var carResp car.DataResponse

	err = json.Unmarshal(msgSendCars.Body, &carResp)
	if err != nil {
		log.Errorln("GetUserCars #5 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		return
	}

	if carResp.Error != nil {
		log.Errorln("GetUserCars #6 ", carResp.Error.Error())

		client.Send <- &util.Response{Error: carResp.Error}
		return
	}

	userCarsResp := &user.UserCarsResponse{
		ID:   userResp.Data.ID,
		Name: userResp.Data.Name,
		Cars: carResp.Data,
	}

	client.Send <- &util.Response{Data: userCarsResp}
}

func (u *Usacase) GetUserEngines(clientID int, userID string) {
	client := u.Room.Clients[clientID]

	err := u.Conn.ProduceMessage([]byte(userID), "GetUserCars", "GetUserCars", "", false, false, "text/plain")
	if err != nil {
		log.Errorln("GetUserCars #1 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		return
	}

	msgSendCar := <-u.Conn.QueueChannel["SendUserCars"].DeliveryChan

	var userResp user.DataResponse

	err = json.Unmarshal(msgSendCar.Body, &userResp)
	if err != nil {
		log.Errorln("GetUserEngines #2 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		return
	}

	if userResp.Error != nil {
		log.Errorln("GetUserEngines #3 ", userResp.Error.Message)

		client.Send <- &util.Response{Error: userResp.Error}
		return
	}

	value, err := json.Marshal(userResp.Data.CarIDs)
	if err != nil {
		log.Errorln("GetUserEngines #4 ", err.Error())

		client.Send <- &util.Response{Error: userResp.Error}
		return
	}

	err = u.Conn.ProduceMessage(value, "GetEngines", "GetEngines", "", false, false, "text/plain")
	if err != nil {
		log.Errorln("GetUserEngines #5 ", err.Error())

		client.Send <- &util.Response{Error: fault.NewResponse(err.Error())}
		return
	}

	msgSendEngines := <-u.Conn.QueueChannel["SendEngines"].DeliveryChan

	var engineResp engine.DataResponse

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

	userEnginesResp := &user.UserEnginesResponse{
		ID:      userResp.Data.ID,
		Name:    userResp.Data.Name,
		Engines: engineResp.Data,
	}

	client.Send <- &util.Response{Data: userEnginesResp}
}
