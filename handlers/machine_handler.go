package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/samber/lo"
	"pkg.formatio/lib"
	"pkg.formatio/services"
	"pkg.formatio/types"
)

type IMachineHandler interface {
	ListMachines(fiber.Ctx) error
	CreateMachine(fiber.Ctx) error
	GetMachine(fiber.Ctx) error
	UpdateMachine(fiber.Ctx) error
	DeleteMachine(fiber.Ctx) error
}

type MachineHandler struct {
	machineService services.IMachineService
}

// ListMachines godoc
//
// @ID			listMachines
// @Tags    machines
// @Accept  json
// @Produce json
// @Param		args			query		types.ListMachineArgs	false "List Machine"
// @Success 200				{array}	db.MachineModel
// @Router  /machine	[get]
func (m *MachineHandler) ListMachines(c fiber.Ctx) error {
	args := types.ListMachineArgs{UserId: lo.ToPtr(fiber.Locals[string](c, "userId"))}

	if err := c.Bind().Query(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	machines, err := m.machineService.ListMachines(args)
	if err != nil {
		return err
	}

	return c.JSON(machines)
}

// CreateMachine godoc
//
// @ID			createMachine
// @Tags    machines
// @Accept  json
// @Produce json
// @Param		args			body			types.CreateMachineArgs	true "List Machine"
// @Success 201				{object}	db.MachineModel
// @Router  /machine	[post]
func (m *MachineHandler) CreateMachine(c fiber.Ctx) error {
	args := types.CreateMachineArgs{OwnerId: fiber.Locals[string](c, "userId")}

	if err := c.Bind().JSON(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	machine, err := m.machineService.CreateMachine(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(machine)
}

// GetMachine godoc
//
// @ID			getMachine
// @Tags    machines
// @Accept  json
// @Produce json
// @Param		machineId							path				string	true "Machine Id"
// @Success 200										{object}		db.MachineModel
// @Router  /machine/{machineId}	[get]
func (m *MachineHandler) GetMachine(c fiber.Ctx) error {
	machineId := fiber.Params[string](c, "machineId")
	machine, err := m.machineService.GetMachine(types.GetMachineArgs{Id: &machineId})

	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(machine)
}

// UpdateMachine godoc
//
// @ID			updateMachine
// @Tags    machines
// @Accept  json
// @Produce json
// @Param		machineId							path			string									true "Machine Id"
// @Param		args									body			types.UpdateMachineArgs	true "Update Machine"
// @Success 202										{object}	nil
// @Router  /machine/{machineId} [patch]
func (m *MachineHandler) UpdateMachine(c fiber.Ctx) error {
	args := types.UpdateMachineArgs{
		ID:      fiber.Params[string](c, "machineId"),
		OwnerId: lo.ToPtr(fiber.Locals[string](c, "userId")),
	}

	if err := c.Bind().JSON(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := m.machineService.UpdateMachine(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusAccepted).JSON(nil)
}

// DeleteMachine godoc
//
// @ID			deleteMachine
// @Tags    machines
// @Accept  json
// @Produce json
// @Param		machineId							path			string	true "Machine Id"
// @Success 204										{object}	nil
// @Router  /machine/{machineId} [delete]
func (m *MachineHandler) DeleteMachine(c fiber.Ctx) error {
	err := m.machineService.DeleteMachine(types.DeleteMachineArgs{Id: fiber.Params[string](c, "machineId")})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusNoContent).JSON(nil)
}

func NewMachineHandler(machineService services.IMachineService) IMachineHandler {
	return &MachineHandler{machineService: machineService}
}
