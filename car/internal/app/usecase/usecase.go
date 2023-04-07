package usecase

import (
	"car/internal/app/domain"
	"car/internal/app/infrastructure/repository"
	"car/internal/app/infrastructure/service/engine"
	engineReq "car/pkg/request/engine"
	"car/pkg/response/car"
	engineRes "car/pkg/response/engine"
)

func GetUserWithCar(carModel *domain.Car) ([]car.Response, error) {
	crs, err := repository.GetCarByUser(carModel)
	if err != nil {
		return nil, err
	}

	return crs, nil
}

func GetUserWithCarEngines(carModel *domain.Car) (*engineRes.DataResponses, error) {
	uwers, err := repository.GetUserCarAndEngine(carModel)
	if err != nil {
		return nil, err
	}

	crs, err := engine.CarEngines(uwers)
	if err != nil {
		return nil, err
	}

	return crs, nil
}

func GetCarWithEnginesByBrand(carModel *domain.Car) (*car.EngineByBrandResponse, error) {
	carsByBrand, err := repository.GetCarByBrand(carModel)
	if err != nil {
		return nil, err
	}

	var engineIDLinks engineReq.EnginesRequest
	for _, cbr := range carsByBrand {
		engineIDLinks.EngineIDs = append(engineIDLinks.EngineIDs, cbr.EngineID)
	}

	crs, err := engine.CarEngines(&engineIDLinks)
	if err != nil {
		return nil, err
	}

	return &car.EngineByBrandResponse{ID: carsByBrand[0].ID, Brand: carsByBrand[0].Brand, EngineResponse: crs.Engines}, nil
}

func GetCarEngine(carModel *domain.Car) (*car.EngineResponse, error) {
	carEngine, err := repository.GetCarEngine(carModel)
	if err != nil {
		return nil, err
	}

	var engineReq engineReq.Request
	engineReq.EngineID = carEngine.EngineID
	engineResp, err := engine.CarEngine(&engineReq)
	if err != nil {
		return nil, err
	}

	return &car.EngineResponse{ID: carEngine.ID, Engines: engineResp.Engines}, nil
}
