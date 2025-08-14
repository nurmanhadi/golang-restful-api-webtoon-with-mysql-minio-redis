package db

import (
	"welltoon/internal/entity"
	"welltoon/internal/repository"

	"gorm.io/gorm"
)

type pageDB struct {
	db *gorm.DB
}

func NewPageDB(db *gorm.DB) repository.PageRepository {
	return &pageDB{db: db}
}
func (r *pageDB) Save(page *entity.Page) error {
	return r.db.Save(page).Error
}
