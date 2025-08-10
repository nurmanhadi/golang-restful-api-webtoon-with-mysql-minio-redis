package entity

import "time"

type Genre struct {
	ID          int64        `gorm:"primaryKey;autoIncrement;type:bigint"`
	Name        string       `gorm:"type:varchar(50);index"`
	CreatedAt   time.Time    `gorm:"autoCreateTime"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime"`
	ComicGenres []ComicGenre `gorm:"foreignKey:GenreID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
