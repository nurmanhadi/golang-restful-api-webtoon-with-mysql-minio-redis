package db

import (
	"fmt"
	"time"
	"welltoon/internal/entity"
	"welltoon/internal/repository"
	"welltoon/pkg/enum"

	"gorm.io/gorm"
)

type userDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) repository.UserRepository {
	return &userDB{db: db}
}
func (r *userDB) CountByUsername(username string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *userDB) CountByID(userID int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.User{}).Where("id = ?", userID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *userDB) Save(user *entity.User) error {
	return r.db.Save(user).Error
}
func (r *userDB) FindByUsername(username string) (*entity.User, error) {
	user := new(entity.User)
	err := r.db.Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *userDB) FindByID(userID int64) (*entity.User, error) {
	user := new(entity.User)
	err := r.db.Where("id = ?", userID).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *userDB) UpdateAvatar(userID int64, avatarFilename string, avatarUrl string) error {
	return r.db.Model(&entity.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"avatar_filename": avatarFilename,
		"avatar_url":      avatarUrl,
	}).Error
}
func (r *userDB) Delete(userID int64) error {
	return r.db.Where("id = ?", userID).Delete(&entity.User{}).Error
}
func (r *userDB) CountTotal(by enum.BY) (int64, error) {
	var count int64
	query := r.db.Model(&entity.User{})
	now := time.Now()

	switch by {
	case enum.BY_DAILY:
		query = query.Where("created_at >= ?", now.AddDate(0, 0, -1))
	case enum.BY_WEEKLY:
		query = query.Where("created_at >= ?", now.AddDate(0, 0, -7))
	case enum.BY_MONTHLY:
		query = query.Where("created_at >= ?", now.AddDate(0, -1, 0))
	case enum.BY_ALL_TIME:
		//without filter
	default:
		return 0, fmt.Errorf("invalid filter by: %s", by)
	}

	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
