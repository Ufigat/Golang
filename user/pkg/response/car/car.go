package car

type Response struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type DataResponse struct {
	Response []Response `json:"data"`
	Error    string     `json:"error"`
}
