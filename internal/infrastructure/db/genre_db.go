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
