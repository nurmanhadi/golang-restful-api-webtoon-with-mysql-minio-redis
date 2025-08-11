package handler

import (
	"welltoon/internal/dto"
	"welltoon/internal/service"
	"welltoon/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	RegisterUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
	UploadAvatar(c *fiber.Ctx) error
}
type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService: userService}
}
func (h *userHandler) RegisterUser(c *fiber.Ctx) error {
	request := new(dto.UserRegisterRequest)
	if err := c.BodyParser(request); err != nil {
		return response.Exception(400, err.Error())
	}
	err := h.userService.RegisterUser(request)
	if err != nil {
		return err
	}
	return response.Success(c, 201, "OK")
}
func (h *userHandler) LoginUser(c *fiber.Ctx) error {
	request := new(dto.UserLoginRequest)
	if err := c.BodyParser(request); err != nil {
		return response.Exception(400, err.Error())
	}
	result, err := h.userService.LoginUser(request)
	if err != nil {
		return err
	}
	return response.Success(c, 200, result)
}
func (h *userHandler) UploadAvatar(c *fiber.Ctx) error {
	userID := c.Params("userID", "0")
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return response.Exception(400, err.Error())
	}
	err = h.userService.UploadAvatar(userID, avatar)
	if err != nil {
		return err
	}
	return response.Success(c, 200, "OK")
}
