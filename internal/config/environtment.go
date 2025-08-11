package config

import (
	"github.com/joho/godotenv"
)

func NewEnvirontment() error {
	return godotenv.Load()
}
