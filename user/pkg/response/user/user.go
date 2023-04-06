package user

import (
	"user/pkg/response/car"
)

type CarsResponse struct {
	ID   int            `json:"id,omitempty"`
	Name string         `json:"name,omitempty"`
	Cars []car.Response `json:"cars,omitempty"`
}
