package routers

import (
	"github.com/gofiber/fiber/v3"

	"pkg.formatio/handlers"
	"pkg.formatio/middlewares"
)

func NewBillingRouter(
	app *fiber.App,
	userMiddleware middlewares.IUserMiddleware,
	jwtMiddleware middlewares.IJwtMiddleware,
	cardHandler handlers.CardHandler,
	invoiceHandler handlers.InvoiceHandler,
) {
	cardGroup := app.Group("/api/v1/billing/cards", jwtMiddleware.Use, userMiddleware.Use)

	cardGroup.Post("pre-authorize", cardHandler.PreAuthorizeCard)
	cardGroup.Post("authorize", cardHandler.AuthorizeCard)
	cardGroup.Get("", cardHandler.ListCards)
	cardGroup.Patch(":cardId", cardHandler.UpdateCard)
	cardGroup.Delete(":cardId", cardHandler.DeleteCard)

	invoiceGroup := app.Group("/api/v1/billing/invoice", jwtMiddleware.Use, userMiddleware.Use)

	invoiceGroup.Get("", invoiceHandler.ListInvoice)
	invoiceGroup.Get(":invoiceId", invoiceHandler.GetInvoice)
}
