package car

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"user/internal/app/domain"
	"user/pkg/response/car"
	"user/pkg/response/engine"
)

func GetCars(userModel *domain.User) ([]car.CarResponse, error) {
	resp, err := http.Get(fmt.Sprint("http://localhost:8081/cars?id=", userModel.ID))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var crs []car.CarResponse
	err = json.Unmarshal(body, &crs)
	if err != nil {
		return nil, err
	}

	return crs, nil
}

func GetCarsWithEngine(userModel *domain.User) ([]engine.EngineResponse, error) {

	fmt.Println("GetCarsWithEngine GetUserWithCar", userModel.ID)

	resp, err := http.Get(fmt.Sprint("http://localhost:8081/car-engines?id=", userModel.ID))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var er []engine.EngineResponse
	err = json.Unmarshal(body, &er)
	if err != nil {
		return nil, err
	}

	return er, nil
}
