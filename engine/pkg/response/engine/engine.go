package engine

import (
	"engine/internal/app/domain"
	"engine/pkg/response/fault"
)

type Response struct {
	Engine *domain.Engine  `json:"engine"`
	Error  *fault.Response `json:"error"`
}

type LinksResponse struct {
	Engines []domain.Engine `json:"engines"`
	Error   *fault.Response `json:"error"`
}
