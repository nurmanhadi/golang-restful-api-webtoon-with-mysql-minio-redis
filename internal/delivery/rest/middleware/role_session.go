package middleware

import (
	"slices"
	"welltoon/pkg/response"

	"github.com/gofiber/fiber/v2"
)

func RoleSession(roles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
		if !ok {
			return response.Exception(403, "you do not have permission to access this resource")
		}
		if slices.Contains(roles, role) {
			return c.Next()
		} else {
			return response.Exception(403, "you do not have permission to access this resource")
		}
	}
}
