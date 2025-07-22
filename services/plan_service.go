package services

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/samber/lo"

	"pkg.formatio/dao"
	"pkg.formatio/prisma/db"
	"pkg.formatio/types"
)

type IMachinePlanService interface {
	ListMachinePlans(types.ListMachinePlansArgs) (*[]db.MachinePlanModel, error)
	CreateMachinePlan(types.CreateMachinePlanArgs) (*db.MachinePlanModel, error)
	GetMachinePlan(types.GetMachinePlanArgs) (*db.MachinePlanModel, error)
	UpdateMachinePlan(types.UpdateMachinePlanArgs) (*db.MachinePlanModel, error)
	DeleteMachinePlan(types.DeleteMachinePlanArgs) error
}

type MachinePlanService struct {
	machinePlanDAO dao.IMachinePlanDao
}

func (m *MachinePlanService) calculateMonthlyRate(rate int) int {
	const hoursPerDay = 24
	const daysPerMonth = 30

	return rate * hoursPerDay * daysPerMonth
}

// DeleteMachinePlan implements MachinePlanServiceInterface.
func (m *MachinePlanService) DeleteMachinePlan(args types.DeleteMachinePlanArgs) error {
	err := m.machinePlanDAO.DeleteMachinePlan(args)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return nil
}

// GetMachinePlan implements MachinePlanServiceInterface.
func (m *MachinePlanService) GetMachinePlan(args types.GetMachinePlanArgs) (machinePlan *db.MachinePlanModel, err error) {
	machinePlan, err = m.machinePlanDAO.GetMachinePlan(args)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return machinePlan, err
}

// UpdateMachinePlan implements MachinePlanServiceInterface.
func (m *MachinePlanService) UpdateMachinePlan(args types.UpdateMachinePlanArgs) (machinePlan *db.MachinePlanModel, err error) {
	machinePlan, err = m.machinePlanDAO.UpdateMachinePlan(args)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_hourlyRate, _ := machinePlan.HourlyRate()
	hourlyRate, _ := strconv.Atoi(_hourlyRate.String())

	if hourlyRate != 0 {
		args.MonthlyRate = lo.ToPtr(m.calculateMonthlyRate(hourlyRate))
	}

	return machinePlan, err
}

// CreateMachinePlan implements MachinePlanServiceInterface.
func (m *MachinePlanService) CreateMachinePlan(args types.CreateMachinePlanArgs) (*db.MachinePlanModel, error) {
	args.HourlyRate = lo.Ternary(args.MonthlyRate != 0, float32(args.MonthlyRate/720), 0)

	machinePlan, err := m.machinePlanDAO.CreateMachinePlan(args)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return machinePlan, nil
}

// ListMachinePlans implements MachinePlanServiceInterface.
func (s *MachinePlanService) ListMachinePlans(args types.ListMachinePlansArgs) (*[]db.MachinePlanModel, error) {
	machinePlans, err := s.machinePlanDAO.ListMachinePlans(types.ListMachinePlansArgs(args))

	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return &machinePlans, nil
}

func NewMachinePlanService(machinePlanDAO dao.IMachinePlanDao) IMachinePlanService {
	return &MachinePlanService{machinePlanDAO: machinePlanDAO}
}
