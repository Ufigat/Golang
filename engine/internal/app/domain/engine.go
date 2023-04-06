package domain

import (
	"errors"
)

type Engine struct {
	ID          int    `json:"engine_id"`
	Designation string `json:"designation"`
}

func (e *Engine) ValidationID() error {
	if e.ID <= 0 {
		return errors.New("field id is not valid")
	}

	return nil
}
