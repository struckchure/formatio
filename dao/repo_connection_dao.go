package dao

import (
	"context"

	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/types"
)

type IRepoConnectionDao interface {
	ListRepoConnections(types.ListRepoConnectionArgs) ([]db.RepoConnectionModel, error)
	CreateRepoConnection(types.CreateRepoConnectionArgs) (*db.RepoConnectionModel, error)
	GetRepoConnection(types.GetRepoConnectionArgs) (*db.RepoConnectionModel, error)
	UpdateRepoConnection(types.UpdateRepoConnectionArgs) (*db.RepoConnectionModel, error)
	DeleteRepoConnection(types.DeleteRepoConnectionArgs) error
}

type RepoConnectionDao struct {
	client *db.PrismaClient
	ctx    context.Context

	machineDao IMachineDao
}

func (p *RepoConnectionDao) ListRepoConnections(args types.ListRepoConnectionArgs) ([]db.RepoConnectionModel, error) {
	return p.client.RepoConnection.
		FindMany(
			db.RepoConnection.MachineID.EqualsIfPresent(args.MachineId),
			db.RepoConnection.Machine.Where(db.Machine.OwnerID.EqualsIfPresent(args.OwnerId)),
			db.RepoConnection.RepoID.EqualsIfPresent(args.RepoId),
		).
		With(db.RepoConnection.Machine.Fetch()).
		OrderBy(db.RepoConnection.CreatedAt.Order(db.DESC)).
		Exec(p.ctx)
}

func (p *RepoConnectionDao) CreateRepoConnection(args types.CreateRepoConnectionArgs) (*db.RepoConnectionModel, error) {
	return p.client.RepoConnection.
		CreateOne(
			db.RepoConnection.Machine.Link(db.Machine.ID.Equals(args.MachineId)),
			db.RepoConnection.RepoName.Set(args.RepoName),
			db.RepoConnection.RepoID.Set(args.RepoId),
		).
		Exec(p.ctx)
}

func (p *RepoConnectionDao) GetRepoConnection(args types.GetRepoConnectionArgs) (*db.RepoConnectionModel, error) {
	return p.client.RepoConnection.
		FindFirst(db.RepoConnection.ID.Equals(args.Id)).
		Exec(p.ctx)
}

func (p *RepoConnectionDao) UpdateRepoConnection(args types.UpdateRepoConnectionArgs) (*db.RepoConnectionModel, error) {
	return p.client.RepoConnection.
		FindUnique(db.RepoConnection.ID.Equals(args.Id)).
		Update(
			db.RepoConnection.RepoID.SetIfPresent(args.RepoId),
			db.RepoConnection.RepoName.SetIfPresent(args.RepoName),
		).
		Exec(p.ctx)
}

func (p *RepoConnectionDao) DeleteRepoConnection(args types.DeleteRepoConnectionArgs) error {
	_, err := p.client.RepoConnection.
		FindUnique(db.RepoConnection.ID.Equals(args.Id)).
		Delete().
		Exec(p.ctx)

	return err
}

func NewRepoConnectionDao(connection *lib.DatabaseConnection, machineDAO IMachineDao) IRepoConnectionDao {
	return &RepoConnectionDao{
		client: connection.Client,
		ctx:    context.Background(),

		machineDao: machineDAO,
	}
}
