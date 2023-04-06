package repository

import (
	"car/internal/app/domain"
	"car/pkg/postgres"
	"car/pkg/request/engine"
	"car/pkg/response/car"
)

func GetCarByUser(carModel *domain.Car) ([]car.Response, error) {
	query := `
		SELECT cars.id as car_id, brands.name, colors.name
		FROM  user_cars
			LEFT JOIN cars ON user_cars.car_id = cars.id
			LEFT JOIN brands ON cars.brand_id = brands.id
			LEFT JOIN colors ON cars.color_id = colors.id
		WHERE user_cars.user_id = $1`

	rows, err := postgres.DB.Query(query, carModel.UserID)
	if err != nil {
		return nil, err
	}

	var carLinksResp []car.Response

	for rows.Next() {
		var carResp car.Response
		err = rows.Scan(&carResp.ID, &carResp.Brand, &carResp.Color)
		if err != nil {

			return nil, err
		}

		carLinksResp = append(carLinksResp, carResp)
	}

	return carLinksResp, nil
}

func GetUserCarAndEngine(userModel *domain.Car) (*engine.EnginesRequest, error) {
	query := `
		SELECT distinct(cars.engine_id)
		FROM user_cars
			LEFT JOIN cars ON user_cars.car_id = cars.id
		WHERE user_cars.user_id = $1`

	rows, err := postgres.DB.Query(query, userModel.UserID)
	if err != nil {

		return nil, err
	}

	var engineReg engine.EnginesRequest

	for rows.Next() {
		var engineID int
		err = rows.Scan(&engineID)
		if err != nil {

			return nil, err
		}

		engineReg.EngineIDs = append(engineReg.EngineIDs, engineID)
	}

	return &engineReg, nil
}

func GetCarByBrand(carModel *domain.Car) ([]car.EngineIDBrandResponse, error) {
	query := `
		SELECT distinct cars.brand_id, cars.engine_id, brands.name
			FROM cars
			Left JOIN brands ON cars.brand_id = brands.id
		WHERE brands.name = $1`

	rows, err := postgres.DB.Query(query, carModel.Brand)
	if err != nil {

		return nil, err
	}

	var carEngineBrandRespLinks []car.EngineIDBrandResponse

	for rows.Next() {
		var carEngineResp car.EngineIDBrandResponse
		err = rows.Scan(&carEngineResp.ID, &carEngineResp.EngineID, &carEngineResp.Brand)
		if err != nil {

			return nil, err
		}

		carEngineBrandRespLinks = append(carEngineBrandRespLinks, carEngineResp)
	}

	return carEngineBrandRespLinks, nil
}

func GetCarEngine(carModel *domain.Car) (*car.EngineIDResponse, error) {
	query := `
		SELECT cars.id as car_id, cars.engine_id
			FROM cars
		WHERE id = $1`

	var carEngineResp car.EngineIDResponse

	err := postgres.DB.QueryRow(query, carModel.ID).Scan(&carEngineResp.ID, &carEngineResp.EngineID)
	if err != nil {

		return nil, err
	}

	return &carEngineResp, nil
}
