package routers

import (
	"github.com/gofiber/fiber/v3"

	"pkg.formatio/handlers"
)

func NewWebhookRouter(app *fiber.App, webhookHandler handlers.IWebhookHandler) {
	webhookGroup := app.Group("/api/v1/")

	webhookGroup.Post("/webhook/github/", webhookHandler.GithubWebhookHandler)
}
