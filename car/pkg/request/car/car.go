package car

import "errors"

type IDsRequest struct {
	CarsIDs []int `json:"cars_ids"`
}

type Request struct {
	ID    int    `json:"id,omitempty"`
	Brand string `json:"brand,omitempty"`
}

func (c *Request) ValidationID() error {
	if c.ID <= 0 {
		return errors.New("field id is not valid")
	}

	return nil
}

func (c *Request) ValidationBrand() error {
	if c.Brand == "" {
		return errors.New("field brand is not valid")
	}

	return nil
}

func (c *IDsRequest) ValidationIDs() error {
	for i := range c.CarsIDs {
		if c.CarsIDs[i] <= 0 {
			return errors.New("field id is not valid")
		}
	}

	return nil
}
