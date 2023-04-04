package fault

type FaultResponse struct {
	Message string `json:"message"`
}

func (fr *FaultResponse) Error() string {
	return fr.Message
}

func NewFaultResponse(message string) *FaultResponse {
	return &FaultResponse{Message: message}
}
