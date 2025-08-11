package repository

import "welltoon/internal/entity"

type UserRepository interface {
	CountByUsername(username string) (int64, error)
	Save(user *entity.User) error
	FindByUsername(username string) (*entity.User, error)
	CountByID(userID int64) (int64, error)
	FindByID(userID int64) (*entity.User, error)
	UpdateAvatar(userID int64, avatarFilename string, avatarUrl string) error
	Delete(userID int64) error
}
