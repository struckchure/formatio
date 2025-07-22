package routers

import (
	"github.com/gofiber/fiber/v3"

	"pkg.formatio/handlers"
	"pkg.formatio/middlewares"
)

func NewUsersRouter(
	app *fiber.App,
	userHandler handlers.IUserHandler,
	userMiddleware middlewares.IUserMiddleware,
	jwtMiddleware middlewares.IJwtMiddleware,
) {
	authGroup := app.Group("/api/v1/auth")
	authGroup.Post("/register/", userHandler.RegisterUser)
	authGroup.Post("/login/", userHandler.LoginUser)
	authGroup.Post("/refresh-access-token/", userHandler.RefreshAccessToken)
	authGroup.Post("/social-connection/", userHandler.AuthSocialConnection)

	userGroup := app.Group("/api/v1/user", jwtMiddleware.Use, userMiddleware.Use)
	userGroup.Get("/profile/", userHandler.GetProfileUser)
	userGroup.Patch("/profile/", userHandler.UpdateProfileUser)
}
