package fault

type Response struct {
	Message string `json:"message"`
}

func (f *Response) Error() string {
	return f.Message
}

func NewResponse(message string) *Response {
	return &Response{Message: message}
}
