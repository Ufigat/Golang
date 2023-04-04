package domain

import "errors"

type Car struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Brand  string `json:"brand"`
	Color  string `json:"color"`
}

func (c *Car) ValidationID() error {
	if c.ID <= 0 {
		return errors.New("field id is not valid")
	}

	return nil
}

func (c *Car) ValidationUserID() error {
	if c.UserID <= 0 {
		return errors.New("field user id is not valid")
	}

	return nil
}

func (c *Car) ValidationBrand() error {
	if c.Brand == "" {
		return errors.New("field brand is not valid")
	}

	return nil
}
