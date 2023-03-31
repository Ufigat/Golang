package usecase

import "user/internal/domain"

func GetUserCars(userModel *domain.User) (*userResponse.GetUserCarsResponse, error) {
	cars, err := repository.GetUserCars(userModel)
	if err != nil {
		return nil, err
	}

	var eqs []carResponse.EquipmentResponse
	var ucr userResponse.GetUserCarsResponse
	for _, car := range cars {
		eqs = append(eqs, carResponse.EquipmentResponse{CarID: car.CarID, Brand: car.Brand, Color: car.Color, Engine: car.Engine})
		ucr.ID = car.ID
		ucr.Name = car.Name
		ucr.Equipments = eqs
	}

	return &ucr, nil
}
