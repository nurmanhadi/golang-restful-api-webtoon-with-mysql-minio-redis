package handler

import (
	"welltoon/internal/dto"
	"welltoon/internal/service"
	"welltoon/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type ComicHandler interface {
	AddComic(c *fiber.Ctx) error
	UpdateComic(c *fiber.Ctx) error
	DeleteComic(c *fiber.Ctx) error
}
type comicHandler struct {
	comicService service.ComicService
}

func NewComicHandler(comicService service.ComicService) ComicHandler {
	return &comicHandler{comicService: comicService}
}
func (h *comicHandler) AddComic(c *fiber.Ctx) error {
	request := new(dto.ComicAddRequest)
	if err := c.BodyParser(request); err != nil {
		return response.Exception(400, err.Error())
	}
	err := h.comicService.AddComic(request)
	if err != nil {
		return err
	}
	return response.Success(c, 201, "OK")
}
func (h *comicHandler) UpdateComic(c *fiber.Ctx) error {
	comicID := c.Params("comicID")
	request := new(dto.ComicUpdateRequest)
	if err := c.BodyParser(request); err != nil {
		return response.Exception(400, err.Error())
	}
	err := h.comicService.UpdateComic(comicID, request)
	if err != nil {
		return err
	}
	return response.Success(c, 200, "OK")
}
func (h *comicHandler) DeleteComic(c *fiber.Ctx) error {
	comicID := c.Params("comicID")
	err := h.comicService.DeleteComic(comicID)
	if err != nil {
		return err
	}
	return response.Success(c, 200, "OK")
}
