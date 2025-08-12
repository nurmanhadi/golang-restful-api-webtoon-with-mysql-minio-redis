package routes

import (
	"welltoon/internal/delivery/rest/handler"
	"welltoon/internal/delivery/rest/middleware"
	"welltoon/pkg/enum"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	App          *fiber.App
	UserHandler  handler.UserHandler
	ComicHandler handler.ComicHandler
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
	comic.Get("/search", r.ComicHandler.SearchComic)    // search comic
	comic.Get("/recent", r.ComicHandler.GetComicRecent) // get comic recent
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
	comic.Get("/:slug", r.ComicHandler.GetComicBySlug) // get comic by slug
	comic.Post("/:comicID/cover",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.ComicHandler.UploadCover) // upload cover
}
