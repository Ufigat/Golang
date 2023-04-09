package user

type Response struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	CarID int    `json:"car_id"`
}
