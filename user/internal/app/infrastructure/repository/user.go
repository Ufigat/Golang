package repository

import (
	"user/internal/app/domain"
	"user/pkg/postgres"
	"user/pkg/response/user"
)

func GetUserWithCar(userModel *domain.User) ([]user.UserWithCarResponse, error) {
	rows, err := postgres.DB.Query(`
	SELECT users.id, name, car_id FROM users
	LEFT JOIN user_cars on users.Id = user_cars.user_id
	WHERE id = $1`, userModel.ID)
	if err != nil {
		return nil, err
	}

	var uwcrs []user.UserWithCarResponse
	for rows.Next() {
		var uwcr user.UserWithCarResponse
		err = rows.Scan(&uwcr.ID, &uwcr.Name, &uwcr.CarID)
		if err != nil {
			return nil, err
		}

		uwcrs = append(uwcrs, uwcr)
	}

	return uwcrs, nil
}
