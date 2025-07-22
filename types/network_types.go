package types

import "pkg.formatio/lib"

type ListNetworksArgs struct {
	lib.BaseListFilterArgs

	MachineId *string `json:"machineId,omitempty" swag-validate:"optional"`
	OwnerId   *string `json:"ownerId,omitempty" swag-validate:"optional" swaggerignore:"true"`
}

type CreateNetworkArgs struct {
	MachineId       string `json:"machineId"`
	Protocol        string `json:"protocol"`
	ListeningPort   int    `json:"listeningPort"`
	DestinationPort int    `json:"destinationPort"`
	HostName        string `swaggerignore:"true"`
	ServiceId       string `swaggerignore:"true"`
	IngressId       string `swaggerignore:"true"`
}

type GetNetworkArgs struct {
	Id        *string
	ServiceId *string
	IngressId *string
}

type UpdateNetworkArgs struct {
	Id              string
	MachineId       *string
	HostName        *string
	Protocol        *string
	ListeningPort   *int
	DestinationPort *int
	ServiceId       *string
	IngressId       *string
}

type DeleteNetworkArgs struct {
	Id string
}
