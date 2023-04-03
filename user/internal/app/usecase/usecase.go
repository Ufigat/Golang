package usecase

import (
	"user/internal/app/domain"
	"user/internal/app/infrastructure/repository"
	"user/internal/app/infrastructure/service/car"
	"user/pkg/response/user"
)

func GetUserWithCar(userModel *domain.User) (*user.UserResponse, error) {
	uwcr, err := repository.GetUserWithCar(userModel)
	if err != nil {
		return nil, err
	}

	gcs, err := car.GetCars(userModel)
	if err != nil {
		return nil, err
	}

	uwcr.Cars = gcs

	return uwcr, nil
}

func GetUserWithCarEngines(userModel *domain.User) (*user.UserEnginesResponse, error) {
	uwcr, err := repository.GetUserWithCar(userModel)
	if err != nil {
		return nil, err
	}

	gcwe, err := car.GetCarsWithEngine(userModel)
	if err != nil {
		return nil, err
	}

	return user.NewUserEnginesResponse(uwcr.ID, uwcr.Name, gcwe), nil
}
