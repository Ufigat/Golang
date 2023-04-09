package repository

import (
	"user/internal/app/domain"
	"user/pkg/postgres"
	"user/pkg/response/user"
)

func GetUser(userModel *domain.User) ([]user.Response, error) {
	query := `
		SELECT users.id, users.name, user_cars.car_id
			FROM users
		JOIN user_cars ON users.id = user_cars.user_id
		WHERE id = $1`

	var resp []user.Response

	rows, err := postgres.DB.Query(query, userModel.ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var elem user.Response

		err = rows.Scan(&elem.ID, &elem.Name, &elem.CarID)
		if err != nil {
			return nil, err
		}

		resp = append(resp, elem)
	}

	return resp, nil
}
