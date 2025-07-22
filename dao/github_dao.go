package dao

import (
	"context"

	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/types"
)

type IGithubAccountConnectionDao interface {
	ListConnections(types.ListGithubAccountConnectionsArgs) ([]db.GithubAccountConnectionModel, error)
	CreateConnection(types.CreateGithubAccountConnectionArgs) (*db.GithubAccountConnectionModel, error)
	GetConnection(types.GetGithubAccountConnectionsArgs) (*db.GithubAccountConnectionModel, error)
	UpdateConnection(types.UpdateGithubAccountConnectionArgs) (*db.GithubAccountConnectionModel, error)
	DeleteConnection(types.DeleteGithubAccountConnectionsArgs) error
}

type GithubAccountConnectionDao struct {
	client *db.PrismaClient
	ctx    context.Context

	userDao IUserDao
}

func (g *GithubAccountConnectionDao) ListConnections(args types.ListGithubAccountConnectionsArgs) (
	[]db.GithubAccountConnectionModel,
	error,
) {
	return g.client.GithubAccountConnection.
		FindMany(db.GithubAccountConnection.User.Where(db.User.ID.EqualsIfPresent(args.UserId))).
		OrderBy(db.GithubAccountConnection.CreatedAt.Order(db.DESC)).
		Exec(g.ctx)
}

func (d *GithubAccountConnectionDao) CreateConnection(args types.CreateGithubAccountConnectionArgs) (
	*db.GithubAccountConnectionModel,
	error,
) {
	return d.client.GithubAccountConnection.
		CreateOne(
			db.GithubAccountConnection.User.Link(db.User.ID.Equals(args.UserId)),
			db.GithubAccountConnection.GithubID.Set(args.GithubId),
			db.GithubAccountConnection.GithubInstallationID.SetIfPresent(args.GithubInstallationId),
			db.GithubAccountConnection.GithubEmail.Set(args.GithubEmail),
			db.GithubAccountConnection.GithubUsername.Set(args.GithubUsername),
		).
		Exec(d.ctx)
}

func (d *GithubAccountConnectionDao) GetConnection(args types.GetGithubAccountConnectionsArgs) (
	*db.GithubAccountConnectionModel,
	error,
) {
	return d.client.GithubAccountConnection.
		FindFirst(
			db.GithubAccountConnection.ID.EqualsIfPresent(args.Id),
			db.GithubAccountConnection.UserID.EqualsIfPresent(args.Id)).
		Exec(d.ctx)
}

func (d *GithubAccountConnectionDao) UpdateConnection(args types.UpdateGithubAccountConnectionArgs) (
	*db.GithubAccountConnectionModel,
	error,
) {
	return d.client.GithubAccountConnection.
		FindUnique(db.GithubAccountConnection.ID.Equals(args.Id)).
		Update(
			db.GithubAccountConnection.GithubID.SetIfPresent(args.GithubId),
			db.GithubAccountConnection.GithubInstallationID.SetIfPresent(args.GithubInstallationId),
			db.GithubAccountConnection.GithubEmail.SetIfPresent(args.GithubEmail),
			db.GithubAccountConnection.GithubUsername.SetIfPresent(args.GithubUsername),
		).
		Exec(d.ctx)
}

func (d *GithubAccountConnectionDao) DeleteConnection(args types.DeleteGithubAccountConnectionsArgs) error {
	_, err := d.client.GithubAccountConnection.
		FindUnique(db.GithubAccountConnection.ID.Equals(args.Id)).
		Delete().
		Exec(d.ctx)

	return err
}

func NewGithubAccountConnectionDao(
	connection *lib.DatabaseConnection,
	userDAO IUserDao,
) IGithubAccountConnectionDao {
	return &GithubAccountConnectionDao{
		client: connection.Client,
		ctx:    context.Background(),

		userDao: userDAO,
	}
}
