package engine

import (
	"bytes"
	"car/pkg/request/engine"
	engineRes "car/pkg/response/engine"
	"car/pkg/response/fault"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CarEngines(engineRequest *engine.UserCarsForEngineRequest) ([]engineRes.EngineResponse, error) {
	value, err := json.Marshal(engineRequest)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://localhost:8082/engines", "application/json", bytes.NewBuffer(value))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		var fault fault.FaultResponse
		err = json.NewDecoder(resp.Body).Decode(&fault)
		if err != nil {
			log.Println("CarEngines ", err.Error())
			return nil, err
		}

		return nil, &fault
	}

	var cers []engineRes.EngineResponse
	err = json.NewDecoder(resp.Body).Decode(&cers)
	if err != nil {
		log.Println("CarEngine ", err.Error())
		return nil, err
	}

	return cers, nil
}

func CarEngine(engineRequest *engine.UserCarForEngineRequest) (*engineRes.EngineResponse, error) {
	resp, err := http.Get(fmt.Sprint("http://localhost:8082/engine?id=", engineRequest.EngineID))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		var fault fault.FaultResponse
		err = json.NewDecoder(resp.Body).Decode(&fault)
		if err != nil {
			log.Println("CarEngine ", err.Error())
			return nil, err
		}

		return nil, &fault
	}

	var cers engineRes.EngineResponse
	err = json.NewDecoder(resp.Body).Decode(&cers)
	if err != nil {
		log.Println("CarEngine ", err.Error())
		return nil, err
	}

	return &cers, nil
}
