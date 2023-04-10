package user

import "errors"

type Request struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Engine string `json:"engine"`
}

func (u *Request) ValidationID() error {
	if u.ID <= 0 {
		return errors.New("field id is not valid")
	}

	return nil
}
