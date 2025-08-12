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
	GetComicBySlug(c *fiber.Ctx) error
	UploadCover(c *fiber.Ctx) error
	GetComicRecent(c *fiber.Ctx) error
	GetTotalComic(c *fiber.Ctx) error
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
func (h *comicHandler) GetComicBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	result, err := h.comicService.GetComicBySlug(slug)
	if err != nil {
		return err
	}
	return response.Success(c, 200, result)
}
func (h *comicHandler) UploadCover(c *fiber.Ctx) error {
	comicID := c.Params("comicID")
	cover, err := c.FormFile("cover")
	if err != nil {
		return response.Exception(400, "cover required")
	}
	err = h.comicService.UploadCover(comicID, cover)
	if err != nil {
		return err
	}
	return response.Success(c, 200, "OK")
}
func (h *comicHandler) GetComicRecent(c *fiber.Ctx) error {
	page := c.Query("page", "1")
	size := c.Query("size", "10")
	result, err := h.comicService.GetComicRecent(page, size)
	if err != nil {
		return err
	}
	return response.Success(c, 200, result)
}
func (h *comicHandler) GetTotalComic(c *fiber.Ctx) error {
	result, err := h.comicService.GetTotalComic()
	if err != nil {
		return err
	}
	return response.Success(c, 200, result)
}
