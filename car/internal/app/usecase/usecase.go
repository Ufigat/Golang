package usecase

import (
	"car/internal/app/domain"
	"car/internal/app/infrastructure/repository"
	"car/internal/app/infrastructure/service/engine"
	engineReq "car/pkg/request/engine"
	"car/pkg/response/car"
	engineRes "car/pkg/response/engine"
)

func GetUserWithCar(userModel *domain.Car) ([]car.CarResponse, error) {
	uwcrs, err := repository.GetUserWithCar(userModel)
	if err != nil {
		return nil, err
	}

	return uwcrs, nil
}

func GetUserWithCarWithEngines(userModel *domain.Car) ([]engineRes.EngineResponse, error) {
	uwers, err := repository.GetUserWithCarAndEngineId(userModel)
	if err != nil {
		return nil, err
	}

	crs, err := engine.GetCarEngine(uwers)
	if err != nil {
		return nil, err
	}

	return crs, nil
}

func GetCarWithEnginesByBrand(userModel *domain.Car) (*car.CarResponseWithEngineByBrand, error) {
	cbrs, err := repository.GetCarByBrand(userModel)
	if err != nil {
		return nil, err
	}

	var uwers engineReq.UserCarsForEngineRequest
	for _, cbr := range cbrs {
		uwers.EngineIDs = append(uwers.EngineIDs, cbr.EngineID)
	}

	crs, err := engine.GetCarEngine(&uwers)
	if err != nil {
		return nil, err
	}

	return car.NewCarResponseWithEngineByBrand(cbrs[0].ID, cbrs[0].Brand, crs), nil
}
