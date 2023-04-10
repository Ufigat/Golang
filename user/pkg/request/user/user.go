package user

import "errors"

func ValidationID(userID int) error {
	if userID <= 0 {
		return errors.New("field id is not valid")
	}

	return nil
}
