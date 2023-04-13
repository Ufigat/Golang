package util

type Response struct {
	Data  any `json:"data"`
	Error any `json:"error"`
}

type Client struct {
	ID int `json:"user_id"`
}
