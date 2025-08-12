package dto

type EnumFilter struct {
	Type   *string `json:"type" validate:"omitempty,oneof=manga manhua manhwa"`
	Status *string `json:"status" validate:"omitempty,oneof=completed hiatus ongoing"`
	By     *string `json:"by" validate:"omitempty,oneof=daily weekly monthly all-time"`
}
