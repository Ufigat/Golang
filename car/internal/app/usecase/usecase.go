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

func GetUserWithCarEngines(carModel *domain.Car) (*engineRes.LinksResponse, error) {
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
	cbrs, err := repository.GetCarByBrand(carModel)
	if err != nil {
		return nil, err
	}

	if len(cbrs) == 0 {
		return car.NewResponseWithEngineByBrand(0, "", nil), nil
	}

	var uwers engineReq.EnginesRequest
	for _, cbr := range cbrs {
		uwers.EngineIDs = append(uwers.EngineIDs, cbr.EngineID)
	}

	crs, err := engine.CarEngines(&uwers)
	if err != nil {
		return nil, err
	}

	return car.NewResponseWithEngineByBrand(cbrs[0].ID, cbrs[0].Brand, crs), nil
}

func GetCarEngine(carModel *domain.Car) (*car.EngineResponse, error) {
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

	return car.NewEngineResponse(crwe.ID, cer), nil
}
