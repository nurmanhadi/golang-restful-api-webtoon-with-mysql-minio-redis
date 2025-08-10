package entity

import (
	"time"
	"welltoon/pkg/enum"
)

type Comic struct {
	ID            int64        `gorm:"primaryKey;autoIncrement;type:bigint"`
	Title         string       `gorm:"type:varchar(200)"`
	Slug          string       `gorm:"type:varchar(200);uniqueIndex"`
	Synopsis      *string      `gorm:"type:text"`
	Author        string       `gorm:"type:varchar(50)"`
	Artist        string       `gorm:"type:varchar(50)"`
	Type          enum.TYPE    `gorm:"type:enum('manga','manhua','manhwa')"`
	Status        enum.STATUS  `gorm:"type:enum('completed','hiatus','ongoing')"`
	CoverFilename *string      `gorm:"type:varchar(255)"`
	CoverUrl      *string      `gorm:"type:varchar(255)"`
	PostOn        time.Time    `gorm:"type:timestamp;index"`
	UpdatedOn     *time.Time   `gorm:"type:timestamp;index"`
	CreatedAt     time.Time    `gorm:"autoCreateTime"`
	UpdatedAt     time.Time    `gorm:"autoUpdateTime"`
	Chapters      []Chapter    `gorm:"foreignKey:ComicID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ComicGenres   []ComicGenre `gorm:"foreignKey:ComicID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Views         []View       `gorm:"foreignKey:ComicID;constraint;onDelete:CASCADE;onUpdate:CASCADE"`
}
