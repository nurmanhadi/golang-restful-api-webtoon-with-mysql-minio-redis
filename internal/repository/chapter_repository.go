package repository

import "welltoon/internal/entity"

type ChapterRepository interface {
	Save(chapter *entity.Chapter) error
}
