package dao

import (
	"context"

	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/types"
)

type INetworkDao interface {
	ListNetworks(types.ListNetworksArgs) ([]db.NetworkModel, error)
	CreateNetwork(types.CreateNetworkArgs) (*db.NetworkModel, error)
	GetNetwork(types.GetNetworkArgs) (*db.NetworkModel, error)
	UpdateNetwork(types.UpdateNetworkArgs) (*db.NetworkModel, error)
	DeleteNetwork(types.DeleteNetworkArgs) error
}

type NetworkDao struct {
	client *db.PrismaClient
	ctx    context.Context

	machineDao IMachineDao
}

func (n *NetworkDao) ListNetworks(args types.ListNetworksArgs) ([]db.NetworkModel, error) {
	args.Skip = lib.UseDefaultValueIf(0, args.Skip, 0)
	args.Take = lib.UseDefaultValueIf(0, args.Take, 10)

	return n.client.Network.
		FindMany(
			db.Network.Machine.Where(db.Machine.OwnerID.EqualsIfPresent(args.OwnerId)),
			db.Network.MachineID.EqualsIfPresent(args.MachineId),
		).
		With(db.Network.Machine.Fetch()).
		Skip(args.Skip).
		Take(args.Take).
		OrderBy(db.Network.CreatedAt.Order(db.DESC)).
		Exec(n.ctx)
}

func (n *NetworkDao) CreateNetwork(args types.CreateNetworkArgs) (*db.NetworkModel, error) {
	return n.client.Network.
		CreateOne(
			db.Network.Machine.Link(db.Machine.ID.Equals(args.MachineId)),
			db.Network.ListeningPort.Set(args.ListeningPort),
			db.Network.DestinationPort.Set(args.DestinationPort),
			db.Network.HostName.Set(args.HostName),
			db.Network.Protocol.Set(args.Protocol),
			db.Network.ServiceID.Set(args.ServiceId),
			db.Network.IngressID.Set(args.IngressId),
		).
		Exec(n.ctx)
}

func (n *NetworkDao) GetNetwork(args types.GetNetworkArgs) (*db.NetworkModel, error) {
	return n.client.Network.
		FindFirst(
			db.Network.ID.EqualsIfPresent(args.Id),
			db.Network.ServiceID.EqualsIfPresent(args.ServiceId),
			db.Network.IngressID.EqualsIfPresent(args.IngressId),
		).
		Exec(n.ctx)
}

func (n *NetworkDao) UpdateNetwork(args types.UpdateNetworkArgs) (*db.NetworkModel, error) {
	return n.client.Network.
		FindUnique(db.Network.ID.Equals(args.Id)).
		Update(
			db.Network.HostName.SetIfPresent(args.HostName),
			db.Network.Protocol.SetIfPresent(args.Protocol),
			db.Network.ListeningPort.SetIfPresent(args.ListeningPort),
			db.Network.DestinationPort.SetIfPresent(args.DestinationPort),
			db.Network.ServiceID.SetIfPresent(args.ServiceId),
			db.Network.IngressID.SetIfPresent(args.IngressId),
		).
		Exec(n.ctx)
}

func (n *NetworkDao) DeleteNetwork(args types.DeleteNetworkArgs) error {
	_, err := n.client.Network.
		FindUnique(db.Network.ID.Equals(args.Id)).
		Delete().
		Exec(n.ctx)

	return err
}

func NewNetworkDao(
	connection *lib.DatabaseConnection,
	machineDAO IMachineDao,
) INetworkDao {
	return &NetworkDao{
		client: connection.Client,
		ctx:    context.Background(),

		machineDao: machineDAO,
	}
}
