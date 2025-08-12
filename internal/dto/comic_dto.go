package dto

import "welltoon/pkg/enum"

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
