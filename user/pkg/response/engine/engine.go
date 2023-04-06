package engine

type Response struct {
	ID          int    `json:"engine_id"`
	Designation string `json:"designation"`
}

type LinksResponse struct {
	ID     int        `json:"id"`
	Name   string     `json:"name"`
	Engine []Response `json:"engines"`
}

type DataResponse struct {
	Response []Response `json:"data"`
	Error    string     `json:"error"`
}
