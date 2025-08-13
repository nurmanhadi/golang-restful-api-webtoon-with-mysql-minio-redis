package repository

import "welltoon/internal/entity"

type ChapterRepository interface {
	Save(chapter *entity.Chapter) error
	FindByID(chapterID int64) (*entity.Chapter, error)
	Delete(chapterID int64) error
}
