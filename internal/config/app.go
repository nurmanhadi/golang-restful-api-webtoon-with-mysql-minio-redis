package config

import (
	"context"
	"welltoon/internal/delivery/rest/handler"
	"welltoon/internal/delivery/rest/routes"
	"welltoon/internal/infrastructure/cache"
	"welltoon/internal/infrastructure/db"
	"welltoon/internal/infrastructure/s3"
	"welltoon/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Configuration struct {
	Ctx        context.Context
	Cache      *redis.Client
	DB         *gorm.DB
	S3         *minio.Client
	Logger     *logrus.Logger
	Validation *validator.Validate
	App        *fiber.App
}

func App(conf *Configuration) {
	// S3
	s3 := s3.NewS3(conf.Ctx, conf.S3)

	// cache
	cache := cache.NewCache(conf.Ctx, conf.Cache)

	// DB
	userDB := db.NewUserDB(conf.DB)
	comicDB := db.NewComicDB(conf.DB)
	chapterDB := db.NewChapterDB(conf.DB)
	pageDB := db.NewPageDB(conf.DB)

	// service
	userServ := service.NewUserService(conf.Validation, conf.Logger, userDB, s3)
	comicServ := service.NewComicService(conf.Logger, conf.Validation, comicDB, s3)
	chapterServ := service.NewChapterService(conf.Logger, conf.Validation, chapterDB, comicDB)
	pageServ := service.NewPageService(conf.Logger, conf.Validation, pageDB, comicDB, chapterDB, cache, s3)

	// handler
	userHand := handler.NewUserHandler(userServ)
	comicHand := handler.NewComicHandler(comicServ)
	chapterHand := handler.NewChapterHandler(chapterServ)
	pageHand := handler.NewPageHandler(pageServ)

	// routes
	route := &routes.Route{
		App:            conf.App,
		UserHandler:    userHand,
		ComicHandler:   comicHand,
		ChapterHandler: chapterHand,
		PageHandler:    pageHand,
	}
	route.Setup()
}
