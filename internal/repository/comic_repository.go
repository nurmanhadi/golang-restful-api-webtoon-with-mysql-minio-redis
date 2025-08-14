package repository

import "welltoon/internal/entity"

type ComicRepository interface {
	Save(comic *entity.Comic) error
	CountBySlug(slug string) (int64, error)
	FindByID(comicID int64) (*entity.Comic, error)
	FindBySlug(slug string) (*entity.Comic, error)
	Delete(comicID int64) error
	UpdateCover(comicID int64, coverFilename string, coverUrl string) error
	FindAllByUpdatedOn(page int, size int) ([]entity.Comic, error)
	CountByUpdatedOn() (int64, error)
	Count() (int64, error)
	FindAllByKeyword(keyword string, page int, size int) ([]entity.Comic, error)
	CountByKeyword(keyword string) (int64, error)
	FindAllByTypeAndStatus(typeComic, status string, page int, size int) ([]entity.Comic, error)
	CountByTypeAndStatus(typeComic, status string) (int64, error)
	FindByTitle(title string) ([]entity.Comic, error)
	FindByCreatedAt() ([]entity.Comic, error)
	CountByID(comicID int64) (int64, error)
	UpdateUpdateOn(comicID int64) error
}
