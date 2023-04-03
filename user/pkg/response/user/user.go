package user

import (
	"user/pkg/response/car"
	"user/pkg/response/engine"
)

type UserResponse struct {
	ID   int               `json:"id"`
	Name string            `json:"name"`
	Cars []car.CarResponse `json:"cars"`
}

type UserEnginesResponse struct {
	ID      int                     `json:"user_id"`
	Name    string                  `json:"name"`
	Engines []engine.EngineResponse `json:"engines"`
}

func NewUserEnginesResponse(id int, name string, engines []engine.EngineResponse) *UserEnginesResponse {
	return &UserEnginesResponse{ID: id, Name: name, Engines: engines}
}
