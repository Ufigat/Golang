package engine

import "car/pkg/response/fault"

type Response struct {
	ID          int    `json:"engine_id"`
	Designation string `json:"designation"`
}

type DataResponse struct {
	Engine Response        `json:"engine,omitempty"`
	Error  *fault.Response `json:"error,omitempty"`
}

type LinksResponse struct {
	Engines []Response      `json:"engines,omitempty"`
	Error   *fault.Response `json:"error,omitempty"`
}
