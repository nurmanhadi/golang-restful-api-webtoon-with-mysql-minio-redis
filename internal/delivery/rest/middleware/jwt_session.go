package middleware

import (
	"errors"
	"strings"
	"welltoon/internal/security"
	"welltoon/pkg/response"

	"github.com/gofiber/fiber/v2"
)

func JwtSession(c *fiber.Ctx) error {
	tokenString, err := getTokenFromHeader(c)
	if err != nil {
		return response.Exception(401, err.Error())
	}

	claims, err := security.JwtVerify(tokenString)
	if err != nil {
		return response.Exception(401, err.Error())
	}
	c.Locals("role", claims.Role)
	return c.Next()
}

func getTokenFromHeader(c *fiber.Ctx) (string, error) {
	header := c.Get("Authorization", "")
	if header == "" {
		return "", errors.New("token null")
	}
	token := strings.Split(header, " ")
	if token[0] != "Bearer" {
		return "", errors.New("value authorization most be Bearer example 'Authorization: Bearer token'")
	}
	return token[1], nil
}
