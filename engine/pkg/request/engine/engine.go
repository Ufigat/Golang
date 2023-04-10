package engine

import "errors"

func ValidationID(engineId int) error {
	if engineId <= 0 {
		return errors.New("field id is not valid")
	}

	return nil
}
