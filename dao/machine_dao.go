package dao

import (
	"context"

	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/types"
)

type IMachineDao interface {
	ListMachines(types.ListMachineArgs) ([]db.MachineModel, error)
	CreateMachine(types.CreateMachineArgs) (*db.MachineModel, error)
	GetMachine(types.GetMachineArgs) (*db.MachineModel, error)
	UpdateMachine(types.UpdateMachineArgs) (*db.MachineModel, error)
	DeleteMachine(types.DeleteMachineArgs) error
}

type MachineDao struct {
	client *db.PrismaClient
	ctx    context.Context
}

func (m *MachineDao) ListMachines(args types.ListMachineArgs) ([]db.MachineModel, error) {
	args.Skip = lib.UseDefaultValueIf(0, args.Skip, 0)
	args.Take = lib.UseDefaultValueIf(0, args.Take, 10)

	return m.client.Machine.
		FindMany(db.Machine.Owner.Where(db.User.ID.EqualsIfPresent(args.UserId))).
		OrderBy(db.Machine.CreatedAt.Order(db.DESC)).
		With(db.Machine.Plan.Fetch()).
		Skip(args.Skip).
		Take(args.Take).
		Exec(m.ctx)
}

func (m *MachineDao) CreateMachine(args types.CreateMachineArgs) (*db.MachineModel, error) {
	return m.client.Machine.
		CreateOne(
			db.Machine.Owner.Link(db.User.ID.Equals(args.OwnerId)),
			db.Machine.Plan.Link(db.MachinePlan.ID.Equals(args.PlanId)),
			db.Machine.ContainerID.Set(args.ContainerId),
			db.Machine.MachineName.Set(args.MachineName),
			db.Machine.MachineImage.Set(args.MachineImage),
			db.Machine.MachineStatus.Set(args.MachineStatus),
		).
		Exec(m.ctx)
}

func (m *MachineDao) GetMachine(args types.GetMachineArgs) (*db.MachineModel, error) {
	return m.client.Machine.
		FindFirst(
			db.Machine.ID.EqualsIfPresent(args.Id),
			db.Machine.ContainerID.EqualsIfPresent(args.ContainerId),
		).
		With(db.Machine.Plan.Fetch()).
		Exec(m.ctx)
}

func (m *MachineDao) UpdateMachine(args types.UpdateMachineArgs) (*db.MachineModel, error) {
	return m.client.Machine.
		FindUnique(db.Machine.ID.Equals(args.ID)).
		Update(
			db.Machine.ContainerID.SetIfPresent(args.ContainerId),
			db.Machine.MachineName.SetIfPresent(args.MachineName),
			db.Machine.MachineImage.SetIfPresent(args.MachineImage),
			db.Machine.MachineStatus.SetIfPresent(args.MachineStatus),
		).
		Exec(m.ctx)
}

func (m *MachineDao) DeleteMachine(args types.DeleteMachineArgs) error {
	_, err := m.client.Machine.
		FindUnique(db.Machine.ID.Equals(args.Id)).
		Delete().
		Exec(m.ctx)

	return err
}

func NewMachineDao(connection *lib.DatabaseConnection) IMachineDao {
	return &MachineDao{
		client: connection.Client,
		ctx:    context.Background(),
	}
}
