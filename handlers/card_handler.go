package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/samber/lo"

	"pkg.formatio/lib"
	"pkg.formatio/services"
	"pkg.formatio/types"
)

type CardHandler struct {
	cardService services.CardService
}

// PreAuthorizeCard godoc
//
// @ID			preAuthorizeCard
// @Tags    billing
// @Accept  json
// @Produce json
// @Param		args													body			types.PreAuthorizeCardArgs	true 	"Pre-Authorize Card"
// @Success 200														{object}	types.PreAuthorizeCardResult
// @Router  /billing/cards/pre-authorize 	[post]
func (h *CardHandler) PreAuthorizeCard(c fiber.Ctx) error {
	args := types.PreAuthorizeCardArgs{UserId: fiber.Locals[string](c, "userId")}

	if err := c.Bind().JSON(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	preAuthCard, err := h.cardService.PreAuthorizeCard(args)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(preAuthCard)
}

// AuthorizeCard godoc
//
// @ID			authorizeCard
// @Tags    billing
// @Accept  json
// @Produce json
// @Param		args								body			types.AuthorizeCardArgs	true 	"Update Repo Connection"
// @Success 202									{object}	db.CardModel
// @Router  /billing/cards/authorize 	[post]
func (h *CardHandler) AuthorizeCard(c fiber.Ctx) error {
	args := types.AuthorizeCardArgs{}

	if err := c.Bind().JSON(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	authCard, err := h.cardService.AuthorizeCard(args)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(authCard)
}

// ListCards godoc
//
// @ID			listCards
// @Tags    billing
// @Accept  json
// @Produce json
// @Param		args						query		types.ListCardArgs	true	"Update Repo Connection"
// @Success 200							{array}	db.CardModel
// @Router  /billing/cards [get]
func (h *CardHandler) ListCards(c fiber.Ctx) error {
	args := types.ListCardArgs{UserId: lo.ToPtr(fiber.Locals[string](c, "userId"))}

	if err := c.Bind().Query(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	cards, err := h.cardService.ListCard(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(cards)
}

// UpdateCard godoc
//
// @ID			UpdateCard
// @Tags    billing
// @Accept  json
// @Produce json
// @Param		cardId									path			string								true	"Card Id"
// @Param		args										body			types.UpdateCardArgs	true	"Update Card"
// @Success 202											{object}	db.CardModel
// @Router  /billing/cards/{cardId} [patch]
func (h *CardHandler) UpdateCard(c fiber.Ctx) error {
	args := types.UpdateCardArgs{Id: fiber.Params[string](c, "cardId")}

	if err := c.Bind().JSON(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	cards, err := h.cardService.UpdateCard(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusAccepted).JSON(cards)
}

// DeleteCard godoc
//
// @ID			deleteCard
// @Tags    billing
// @Accept  json
// @Produce json
// @Param		cardId									path			string	true "Card Id"
// @Success 204											{object}	nil
// @Router  /billing/cards/{cardId} [delete]
func (h *CardHandler) DeleteCard(c fiber.Ctx) error {
	cardId := fiber.Params[string](c, "cardId")

	err := h.cardService.DeleteCard(types.DeleteCardArgs{Id: cardId})
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusNoContent).JSON(nil)
}

func NewCardHandler(cardService services.CardService) CardHandler {
	return CardHandler{cardService: cardService}
}
