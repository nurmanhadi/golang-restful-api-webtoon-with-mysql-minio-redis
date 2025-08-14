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
func (r *chapterDB) FindByID(chapterID int64) (*entity.Chapter, error) {
	chapter := new(entity.Chapter)
	err := r.db.Where("id = ?", chapterID).First(chapter).Error
	if err != nil {
		return nil, err
	}
	return chapter, nil
}
func (r *chapterDB) Delete(chapterID int64) error {
	return r.db.Where("id = ?", chapterID).Delete(&entity.Chapter{}).Error
}
func (r *chapterDB) FindByComicIDAndNumber(comicID int64, number int) (*entity.Chapter, error) {
	chapter := new(entity.Chapter)
	err := r.db.
		Where("comic_id = ? AND number = ?", comicID, number).
		Preload("Pages").
		First(chapter).Error
	if err != nil {
		return nil, err
	}
	return chapter, nil
}
func (r *chapterDB) CountByID(chapterID int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Chapter{}).Where("id = ?", chapterID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
