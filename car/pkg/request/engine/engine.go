package engine

type EnginesRequest struct {
	EngineIDs []int `json:"engine_ids"`
}

type Request struct {
	EngineID int `json:"engine_id"`
}
