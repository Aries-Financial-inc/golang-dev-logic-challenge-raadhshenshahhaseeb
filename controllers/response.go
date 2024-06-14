package controllers

type StandardResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewStandardResponse(code int, msg string, data interface{}) *StandardResponse {
	// Nil data would be set to empty
	if data == nil {
		data = ""
	}

	return &StandardResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}
