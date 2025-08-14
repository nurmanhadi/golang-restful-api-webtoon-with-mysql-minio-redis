package repository

import "welltoon/internal/entity"

type PageRepository interface {
	Save(page *entity.Page) error
	FindByID(pageID int64) (*entity.Page, error)
	Delete(pageID int64) error
}
