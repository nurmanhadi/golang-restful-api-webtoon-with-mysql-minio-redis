package handler

import (
	"welltoon/internal/dto"
	"welltoon/internal/service"
	"welltoon/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type ChapterHandler interface {
	AddChapter(c *fiber.Ctx) error
	UpdateChapter(c *fiber.Ctx) error
	DeleteChapter(c *fiber.Ctx) error
	GetChapterBySlugAndNumber(c *fiber.Ctx) error
}
type chapterHandler struct {
	chapterService service.ChapterService
}

func NewChapterHandler(chapterService service.ChapterService) ChapterHandler {
	return &chapterHandler{chapterService: chapterService}
}
func (h *chapterHandler) AddChapter(c *fiber.Ctx) error {
	comicID := c.Params("comicID")
	request := new(dto.ChapterAddRequest)
	if err := c.BodyParser(request); err != nil {
		return response.Exception(400, err.Error())
	}
	err := h.chapterService.AddChapter(comicID, request)
	if err != nil {
		return err
	}
	return response.Success(c, 201, "OK")
}
func (h *chapterHandler) UpdateChapter(c *fiber.Ctx) error {
	comicID := c.Params("comicID")
	chapterID := c.Params("chapterID")
	request := new(dto.ChapterUpdateRequest)
	if err := c.BodyParser(request); err != nil {
		return response.Exception(400, err.Error())
	}
	err := h.chapterService.UpdateChapter(comicID, chapterID, request)
	if err != nil {
		return err
	}
	return response.Success(c, 200, "OK")
}
func (h *chapterHandler) DeleteChapter(c *fiber.Ctx) error {
	comicID := c.Params("comicID")
	chapterID := c.Params("chapterID")
	err := h.chapterService.DeleteChapter(comicID, chapterID)
	if err != nil {
		return err
	}
	return response.Success(c, 200, "OK")
}
func (h *chapterHandler) GetChapterBySlugAndNumber(c *fiber.Ctx) error {
	slug := c.Params("slug")
	number := c.Params("number")
	result, err := h.chapterService.GetChapterBySlugAndNumber(slug, number)
	if err != nil {
		return err
	}
	return response.Success(c, 200, result)
}
