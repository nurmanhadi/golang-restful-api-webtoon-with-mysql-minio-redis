package repository

import "welltoon/internal/entity"

type GenreRepository interface {
	Save(genre *entity.Genre) error
	CountByID(genreID int64) (int64, error)
	UpdateName(genreID int64, name string) error
	Delete(genreID int64) error
	FindAll() ([]entity.Genre, error)
	FindByName(name string) (*entity.Genre, error)
}
