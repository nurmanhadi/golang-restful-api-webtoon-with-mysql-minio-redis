package db

import (
	"welltoon/internal/entity"
	"welltoon/internal/repository"

	"gorm.io/gorm"
)

type comicDB struct {
	db *gorm.DB
}

func NewComicDB(gb *gorm.DB) repository.ComicRepository {
	return &comicDB{db: gb}
}

func (r *comicDB) Save(comic *entity.Comic) error {
	return r.db.Save(comic).Error
}
func (r *comicDB) CountBySlug(slug string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Comic{}).Where("slug = ?", slug).Count(&count).Error
	if err != nil {
		return 0, nil
	}
	return count, nil
}
func (r *comicDB) FindByID(comicID int64) (*entity.Comic, error) {
	comic := new(entity.Comic)
	err := r.db.Where("id = ?", comicID).First(comic).Error
	if err != nil {
		return nil, err
	}
	return comic, nil
}
func (r *comicDB) FindBySlug(slug string) (*entity.Comic, error) {
	comic := new(entity.Comic)
	err := r.db.Where("slug = ?", slug).First(comic).Error
	if err != nil {
		return nil, err
	}
	return comic, nil
}
func (r *comicDB) Delete(comicID int64) error {
	return r.db.Where("id = ?", comicID).Delete(&entity.Comic{}).Error
}
func (r *comicDB) UpdateCover(comicID int64, coverFilename string, coverUrl string) error {
	return r.db.Model(&entity.Comic{}).Where("id = ?", comicID).Updates(map[string]interface{}{
		"cover_filename": coverFilename,
		"cover_url":      coverUrl,
	}).Error
}
func (r *comicDB) FindAllByUpdatedOn(page int, size int) ([]entity.Comic, error) {
	var comics []entity.Comic
	err := r.db.
		Offset((page - 1) * size).
		Limit(size).
		Where("updated_on IS NOT NULL").
		Order("updated_on DESC").
		Preload("Chapters").
		Find(&comics).
		Error
	if err != nil {
		return nil, err
	}
	return comics, nil
}
func (r *comicDB) CountByUpdatedOn() (int64, error) {
	var count int64
	err := r.db.
		Model(&entity.Comic{}).
		Where("updated_on IS NOT NULL").
		Order("updated_on DESC").
		Count(&count).
		Error
	if err != nil {
		return 0, nil
	}
	return count, nil
}
