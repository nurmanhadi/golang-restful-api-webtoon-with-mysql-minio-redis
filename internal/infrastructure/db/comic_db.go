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
		Count(&count).
		Error
	if err != nil {
		return 0, nil
	}
	return count, nil
}
func (r *comicDB) Count() (int64, error) {
	var count int64
	err := r.db.
		Model(&entity.Comic{}).
		Count(&count).
		Error
	if err != nil {
		return 0, nil
	}
	return count, nil
}
func (r *comicDB) FindAllByKeyword(keyword string, page int, size int) ([]entity.Comic, error) {
	var comics []entity.Comic
	key := "%" + keyword + "%"
	err := r.db.
		Offset((page-1)*size).
		Limit(size).
		Where("updated_on IS NOT NULL AND (title LIKE ? OR synopsis LIKE ?)", key, key).
		Find(&comics).
		Error
	if err != nil {
		return nil, err
	}
	return comics, nil
}
func (r *comicDB) CountByKeyword(keyword string) (int64, error) {
	var count int64
	key := "%" + keyword + "%"
	err := r.db.
		Model(&entity.Comic{}).
		Where("updated_on IS NOT NULL AND (title LIKE ? OR synopsis LIKE ?)", key, key).
		Count(&count).
		Error
	if err != nil {
		return 0, nil
	}
	return count, nil
}
func (r *comicDB) FindAllByTypeAndStatus(typeComic, status string, page int, size int) ([]entity.Comic, error) {
	var comics []entity.Comic
	err := r.db.
		Offset((page-1)*size).
		Limit(size).
		Where("updated_on IS NOT NULL AND type = ? AND status = ?", typeComic, status).
		Order("updated_on DESC").
		Find(&comics).
		Error
	if err != nil {
		return nil, err
	}
	return comics, nil
}
func (r *comicDB) CountByTypeAndStatus(typeComic, status string) (int64, error) {
	var count int64
	err := r.db.
		Model(&entity.Comic{}).
		Where("updated_on IS NOT NULL AND type = ? AND status = ?", typeComic, status).
		Count(&count).
		Error
	if err != nil {
		return 0, nil
	}
	return count, nil
}
func (r *comicDB) FindByTitle(title string) ([]entity.Comic, error) {
	var comics []entity.Comic
	key := "%" + title + "%"
	err := r.db.
		Limit(6).
		Where("updated_on IS NOT NULL AND title != ? AND title LIKE ?", title, key).
		Find(&comics).
		Error
	if err != nil {
		return nil, err
	}
	return comics, nil
}
func (r *comicDB) FindByCreatedAt() ([]entity.Comic, error) {
	var comics []entity.Comic
	err := r.db.
		Limit(10).
		Order("created_at DESC").
		Find(&comics).
		Error
	if err != nil {
		return nil, err
	}
	return comics, nil
}
func (r *comicDB) CountByID(comicID int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Comic{}).Where("id = ?", comicID).Count(&count).Error
	if err != nil {
		return 0, nil
	}
	return count, nil
}
