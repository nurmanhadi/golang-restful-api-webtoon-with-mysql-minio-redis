package entity

import "time"

type Page struct {
	ID            int64     `gorm:"primaryKey;autoIncrement;type:bigint"`
	ChapterID     int64     `gorm:"type:bigint"`
	ImageFilename string    `gorm:"type:varchar(255)"`
	ImageUrl      string    `gorm:"type:varchar(255)"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	Chapter       *Chapter  `gorm:"foreignKey:ChapterID;constraint;onDelete:CASCADE;onUpdate:CASCADE"`
}
