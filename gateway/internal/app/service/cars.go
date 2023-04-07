package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	carReq "gateway/pkg/request/car"
	"gateway/pkg/response/car"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetCars(carIDs []int) (*car.DataResponse, error) {

	carIDsReq := &carReq.CarsRequest{CarsIDs: carIDs}

	value, err := json.Marshal(carIDsReq)
	if err != nil {
		log.Errorln("GetUserCars #3 ", err.Error())

		return nil, err
	}

	carResp, err := http.Post(fmt.Sprint(viper.GetString("carService"), "/cars/"), "application/json", bytes.NewBuffer(value))
	if err != nil {
		log.Errorln("GetUserCars #4 ", err.Error())

		return nil, err
	}

	defer carResp.Body.Close()

	var carDataResp car.DataResponse

	err = json.NewDecoder(carResp.Body).Decode(&carDataResp)
	if err != nil {
		log.Errorln("GetUserCars #5 ", err.Error())

		return nil, err
	}

	fmt.Println("carDataResp", carDataResp)
	return &carDataResp, nil
}

func GetCarsEngine(carIDs []int) (*car.DataResponse, error) {

	carIDsReq := &carReq.CarsRequest{CarsIDs: carIDs}
	fmt.Println("GetUsersEngine", carIDsReq)
	value, err := json.Marshal(carIDsReq)
	if err != nil {
		log.Errorln("GetUserCars #3 ", err.Error())

		return nil, err
	}
	// fmt.Println(fmt.Sprint(viper.GetString("carService"), "/cars/engines"))
	carResp, err := http.Post(fmt.Sprint(viper.GetString("carService"), "/cars/engines"), "application/json", bytes.NewBuffer(value))
	if err != nil {
		log.Errorln("GetUserCars #4 ", err.Error())

		return nil, err
	}

	defer carResp.Body.Close()

	var carDataResp car.DataResponse

	err = json.NewDecoder(carResp.Body).Decode(&carDataResp)
	if err != nil {
		log.Errorln("GetUserCars #5 ", err.Error())

		return nil, err
	}

	return &carDataResp, nil
}
