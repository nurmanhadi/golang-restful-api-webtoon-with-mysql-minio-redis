package dto

import "time"

type GenreRequest struct {
	Name string `json:"name" validate:"required,max=50"`
}
type GenreResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
