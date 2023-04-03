package engine

import "engine/internal/app/domain"

type EngineResponse struct {
	ID          int             `json:"brand_id"`
	Brand       string          `json:"brand"`
	Designation []domain.Engine `json:"designation"`
}

type EngineWithBrand struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	EngineID    int    `json:"engine_id"`
	Designation string `json:"designation"`
}
