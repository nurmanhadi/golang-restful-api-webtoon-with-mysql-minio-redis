package entity

import "time"

type ComicGenre struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;type:bigint"`
	ComicID   int64     `gorm:"type:bigint"`
	GenreID   int64     `gorm:"type:bigint"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Comic     *Comic    `gorm:"foreignKey:ComicID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Genre     *Genre    `gorm:"foreignKey:GenreID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
