package dto

type UserRegisterRequest struct {
	Username string `validate:"required,max=100" json:"username"`
	Password string `validate:"required,max=100" json:"password"`
}
type UserLoginRequest struct {
	Username string `validate:"required,max=100" json:"username"`
	Password string `validate:"required,max=100" json:"password"`
}
