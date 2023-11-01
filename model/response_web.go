package model

type ResponseWebSuccess struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
type ResponseFailWeb struct {
	Error      any    `json:"error"`
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
}
