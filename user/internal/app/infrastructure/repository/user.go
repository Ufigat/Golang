package repository

import (
	"user/pkg/postgres"
	res "user/pkg/response/user"

	"github.com/lib/pq"
)

func GetUser(userID int) (*res.Response, error) {
	query := `
		SELECT users.id, users.name, array_agg(user_cars.car_id)
			FROM users
		JOIN user_cars ON users.id = user_cars.user_id
		WHERE id = $1
		GROUP BY users.id, users.name`

	var resp res.Response

	err := postgres.DB.QueryRow(query, userID).Scan(&resp.ID, &resp.Name, pq.Array(&resp.CarIDs))
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
