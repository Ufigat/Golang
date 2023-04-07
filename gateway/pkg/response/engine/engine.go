package engine

type Response struct {
	ID          int    `json:"engine_id"`
	Designation string `json:"designation"`
}

type DataResponse struct {
	Data  []Response `json:"data"`
	Error string     `json:"error"`
}
