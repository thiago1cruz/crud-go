package exceptions

import "net/http"

type ErrorResponse struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int64    `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *ErrorResponse) Error() string {
	return r.Message
}

func NewErrorResponse(message, err string, code int64, causes []Causes) *ErrorResponse {
	return &ErrorResponse{Message: message, Err: err, Code: code, Causes: causes}
}

func NewBadRequestError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Err:     "Validation error",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}

}
func NewNotFoundError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Err:     "not found",
		Code:    http.StatusNotFound,
	}
}

func NewForBiddenError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
