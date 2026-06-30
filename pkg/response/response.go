package response

type Response struct {
	Success bool           `json:"success"`
	Message string         `json:"message,omitempty"`
	Data    interface{}    `json:"data,omitempty"`
	Meta    interface{}    `json:"meta,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
}
