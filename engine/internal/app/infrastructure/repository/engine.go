package repository

import (
	"engine/internal/app/domain"
	"engine/pkg/postgres"
	"engine/pkg/request/engine"
	"engine/pkg/response/fault"

	"github.com/lib/pq"
)

func GetEngines(er *engine.IDsRequest) ([]domain.Engine, error) {
	query := `
		SELECT id as engine_id, engines.designation as designation
			FROM engines
		WHERE id = any($1)`

	rows, err := postgres.DB.Query(query, pq.Array(er.EngineID))

	if err != nil {
		return nil, err
	}

	var engineLinks []domain.Engine

	for rows.Next() {
		var en domain.Engine
		err = rows.Scan(&en.ID, &en.Designation)
		if err != nil {
			return nil, err
		}

		engineLinks = append(engineLinks, en)
	}

	if len(engineLinks) == 0 {
		return nil, fault.NewResponse("no rows in result set")
	}

	return engineLinks, nil
}

func GetEngine(er *domain.Engine) (*domain.Engine, error) {
	query := `
		SELECT id as engine_id, engines.designation as designation
			FROM engines
		WHERE id = $1`

	var engine domain.Engine

	if err := postgres.DB.QueryRow(query, er.ID).Scan(&engine.ID, &engine.Designation); err != nil {
		return nil, err
	}

	return &engine, nil
}
