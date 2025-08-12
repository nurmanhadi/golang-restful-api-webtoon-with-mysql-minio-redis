package service

import (
	"math"
	"mime/multipart"
	"sort"
	"strconv"
	"time"
	"welltoon/internal/dto"
	"welltoon/internal/entity"
	"welltoon/internal/repository"
	"welltoon/pkg"
	"welltoon/pkg/image"
	"welltoon/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ComicService interface {
	AddComic(request *dto.ComicAddRequest) error
	UpdateComic(comicID string, request *dto.ComicUpdateRequest) error
	DeleteComic(comicID string) error
	GetComicBySlug(slug string) (*dto.ComicResponse, error)
	UploadCover(comicID string, cover *multipart.FileHeader) error
	GetComicRecent(page string, size string) (*dto.Pagination[[]dto.ComicResponse], error)
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
func (s *comicService) UpdateComic(comicID string, request *dto.ComicUpdateRequest) error {
	newComicID, err := strconv.ParseInt(comicID, 10, 64)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"comic_id": comicID,
		}).Warn("comicID most be number")
		return response.Exception(400, "comicID most be number")
	}
	if err := s.validation.Struct(request); err != nil {
		s.logger.WithField("data", request).Warn("validation failed")
		return err
	}
	comic, err := s.comicRepository.FindByID(newComicID)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"comic_id": newComicID,
		}).Warn("comic not found")
		return response.Exception(404, "comic not found")
	}
	if request.Title != nil {
		slug := pkg.GenerateSlug(*request.Title)
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
		comic.Title = *request.Title
		comic.Slug = slug
	}
	if request.Synopsis != nil {
		comic.Synopsis = request.Synopsis
	}
	if request.Author != nil {
		comic.Author = *request.Author
	}
	if request.Artist != nil {
		comic.Artist = *request.Artist
	}
	if request.Type != nil {
		comic.Type = *request.Type
	}
	if request.Status != nil {
		comic.Status = *request.Status
	}
	err = s.comicRepository.Save(comic)
	if err != nil {
		s.logger.WithError(err).Error("save comic to database failed")
		return err
	}
	s.logger.WithField("data", comic).Info("add comic success")
	return nil
}
func (s *comicService) DeleteComic(comicID string) error {
	newComicID, err := strconv.ParseInt(comicID, 10, 64)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"comic_id": comicID,
		}).Warn("comicID most be number")
		return response.Exception(400, "comicID most be number")
	}
	comic, err := s.comicRepository.FindByID(newComicID)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"comic_id": newComicID,
		}).Warn("comic not found")
		return response.Exception(404, "comic not found")
	}
	if comic.CoverFilename != nil && comic.CoverUrl != nil {
		err := s.s3Repository.RemoveObject(*comic.CoverFilename)
		if err != nil {
			s.logger.WithError(err).Error("s3 remove object failed")
			return err
		}
	}
	err = s.comicRepository.Delete(newComicID)
	if err != nil {
		s.logger.WithError(err).Error("delete comic to database failed")
		return err
	}
	s.logger.WithField("data", fiber.Map{
		"comic_id": newComicID,
	}).Info("delete comic success")
	return nil
}
func (s *comicService) GetComicBySlug(slug string) (*dto.ComicResponse, error) {
	comic, err := s.comicRepository.FindBySlug(slug)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"slug": slug,
		}).Warn("comic not found")
		return nil, response.Exception(404, "comic not found")
	}
	result := &dto.ComicResponse{
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
	}
	return result, nil
}
func (s *comicService) UploadCover(comicID string, cover *multipart.FileHeader) error {
	newComicID, err := strconv.ParseInt(comicID, 10, 64)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"comic_id": comicID,
		}).Warn("comicID most be number")
		return response.Exception(400, "comicID most be number")
	}
	err = image.Validate(cover.Filename)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"cover_filename": cover.Filename,
		}).Warn(err.Error())
		return response.Exception(400, err.Error())
	}
	comic, err := s.comicRepository.FindByID(newComicID)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"comic_id": newComicID,
		}).Warn("comic not found")
		return response.Exception(404, "comic not found")
	}
	if comic.CoverFilename != nil && comic.CoverUrl != nil {
		file, err := image.CompressToWebp(cover)
		if err != nil {
			s.logger.WithError(err).Error("compress image to webp failed")
			return err
		}
		file.Filename = *comic.CoverFilename
		err = s.s3Repository.PutObject(file)
		if err != nil {
			s.logger.WithError(err).Error("s3 put object failed")
			return err
		}
		s.logger.WithField("data", fiber.Map{
			"cover_url": comic.CoverUrl,
		}).Info("upload cover success")
	} else {
		file, err := image.CompressToWebp(cover)
		if err != nil {
			s.logger.WithError(err).Error("compress image to webp failed")
			return err
		}
		coverUrl, err := pkg.S3GenerateUrl(file.Filename)
		if err != nil {
			s.logger.WithError(err).Error("s3 generate url failed")
			return err
		}
		err = s.s3Repository.PutObject(file)
		if err != nil {
			s.logger.WithError(err).Error("s3 put object failed")
			return err
		}
		err = s.comicRepository.UpdateCover(newComicID, file.Filename, coverUrl)
		if err != nil {
			err = s.s3Repository.RemoveObject(file.Filename)
			if err != nil {
				s.logger.WithError(err).Error("s3 remove object failed")
				return err
			}
			s.logger.WithError(err).Error("update cover failed")
			return err
		}
		s.logger.WithField("data", fiber.Map{
			"cover_url": coverUrl,
		}).Info("upload cover success")
	}
	return nil
}
func (s *comicService) GetComicRecent(page string, size string) (*dto.Pagination[[]dto.ComicResponse], error) {
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
	comics, err := s.comicRepository.FindAllByUpdatedOn(newPage, newSize)
	if err != nil {
		s.logger.WithError(err).Error("find all by update_on to database failed")
		return nil, err
	}
	contents := make([]dto.ComicResponse, 0, len(comics))
	var totalPage int
	var totalElement int
	if len(comics) != 0 {
		for _, comic := range comics {
			chapters := make([]dto.ChapterResponse, 0, len(comic.Chapters))
			if len(comic.Chapters) != 0 {
				for _, chapter := range comic.Chapters {
					if chapter.Publish {
						chapters = append(chapters, dto.ChapterResponse{
							ID:        chapter.ID,
							ComicID:   chapter.ComicID,
							Number:    chapter.Number,
							Publish:   chapter.Publish,
							CreatedAt: chapter.CreatedAt,
							UpdatedAt: chapter.UpdatedAt,
						})
					}
				}
				if len(chapters) > 0 {
					sort.Slice(chapters, func(i, j int) bool {
						return chapters[i].Number > chapters[j].Number // DESC
					})
					chapters = chapters[:1]
				}
			}
			contents = append(contents, dto.ComicResponse{
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
			})
		}
		countTotalElement, err := s.comicRepository.CountByUpdatedOn()
		if err != nil {
			s.logger.WithError(err).Error("count by total updated_on to database failed")
			return nil, err
		}
		totalElement = int(countTotalElement)
		totalPage = int(math.Ceil(float64(countTotalElement) / float64(newSize)))
	}
	result := &dto.Pagination[[]dto.ComicResponse]{
		Contents:     contents,
		Page:         newPage,
		Size:         newSize,
		TotalPage:    totalPage,
		TotalElement: totalElement,
	}
	s.logger.WithField("data", fiber.Map{
		"total_element": totalElement,
	}).Info("get comic recent success")
	return result, nil
}
