package dto

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
