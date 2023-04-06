package usecase

import (
	"user/internal/app/domain"
	"user/internal/app/infrastructure/repository"
	"user/internal/app/infrastructure/service/car"
	"user/pkg/response/engine"
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

func GetUserWithCarEngines(userModel *domain.User) (*engine.LinksResponse, error) {
	users, err := repository.GetUser(userModel)
	if err != nil {

		return nil, err
	}

	linkEngines, err := car.GetCarsWithEngine(userModel)
	if err != nil {

		return nil, err
	}

	//fmt.Println(&engine.LinksResponse{users.ID, users.Name, linkEngines.Response})

	return &engine.LinksResponse{ID: users.ID, Name: users.Name, Engine: linkEngines.Response}, nil
}
