package db

import (
	"welltoon/internal/entity"
	"welltoon/internal/repository"

	"gorm.io/gorm"
)

type comicDB struct {
	db *gorm.DB
}

func NewComicDB(gb *gorm.DB) repository.ComicRepository {
	return &comicDB{db: gb}
}

func (r *comicDB) Save(comic *entity.Comic) error {
	return r.db.Save(comic).Error
}
func (r *comicDB) CountBySlug(slug string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Comic{}).Where("slug = ?", slug).Count(&count).Error
	if err != nil {
		return 0, nil
	}
	return count, nil
}
