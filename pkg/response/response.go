package response

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

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

func Success[T any](c *fiber.Ctx, code int, data T) error {
	return c.Status(code).JSON(fiber.Map{
		"data": data,
		"path": c.OriginalURL(),
	})
}
