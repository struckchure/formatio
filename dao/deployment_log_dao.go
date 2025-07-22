package dao

import (
	"context"

	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/types"
)

type IDeploymentLogDao interface {
	ListLogs(types.ListDeploymentLogArgs) ([]db.DeploymentLogModel, error)
	CreateLog(types.CreateDeploymentLogArgs) (*db.DeploymentLogModel, error)
	Getlog(types.GetDeploymentLogArgs) (*db.DeploymentLogModel, error)
	UpdateLog(types.UpdateDeploymentLogArgs) (*db.DeploymentLogModel, error)
	DeleteLog(types.DeleteDeploymentLogArgs) error
}

type DeploymentLogDao struct {
	client *db.PrismaClient
	ctx    context.Context

	deploymentDao IDeploymentDao
}

func (d *DeploymentLogDao) ListLogs(args types.ListDeploymentLogArgs) ([]db.DeploymentLogModel, error) {
	args.Skip = lib.UseDefaultValueIf(0, args.Skip, 0)
	args.Take = lib.UseDefaultValueIf(0, args.Take, 10)

	return d.client.DeploymentLog.
		FindMany(
			db.DeploymentLog.DeploymentID.EqualsIfPresent(args.DeploymentId),
			db.DeploymentLog.JobID.EqualsIfPresent(args.JobId),
			db.DeploymentLog.Message.EqualsIfPresent(args.Message),
		).
		OrderBy(db.DeploymentLog.CreatedAt.Order(db.ASC)).
		// Skip(args.Skip).
		// Take(args.Take).
		Exec(d.ctx)
}

func (d *DeploymentLogDao) CreateLog(args types.CreateDeploymentLogArgs) (*db.DeploymentLogModel, error) {
	return d.client.DeploymentLog.
		CreateOne(
			db.DeploymentLog.Deployment.Link(db.Deployment.ID.Equals(args.DeploymentId)),
			db.DeploymentLog.Message.Set(args.Message),
			db.DeploymentLog.JobID.Set(args.JobId),
		).
		Exec(d.ctx)
}

func (d *DeploymentLogDao) Getlog(args types.GetDeploymentLogArgs) (*db.DeploymentLogModel, error) {
	return d.client.DeploymentLog.
		FindFirst(db.DeploymentLog.ID.Equals(args.Id)).
		Exec(d.ctx)
}

func (d *DeploymentLogDao) UpdateLog(args types.UpdateDeploymentLogArgs) (*db.DeploymentLogModel, error) {
	return d.client.DeploymentLog.
		FindUnique(db.DeploymentLog.ID.Equals(args.Id)).
		Update(
			db.DeploymentLog.Message.SetIfPresent(args.Message),
			db.DeploymentLog.JobID.SetIfPresent(args.JobId),
		).
		Exec(d.ctx)
}

func (d *DeploymentLogDao) DeleteLog(args types.DeleteDeploymentLogArgs) error {
	_, err := d.client.DeploymentLog.
		FindUnique(db.DeploymentLog.ID.Equals(args.Id)).
		Delete().
		Exec(d.ctx)

	return err
}

func NewDeploymentLogDao(
	connection *lib.DatabaseConnection,
	deploymentDao IDeploymentDao,
) IDeploymentLogDao {
	return &DeploymentLogDao{
		client: connection.Client,
		ctx:    context.Background(),

		deploymentDao: deploymentDao,
	}
}
