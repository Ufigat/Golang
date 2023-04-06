package user

import (
	"gateway/pkg/response/engine"
	"gateway/pkg/response/fault"
)

type CarsResponse struct {
	ID   int        `json:"id"`
	Name string     `json:"name"`
	Cars []Response `json:"cars"`
}

type DataResponse struct {
	Response *CarsResponse   `json:"cars"`
	Error    *fault.Response `json:"error"`
}

type Response struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type EnginesResponse struct {
	ID     int                     `json:"user_id"`
	Name   string                  `json:"name"`
	Engine []engine.EngineResponse `json:"engines"`
}
