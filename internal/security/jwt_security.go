package security

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtCustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func JwtCreateToken(userID int64, role string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	claims := &JwtCustomClaims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.NewString(),
			Subject:   fmt.Sprint(userID),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add((24 * time.Hour) * 7)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "user",
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
func JwtVerify(tokenString string) (*JwtCustomClaims, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(t *jwt.Token) (any, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}
