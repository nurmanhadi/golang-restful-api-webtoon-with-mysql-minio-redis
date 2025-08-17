package db

import (
	"welltoon/internal/entity"
	"welltoon/internal/repository"

	"gorm.io/gorm"
)

type viewDB struct {
	db *gorm.DB
}

func NewViewDB(db *gorm.DB) repository.ViewRepository {
	return &viewDB{db: db}
}
func (r *viewDB) Save(view *entity.View) error {
	return r.db.Save(view).Error
}
func (r *viewDB) FindByID(comicID int64) (*entity.View, error) {
	view := new(entity.View)
	err := r.db.Where("comic_id = ?", comicID).First(view).Error
	if err != nil {
		return nil, err
	}
	return view, nil
}
func (r *viewDB) CountByComicIDIsNull() (int64, error) {
	var count int64
	err := r.db.Model(&entity.View{}).Where("comic_id = NULL").Count(&count).Error
	if err != nil {
		return 0, nil
	}
	return count, nil
}
func (r *viewDB) FindByComicIDIsNull() (*entity.View, error) {
	view := new(entity.View)
	err := r.db.Where("comic_id = NULL").First(view).Error
	if err != nil {
		return nil, err
	}
	return view, nil
}
