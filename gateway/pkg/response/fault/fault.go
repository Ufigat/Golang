package fault

type FaultResponse struct {
	Message string `json:"message"`
}

func NewFaultResponse(message string) *FaultResponse {
	return &FaultResponse{Message: message}
}
