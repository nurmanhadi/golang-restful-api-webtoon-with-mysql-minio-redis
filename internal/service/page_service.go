package service

import (
	"mime/multipart"
	"strconv"
	"welltoon/internal/entity"
	"welltoon/internal/repository"
	"welltoon/pkg"
	"welltoon/pkg/image"
	"welltoon/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type PageService interface {
	AddBulkPage(comicID, chapterID string, files []*multipart.FileHeader) error
}
type pageService struct {
	logger            *logrus.Logger
	validation        *validator.Validate
	pageRepository    repository.PageRepository
	comicRepository   repository.ComicRepository
	chapterRepository repository.ChapterRepository
	cacheRepository   repository.CacheRepository
	s3Repository      repository.S3Repository
}

func NewPageService(logger *logrus.Logger,
	validation *validator.Validate,
	pageRepository repository.PageRepository,
	comicRepository repository.ComicRepository,
	chapterRepository repository.ChapterRepository,
	cacheRepository repository.CacheRepository,
	s3Repository repository.S3Repository) PageService {
	return &pageService{
		logger:            logger,
		validation:        validation,
		pageRepository:    pageRepository,
		comicRepository:   comicRepository,
		chapterRepository: chapterRepository,
		cacheRepository:   cacheRepository,
		s3Repository:      s3Repository,
	}
}
func (s *pageService) AddBulkPage(comicID, chapterID string, files []*multipart.FileHeader) error {
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
	countChapter, err := s.chapterRepository.CountByID(newChapterID)
	if err != nil {
		s.logger.WithError(err).Error("count by id to database failed")
		return err
	}
	if countChapter < 1 {
		s.logger.WithField("data", fiber.Map{
			"chapter_id": newChapterID,
		}).Warn("chapter not found")
		return response.Exception(404, "chapter not found")
	}
	for _, file := range files {
		if err := image.Validate(file.Filename); err != nil {
			s.logger.WithField("data", fiber.Map{
				"filename": file.Filename,
			}).Warn(err.Error())
			return response.Exception(400, err.Error())
		}
	}
	for _, file := range files {
		webp, err := image.CompressToWebp(file)
		if err != nil {
			s.logger.WithError(err).Error("compress image to webp failes")
			return err
		}
		err = s.s3Repository.PutObject(webp)
		if err != nil {
			s.logger.WithError(err).Error("s3 put object failed")
			return err
		}
		imageUrl, err := pkg.S3GenerateUrl(webp.Filename)
		if err != nil {
			err = s.s3Repository.RemoveObject(webp.Filename)
			if err != nil {
				s.logger.WithError(err).Error("s3 remove object failed")
				return err
			}
			s.logger.WithError(err).Error("s3 generate url failed")
			return err
		}
		page := &entity.Page{
			ChapterID:     newChapterID,
			ImageFilename: webp.Filename,
			ImageUrl:      imageUrl,
		}
		err = s.pageRepository.Save(page)
		if err != nil {
			err = s.s3Repository.RemoveObject(webp.Filename)
			if err != nil {
				s.logger.WithError(err).Error("s3 remove object failed")
				return err
			}
			s.logger.WithError(err).Error("save page to database failed")
			return err
		}
	}
	s.logger.WithField("data", fiber.Map{
		"total_file": len(files),
	}).Info("upload bulk page success")
	return nil
}
