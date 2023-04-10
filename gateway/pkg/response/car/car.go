package car

import (
	"gateway/pkg/response/engine"
	"gateway/pkg/response/fault"
)

type Response struct {
	ID       int    `json:"id,omitempty"`
	Brand    string `json:"name,omitempty"`
	Color    string `json:"color,omitempty"`
	EngineID int    `json:"engine_id,omitempty"`
}

type DataResponse struct {
	Data  []Response      `json:"data"`
	Error *fault.Response `json:"error"`
}

type EngineByBrandResponse struct {
	Name    string            `json:"brand"`
	Engines []engine.Response `json:"engines"`
}

type CarEngineResponse struct {
	Data  *Response       `json:"data"`
	Error *fault.Response `json:"error"`
}
