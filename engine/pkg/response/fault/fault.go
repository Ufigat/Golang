package fault

type Response struct {
	Message string `json:"error"`
}

func NewResponse(message string) *Response {
	return &Response{Message: message}
}
