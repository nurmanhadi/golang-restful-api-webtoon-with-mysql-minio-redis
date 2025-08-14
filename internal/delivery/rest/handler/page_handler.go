package handler

import (
	"welltoon/internal/service"
	"welltoon/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type PageHandler interface {
	AddBulkPage(c *fiber.Ctx) error
	DeletePage(c *fiber.Ctx) error
}
type pageHandler struct {
	pageService service.PageService
}

func NewPageHandler(pageService service.PageService) PageHandler {
	return &pageHandler{pageService: pageService}
}
func (h *pageHandler) AddBulkPage(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return response.Exception(400, err.Error())
	}
	chapterID := c.FormValue("chapter_id", "none")
	if chapterID == "none" {
		return response.Exception(400, "chapter_id required")
	}
	files := form.File["pages"]
	if len(files) > 20 {
		return response.Exception(400, "pages max 20")
	}
	err = h.pageService.AddBulkPage(chapterID, files)
	if err != nil {
		return err
	}
	return response.Success(c, 201, "OK")
}
func (h *pageHandler) DeletePage(c *fiber.Ctx) error {
	pageID := c.Params("pageID")
	err := h.pageService.DeletePage(pageID)
	if err != nil {
		return err
	}
	return response.Success(c, 200, "OK")
}
