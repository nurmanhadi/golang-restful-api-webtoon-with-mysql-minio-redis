package entity

import "time"

type View struct {
	ID        int64      `gorm:"primaryKey;autoIncrement;type:bigint"`
	ComicID   *int64     `gorm:"type:bigint;index"`
	Daily     int        `gorm:"type:int"`
	Weekly    int        `gorm:"type:int"`
	Monthly   int        `gorm:"type:int"`
	AllTime   int        `gorm:"type:int"`
	ViewedAt  *time.Time `gorm:"type:timestamp;index"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	Comic     *Comic     `gorm:"foreignKey:ComicID;constraint;onDelete:CASCADE;onUpdate:CASCADE"`
}
