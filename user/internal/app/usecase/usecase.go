package usecase

import (
	"user/internal/app/domain"
	"user/internal/app/infrastructure/repository"
	"user/internal/app/infrastructure/service/car"
	"user/pkg/response/user"
)

func GetUserWithCar(userModel *domain.User) (*user.UserWithCarsResponse, error) {
	user, err := repository.GetUser(userModel)
	if err != nil {

		return nil, err
	}

	gcs, err := car.GetCars(userModel)
	if err != nil {

		return nil, err
	}

	user.Cars = gcs

	return user, nil
}

func GetUserWithCarEngines(userModel *domain.User) (*user.UserEnginesResponse, error) {
	users, err := repository.GetUser(userModel)
	if err != nil {

		return nil, err
	}

	linkEngines, err := car.GetCarsWithEngine(userModel)
	if err != nil {

		return nil, err
	}

	return user.NewUserEnginesResponse(users.ID, users.Name, linkEngines), nil
}
