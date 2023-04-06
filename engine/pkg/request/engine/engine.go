package engine

type IDsRequest struct {
	EngineID []int `json:"engine_ids"`
}

type Request struct {
	EngineID int `json:"engine_id"`
}
