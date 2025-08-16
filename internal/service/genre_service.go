package service

import (
	"math"
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
	GetComicByGenreName(name, page, size string) (*dto.Pagination[[]dto.ComicResponse], error)
}
type genreService struct {
	logger               *logrus.Logger
	validation           *validator.Validate
	genreRepository      repository.GenreRepository
	comicGenreRepository repository.ComicGenreRepository
}

func NewGenreService(logger *logrus.Logger,
	validation *validator.Validate,
	genreRepository repository.GenreRepository,
	comicGenreRepository repository.ComicGenreRepository) GenreService {
	return &genreService{
		logger:               logger,
		validation:           validation,
		genreRepository:      genreRepository,
		comicGenreRepository: comicGenreRepository,
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
func (s *genreService) GetComicByGenreName(name, page, size string) (*dto.Pagination[[]dto.ComicResponse], error) {
	newPage, err := strconv.Atoi(page)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"page": page,
		}).Warn("page most be number")
		return nil, response.Exception(400, "page most be number")
	}
	newSize, err := strconv.Atoi(size)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"size": size,
		}).Warn("size most be number")
		return nil, response.Exception(400, "size most be number")
	}
	genre, err := s.genreRepository.FindByName(name)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"name": name,
		}).Warn("genre not found")
		return nil, response.Exception(404, "genre not found")
	}
	comicGenres, err := s.comicGenreRepository.FindAllByGenreID(genre.ID, newPage, newSize)
	if err != nil {
		s.logger.WithError(err).Error("find all by genre id to database failed")
		return nil, err
	}
	comics := make([]dto.ComicResponse, 0, len(comicGenres))
	var totalElement int
	var totalPage int
	if len(comicGenres) != 0 {
		for _, comicGenre := range comicGenres {
			comics = append(comics, dto.ComicResponse{
				ID:            comicGenre.Comic.ID,
				Title:         comicGenre.Comic.Title,
				Slug:          comicGenre.Comic.Slug,
				Synopsis:      comicGenre.Comic.Synopsis,
				Author:        comicGenre.Comic.Author,
				Artist:        comicGenre.Comic.Artist,
				Type:          comicGenre.Comic.Type,
				Status:        comicGenre.Comic.Status,
				CoverFilename: comicGenre.Comic.CoverFilename,
				CoverUrl:      comicGenre.Comic.CoverUrl,
				PostOn:        comicGenre.Comic.PostOn,
				UpdatedOn:     comicGenre.Comic.UpdatedOn,
				CreatedAt:     comicGenre.Comic.CreatedAt,
				UpdatedAt:     comicGenre.Comic.UpdatedAt,
			})
		}
		count, err := s.comicGenreRepository.CountByGenreID(genre.ID)
		if err != nil {
			s.logger.WithError(err).Error("count by genre id to database failed")
			return nil, err
		}
		totalElement = int(count)
		totalPage = int(math.Ceil(float64(count) / float64(newSize)))
	}
	result := &dto.Pagination[[]dto.ComicResponse]{
		Contents:     comics,
		Page:         newPage,
		Size:         newSize,
		TotalPage:    totalPage,
		TotalElement: totalElement,
	}
	return result, nil
}
