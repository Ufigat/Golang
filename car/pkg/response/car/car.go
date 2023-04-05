package car

import "car/pkg/response/engine"

type CarResponse struct {
	ID    int    `json:"id"`
	Brand string `json:"name"`
	Color string `json:"color"`
}

type CarWithEngineIDResponse struct {
	ID       int    `json:"id"`
	Brand    string `json:"brand"`
	EngineID int    `json:"engine_id"`
}

type CarIDWithEngineIDResponse struct {
	ID       int `json:"id"`
	EngineID int `json:"engine_id"`
}

type CarWithEngineByBrandResponse struct {
	ID             int                     `json:"brand_id"`
	Brand          string                  `json:"brand"`
	EngineResponse []engine.EngineResponse `json:"engines"`
}

func NewCarResponseWithEngineByBrand(id int, brand string, engineResponse []engine.EngineResponse) *CarWithEngineByBrandResponse {
	return &CarWithEngineByBrandResponse{
		ID:             id,
		Brand:          brand,
		EngineResponse: engineResponse,
	}
}

type CarWithEngineResponse struct {
	ID             int                   `json:"car_id"`
	EngineResponse engine.EngineResponse `json:"engine"`
}

func NewCarWithEngineResponse(id int, engineResponse *engine.EngineResponse) *CarWithEngineResponse {
	return &CarWithEngineResponse{
		ID:             id,
		EngineResponse: *engineResponse,
	}
}
