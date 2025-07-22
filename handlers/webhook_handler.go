package handlers

import (
	"github.com/gofiber/fiber/v3"

	"pkg.formatio/services"
	"pkg.formatio/types"
)

type IWebhookHandler interface {
	GithubWebhookHandler(fiber.Ctx) error
}

type WebhookHandler struct {
	githubService services.IGithubService
}

// GithubWebhookHandler implements WebhookHandlerInterface.
func (w *WebhookHandler) GithubWebhookHandler(c fiber.Ctx) error {
	var pushEvent types.PushEvent
	if err := c.Bind().JSON(&pushEvent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := w.githubService.DeployRepo(types.DeployRepoArgs{
		InstallationId: pushEvent.Installation.ID,
		RepoId:         pushEvent.Repository.ID,
		RepoFullName:   pushEvent.Repository.FullName,
		Ref:            pushEvent.Ref,
		CommitHash:     &pushEvent.HeadCommit.ID,
		CommitMessage:  &pushEvent.HeadCommit.Message,
		Author:         &pushEvent.HeadCommit.Author.Name,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

func NewWebhookHandler(githubService services.IGithubService) IWebhookHandler {
	return &WebhookHandler{githubService: githubService}
}
