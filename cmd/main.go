package main

import (
	"context"
	"log"
	"welltoon/internal/config"
)

func main() {
	ctx := context.Background()
	config.NewEnvirontment()
	config.NewRedis(ctx)
	config.NewGorm()
	config.NewMinio()
	config.NewLogger()
	config.NewValidator()
	app := config.NewFiber()

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("failed start server: %s", err)
	}
}
