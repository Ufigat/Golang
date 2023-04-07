package car

import (
	"car/pkg/response/engine"
)

type Response struct {
	ID    int    `json:"id"`
	Brand string `json:"name"`
	Color string `json:"color"`
}

type EngineIDBrandResponse struct {
	ID       int    `json:"id"`
	Brand    string `json:"brand"`
	EngineID int    `json:"engine_id"`
}

type EngineIDResponse struct {
	ID       int `json:"id"`
	EngineID int `json:"engine_id"`
}

type EngineByBrandResponse struct {
	ID             int               `json:"brand_id"`
	Brand          string            `json:"brand"`
	EngineResponse []engine.Response `json:"engines"`
}

type EngineResponse struct {
	ID      int              `json:"id"`
	Engines *engine.Response `json:"engine"`
}
