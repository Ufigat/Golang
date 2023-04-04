package repository

import (
	"user/internal/app/domain"
	"user/pkg/postgres"
	"user/pkg/response/user"
)

func GetUserWithCar(userModel *domain.User) (*user.UserWithCarsResponse, error) {
	var ur user.UserWithCarsResponse
	if err := postgres.DB.QueryRow(`
	SELECT users.id, name
		FROM users
	WHERE id = $1`, userModel.ID).Scan(&ur.ID, &ur.Name); err != nil {
		return nil, err
	}

	return &ur, nil
}
