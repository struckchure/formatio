package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/samber/lo"
	"pkg.formatio/lib"
	"pkg.formatio/services"
	"pkg.formatio/types"
)

type IRepoConnectionHandler interface {
	ListRepoConnections(fiber.Ctx) error
	CreateRepoConnection(fiber.Ctx) error
	GetRepoConnection(fiber.Ctx) error
	UpdateRepoConnection(fiber.Ctx) error
	DeleteRepoConnection(fiber.Ctx) error
}

type RepoConnectionHandler struct {
	repoConnectionService *services.RepoConnectionService
}

// ListRepoConnections godoc
//
// @ID			listRepoConnections
// @Tags    repo-connection
// @Accept  json
// @Produce json
// @Param		args							query		types.ListRepoConnectionArgs	false "List Repo Connection"
// @Success 200								{array}	db.RepoConnectionModel
// @Router  /repo-connection	[get]
func (r *RepoConnectionHandler) ListRepoConnections(c fiber.Ctx) error {
	args := types.ListRepoConnectionArgs{OwnerId: lo.ToPtr(fiber.Locals[string](c, "userId"))}

	if err := c.Bind().Query(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	repoConnections, err := r.repoConnectionService.ListRepoConnections(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(repoConnections)
}

// CreateRepoConnection godoc
//
// @ID			createRepoConnection
// @Tags    repo-connection
// @Accept  json
// @Produce json
// @Param		args									body			types.CreateRepoConnectionArgs	true "Create Repo Connection"
// @Success 201										{object}	db.RepoConnectionModel
// @Router  /repo-connection			[post]
func (r *RepoConnectionHandler) CreateRepoConnection(c fiber.Ctx) error {
	args := types.CreateRepoConnectionArgs{UserId: fiber.Locals[string](c, "userId")}

	if err := c.Bind().JSON(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	connection, err := r.repoConnectionService.CreateRepoConnection(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(connection)
}

// GetRepoConnection godoc
//
// @ID			getRepoConnection
// @Tags    repo-connection
// @Accept  json
// @Produce json
// @Param		connectionId										path			string	true "Repo Connection Id"
// @Success 204															{object}	db.RepoConnectionModel
// @Router  /repo-connection/{connectionId} [get]
func (r *RepoConnectionHandler) GetRepoConnection(c fiber.Ctx) error {
	repoConnection, err := r.repoConnectionService.GetRepoConnection(
		types.GetRepoConnectionArgs{Id: fiber.Params[string](c, "connectionId")},
	)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(repoConnection)
}

// UpdateRepoConnection godoc
//
// @ID			updateRepoConnection
// @Tags    repo-connection
// @Accept  json
// @Produce json
// @Param		connectionId										path			string													true "Repo Connection Id"
// @Param		args														body			types.UpdateRepoConnectionArgs	true 	"Update Repo Connection"
// @Success 202															{object}	db.RepoConnectionModel
// @Router  /repo-connection/{connectionId} [patch]
func (r *RepoConnectionHandler) UpdateRepoConnection(c fiber.Ctx) error {
	args := types.UpdateRepoConnectionArgs{Id: fiber.Params[string](c, "connectionId")}

	if err := c.Bind().JSON(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	projects, err := r.repoConnectionService.UpdateRepoConnection(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusAccepted).JSON(projects)
}

// DeleteRepoConnection godoc
//
// @ID			deleteRepoConnection
// @Tags    repo-connection
// @Accept  json
// @Produce json
// @Param		connectionId										path			string	true "Repo Connection Id"
// @Success 204															{object}	nil
// @Router  /repo-connection/{connectionId} [delete]
func (r *RepoConnectionHandler) DeleteRepoConnection(c fiber.Ctx) error {
	err := r.repoConnectionService.DeleteRepoConnection(
		types.DeleteRepoConnectionArgs{Id: fiber.Params[string](c, "connectionId")},
	)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusNoContent).JSON(nil)
}

func NewRepoConnectionHandler(repoConnectionService *services.RepoConnectionService) IRepoConnectionHandler {
	return &RepoConnectionHandler{repoConnectionService: repoConnectionService}
}
