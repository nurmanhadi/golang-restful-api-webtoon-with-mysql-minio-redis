package service

import (
	"strconv"
	"welltoon/internal/dto"
	"welltoon/internal/entity"
	"welltoon/internal/repository"
	"welltoon/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ChapterService interface {
	AddChapter(comicID string, request *dto.ChapterAddRequest) error
}
type chapterService struct {
	logger            *logrus.Logger
	validation        *validator.Validate
	chapterRepository repository.ChapterRepository
	comicRepository   repository.ComicRepository
}

func NewChapterService(logger *logrus.Logger,
	validation *validator.Validate,
	chapterRepository repository.ChapterRepository,
	comicRepository repository.ComicRepository) ChapterService {
	return &chapterService{
		logger:            logger,
		validation:        validation,
		chapterRepository: chapterRepository,
		comicRepository:   comicRepository,
	}
}
func (s *chapterService) AddChapter(comicID string, request *dto.ChapterAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.logger.WithField("data", request).Warn("validation failed")
		return err
	}
	newComicID, err := strconv.ParseInt(comicID, 10, 64)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"comic_id": comicID,
		}).Warn("comicID most be number")
		return response.Exception(400, "comicID most be number")
	}
	countComic, err := s.comicRepository.CountByID(newComicID)
	if err != nil {
		s.logger.WithError(err).Error("count by id to database failed")
		return err
	}
	if countComic < 1 {
		s.logger.WithField("data", fiber.Map{
			"comic_id": newComicID,
		}).Warn("comic not found")
		return response.Exception(404, "comic not found")
	}
	chapter := &entity.Chapter{
		ComicID: newComicID,
		Number:  request.Number,
		Publish: false,
	}
	err = s.chapterRepository.Save(chapter)
	if err != nil {
		s.logger.WithError(err).Error("save chapter to database failed")
		return err
	}
	s.logger.WithField("data", fiber.Map{
		"number": request.Number,
	}).Info("add chapter success")
	return nil
}
