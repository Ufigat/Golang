package engine

import (
	"bytes"
	"car/pkg/request/engine"
	engineRes "car/pkg/response/engine"
	"encoding/json"
	"io"
	"net/http"
)

func GetCarEngine(engineRequest *engine.UserCarsForEngineRequest) ([]engineRes.EngineResponse, error) {
	value, err := json.Marshal(engineRequest)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://localhost:8082/engines", "application/json", bytes.NewBuffer(value))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var cers []engineRes.EngineResponse
	err = json.Unmarshal(body, &cers)
	if err != nil {
		return nil, err
	}

	return cers, nil
}
