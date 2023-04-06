package repository

import (
	"user/internal/app/domain"
	"user/pkg/postgres"
	"user/pkg/response/user"
)

func GetUser(userModel *domain.User) (*user.CarsResponse, error) {
	query := `
		SELECT users.id, name
			FROM users
		WHERE id = $1`

	var userCarsResp user.CarsResponse

	err := postgres.DB.QueryRow(query, userModel.ID).Scan(&userCarsResp.ID, &userCarsResp.Name)
	if err != nil {

		return nil, err
	}

	return &userCarsResp, nil
}
