package engine

type UserCarsForEnginesRequest struct {
	EngineID []int `json:"engine_ids"`
}

type CarForEngineRequest struct {
	EngineID int `json:"engine_id"`
}
