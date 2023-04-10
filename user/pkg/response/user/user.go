package user

type Response struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	CarIDs []int64 `json:"car_ids"`
}
