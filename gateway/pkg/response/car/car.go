package car

import (
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

type CarEngineResponse struct {
	Data  *Response       `json:"data"`
	Error *fault.Response `json:"error"`
}
