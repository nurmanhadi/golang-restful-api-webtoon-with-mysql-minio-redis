package repository

import "welltoon/internal/entity"

type ComicRepository interface {
	Save(comic *entity.Comic) error
	CountBySlug(slug string) (int64, error)
	FindByID(comicID int64) (*entity.Comic, error)
	FindBySlug(slug string) (*entity.Comic, error)
	Delete(comicID int64) error
}
