package car

type Response struct {
	ID       int    `json:"id,omitempty"`
	Brand    string `json:"name,omitempty"`
	Color    string `json:"color,omitempty"`
	EngineID int    `json:"engine_id,omitempty"`
}

type EngineIDResponse struct {
	ID       int `json:"id"`
	EngineID int `json:"engine_id"`
}
