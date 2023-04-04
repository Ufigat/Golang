package car

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"user/internal/app/domain"
	"user/pkg/response/car"
	"user/pkg/response/engine"
	"user/pkg/response/fault"
)

func GetCars(userModel *domain.User) ([]car.CarResponse, error) {
	resp, err := http.Get(fmt.Sprint("http://localhost:8081/cars?id=", userModel.ID))
	if err != nil {
		log.Println("GetCars ", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		var fault fault.FaultResponse
		err = json.NewDecoder(resp.Body).Decode(&fault)
		if err != nil {
			log.Println("GetCars ", err.Error())
			return nil, err
		}

		return nil, &fault
	}

	var crs []car.CarResponse
	err = json.NewDecoder(resp.Body).Decode(&crs)
	if err != nil {
		log.Println("GetCars ", err.Error())
		return nil, err
	}

	return crs, nil
}

func GetCarsWithEngine(userModel *domain.User) ([]engine.EngineResponse, error) {
	resp, err := http.Get(fmt.Sprint("http://localhost:8081/car-user-engines?id=", userModel.ID))
	if err != nil {
		log.Println("GetCarsWithEngine ", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		var fault fault.FaultResponse
		err = json.NewDecoder(resp.Body).Decode(&fault)
		if err != nil {
			log.Println("GetCarsWithEngine ", err.Error())
			return nil, err
		}

		return nil, &fault
	}

	var er []engine.EngineResponse
	err = json.NewDecoder(resp.Body).Decode(&er)
	if err != nil {
		log.Println("GetCarsWithEngine ", err.Error())
		return nil, err
	}

	return er, nil
}
