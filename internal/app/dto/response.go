package dto

type Response struct {
	Error  bool        `json:"error"`
	Result interface{} `json:"result"`
}

type ErrorResponse struct {
	Error  bool        `json:"error"`
	Result interface{} `json:"result"`
}
