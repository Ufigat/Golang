package user

type UserWithCarResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	CarID int    `json:"car_id"`
}
