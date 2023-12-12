package response

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(code int, message string) *errorResponse {
	return &errorResponse{
		Code:    code,
		Message: message,
	}
}
