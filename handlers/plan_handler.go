package handlers

import (
	"github.com/gofiber/fiber/v3"
	"pkg.formatio/lib"
	"pkg.formatio/services"
	"pkg.formatio/types"
)

type IMachinePlanHandler interface {
	ListMachinePlans(fiber.Ctx) error
	CreateMachinePlan(fiber.Ctx) error
	GetMachinePlan(fiber.Ctx) error
	UpdateMachinePlan(fiber.Ctx) error
	DeleteMachinePlan(fiber.Ctx) error
}

type MachinePlanHandler struct {
	machinePlanService services.IMachinePlanService
}

// ListMachinePlans godoc
//
// @ID			listMachinePlans
// @Tags    plans
// @Accept  json
// @Produce json
// @Param		args		query		types.ListMachinePlansArgs	true "List Machine"
// @Success 200			{array}	db.MachinePlanModel
// @Router  /plans 	[get]
func (m *MachinePlanHandler) ListMachinePlans(c fiber.Ctx) error {
	args := types.ListMachinePlansArgs{}

	if err := c.Bind().Query(&args); err != nil {
		return err
	}

	machinePlans, err := m.machinePlanService.ListMachinePlans(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(machinePlans)
}

// CreateMachinePlan godoc
//
// @ID			createMachinePlan
// @Tags    plans
// @Accept  json
// @Produce json
// @Param		args		body			types.CreateMachinePlanArgs	true "Create Machine Plan"
// @Success 201			{object}	db.MachinePlanModel
// @Router  /plans 	[post]
func (m *MachinePlanHandler) CreateMachinePlan(c fiber.Ctx) error {
	args := types.CreateMachinePlanArgs{}
	if err := c.Bind().JSON(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	machinePlan, err := m.machinePlanService.CreateMachinePlan(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(machinePlan)
}

// GetMachinePlan godoc
//
// @ID			getMachinePlan
// @Tags    plans
// @Accept  json
// @Produce json
// @Param		machinePlanId						path			string	true	"Machine Plan Id"
// @Success 200											{object}	db.MachinePlanModel
// @Router  /plans/{machinePlanId} 	[get]
func (m *MachinePlanHandler) GetMachinePlan(c fiber.Ctx) error {
	machinePlan, err := m.machinePlanService.GetMachinePlan(
		types.GetMachinePlanArgs{Id: fiber.Params[string](c, "machinePlanId")},
	)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.JSON(machinePlan)
}

// UpdateMachinePlan godoc
//
// @ID			updateMachinePlan
// @Tags    plans
// @Accept  json
// @Produce json
// @Param		machinePlanId						path			string											true "Machine Plan Id"
// @Param		args										body			types.UpdateMachinePlanArgs	true 	"Update Machine"
// @Success 202											{object}	db.MachinePlanModel
// @Router  /plans/{machinePlanId} 	[patch]
func (m *MachinePlanHandler) UpdateMachinePlan(c fiber.Ctx) error {
	args := types.UpdateMachinePlanArgs{Id: fiber.Params[string](c, "machinePlanId")}

	if err := c.Bind().JSON(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	machinePlan, err := m.machinePlanService.UpdateMachinePlan(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusAccepted).JSON(machinePlan)
}

// DeleteMachinePlan godoc
//
// @ID			deleteMachinePlan
// @Tags    plans
// @Accept  json
// @Produce json
// @Param		machinePlanId						path			string	true "Machine Plan Id"
// @Success 204											{object}	nil
// @Router  /plans/{machinePlanId} 	[delete]
func (m *MachinePlanHandler) DeleteMachinePlan(c fiber.Ctx) error {
	machinePlanId := fiber.Params[string](c, "machinePlanId")
	err := m.machinePlanService.DeleteMachinePlan(types.DeleteMachinePlanArgs{Id: machinePlanId})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}

	return c.Status(fiber.StatusNoContent).JSON(nil)
}

func NewMachinePlanHandler(machinePlanService services.IMachinePlanService) IMachinePlanHandler {
	return &MachinePlanHandler{machinePlanService: machinePlanService}
}
