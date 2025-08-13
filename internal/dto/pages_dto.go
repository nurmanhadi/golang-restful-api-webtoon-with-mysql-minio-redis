package dto

import "time"

type PagesResponse struct {
	ID            int64     `json:"id"`
	ChapterID     int64     `json:"chapter_id"`
	ImageFilename string    `json:"image_filename"`
	ImageUrl      string    `json:"image_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
