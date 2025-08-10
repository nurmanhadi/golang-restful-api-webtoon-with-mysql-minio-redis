package entity

import "time"

type Chapter struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;type:bigint"`
	ComicID   int64     `gorm:"type:bigint;index"`
	Number    int       `gorm:"type:int;index"`
	Publish   bool      `gorm:"type:bool"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Comic     *Comic    `gorm:"foreignKey:ComicID;contraint;onDelete:CASCADE;onUpdate:CASCADE"`
	Pages     []Page    `gorm:"foreignKey:ChapterID;constraint;onDelete:CASCADE;onUpdate:CASCADE"`
}
