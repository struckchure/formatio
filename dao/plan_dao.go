package dao

import (
	"context"

	"github.com/shopspring/decimal"
	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/types"
)

type IMachinePlanDao interface {
	ListMachinePlans(types.ListMachinePlansArgs) ([]db.MachinePlanModel, error)
	CreateMachinePlan(types.CreateMachinePlanArgs) (*db.MachinePlanModel, error)
	GetMachinePlan(types.GetMachinePlanArgs) (*db.MachinePlanModel, error)
	UpdateMachinePlan(types.UpdateMachinePlanArgs) (*db.MachinePlanModel, error)
	DeleteMachinePlan(types.DeleteMachinePlanArgs) error
}

type MachinePlanDao struct {
	client *db.PrismaClient
	ctx    context.Context
}

func (m *MachinePlanDao) ListMachinePlans(args types.ListMachinePlansArgs) ([]db.MachinePlanModel, error) {
	args.Skip = lib.UseDefaultValueIf(0, args.Skip, 0)
	args.Take = lib.UseDefaultValueIf(0, args.Take, 10)

	return m.client.MachinePlan.
		FindMany().
		Skip(args.Skip).
		Take(args.Take).
		OrderBy(db.MachinePlan.CreatedAt.Order(db.DESC)).
		Exec(m.ctx)
}

func (m *MachinePlanDao) CreateMachinePlan(args types.CreateMachinePlanArgs) (*db.MachinePlanModel, error) {
	return m.client.MachinePlan.
		CreateOne(
			db.MachinePlan.CPU.Set(args.Cpu),
			db.MachinePlan.Memory.Set(args.Memory),
			db.MachinePlan.Name.Set(args.Name),
			db.MachinePlan.MonthlyRate.Set(decimal.NewFromInt(int64(args.MonthlyRate))),
			// db.MachinePlan.HourlyRate.Set(decimal.NewFromInt(int64(args.HourlyRate))),
		).
		Exec(m.ctx)
}

func (m *MachinePlanDao) GetMachinePlan(args types.GetMachinePlanArgs) (*db.MachinePlanModel, error) {
	return m.client.MachinePlan.
		FindFirst(db.MachinePlan.ID.Equals(args.Id)).
		Exec(m.ctx)
}

func (m *MachinePlanDao) UpdateMachinePlan(args types.UpdateMachinePlanArgs) (*db.MachinePlanModel, error) {
	return m.client.MachinePlan.
		FindUnique(db.MachinePlan.ID.Equals(args.Id)).
		Update(
			db.MachinePlan.Name.SetIfPresent(args.Name),
			// db.MachinePlan.Currency.SetIfPresent(args.Currency),
			// db.MachinePlan.MonthlyRate.SetIfPresent(decimal.NewFromInt(int64(args.MonthlyRate))),
			// db.MachinePlan.HourlyRate.SetIfPresent(args.HourlyRate),
			db.MachinePlan.CPU.SetIfPresent(args.Cpu),
			db.MachinePlan.Memory.SetIfPresent(args.Memory),
		).
		Exec(m.ctx)
}

func (m *MachinePlanDao) DeleteMachinePlan(args types.DeleteMachinePlanArgs) error {
	_, err := m.client.MachinePlan.
		FindUnique(db.MachinePlan.ID.Equals(args.Id)).
		Delete().
		Exec(m.ctx)

	return err
}

func NewMachinePlanDao(connection *lib.DatabaseConnection) IMachinePlanDao {
	return &MachinePlanDao{
		client: connection.Client,
		ctx:    context.Background(),
	}
}
