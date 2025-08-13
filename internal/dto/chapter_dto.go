package dto

import "time"

type ChapterAddRequest struct {
	Number int `json:"number" validate:"required"`
}

type ChapterResponse struct {
	ID        int64     `json:"id"`
	ComicID   int64     `json:"comic_id"`
	Number    int       `json:"number"`
	Publish   bool      `json:"publish"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
