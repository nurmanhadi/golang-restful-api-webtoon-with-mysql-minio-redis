package service

import (
	"time"
	"welltoon/internal/dto"
	"welltoon/internal/entity"
	"welltoon/internal/repository"
	"welltoon/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ViewService interface {
	AddView(request *dto.ViewAddRequest) error
	GetView() (*dto.ViewResponse, error)
}
type viewService struct {
	logger          *logrus.Logger
	validation      *validator.Validate
	viewRepository  repository.ViewRepository
	cacheRepository repository.CacheRepository
	comicRepository repository.ComicRepository
}

func NewViewService(logger *logrus.Logger,
	validation *validator.Validate,
	viewRepository repository.ViewRepository,
	cacheRepository repository.CacheRepository,
	comicRepository repository.ComicRepository) ViewService {
	return &viewService{
		logger:          logger,
		validation:      validation,
		viewRepository:  viewRepository,
		cacheRepository: cacheRepository,
		comicRepository: comicRepository,
	}
}
func (s *viewService) AddView(request *dto.ViewAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.logger.WithField("data", request).Warn("validation failed")
		return err
	}
	countComic, err := s.comicRepository.CountByID(request.ComicID)
	if err != nil {
		s.logger.WithError(err).Error("count by id to database failed")
		return err
	}
	if countComic < 1 {
		s.logger.WithField("data", fiber.Map{
			"comic_id": request.ComicID,
		}).Warn("comic not found")
		return response.Exception(404, "comic not found")
	}
	viewedAt := time.Now()
	view := &entity.View{
		ComicID:  &request.ComicID,
		Daily:    1,
		Weekly:   1,
		Monthly:  1,
		AllTime:  1,
		ViewedAt: &viewedAt,
	}
	if err := s.viewRepository.Save(view); err != nil {
		s.logger.WithError(err).Error("save view failed")
		return err
	}
	if err := s.cacheRepository.SetView(); err != nil {
		s.logger.WithError(err).Error("cache set view vailed")
		return err
	}
	s.logger.WithField("data", fiber.Map{
		"comic_id": request.ComicID,
	}).Info("add view success")
	return nil
}
func (s *viewService) GetView() (*dto.ViewResponse, error) {
	view, err := s.viewRepository.FindByComicIDIsNull()
	if err != nil {
		s.logger.WithError(err).Warn("view not found")
		return nil, response.Exception(404, "view not found")
	}
	result := &dto.ViewResponse{
		Daily:   view.Daily,
		Weekly:  view.Weekly,
		Monthly: view.Monthly,
		AllTime: view.AllTime,
	}
	return result, nil
}
