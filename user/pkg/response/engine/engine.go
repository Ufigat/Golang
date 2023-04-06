package engine

import "user/pkg/response/fault"

type Response struct {
	ID          int    `json:"engine_id"`
	Designation string `json:"designation"`
}

type LinksResponse struct {
	ID     int             `json:"id"`
	Name   string          `json:"name"`
	Engine []Response      `json:"engines"`
	Error  *fault.Response `json:"error"`
}

type DataResponse struct {
	Engine Response        `json:"engine,omitempty"`
	Error  *fault.Response `json:"error,omitempty"`
}
