package routers

import (
	"github.com/gofiber/fiber/v3"

	"pkg.formatio/handlers"
	"pkg.formatio/middlewares"
)

func NewProjectRouter(
	app *fiber.App,

	repoConnectionHandler handlers.IRepoConnectionHandler,
	deploymentHandler *handlers.DeploymentHandler,
	userMiddleware middlewares.IUserMiddleware,
	jwtMiddleware middlewares.IJwtMiddleware,
) {
	repoConnectionGroup := app.Group("/api/v1/repo-connection", jwtMiddleware.Use, userMiddleware.Use)

	repoConnectionGroup.Post("/", repoConnectionHandler.CreateRepoConnection)
	repoConnectionGroup.Get("/", repoConnectionHandler.ListRepoConnections)
	repoConnectionGroup.Get(":connectionId", repoConnectionHandler.GetRepoConnection)
	repoConnectionGroup.Patch(":connectionId", repoConnectionHandler.UpdateRepoConnection)
	repoConnectionGroup.Delete(":connectionId", repoConnectionHandler.DeleteRepoConnection)

	deploymentGroup := app.Group("/api/v1/deployments", jwtMiddleware.Use, userMiddleware.Use)

	deploymentGroup.Get("/", deploymentHandler.ListDeployments)
	deploymentGroup.Post("/deploy", deploymentHandler.DeployRepo)
	deploymentGroup.Get("/:deploymentId", deploymentHandler.GetDeployment)
	deploymentGroup.Get("/:deploymentId/logs", deploymentHandler.ListDeploymentLogs)
}
