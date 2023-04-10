package repository

import (
	"user/pkg/postgres"
	"user/pkg/request/user"
	res "user/pkg/response/user"
)

func GetUser(req *user.Request) ([]res.Response, error) {
	query := `
		SELECT users.id, users.name, user_cars.car_id
			FROM users
		JOIN user_cars ON users.id = user_cars.user_id
		WHERE id = $1`

	var userLinks []res.Response

	rows, err := postgres.DB.Query(query, req.ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var elem res.Response

		err = rows.Scan(&elem.ID, &elem.Name, &elem.CarID)
		if err != nil {
			return nil, err
		}

		userLinks = append(userLinks, elem)
	}

	return userLinks, nil
}
