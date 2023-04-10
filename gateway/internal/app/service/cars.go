package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	carReq "gateway/pkg/request/car"
	"gateway/pkg/response/car"
	"gateway/pkg/response/fault"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetCars(carIDs []int) (*car.DataResponse, error) {
	carIDsReq := &carReq.CarsRequest{CarsIDs: carIDs}
	value, err := json.Marshal(carIDsReq)
	if err != nil {
		log.Errorln("GetCars #1 ", err.Error())

		return nil, err
	}

	carsResp, err := http.Post(fmt.Sprint(viper.GetString("services.car"), "/cars/"), "application/json", bytes.NewBuffer(value))
	if err != nil {
		log.Errorln("GetCars #2 ", err.Error())

		return nil, err
	}

	defer carsResp.Body.Close()

	var carDataResp car.DataResponse

	err = json.NewDecoder(carsResp.Body).Decode(&carDataResp)
	if err != nil {
		log.Errorln("GetCars #3 ", err.Error())

		return nil, err
	}

	if carDataResp.Error != nil {
		log.Errorln("GetCars #4 ", carDataResp.Error.Message)

		return nil, &fault.Response{Message: carDataResp.Error.Message}
	}

	return &carDataResp, nil
}

func GetCarsEngine(carIDs []int) (*car.DataResponse, error) {
	carIDsReq := &carReq.CarsRequest{CarsIDs: carIDs}
	value, err := json.Marshal(carIDsReq)
	if err != nil {
		log.Errorln("GetCarsEngine #1 ", err.Error())

		return nil, err
	}

	carResp, err := http.Post(fmt.Sprint(viper.GetString("services.car"), "/cars/engines"), "application/json", bytes.NewBuffer(value))
	if err != nil {
		log.Errorln("GetCarsEngine #2 ", err.Error())

		return nil, err
	}

	defer carResp.Body.Close()

	var carDataResp car.DataResponse

	err = json.NewDecoder(carResp.Body).Decode(&carDataResp)
	if err != nil {
		log.Errorln("GetCarsEngine #3 ", err.Error())

		return nil, err
	}

	if carDataResp.Error != nil {
		log.Errorln("GetCarsEngine #4 ", carDataResp.Error.Message)

		return nil, &fault.Response{Message: carDataResp.Error.Message}
	}

	return &carDataResp, nil
}

func GetCarsEngineByBrand(brand string) (*car.DataResponse, error) {
	resp, err := http.Get(fmt.Sprint(viper.GetString("services.car"), "/cars/", brand, "/engines-brand"))
	if err != nil {
		log.Errorln("GetCarsEngineByBrand #1 ", err.Error())

		return nil, err
	}

	defer resp.Body.Close()

	var carsByBrandResp car.DataResponse

	err = json.NewDecoder(resp.Body).Decode(&carsByBrandResp)
	if err != nil {
		log.Errorln("GetCarsEngineByBrand #2 ", err.Error())

		return nil, err
	}

	if carsByBrandResp.Error != nil {
		log.Errorln("GetCarsEngineByBrand #3 ", carsByBrandResp.Error.Message)

		return nil, &fault.Response{Message: carsByBrandResp.Error.Message}
	}

	return &carsByBrandResp, nil
}

func GetCar(carID string) (*car.CarEngineResponse, error) {
	resp, err := http.Get(fmt.Sprint(viper.GetString("services.car"), "/cars/", carID, "/engine"))
	if err != nil {
		log.Errorln("GetCar #1 ", err.Error())

		return nil, err
	}

	defer resp.Body.Close()

	var carDataResp car.CarEngineResponse

	err = json.NewDecoder(resp.Body).Decode(&carDataResp)
	if err != nil {
		log.Errorln("GetCar #2 ", err.Error())

		return nil, err
	}

	if carDataResp.Error != nil {
		log.Errorln("GetCar #3 ", carDataResp.Error.Message)

		return nil, &fault.Response{Message: carDataResp.Error.Message}
	}

	return &carDataResp, nil
}
