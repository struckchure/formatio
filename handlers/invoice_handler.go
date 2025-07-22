package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/samber/lo"

	"pkg.formatio/lib"
	"pkg.formatio/services"
	"pkg.formatio/types"
)

type InvoiceHandler struct {
	invoiceService services.InvoiceService
}

// ListInvoice godoc
//
// @ID			listInvoice
// @Tags    billing
// @Accept  json
// @Produce json
// @Param		args							body		types.ListInvoiceArgs	true "List Invoice"
// @Success 202								{array}	db.InvoiceModel
// @Router  /billing/invoice 	[get]
func (h *InvoiceHandler) ListInvoice(c fiber.Ctx) error {
	cards, err := h.invoiceService.ListInvoice(
		types.ListInvoiceArgs{UserId: lo.ToPtr(fiber.Locals[string](c, "userId"))},
	)

	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(cards)
}

// GetInvoice godoc
//
// @ID			getInvoice
// @Tags    billing
// @Accept  json
// @Produce json
// @Param		invoiceId											path			string	true "Invoice Id"
// @Success 200														{object}	db.InvoiceModel
// @Router  /billing/invoice/{invoiceId} 	[get]
func (h *InvoiceHandler) GetInvoice(c fiber.Ctx) error {
	cardId := fiber.Params[string](c, "invoiceId")

	invoice, err := h.invoiceService.GetInvoice(types.GetInvoiceArgs{Id: cardId})
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(invoice)
}

func NewInvoiceHandler(invoiceService services.InvoiceService) InvoiceHandler {
	return InvoiceHandler{invoiceService: invoiceService}
}
