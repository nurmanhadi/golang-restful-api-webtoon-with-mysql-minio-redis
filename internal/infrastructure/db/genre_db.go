package db

import (
	"welltoon/internal/entity"
	"welltoon/internal/repository"

	"gorm.io/gorm"
)

type genreDB struct {
	db *gorm.DB
}

func NewGenreDB(db *gorm.DB) repository.GenreRepository {
	return &genreDB{db: db}
}
func (r *genreDB) Save(genre *entity.Genre) error {
	return r.db.Save(genre).Error
}
func (r *genreDB) CountByID(genreID int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Genre{}).Where("id = ?", genreID).Count(&count).Error
	if err != nil {
		return 0, nil
	}
	return count, nil
}
func (r *genreDB) UpdateName(genreID int64, name string) error {
	return r.db.Model(&entity.Genre{}).Where("id = ?", genreID).Update("name", name).Error
}
func (r *genreDB) Delete(genreID int64) error {
	return r.db.Where("id = ?", genreID).Delete(&entity.Genre{}).Error
}
