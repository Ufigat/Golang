package engine

type Response struct {
	ID          int    `json:"engine_id"`
	Designation string `json:"designation"`
}

type DataResponse struct {
	Data  []Response `json:"data"`
	Error string     `json:"error"`
}

type EnigneResponse struct {
	Data  *Response `json:"data"`
	Error string    `json:"error"`
}

type ForCar struct {
	ID     int       `json:"car_id"`
	Engine *Response `json:"engine"`
}

type ForCarResponse struct {
	Data  *ForCar `json:"data"`
	Error string  `json:"error"`
}
