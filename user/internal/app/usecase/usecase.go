package usecase

import (
	"user/internal/app/domain"
	"user/internal/app/infrastructure/repository"
	"user/pkg/response/user"
)

func GetUserWithCar(userModel *domain.User) ([]user.UserWithCarResponse, error) {
	uwcrs, err := repository.GetUserWithCar(userModel)
	if err != nil {
		return nil, err
	}

	return uwcrs, nil
}
