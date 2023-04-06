package usecase

import (
	"user/internal/app/domain"
	"user/internal/app/infrastructure/repository"
	"user/internal/app/infrastructure/service/car"
	"user/pkg/response/user"
)

func GetUserWithCar(userModel *domain.User) (*user.CarsResponse, error) {
	user, err := repository.GetUser(userModel)
	if err != nil {

		return nil, err
	}

	carDataResp, err := car.GetCars(userModel)
	if err != nil {

		return nil, err
	}

	user.Cars = carDataResp.Response

	return user, nil
}

// func GetUserWithCarEngines(userModel *domain.User) (*user.EnginesResponse, error) {
// 	users, err := repository.GetUser(userModel)
// 	if err != nil {

// 		return nil, err
// 	}

// 	linkEngines, err := car.GetCarsWithEngine(userModel)
// 	if err != nil {

// 		return nil, err
// 	}

// 	return user.NewEnginesResponse(users.ID, users.Name, linkEngines), nil
// }
