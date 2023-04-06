package user

import (
	"user/pkg/response/car"
)

type CarsResponse struct {
	ID   int            `json:"id"`
	Name string         `json:"name"`
	Cars []car.Response `json:"cars"`
}
