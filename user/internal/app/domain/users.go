package domain

import "errors"

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Engine string `json:"engine"`
}

func (u *User) ValidationID() error {
	if u.ID <= 0 {
		return errors.New("field id is not valid")
	}

	return nil
}
