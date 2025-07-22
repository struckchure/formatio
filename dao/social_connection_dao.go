package dao

import (
	"context"

	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/types"
)

type ISocialConnectionDao interface {
	ListConnections(types.ListConnectionsArgs) ([]db.SocialConnectionModel, error)
	CreateConnection(types.CreateConnectionArgs) (*db.SocialConnectionModel, error)
	GetConnection(types.GetConnectionArgs) (*db.SocialConnectionModel, error)
	UpdateConnection(types.UpdateConnectionArgs) (*db.SocialConnectionModel, error)
	DeleteConnection(types.DeleteConnectionArgs) error
}

type SocialConnectionDao struct {
	client *db.PrismaClient
	ctx    context.Context

	userDao IUserDao
}

func (s *SocialConnectionDao) ListConnections(args types.ListConnectionsArgs) ([]db.SocialConnectionModel, error) {
	args.Skip = lib.UseDefaultValueIf(0, args.Skip, 0)
	args.Take = lib.UseDefaultValueIf(0, args.Take, 10)

	return s.client.SocialConnection.
		FindMany().
		Skip(args.Skip).
		Take(args.Take).
		OrderBy(db.SocialConnection.CreatedAt.Order(db.DESC)).
		Exec(s.ctx)
}

func (s *SocialConnectionDao) CreateConnection(args types.CreateConnectionArgs) (*db.SocialConnectionModel, error) {
	return s.client.SocialConnection.
		CreateOne(
			db.SocialConnection.User.Link(
				db.User.ID.Equals(args.UserId),
			),
			db.SocialConnection.ConnectionID.Set(args.ConnectionId),
			db.SocialConnection.ConnectionType.Set(args.ConnectionType),
		).Exec(s.ctx)
}

func (*SocialConnectionDao) GetConnection(types.GetConnectionArgs) (*db.SocialConnectionModel, error) {
	panic("unimplemented")
}

func (*SocialConnectionDao) UpdateConnection(types.UpdateConnectionArgs) (*db.SocialConnectionModel, error) {
	panic("unimplemented")
}

func (*SocialConnectionDao) DeleteConnection(types.DeleteConnectionArgs) error {
	panic("unimplemented")
}

func NewSocialConnectionDao(
	connection *lib.DatabaseConnection,
	userDAO IUserDao,
) ISocialConnectionDao {
	return &SocialConnectionDao{
		client: connection.Client,
		ctx:    context.Background(),

		userDao: userDAO,
	}
}
