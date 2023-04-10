package engine

import (
	"engine/pkg/response/fault"
)

type Engine struct {
	ID          int    `json:"engine_id"`
	Designation string `json:"designation"`
}
type Response struct {
	Engine *Engine         `json:"engine"`
	Error  *fault.Response `json:"error"`
}

type LinksResponse struct {
	Engines []Engine        `json:"engines"`
	Error   *fault.Response `json:"error"`
}
