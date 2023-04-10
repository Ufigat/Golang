package repository

import (
	"car/pkg/postgres"
	carReq "car/pkg/request/car"
	"car/pkg/response/car"

	"github.com/lib/pq"
)

func GetCarByUser(carsIDs []int) ([]car.Response, error) {
	query := `
		SELECT cars.id, brands.name, colors.name
			FROM cars
			JOIN brands ON cars.brand_id = brands.id
			JOIN colors ON cars.color_id = colors.id
		WHERE cars.id = any($1)`

	rows, err := postgres.DB.Query(query, pq.Array(carsIDs))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

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

func GetCarEngineByUser(carsIDs []int) ([]car.Response, error) {
	query := `
		SELECT cars.id, cars.engine_id
			FROM cars
			JOIN brands ON cars.brand_id = brands.id
		WHERE cars.id = any($1)`

	rows, err := postgres.DB.Query(query, pq.Array(carsIDs))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var carLinksResp []car.Response

	for rows.Next() {
		var carResp car.Response

		err = rows.Scan(&carResp.ID, &carResp.EngineID)
		if err != nil {
			return nil, err
		}

		carLinksResp = append(carLinksResp, carResp)
	}

	return carLinksResp, nil
}

func GetCarEngineByBrand(req *carReq.Request) ([]car.Response, error) {
	query := `
		SELECT distinct cars.engine_id
			FROM cars
			JOIN brands ON cars.brand_id = brands.id
		WHERE brands.name = $1`

	rows, err := postgres.DB.Query(query, req.Brand)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var carLinksResp []car.Response

	for rows.Next() {
		var carResp car.Response

		err = rows.Scan(&carResp.EngineID)
		if err != nil {
			return nil, err
		}

		carLinksResp = append(carLinksResp, carResp)
	}

	return carLinksResp, nil
}

func GetCarEngine(req *carReq.Request) (*car.EngineIDResponse, error) {
	query := `
		SELECT cars.id as car_id, cars.engine_id
			FROM cars
		WHERE id = $1`

	var carEngineResp car.EngineIDResponse

	err := postgres.DB.QueryRow(query, req.ID).Scan(&carEngineResp.ID, &carEngineResp.EngineID)
	if err != nil {
		return nil, err
	}

	return &carEngineResp, nil
}
