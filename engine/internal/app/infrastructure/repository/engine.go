package repository

import (
	"engine/internal/app/domain"
	"engine/pkg/postgres"
	"engine/pkg/request/engine"

	"github.com/lib/pq"
)

func GetEngines(er *engine.UserCarsForEnginesRequest) ([]domain.Engine, error) {
	query := `
		SELECT id as engine_id, engines.designation as designation
			FROM engines
		WHERE id = any($1)`

	rows, err := postgres.DB.Query(query, pq.Array(er.EngineID))

	if err != nil {
		return nil, err
	}

	var ens []domain.Engine

	for rows.Next() {
		var en domain.Engine
		err = rows.Scan(&en.ID, &en.Designation)
		if err != nil {
			return nil, err
		}

		ens = append(ens, en)
	}

	return ens, nil
}

func GetEngine(er *domain.Engine) (*domain.Engine, error) {
	query := `
		SELECT id as engine_id, engines.designation as designation
			FROM engines
		WHERE id = $1`

	var en domain.Engine

	if err := postgres.DB.QueryRow(query, er.ID).Scan(&en.ID, &en.Designation); err != nil {
		return nil, err
	}

	return &en, nil
}
