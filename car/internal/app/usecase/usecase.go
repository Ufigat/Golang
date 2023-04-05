package usecase

import (
	"car/internal/app/domain"
	"car/internal/app/infrastructure/repository"
	"car/internal/app/infrastructure/service/engine"
	engineReq "car/pkg/request/engine"
	"car/pkg/response/car"
	engineRes "car/pkg/response/engine"
)

func GetUserWithCar(carModel *domain.Car) ([]car.CarResponse, error) {
	crs, err := repository.GetCarByUser(carModel)
	if err != nil {
		return nil, err
	}

	return crs, nil
}

func GetUserWithCarWithEngines(carModel *domain.Car) ([]engineRes.EngineResponse, error) {
	uwers, err := repository.GetCarWithUserAndEngineId(carModel)
	if err != nil {
		return nil, err
	}

	crs, err := engine.CarEngines(uwers)
	if err != nil {
		return nil, err
	}

	return crs, nil
}

func GetCarWithEnginesByBrand(carModel *domain.Car) (*car.CarWithEngineByBrandResponse, error) {
	cbrs, err := repository.GetCarByBrand(carModel)
	if err != nil {
		return nil, err
	}

	if len(cbrs) == 0 {
		return car.NewCarResponseWithEngineByBrand(0, "", nil), nil
	}

	var uwers engineReq.EnginesRequest
	for _, cbr := range cbrs {
		uwers.EngineIDs = append(uwers.EngineIDs, cbr.EngineID)
	}

	crs, err := engine.CarEngines(&uwers)
	if err != nil {
		return nil, err
	}

	return car.NewCarResponseWithEngineByBrand(cbrs[0].ID, cbrs[0].Brand, crs), nil
}

func GetCarEngine(carModel *domain.Car) (*car.CarWithEngineResponse, error) {
	crwe, err := repository.GetCarEngine(carModel)
	if err != nil {
		return nil, err
	}

	var uwers engineReq.EngineRequest
	uwers.EngineID = crwe.EngineID
	cer, err := engine.CarEngine(&uwers)
	if err != nil {
		return nil, err
	}

	return car.NewCarWithEngineResponse(crwe.ID, cer), nil
}
