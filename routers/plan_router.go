package routers

import (
	"github.com/gofiber/fiber/v3"

	"pkg.formatio/handlers"
	"pkg.formatio/middlewares"
)

func NewPlanRouter(
	app *fiber.App,
	machinePlanHandler handlers.IMachinePlanHandler,
	userMiddleware middlewares.IUserMiddleware,
	jwtMiddleware middlewares.IJwtMiddleware,
) {
	machinePlanGroup := app.Group("/api/v1/plans")
	machinePlanGroup.Get("/", machinePlanHandler.ListMachinePlans)
	machinePlanGroup.Get("/:machinePlanId", machinePlanHandler.GetMachinePlan)

	machinePlanGroup.Use(jwtMiddleware.Use, userMiddleware.Use, middlewares.UserRolesMiddleware([]string{"ADMIN"})).
		Post("/", machinePlanHandler.CreateMachinePlan).
		Patch("/:machinePlanId", machinePlanHandler.UpdateMachinePlan).
		Delete("/:machinePlanId", machinePlanHandler.DeleteMachinePlan)
}
