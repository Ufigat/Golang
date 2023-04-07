package car

type Response struct {
	ID       int    `json:"id"`
	Brand    string `json:"name,omitempty"`
	Color    string `json:"color,omitempty"`
	EngineID int    `json:"engine_id,omitempty"`
}

type DataResponse struct {
	Data  []Response `json:"data"`
	Error string     `json:"error"`
}
