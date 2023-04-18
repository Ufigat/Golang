package engine

import "gateway/pkg/response/fault"

type Response struct {
	ID          int    `json:"engine_id"`
	Designation string `json:"designation"`
}

type DataResponse struct {
	Data  []Response      `json:"data"`
	Error *fault.Response `json:"error"`
}

type EnigneResponse struct {
	Data  *Response       `json:"data"`
	Error *fault.Response `json:"error"`
}

type ForCar struct {
	ID     string    `json:"car_id"`
	Engine *Response `json:"engine"`
}

type ForCarResponse struct {
	Data  *ForCar         `json:"data"`
	Error *fault.Response `json:"error"`
}

type ByBrandResponse struct {
	Brand   string     `json:"brand"`
	Engines []Response `json:"engines"`
}
