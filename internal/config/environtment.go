package config

import (
	"github.com/joho/godotenv"
)

func NewEnvirontment() error {
	godotenv.Load()
	return nil
}
