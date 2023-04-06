package car

import (
	"car/pkg/response/engine"
)

type Response struct {
	ID    int    `json:"id,omitempty"`
	Brand string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
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
	ID             int                  `json:"brand_id"`
	Brand          string               `json:"brand"`
	EngineResponse *engine.DataResponse `json:"engines"`
}

// func NewResponseWithEngineByBrand(id int, brand string, engineResponse *engine.DataResponse) *EngineByBrandResponse {
// 	return &EngineByBrandResponse{
// 		ID:             id,
// 		Brand:          brand,
// 		EngineResponse: engineResponse,
// 	}
// }

// type EngineResponse struct {
// 	ID             int                  `json:"car_id"`
// 	EngineResponse *engine.DataResponse `json:"engine"`
// }

// func NewEngineResponse(id int, engineResponse *engine.DataResponse) *EngineResponse {
// 	return &EngineResponse{
// 		ID:             id,
// 		EngineResponse: engineResponse,
// 	}
// }
