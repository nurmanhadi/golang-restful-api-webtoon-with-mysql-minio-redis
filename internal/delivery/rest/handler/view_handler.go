package handler

import (
	"welltoon/internal/dto"
	"welltoon/internal/service"
	"welltoon/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type ViewHandler interface {
	AddView(c *fiber.Ctx) error
	GetView(c *fiber.Ctx) error
}
type viewHandler struct {
	viewService service.ViewService
}

func NewViewHandler(viewService service.ViewService) ViewHandler {
	return &viewHandler{viewService: viewService}
}
func (h *viewHandler) AddView(c *fiber.Ctx) error {
	request := new(dto.ViewAddRequest)
	if err := c.BodyParser(request); err != nil {
		return response.Exception(400, err.Error())
	}
	err := h.viewService.AddView(request)
	if err != nil {
		return err
	}
	return response.Success(c, 201, "OK")
}
func (h *viewHandler) GetView(c *fiber.Ctx) error {
	result, err := h.viewService.GetView()
	if err != nil {
		return err
	}
	return response.Success(c, 200, result)
}
