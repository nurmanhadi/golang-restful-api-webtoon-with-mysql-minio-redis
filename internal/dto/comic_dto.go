package dto

import (
	"time"
	"welltoon/pkg/enum"
)

type ComicAddRequest struct {
	Title    string      `validate:"required,max=200" json:"title"`
	Synopsis *string     `validate:"omitempty" json:"synopsis"`
	Author   string      `validate:"required,max=50" json:"author"`
	Artist   string      `validate:"required,max=50" json:"artist"`
	Type     enum.TYPE   `validate:"required,oneof=manga manhua manhwa" json:"type"`
	Status   enum.STATUS `validate:"required,oneof=completed hiatus ongoing" json:"status"`
}
type ComicUpdateRequest struct {
	Title    *string      `validate:"omitempty,max=200" json:"title"`
	Synopsis *string      `validate:"omitempty" json:"synopsis"`
	Author   *string      `validate:"omitempty,max=50" json:"author"`
	Artist   *string      `validate:"omitempty,max=50" json:"artist"`
	Type     *enum.TYPE   `validate:"omitempty,oneof=manga manhua manhwa" json:"type"`
	Status   *enum.STATUS `validate:"omitempty,oneof=completed hiatus ongoing" json:"status"`
}
type ComicResponse struct {
	ID            int64             `json:"id"`
	Title         string            `json:"title"`
	Slug          string            `json:"slug"`
	Synopsis      *string           `json:"synopsis"`
	Author        string            `json:"author"`
	Artist        string            `json:"artist"`
	Type          enum.TYPE         `json:"type"`
	Status        enum.STATUS       `json:"status"`
	CoverFilename *string           `json:"cover_filename"`
	CoverUrl      *string           `json:"cover_url"`
	PostOn        time.Time         `json:"post_on"`
	UpdatedOn     *time.Time        `json:"updated_on"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
	Chapters      []ChapterResponse `json:"chapters"`
	Genres        []GenreResponse   `json:"genres"`
}
type ComicTotalResponse struct {
	TotalComic int `json:"total_comic"`
}
