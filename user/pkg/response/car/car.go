package car

import "user/pkg/response/engine"

type CarResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type UserCarEnginesResponse struct {
	ID     int                     `json:"id"`
	Name   string                  `json:"name"`
	Engine []engine.EngineResponse `json:"engines"`
}
