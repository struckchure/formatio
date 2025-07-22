package types

import (
	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
)

type ListMachineArgs struct {
	lib.BaseListFilterArgs

	UserId *string `swag-validate:"optional"`
}

type CreateMachineArgs struct {
	OwnerId       string           `swaggerignore:"true"`
	PlanId        string           `json:"planId"`
	MachineName   string           `json:"machineName"`
	MachineImage  string           `json:"machineImage"`
	ContainerId   string           `swaggerignore:"true"`
	MachineStatus db.MachineStatus `swaggerignore:"true"`
}

type GetMachineArgs struct {
	Id          *string
	ContainerId *string
}

type UpdateMachineArgs struct {
	ID            string             `swaggerignore:"true"`
	OwnerId       *string            `json:"ownerId" swag-validate:"optional"`
	PlanId        *string            `json:"planId" swag-validate:"optional"`
	ContainerId   *string            `swaggerignore:"true"`
	MachineName   *string            `json:"machineName" swag-validate:"optional"`
	MachineImage  *string            `json:"machineImage" swag-validate:"optional"`
	MachineStatus *db.MachineStatus  `swaggerignore:"true"`
	Ports         *[]lib.NetworkPort `swaggerignore:"true"` // TODO: fix type
}

type DeleteMachineArgs struct {
	Id string
}

type CreateMachineEventHandlerArgs struct {
	Id string
}

type DeleteMachineEventHandlerArgs struct {
	Id          string
	ContainerId *string
}
