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

type GenreService interface {
	AddGenre(request *dto.GenreRequest) error
	UpdateGenre(genreID string, request *dto.GenreRequest) error
	DeleteGenre(genreID string) error
	GetAllGenre() ([]dto.GenreResponse, error)
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
func (s *genreService) UpdateGenre(genreID string, request *dto.GenreRequest) error {
	newGenreID, err := strconv.ParseInt(genreID, 10, 64)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"genre_id": genreID,
		}).Warn("genreID most be number")
		return response.Exception(400, "genreID most be number")
	}
	if err := s.validation.Struct(request); err != nil {
		s.logger.WithField("data", request).Warn("validation failed")
		return err
	}
	countGenre, err := s.genreRepository.CountByID(newGenreID)
	if err != nil {
		s.logger.WithError(err).Error("count by id to database failed")
		return err
	}
	if countGenre < 1 {
		s.logger.WithField("data", fiber.Map{
			"genre_id": newGenreID,
		}).Warn("genre not found")
		return response.Exception(404, "genre not found")
	}
	if err := s.genreRepository.UpdateName(newGenreID, request.Name); err != nil {
		s.logger.WithError(err).Error("genre update name to database failed")
		return err
	}
	s.logger.WithField("data", request).Info("update genre success")
	return nil
}
func (s *genreService) DeleteGenre(genreID string) error {
	newGenreID, err := strconv.ParseInt(genreID, 10, 64)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"genre_id": genreID,
		}).Warn("genreID most be number")
		return response.Exception(400, "genreID most be number")
	}
	countGenre, err := s.genreRepository.CountByID(newGenreID)
	if err != nil {
		s.logger.WithError(err).Error("count by id to database failed")
		return err
	}
	if countGenre < 1 {
		s.logger.WithField("data", fiber.Map{
			"genre_id": newGenreID,
		}).Warn("genre not found")
		return response.Exception(404, "genre not found")
	}
	if err := s.genreRepository.Delete(newGenreID); err != nil {
		s.logger.WithError(err).Error("genre delete to database failed")
		return err
	}
	s.logger.WithField("data", fiber.Map{
		"genre_id": newGenreID,
	}).Info("delete genre success")
	return nil
}
func (s *genreService) GetAllGenre() ([]dto.GenreResponse, error) {
	genres, err := s.genreRepository.FindAll()
	if err != nil {
		s.logger.WithError(err).Error("genre find all to database failed")
		return nil, err
	}
	result := make([]dto.GenreResponse, 0, len(genres))
	if len(genres) != 0 {
		for _, genre := range genres {
			result = append(result, dto.GenreResponse{
				ID:        genre.ID,
				Name:      genre.Name,
				CreatedAt: genre.CreatedAt,
				UpdatedAt: genre.UpdatedAt,
			})
		}
	}
	s.logger.WithField("data", fiber.Map{
		"total_genre": len(result),
	}).Info("get all genre success")
	return result, nil
}
