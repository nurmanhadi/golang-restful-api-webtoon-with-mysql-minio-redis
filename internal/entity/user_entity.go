package entity

import "time"

type User struct {
	ID             int64     `gorm:"primaryKey;autoIncrement;type:bigint"`
	Username       string    `gorm:"type:varchar(100);uniqueIndex"`
	Password       string    `gorm:"type:varchar(100)"`
	AvatarFilename *string   `gorm:"type:varchar(255)"`
	AvatarUrl      *string   `gorm:"type:varchar(255)"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
