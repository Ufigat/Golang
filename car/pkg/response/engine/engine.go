package engine

type Response struct {
	ID          int    `json:"engine_id"`
	Designation string `json:"designation"`
}

type DataResponses struct {
	Engines []Response `json:"data"`
	Error   string     `json:"error"`
}

type DataResponse struct {
	Engines *Response `json:"data"`
	Error   string    `json:"error"`
}
