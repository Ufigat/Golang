package usecase

import (
	"engine/internal/app/domain"
	"engine/internal/app/infrastructure/repository"
	"engine/pkg/request/engine"
)

func GetEngines(er *engine.UserCarsForEngineRequest) ([]domain.Engine, error) {
	engines, err := repository.GetEngines(er)
	if err != nil {
		return nil, err
	}

	return engines, nil
}
