package handler

import (
	"welltoon/internal/dto"
	"welltoon/internal/service"
	"welltoon/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type GenreHandler interface {
	AddGenre(c *fiber.Ctx) error
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
