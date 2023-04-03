package car

import "car/pkg/response/engine"

type CarResponse struct {
	ID    int    `json:"id"`
	Brand string `json:"name"`
	Color string `json:"color"`
}

type CarResponseWithEngineID struct {
	ID       int    `json:"id,omitempty"`
	Brand    string `json:"brand,omitempty"`
	EngineID int    `json:"engine_id"`
}

type CarResponseWithEngineByBrand struct {
	ID             int                     `json:"brand_id"`
	Brand          string                  `json:"brand"`
	EngineResponse []engine.EngineResponse `json:"engines"`
}

func NewCarResponseWithEngineByBrand(id int, brand string, engineResponse []engine.EngineResponse) *CarResponseWithEngineByBrand {
	return &CarResponseWithEngineByBrand{
		ID:             id,
		Brand:          brand,
		EngineResponse: engineResponse}
}
