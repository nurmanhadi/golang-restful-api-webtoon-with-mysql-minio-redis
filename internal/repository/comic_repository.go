package repository

import "welltoon/internal/entity"

type ComicRepository interface {
	Save(comic *entity.Comic) error
	CountBySlug(slug string) (int64, error)
}
