package engine

import "errors"

type IDsRequest struct {
	EngineID []int `json:"engine_ids"`
}

type Request struct {
	ID int `json:"engine_id"`
}

func (r *Request) ValidationID() error {
	if r.ID <= 0 {
		return errors.New("field id is not valid")
	}

	return nil
}
