package routes

import (
	"welltoon/internal/delivery/rest/handler"
	"welltoon/internal/delivery/rest/middleware"
	"welltoon/pkg/enum"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	App         *fiber.App
	UserHandler handler.UserHandler
}

func (r *Route) Setup() {
	api := r.App.Group("/api")

	// user
	user := api.Group("/users")
	user.Post("/register", r.UserHandler.RegisterUser) // register user
	user.Post("/login", r.UserHandler.LoginUser)       // login user
	user.Post("/:userID/avatar",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_USER), string(enum.ROLE_ADMIN)}),
		r.UserHandler.UploadAvatar) // upload avatar
	user.Patch("/:userID",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_USER), string(enum.ROLE_ADMIN)}),
		r.UserHandler.UpdateUser) // update user
	user.Get("/:userID",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_USER), string(enum.ROLE_ADMIN)}),
		r.UserHandler.GetUserByID) // get user by id
	user.Post("/admins",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.UserHandler.AddAdmin) // add admin
	user.Delete("/:userID",
		middleware.JwtSession,
		middleware.RoleSession([]string{string(enum.ROLE_ADMIN)}),
		r.UserHandler.DeleteUser) // delete user
}
