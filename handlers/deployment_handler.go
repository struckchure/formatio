package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/samber/lo"

	"pkg.formatio/dao"
	"pkg.formatio/lib"
	"pkg.formatio/services"
	"pkg.formatio/types"
)

type DeploymentHandler struct {
	rmq lib.RabbitMQ

	deploymentService    *services.DeploymentService
	deploymentLogService *services.DeploymentLogService
}

// ListDeployments godoc
//
// @ID					listDeployments
// @Tags        deployments
// @Accept      json
// @Produce     json
// @Param   		machineId			query 	string 										false "Machine Id"
// @Param				args					query		types.ListDeploymentArgs	true	"List Deployment"
// @Success     200						{array}	db.DeploymentModel
// @Router      /deployments	[get]
func (h *DeploymentHandler) ListDeployments(c fiber.Ctx) error {
	deployments, err := h.deploymentService.ListDeployments(types.ListDeploymentArgs{
		MachineId: lo.ToPtr(fiber.Query[string](c, "machineId")),
	})
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(deployments)
}

// GetDeployment godoc
//
// @ID					getDeploymentById
// @Tags        deployments
// @Accept      json
// @Produce     json
// @Param   		deploymentId			path 			string	true "Deployment Id"
// @Success     200								{object}	db.DeploymentModel
// @Router      /deployments/{deploymentId}	[get]
func (h *DeploymentHandler) GetDeployment(c fiber.Ctx) error {
	args := types.GetDeploymentArgs{Id: fiber.Params[string](c, "deploymentId")}

	deployments, err := h.deploymentService.GetDeployment(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(deployments)
}

// ListDeploymentLogs godoc
//
// @ID			listDeploymentLogsById
// @Tags    deployments
// @Accept  json
// @Produce json
// @Param   deploymentId											path 			string	true "Deployment Id"
// @Success 200																{array}		db.DeploymentLogModel
// @Router	/deployments/{deploymentId}/logs	[get]
func (h *DeploymentHandler) ListDeploymentLogs(c fiber.Ctx) error {
	deploymentId := fiber.Params[string](c, "deploymentId")
	deploymentLogs, err := h.deploymentLogService.ListDeploymentLogs(types.GetDeploymentArgs{
		Id: deploymentId,
	})
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(deploymentLogs)
}

// DeployRepo godoc
//
// @ID			deployRepo
// @Tags    deployments
// @Accept  json
// @Produce json
// @Param		args								body			services.DeployRepoArgs	true "Deploy Repo"
// @Success	200									{object}	nil
// @Router  /deployments/deploy [post]
func (h *DeploymentHandler) DeployRepo(c fiber.Ctx) error {
	args := services.DeployRepoArgs{UserId: fiber.Locals[string](c, "userId")}

	if err := c.Bind().JSON(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := h.deploymentService.DeployRepo(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(nil)
}

func NewDeploymentHandler(
	rmq lib.RabbitMQ,
	deploymentService *services.DeploymentService,
	deploymentLogDAO dao.IDeploymentLogDao,
	deploymentLogService *services.DeploymentLogService,
) *DeploymentHandler {
	return &DeploymentHandler{
		rmq: rmq,

		deploymentService:    deploymentService,
		deploymentLogService: deploymentLogService,
	}
}
