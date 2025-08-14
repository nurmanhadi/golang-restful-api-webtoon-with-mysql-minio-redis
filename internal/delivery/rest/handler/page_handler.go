package handler

import (
	"welltoon/internal/service"
	"welltoon/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type PageHandler interface {
	AddBulkPage(c *fiber.Ctx) error
}
type pageHandler struct {
	pageService service.PageService
}

func NewPageHandler(pageService service.PageService) PageHandler {
	return &pageHandler{pageService: pageService}
}
func (h *pageHandler) AddBulkPage(c *fiber.Ctx) error {
	comicID := c.Params("comicID")
	chapterID := c.Params("chapterID")
	form, err := c.MultipartForm()
	if err != nil {
		return response.Exception(400, err.Error())
	}
	files := form.File["pages"]
	if len(files) > 20 {
		return response.Exception(400, "pages max 20")
	}
	err = h.pageService.AddBulkPage(comicID, chapterID, files)
	if err != nil {
		return err
	}
	return response.Success(c, 201, "OK")
}
