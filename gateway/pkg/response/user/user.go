package user

import (
	"gateway/pkg/response/engine"
	"user/pkg/response/car"
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
