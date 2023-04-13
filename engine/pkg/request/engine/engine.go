package engine

import (
	"errors"
	"fmt"
)

func ValidationID(engineId int) error {
	if engineId <= 0 {
		return errors.New("field id is not valid")
	}

	return nil
}

func ValidationIDs(engineIDsReq []int) error {
	for i := range engineIDsReq {
		if engineIDsReq[i] <= 0 {
			return errors.New(fmt.Sprintln("ID is not valid: ", engineIDsReq[i]))
		}
	}

	return nil
}
