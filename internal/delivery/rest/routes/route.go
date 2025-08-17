package routes

import (
	"welltoon/internal/delivery/rest/handler"
	"welltoon/internal/delivery/rest/middleware"
	"welltoon/pkg/enum"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	App            *fiber.App
	UserHandler    handler.UserHandler
	ComicHandler   handler.ComicHandler
	ChapterHandler handler.ChapterHandler
	PageHandler    handler.PageHandler
	GenreHandler   handler.GenreHandler
	ViewHandler    handler.ViewHandler
}

func (r *Route) Setup() {
	api := r.App.Group("/api")

	// user
	user := api.Group("/users")
	user.Post("/register", r.UserHandler.RegisterUser) // register user
	user.Post("/login", r.UserHandler.LoginUser)       // login user
	user.Post("/admins",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.UserHandler.AddAdmin) // add admin
	user.Get("/total",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.UserHandler.GetTotalUser) // get total user
	user.Patch("/:userID",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_USER), string(enum.ROLE_ADMIN)}),
		r.UserHandler.UpdateUser) // update user
	user.Get("/:userID",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_USER), string(enum.ROLE_ADMIN)}),
		r.UserHandler.GetUserByID) // get user by id
	user.Delete("/:userID",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.UserHandler.DeleteUser) // delete user
	user.Post("/:userID/avatar",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_USER), string(enum.ROLE_ADMIN)}),
		r.UserHandler.UploadAvatar) // upload avatar
	user.Post("/:userID/logout",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_USER), string(enum.ROLE_ADMIN)}),
		r.UserHandler.LogoutUser) // logout user

	// comic
	comic := api.Group("/comics")
	comic.Post("/",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.ComicHandler.AddComic) // add comic
	comic.Get("/", r.ComicHandler.GetComicByTypeAndStatus) // get comic by type and status
	comic.Get("/search", r.ComicHandler.SearchComic)       // search comic
	comic.Get("/recent", r.ComicHandler.GetComicRecent)    // get comic recent
	comic.Get("/new", r.ComicHandler.GetComicNew)          // get comic new
	comic.Get("/total",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.ComicHandler.GetTotalComic) // get total comic
	comic.Patch("/:comicID",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.ComicHandler.UpdateComic) // update comic
	comic.Delete("/:comicID",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.ComicHandler.DeleteComic) // delete comic
	comic.Get("/:slug", r.ComicHandler.GetComicBySlug)          // get comic by slug
	comic.Get("/:slug/related", r.ComicHandler.GetComicRelated) // get comic related
	comic.Post("/:comicID/cover",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.ComicHandler.UploadCover) // upload cover
	comic.Post("/:comicID/genre",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.ComicHandler.ComicAddGenre) // comic add genre

	// chapter
	chapter := api.Group("/chapters")
	comic.Get("/:slug/chapters/:number", r.ChapterHandler.GetChapterBySlugAndNumber) // get chapter by slug and number
	chapter.Post("/",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.ChapterHandler.AddChapter) // add chapter
	chapter.Patch("/:chapterID",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.ChapterHandler.UpdateChapter) // update chapter
	chapter.Delete("/:chapterID",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.ChapterHandler.DeleteChapter) // delete chapter

	// page
	page := api.Group("/pages")
	page.Post("/",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.PageHandler.AddBulkPage) // add bulk page
	page.Delete("/:pageID",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.PageHandler.DeletePage) // delete page

	// genre
	genre := api.Group("/genres")
	genre.Get("/", r.GenreHandler.GetAllGenre) // get all genre
	genre.Post("/",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.GenreHandler.AddGenre) // add genre
	genre.Patch("/:genreID",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.GenreHandler.UpdateGenre) // update genre
	genre.Delete("/:genreID",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.GenreHandler.DeleteGenre) // delete genre
	genre.Get("/:name", r.GenreHandler.GetComicByGenreName) // get comic by genre name

	// view
	view := api.Group("/views")
	view.Post("/", r.ViewHandler.AddView) // add view
	view.Get("/", r.ViewHandler.GetView)  // get view
}
