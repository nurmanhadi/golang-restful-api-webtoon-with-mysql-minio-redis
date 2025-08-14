package repository

import "welltoon/internal/entity"

type PageRepository interface {
	Save(page *entity.Page) error
}
