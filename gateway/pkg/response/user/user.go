package user

import (
	"gateway/pkg/response/car"
	"gateway/pkg/response/engine"
)

type UserCarsResponse struct {
	ID   int            `json:"id"`
	Name string         `json:"name"`
	Cars []car.Response `json:"cars"`
}

type UserEnginesResponse struct {
	ID      int               `json:"id"`
	Name    string            `json:"name"`
	Engines []engine.Response `json:"engines"`
}

type Response struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	CarID int    `json:"car_id"`
}

type DataResponse struct {
	Data  []Response `json:"data"`
	Error string     `json:"error"`
}
