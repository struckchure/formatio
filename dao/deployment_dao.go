package dao

import (
	"context"

	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/types"
)

type IDeploymentDao interface {
	ListDeployments(types.ListDeploymentArgs) ([]db.DeploymentModel, error)
	CreateDeployment(types.CreateDeploymentArgs) (*db.DeploymentModel, error)
	GetDeployment(types.GetDeploymentArgs) (*db.DeploymentModel, error)
	UpdateDeployment(types.UpdateDeploymentArgs) (*db.DeploymentModel, error)
	DeleteDeployment(types.DeleteDeploymentArgs) error
}

type DeploymentDao struct {
	client *db.PrismaClient
	ctx    context.Context

	machineDao        IMachineDao
	repoConnectionDao IRepoConnectionDao
}

func (d *DeploymentDao) ListDeployments(args types.ListDeploymentArgs) (deployments []db.DeploymentModel, err error) {
	return d.client.Deployment.
		FindMany(
			db.Deployment.MachineID.EqualsIfPresent(args.MachineId),
			db.Deployment.RepoConnectionID.EqualsIfPresent(args.RepoConnectionId),
		).
		OrderBy(db.Deployment.CreatedAt.Order(db.DESC)).
		Exec(d.ctx)
}

func (d *DeploymentDao) CreateDeployment(args types.CreateDeploymentArgs) (*db.DeploymentModel, error) {
	return d.client.Deployment.
		CreateOne(
			db.Deployment.Machine.Link(db.Machine.ID.Equals(args.MachineId)),
			db.Deployment.RepoConnection.Link(db.RepoConnection.ID.Equals(args.RepoConnectionId)),
			db.Deployment.CommitHash.Set(args.CommitHash),
			db.Deployment.CommitMessage.Set(args.CommitMessage),
			db.Deployment.Status.Set(args.Status),
			db.Deployment.Actor.Set(args.Actor),
		).
		Exec(d.ctx)
}

func (d *DeploymentDao) GetDeployment(args types.GetDeploymentArgs) (*db.DeploymentModel, error) {
	return d.client.Deployment.
		FindFirst(db.Deployment.ID.Equals(args.Id)).
		Exec(d.ctx)
}

func (d *DeploymentDao) UpdateDeployment(args types.UpdateDeploymentArgs) (*db.DeploymentModel, error) {
	return d.client.Deployment.
		FindUnique(db.Deployment.ID.Equals(args.Id)).
		Update(
			db.Deployment.CommitHash.SetIfPresent(args.CommitHash),
			db.Deployment.CommitMessage.SetIfPresent(args.CommitMessage),
			db.Deployment.Actor.SetIfPresent(args.Actor),
			db.Deployment.Status.SetIfPresent(args.Status),
		).
		Exec(d.ctx)
}

func (d *DeploymentDao) DeleteDeployment(args types.DeleteDeploymentArgs) error {
	_, err := d.client.Deployment.
		FindUnique(db.Deployment.ID.Equals(args.Id)).
		Exec(d.ctx)

	return err
}

func NewDeploymentDao(
	connection *lib.DatabaseConnection,
	machineDao IMachineDao,
	repoConnectionDao IRepoConnectionDao,
) IDeploymentDao {
	return &DeploymentDao{
		client: connection.Client,
		ctx:    context.Background(),

		machineDao:        machineDao,
		repoConnectionDao: repoConnectionDao,
	}
}
