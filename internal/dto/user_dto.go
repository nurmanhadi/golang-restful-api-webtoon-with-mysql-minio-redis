package dto

import (
	"time"
	"welltoon/pkg/enum"
)

type UserRegisterRequest struct {
	Username string `validate:"required,max=100" json:"username"`
	Password string `validate:"required,max=100" json:"password"`
}
type UserLoginRequest struct {
	Username string `validate:"required,max=100" json:"username"`
	Password string `validate:"required,max=100" json:"password"`
}
type UserUpdateRequest struct {
	Username    *string `validate:"omitempty,max=100" json:"username"`
	OldPassword *string `validate:"omitempty,max=100" json:"old_password"`
	NewPassword *string `validate:"omitempty,max=100" json:"new_password"`
}
type UserResponse struct {
	ID             int64     `json:"id"`
	Username       string    `json:"username"`
	Role           enum.ROLE `json:"role"`
	AvatarFilename *string   `json:"avatar_filename"`
	AvatarUrl      *string   `json:"avatar_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
type UserAddAdminRequest struct {
	Username string `validate:"required,max=100" json:"username"`
	Password string `validate:"required,max=100" json:"password"`
}
type UserTotalResponse struct {
	TotalUser int64 `json:"total_user"`
}
