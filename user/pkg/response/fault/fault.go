package fault

type Response struct {
	Message string `json:"message"`
}

func (fr *Response) Error() string {
	return fr.Message
}

func NewResponse(message string) *Response {
	return &Response{Message: message}
}
