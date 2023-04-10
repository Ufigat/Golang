package repository

import (
	"engine/pkg/postgres"
	resp "engine/pkg/response/engine"
	"engine/pkg/response/fault"

	"github.com/lib/pq"
)

func GetEngines(reqIDs []int) ([]resp.Engine, error) {
	query := `
		SELECT id, designation
			FROM engines
		WHERE id = any($1)`

	rows, err := postgres.DB.Query(query, pq.Array(reqIDs))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var respLinks []resp.Engine

	for rows.Next() {
		var resp resp.Engine

		err = rows.Scan(&resp.ID, &resp.Designation)
		if err != nil {
			return nil, err
		}

		respLinks = append(respLinks, resp)
	}

	if len(respLinks) == 0 {
		return nil, fault.NewResponse("no rows in result set")
	}

	return respLinks, nil
}

func GetEngine(engineID int) (*resp.Engine, error) {
	query := `
		SELECT id, designation
			FROM engines
		WHERE id = $1`

	var resp resp.Engine

	if err := postgres.DB.QueryRow(query, engineID).Scan(&resp.ID, &resp.Designation); err != nil {
		return nil, err
	}

	return &resp, nil
}
