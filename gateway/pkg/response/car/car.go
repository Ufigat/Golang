package car

import (
	"gateway/pkg/response/engine"
	"gateway/pkg/response/fault"
)

// type Response struct {
// 	ID            int                 `json:"id"`
// 	Name          string              `json:"name"`
// 	Color         string              `json:"color"`
// 	FaultResponse fault.FaultResponse `json:"error"`
// }

type BrandEngineResponse struct {
	ID             int                     `json:"brand_id"`
	Brand          string                  `json:"brand"`
	EngineResponse []engine.EngineResponse `json:"engines"`
	FaultResponse  fault.Response          `json:"error"`
}

type EngineResponse struct {
	ID             int                   `json:"car_id"`
	EngineResponse engine.EngineResponse `json:"engine"`
}
