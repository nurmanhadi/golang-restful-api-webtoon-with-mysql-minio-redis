package db

import (
	"welltoon/internal/entity"
	"welltoon/internal/repository"

	"gorm.io/gorm"
)

type comicGenreDB struct {
	db *gorm.DB
}

func NewComicGenreDB(db *gorm.DB) repository.ComicGenreRepository {
	return &comicGenreDB{db: db}
}
func (r *comicGenreDB) FindAllByGenreID(genreID int64, page, size int) ([]entity.ComicGenre, error) {
	var comicGenres []entity.ComicGenre
	err := r.db.
		Offset((page-1)*size).
		Limit(size).
		Where("genre_id = ?", genreID).
		Preload("Comic").
		Find(&comicGenres).
		Error
	if err != nil {
		return nil, err
	}
	return comicGenres, nil
}
func (r *comicGenreDB) CountByGenreID(genreID int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.ComicGenre{}).Where("genre_id = ?", genreID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
