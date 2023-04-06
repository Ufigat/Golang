package engine

import (
	"engine/internal/app/domain"
	"engine/pkg/response/fault"
)

type Response struct {
	Engine domain.Engine   `json:"engine,omitempty"`
	Error  *fault.Response `json:"error,omitempty"`
}

type LinksResponse struct {
	Engines []domain.Engine `json:"engines,omitempty"`
	Error   *fault.Response `json:"error,omitempty"`
}
