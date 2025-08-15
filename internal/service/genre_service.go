package service

import (
	"welltoon/internal/dto"
	"welltoon/internal/entity"
	"welltoon/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type GenreService interface {
	AddGenre(request *dto.GenreRequest) error
}
type genreService struct {
	logger          *logrus.Logger
	validation      *validator.Validate
	genreRepository repository.GenreRepository
}

func NewGenreService(logger *logrus.Logger,
	validation *validator.Validate,
	genreRepository repository.GenreRepository) GenreService {
	return &genreService{
		logger:          logger,
		validation:      validation,
		genreRepository: genreRepository,
	}
}
func (s *genreService) AddGenre(request *dto.GenreRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.logger.WithField("data", request).Warn("validation failed")
		return err
	}
	genre := &entity.Genre{
		Name: request.Name,
	}
	if err := s.genreRepository.Save(genre); err != nil {
		s.logger.WithError(err).Error("genre save to database failed")
		return err
	}
	s.logger.WithField("data", request).Info("add genre success")
	return nil
}
