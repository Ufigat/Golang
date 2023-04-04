package car

import "gateway/pkg/response/engine"

type CarResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type CarBrandWithEngineResponse struct {
	ID             int                     `json:"brand_id"`
	Brand          string                  `json:"brand"`
	EngineResponse []engine.EngineResponse `json:"engines"`
}

type CarWithEngineResponse struct {
	ID             int                   `json:"car_id"`
	EngineResponse engine.EngineResponse `json:"engine"`
}
