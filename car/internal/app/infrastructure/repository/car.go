package repository

import (
	"car/internal/app/domain"
	"car/pkg/postgres"
	"car/pkg/request/engine"
	"car/pkg/response/car"
)

func GetCarByUser(carModel *domain.Car) ([]car.CarResponse, error) {
	rows, err := postgres.DB.Query(`
	Select cars.id as car_id, brands.name, colors.name
	FROM  user_cars
		LEFT JOIN cars on user_cars.car_id = cars.id
		LEFT JOIN brands on cars.brand_id = brands.id
		LEFT JOIN colors on cars.color_id = colors.id
	WHERE user_cars.user_id = $1`, carModel.UserID)
	if err != nil {
		return nil, err
	}

	var crs []car.CarResponse
	for rows.Next() {
		var cr car.CarResponse
		err = rows.Scan(&cr.ID, &cr.Brand, &cr.Color)
		if err != nil {
			return nil, err
		}

		crs = append(crs, cr)
	}

	return crs, nil
}

func GetCarWithUserAndEngineId(userModel *domain.Car) (*engine.UserCarsForEngineRequest, error) {
	rows, err := postgres.DB.Query(`
	Select distinct(cars.engine_id)
	FROM user_cars
		LEFT JOIN cars on user_cars.car_id = cars.id
	WHERE user_cars.user_id = $1`, userModel.UserID)
	if err != nil {
		return nil, err
	}

	var ucfer engine.UserCarsForEngineRequest
	for rows.Next() {
		var engineID int
		err = rows.Scan(&engineID)
		if err != nil {
			return nil, err
		}

		ucfer.EngineIDs = append(ucfer.EngineIDs, engineID)
	}

	return &ucfer, nil
}

func GetCarByBrand(carModel *domain.Car) ([]car.CarWithEngineIDResponse, error) {
	rows, err := postgres.DB.Query(`
	Select distinct cars.brand_id, cars.engine_id, brands.name
		FROM cars
		Left JOIN brands on cars.brand_id = brands.id
	WHERE brands.name = $1`, carModel.Brand)
	if err != nil {
		return nil, err
	}

	var crwes []car.CarWithEngineIDResponse
	for rows.Next() {
		var crwe car.CarWithEngineIDResponse
		err = rows.Scan(&crwe.ID, &crwe.EngineID, &crwe.Brand)
		if err != nil {
			return nil, err
		}

		crwes = append(crwes, crwe)
	}

	return crwes, nil
}

func GetCarEngine(carModel *domain.Car) (*car.CarIDWithEngineIDResponse, error) {
	var cr car.CarIDWithEngineIDResponse
	if err := postgres.DB.QueryRow(`
	SELECT cars.id as car_id, cars.engine_id
		FROM cars
	WHERE id = $1`, carModel.ID).Scan(&cr.ID, &cr.EngineID); err != nil {
		return nil, err
	}

	return &cr, nil
}
