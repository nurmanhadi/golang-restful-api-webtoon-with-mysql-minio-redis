package service

import (
	"time"
	"welltoon/internal/dto"
	"welltoon/internal/entity"
	"welltoon/internal/repository"
	"welltoon/pkg"
	"welltoon/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ComicService interface {
	AddComic(request *dto.ComicAddRequest) error
}

type comicService struct {
	logger          *logrus.Logger
	validation      *validator.Validate
	comicRepository repository.ComicRepository
	s3Repository    repository.S3Repository
}

func NewComicService(logger *logrus.Logger,
	validation *validator.Validate,
	comicRepository repository.ComicRepository,
	s3Repository repository.S3Repository) ComicService {
	return &comicService{
		logger:          logger,
		validation:      validation,
		comicRepository: comicRepository,
		s3Repository:    s3Repository,
	}
}

func (s *comicService) AddComic(request *dto.ComicAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.logger.WithField("data", request).Warn("validation failed")
		return err
	}
	slug := pkg.GenerateSlug(request.Title)
	countSlug, err := s.comicRepository.CountBySlug(slug)
	if err != nil {
		s.logger.WithError(err).Error("count slug to database failed")
		return err
	}
	if countSlug > 0 {
		s.logger.WithField("data", fiber.Map{
			"title": request.Title,
			"slug":  slug,
		}).Warn("comic already exists")
		return response.Exception(400, "comic already exists")
	}
	comic := &entity.Comic{
		Title:  request.Title,
		Slug:   slug,
		Author: request.Author,
		Artist: request.Artist,
		Type:   request.Type,
		Status: request.Status,
		PostOn: time.Now(),
	}
	if request.Synopsis != nil {
		comic.Synopsis = request.Synopsis
	}
	err = s.comicRepository.Save(comic)
	if err != nil {
		s.logger.WithError(err).Error("save comic to database failed")
		return err
	}

	s.logger.WithField("data", comic).Info("add comic success")
	return nil
}
