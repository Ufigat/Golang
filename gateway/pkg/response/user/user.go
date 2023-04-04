package user

import (
	"gateway/pkg/response/car"
	"gateway/pkg/response/engine"
)

type UserCarsResponse struct {
	ID   int               `json:"id"`
	Name string            `json:"name"`
	Cars []car.CarResponse `json:"cars"`
}

type UserEnginesResponse struct {
	ID     int                     `json:"user_id"`
	Name   string                  `json:"name"`
	Engine []engine.EngineResponse `json:"engines"`
}

type CarResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
