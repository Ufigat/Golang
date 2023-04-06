package fault

type Response struct {
	Message string `json:"message"`
}

func (fr *Response) Error() string {
	return fr.Message
}
