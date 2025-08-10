package response

import "fmt"

type ErrorResponse struct {
	Code    int
	Message string
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Message)
}
func Exception(code int, message string) error {
	return &ErrorResponse{
		Code:    code,
		Message: message,
	}
}
