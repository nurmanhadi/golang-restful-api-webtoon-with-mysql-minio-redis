package repository

import "welltoon/internal/entity"

type ComicGenreRepository interface {
	FindAllByGenreID(genreID int64, page, size int) ([]entity.ComicGenre, error)
	CountByGenreID(genreID int64) (int64, error)
	Save(comicGenre *entity.ComicGenre) error
}
