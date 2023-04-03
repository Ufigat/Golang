package domain

import "errors"

type Car struct {
	ID    int    `json:"id"`
	Owner string `json:"owner"`
	Brand string `json:"brand"`
	Color string `json:"color"`
}

func (c *Car) ValidationID() error {
	if c.ID <= 0 {
		return errors.New("field id is not valid")
	}

	return nil
}

func (c *Car) ValidationBrand() error {
	if c.Brand == "" {
		return errors.New("field brand is not valid")
	}

	return nil
}
