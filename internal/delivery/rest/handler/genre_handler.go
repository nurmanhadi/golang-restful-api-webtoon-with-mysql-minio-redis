package handler

import (
	"welltoon/internal/dto"
	"welltoon/internal/service"
	"welltoon/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type GenreHandler interface {
	AddGenre(c *fiber.Ctx) error
	UpdateGenre(c *fiber.Ctx) error
	DeleteGenre(c *fiber.Ctx) error
	GetAllGenre(c *fiber.Ctx) error
}
type genreHandler struct {
	genreService service.GenreService
}

func NewGenreHandler(genreService service.GenreService) GenreHandler {
	return &genreHandler{genreService: genreService}
}
func (h *genreHandler) AddGenre(c *fiber.Ctx) error {
	request := new(dto.GenreRequest)
	if err := c.BodyParser(request); err != nil {
		return response.Exception(400, err.Error())
	}
	err := h.genreService.AddGenre(request)
	if err != nil {
		return err
	}
	return response.Success(c, 201, "OK")
}
func (h *genreHandler) UpdateGenre(c *fiber.Ctx) error {
	genreID := c.Params("genreID")
	request := new(dto.GenreRequest)
	if err := c.BodyParser(request); err != nil {
		return response.Exception(400, err.Error())
	}
	err := h.genreService.UpdateGenre(genreID, request)
	if err != nil {
		return err
	}
	return response.Success(c, 200, "OK")
}
func (h *genreHandler) DeleteGenre(c *fiber.Ctx) error {
	genreID := c.Params("genreID")
	err := h.genreService.DeleteGenre(genreID)
	if err != nil {
		return err
	}
	return response.Success(c, 200, "OK")
}
func (h *genreHandler) GetAllGenre(c *fiber.Ctx) error {
	result, err := h.genreService.GetAllGenre()
	if err != nil {
		return err
	}
	return response.Success(c, 200, result)
}
