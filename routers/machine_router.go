package routers

import (
	"github.com/gofiber/fiber/v3"

	"pkg.formatio/handlers"
	"pkg.formatio/middlewares"
)

func NewMachineRouter(
	app *fiber.App,
	machineHandler handlers.IMachineHandler,
	networkHandler handlers.INetworkHandler,
	userMiddleware middlewares.IUserMiddleware,
	jwtMiddleware middlewares.IJwtMiddleware,
) {
	// TODO: implement object owner permissions

	machineGroup := app.Group("/api/v1/machine", jwtMiddleware.Use, userMiddleware.Use)

	machineGroup.Get("/", machineHandler.ListMachines)
	machineGroup.Post("/", machineHandler.CreateMachine)
	machineGroup.Get("/:machineId", machineHandler.GetMachine)
	machineGroup.Patch("/:machineId", machineHandler.UpdateMachine)
	machineGroup.Delete("/:machineId", machineHandler.DeleteMachine)

	networkGroup := app.Group("/api/v1/network", jwtMiddleware.Use, userMiddleware.Use)

	networkGroup.Get("/", networkHandler.ListNetworks)
	networkGroup.Post("/", networkHandler.CreateNetwork)
	networkGroup.Delete("/:networkId", networkHandler.DeleteNetwork)
}
