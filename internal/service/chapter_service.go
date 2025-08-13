package service

import (
	"sort"
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
	UpdateChapter(comicID, chapterID string, request *dto.ChapterUpdateRequest) error
	DeleteChapter(comicID, chapterID string) error
	GetChapterBySlugAndNumber(slug string, chapterNumber string) (*dto.ChapterResponse, error)
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
func (s *chapterService) UpdateChapter(comicID, chapterID string, request *dto.ChapterUpdateRequest) error {
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
	newChapterID, err := strconv.ParseInt(chapterID, 10, 64)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"chapter_id": chapterID,
		}).Warn("chapterID most be number")
		return response.Exception(400, "chapterID most be number")
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
	chapter, err := s.chapterRepository.FindByID(newChapterID)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"chapter_id": newChapterID,
		}).Warn("chapter not found")
		return response.Exception(404, "chapter not found")
	}
	if request.Number != nil {
		chapter.Number = *request.Number
	}
	if request.Publish != nil {
		chapter.Publish = *request.Publish
	}
	err = s.chapterRepository.Save(chapter)
	if err != nil {
		s.logger.WithError(err).Error("save chapter to database failed")
		return err
	}
	s.logger.WithField("data", request).Info("update chapter success")
	return nil
}
func (s *chapterService) DeleteChapter(comicID, chapterID string) error {
	newComicID, err := strconv.ParseInt(comicID, 10, 64)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"comic_id": comicID,
		}).Warn("comicID most be number")
		return response.Exception(400, "comicID most be number")
	}
	newChapterID, err := strconv.ParseInt(chapterID, 10, 64)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"chapter_id": chapterID,
		}).Warn("chapterID most be number")
		return response.Exception(400, "chapterID most be number")
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
	chapter, err := s.chapterRepository.FindByID(newChapterID)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"chapter_id": newChapterID,
		}).Warn("chapter not found")
		return response.Exception(404, "chapter not found")
	}
	err = s.chapterRepository.Delete(chapter.ID)
	if err != nil {
		s.logger.WithError(err).Error("delete chapter to database failed")
		return err
	}
	s.logger.WithField("data", fiber.Map{
		"chapter_id": newChapterID,
	}).Info("delete chapter success")
	return nil
}
func (s *chapterService) GetChapterBySlugAndNumber(slug string, chapterNumber string) (*dto.ChapterResponse, error) {
	newChapterNumber, err := strconv.Atoi(chapterNumber)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"number": chapterNumber,
		}).Warn("number most be number")
		return nil, response.Exception(400, "number most be number")
	}
	comic, err := s.comicRepository.FindBySlug(slug)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"slug": slug,
		}).Warn("comic not found")
		return nil, response.Exception(404, "comic not found")
	}
	chapter, err := s.chapterRepository.FindByComicIDAndNumber(comic.ID, newChapterNumber)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"number": newChapterNumber,
		}).Warn("chapter not found")
		return nil, response.Exception(404, "chapter not found")
	}
	chapters := make([]dto.ChapterResponse, 0, len(comic.Chapters))
	if len(comic.Chapters) != 0 {
		for _, ch := range comic.Chapters {
			if ch.Publish {
				chapters = append(chapters, dto.ChapterResponse{
					ID:        ch.ID,
					ComicID:   ch.ComicID,
					Number:    ch.Number,
					Publish:   ch.Publish,
					CreatedAt: ch.CreatedAt,
					UpdatedAt: ch.UpdatedAt,
				})
			}
		}
		if len(chapters) > 0 {
			sort.Slice(chapters, func(i, j int) bool {
				return chapters[i].Number < chapters[j].Number // DESC
			})
		}
	}
	comicResponse := &dto.ComicResponse{
		ID:            comic.ID,
		Title:         comic.Title,
		Slug:          comic.Slug,
		Synopsis:      comic.Synopsis,
		Author:        comic.Author,
		Artist:        comic.Artist,
		Type:          comic.Type,
		Status:        comic.Status,
		CoverFilename: comic.CoverFilename,
		CoverUrl:      comic.CoverUrl,
		PostOn:        comic.PostOn,
		UpdatedOn:     comic.UpdatedOn,
		CreatedAt:     comic.CreatedAt,
		UpdatedAt:     comic.UpdatedAt,
		Chapters:      chapters,
	}
	pages := make([]dto.PagesResponse, 0, len(chapter.Pages))
	if len(chapter.Pages) != 0 {
		for _, page := range chapter.Pages {
			pages = append(pages, dto.PagesResponse{
				ID:            page.ID,
				ChapterID:     page.ChapterID,
				ImageFilename: page.ImageFilename,
				ImageUrl:      page.ImageUrl,
				CreatedAt:     page.CreatedAt,
				UpdatedAt:     page.UpdatedAt,
			})
		}
	}
	result := &dto.ChapterResponse{
		ID:        chapter.ID,
		ComicID:   chapter.ComicID,
		Number:    chapter.Number,
		Publish:   chapter.Publish,
		CreatedAt: chapter.CreatedAt,
		UpdatedAt: chapter.UpdatedAt,
		Comic:     comicResponse,
		Pages:     pages,
	}
	s.logger.WithField("data", fiber.Map{
		"slug":   slug,
		"number": newChapterNumber,
	}).Info("get chapter by slug and number success")
	return result, nil
}
