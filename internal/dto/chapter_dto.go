package dto

import "time"

type ChapterAddRequest struct {
	Number int `json:"number" validate:"required"`
}
type ChapterUpdateRequest struct {
	Number  *int  `json:"number" validate:"omitempty"`
	Publish *bool `json:"publish" validate:"omitempty"`
}
type ChapterResponse struct {
	ID        int64           `json:"id"`
	ComicID   int64           `json:"comic_id"`
	Number    int             `json:"number"`
	Publish   bool            `json:"publish"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Comic     *ComicResponse  `json:"comic"`
	Pages     []PagesResponse `json:"pages"`
}
