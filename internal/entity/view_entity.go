package entity

import "time"

type View struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;type:bigint"`
	ComicID   int64     `gorm:"type:bigint;index"`
	ViewedAt  time.Time `gorm:"type:timestamp;index"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Comic     *Comic    `gorm:"foreignKey:ComicID;constraint;onDelete:CASCADE;onUpdate:CASCADE"`
}
