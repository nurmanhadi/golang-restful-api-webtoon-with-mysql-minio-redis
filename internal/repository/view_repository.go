package repository

import "welltoon/internal/entity"

type ViewRepository interface {
	Save(view *entity.View) error
	FindByID(comicID int64) (*entity.View, error)
	FindByComicIDIsNull() (*entity.View, error)
}
