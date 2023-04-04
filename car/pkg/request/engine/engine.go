package engine

type UserCarsForEngineRequest struct {
	EngineIDs []int `json:"engine_ids"`
}

type UserCarForEngineRequest struct {
	EngineID int `json:"engine_id"`
}
