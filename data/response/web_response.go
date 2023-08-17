package response

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErorrResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
}