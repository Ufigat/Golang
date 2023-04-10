package domain

type Car struct {
	ID       int    `json:"id"`
	BrandID  string `json:"brand_id"`
	ColorID  string `json:"color_id"`
	EngineID string `json:"engine_id"`
}
