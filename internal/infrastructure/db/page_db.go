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
func (r *pageDB) FindByID(pageID int64) (*entity.Page, error) {
	page := new(entity.Page)
	err := r.db.Where("id = ?", pageID).First(page).Error
	if err != nil {
		return nil, err
	}
	return page, nil
}
func (r *pageDB) Delete(pageID int64) error {
	return r.db.Where("id = ?", pageID).Delete(&entity.Page{}).Error
}
