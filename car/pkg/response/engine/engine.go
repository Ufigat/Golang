package engine

type Response struct {
	ID          int    `json:"engine_id"`
	Designation string `json:"designation"`
}

// type DataResponse struct {
// 	Engine Response        `json:"engine"`
// 	Error  *fault.Response `json:"error"`
// }

type DataResponse struct {
	Engines []Response `json:"data"`
	Error   string     `json:"error"`
}
