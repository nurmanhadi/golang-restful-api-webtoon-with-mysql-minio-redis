package entity

import (
	"time"
	"welltoon/pkg/enum"
)

type User struct {
	ID             int64     `gorm:"primaryKey;autoIncrement;type:bigint"`
	Username       string    `gorm:"type:varchar(100);uniqueIndex"`
	Password       string    `gorm:"type:varchar(100)"`
	Role           enum.ROLE `gorm:"type:enum('admin','user')"`
	AvatarFilename *string   `gorm:"type:varchar(255)"`
	AvatarUrl      *string   `gorm:"type:varchar(255)"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
