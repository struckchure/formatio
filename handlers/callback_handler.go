package handlers

import (
	"github.com/gofiber/fiber/v3"

	"pkg.formatio/lib"
	"pkg.formatio/services"
	"pkg.formatio/types"
)

type ICallbackHandler interface {
	GithubCallbackHandler(fiber.Ctx) error
}

type CallbackHandler struct {
	githubService services.IGithubService
}

// GithubCallbackHandler implements CallbackHandlerInterface.
func (w *CallbackHandler) GithubCallbackHandler(c fiber.Ctx) error {
	userId := fiber.Query[string](c, "state")
	code := fiber.Query[string](c, "code")

	redirectUrl, err := w.githubService.ConnectGithubAccount(types.ConnectGithubAccountArgs{
		UserId: userId,
		Code:   code,
	})
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Redirect().Status(fiber.StatusMovedPermanently).To(*redirectUrl)
}

func NewCallbackHandler(githubService services.IGithubService) ICallbackHandler {
	return &CallbackHandler{githubService: githubService}
}
