package db

import (
	"welltoon/internal/entity"
	"welltoon/internal/repository"

	"gorm.io/gorm"
)

type chapterDB struct {
	db *gorm.DB
}

func NewChapterDB(db *gorm.DB) repository.ChapterRepository {
	return &chapterDB{db: db}
}
func (r *chapterDB) Save(chapter *entity.Chapter) error {
	return r.db.Save(chapter).Error
}
