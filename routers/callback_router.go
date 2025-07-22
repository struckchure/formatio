package routers

import (
	"github.com/gofiber/fiber/v3"

	"pkg.formatio/handlers"
)

func NewCallbackRouter(app *fiber.App, callbackHandler handlers.ICallbackHandler) {
	callbackGroup := app.Group("/api/v1/callback")

	callbackGroup.Get("/github/", callbackHandler.GithubCallbackHandler)
	callbackGroup.Get("/flutterwave/", func(c fiber.Ctx) error {
		return nil
	})
}
