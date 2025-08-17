package main

import (
	"context"
	"log"
	"welltoon/internal/config"
)

func main() {
	ctx := context.Background()
	err := config.NewEnvirontment()
	if err != nil {
		log.Fatalf("environtment error: %s", err)
	}
	cache := config.NewRedis(ctx)
	db := config.NewGorm()
	s3 := config.NewMinio()
	logger := config.NewLogger()
	validation := config.NewValidator()
	app := config.NewFiber()
	schedule := config.NewGocron()

	config.App(&config.Configuration{
		Ctx:        ctx,
		Cache:      cache,
		DB:         db,
		S3:         s3,
		Logger:     logger,
		Validation: validation,
		App:        app,
		Schedule:   schedule,
	})

	schedule.Start()
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("failed start server: %s", err)
	}
}
