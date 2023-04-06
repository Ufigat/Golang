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

func GetUserWithCarEngines(carModel *domain.Car) (*engineRes.DataResponse, error) {
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

	if len(carsByBrand) == 0 {

		return &car.EngineByBrandResponse{ID: carsByBrand[0].ID, Brand: carsByBrand[0].Brand}, nil
	}

	var engineIDLinks engineReq.EnginesRequest
	for _, cbr := range carsByBrand {
		engineIDLinks.EngineIDs = append(engineIDLinks.EngineIDs, cbr.EngineID)
	}

	crs, err := engine.CarEngines(&engineIDLinks)
	if err != nil {
		return nil, err
	}

	return &car.EngineByBrandResponse{ID: carsByBrand[0].ID, Brand: carsByBrand[0].Brand, EngineResponse: crs}, nil
}

// func GetCarEngine(carModel *domain.Car) (*car.EngineResponse, error) {
// 	crwe, err := repository.GetCarEngine(carModel)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var uwers engineReq.EngineRequest
// 	uwers.EngineID = crwe.EngineID
// 	cer, err := engine.CarEngine(&uwers)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return car.NewEngineResponse(crwe.ID, cer), nil
// }
