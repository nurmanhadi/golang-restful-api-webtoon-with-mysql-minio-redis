package repository

import "welltoon/internal/entity"

type GenreRepository interface {
	Save(genre *entity.Genre) error
}
