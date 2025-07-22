package handlers

import (
	"github.com/gofiber/fiber/v3"

	"pkg.formatio/lib"
	"pkg.formatio/services"
	"pkg.formatio/types"
)

type IGithubHandler interface {
	ListRepositories(fiber.Ctx) error
	AuthorizeGithubAccount(fiber.Ctx) error
	UpdateAppAccess(fiber.Ctx) error
	ListAccountConnections(fiber.Ctx) error
}

type GithubHandler struct {
	githubService services.IGithubService
}

// ListRepositories godoc
//
// @ID			ListRepositories
// @Tags    github
// @Accept  json
// @Produce json
// @Param   args						query 	types.ListRepositoriesArgs	true "List Repositories"
// @Success 200							{array}	types.Repository
// @Router	/gh/repos				[get]
func (g *GithubHandler) ListRepositories(c fiber.Ctx) error {
	args := types.ListRepositoriesArgs{UserId: fiber.Locals[string](c, "userId")}

	if err := c.Bind().Query(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	repos, err := g.githubService.ListRepositories(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(repos)
}

type RedirectResult struct {
	Link string `json:"link"`
}

// AuthorizeGithubAccount godoc
//
// @ID			authorizeGithubAccount
// @Tags    github
// @Accept  json
// @Produce json
// @Param   args									query 		types.AuthorizeGithubAccountArgs	true "Authorize Github Account"
// @Success 200										{object}	RedirectResult
// @Router	/gh/authorize					[get]
func (g *GithubHandler) AuthorizeGithubAccount(c fiber.Ctx) error {
	args := types.AuthorizeGithubAccountArgs{UserId: fiber.Locals[string](c, "userId")}

	if err := c.Bind().Query(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	link, err := g.githubService.AuthorizeGithubAccount(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(RedirectResult{Link: link})
}

// UpdateAppAccess godoc
//
// @ID			updateAppAccess
// @Tags    github
// @Accept  json
// @Produce json
// @Success 200										{object}	RedirectResult
// @Router	/gh/update-app-access	[get]
func (g *GithubHandler) UpdateAppAccess(c fiber.Ctx) error {
	args := types.AuthorizeGithubAccountArgs{UserId: fiber.Locals[string](c, "userId")}

	link, err := g.githubService.UpdateAppAccess(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(RedirectResult{Link: *link})
}

// ListAccountConnections godoc
//
// @ID			listAccountConnections
// @Tags    github
// @Accept  json
// @Produce json
// @Success 200											{array}	db.GithubAccountConnectionModel
// @Router	/gh/account-connections	[get]
func (g *GithubHandler) ListAccountConnections(c fiber.Ctx) error {
	args := types.ListGithubAccountConnectionsArgs{UserId: fiber.Locals[*string](c, "userId")}

	accountConnections, err := g.githubService.ListAccountConnections(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(accountConnections)
}

func NewGithubHandler(githubService services.IGithubService) IGithubHandler {
	return &GithubHandler{githubService: githubService}
}
