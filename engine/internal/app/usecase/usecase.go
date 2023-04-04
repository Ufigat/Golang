package usecase

import (
	"engine/internal/app/domain"
	"engine/internal/app/infrastructure/repository"
	"engine/pkg/request/engine"
)

func GetEngines(er *engine.UserCarsForEnginesRequest) ([]domain.Engine, error) {
	engines, err := repository.GetEngines(er)
	if err != nil {
		return nil, err
	}

	return engines, nil
}

func GetEngine(er *domain.Engine) (*domain.Engine, error) {
	engines, err := repository.GetEngine(er)
	if err != nil {
		return nil, err
	}

	return engines, nil
}
