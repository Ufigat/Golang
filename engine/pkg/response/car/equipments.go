package car

type EquipmentResponse struct {
	CarID  int    `json:"car_id"`
	Brand  string `json:"brand"`
	Color  string `json:"color"`
	Engine string `json:"engine"`
}

type CarResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	CarID  int    `json:"car_id"`
	Brand  string `json:"brand"`
	Engine string `json:"engine"`
	Color  string `json:"color"`
}
