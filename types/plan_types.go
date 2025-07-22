package types

import (
	"pkg.formatio/lib"
)

type ListMachinePlansArgs struct {
	lib.BaseListFilterArgs
}

type CreateMachinePlanArgs struct {
	Name        string  `json:"name"`
	Currency    string  `json:"currency"`
	MonthlyRate int     `json:"monthlyRate"`
	HourlyRate  float32 `swag-validate:"optional" json:"-" swaggerignore:"true"`
	Cpu         string  `json:"cpu"`
	Memory      string  `json:"memory"`
}

type GetMachinePlanArgs struct {
	Id string
}

type UpdateMachinePlanArgs struct {
	Id          string  `swaggerignore:"true"`
	Name        *string `json:"name" swag-validate:"optional"`
	Currency    *string `json:"currency" swag-validate:"optional"`
	MonthlyRate *int    `json:"monthlyRate" swag-validate:"optional"`
	// HourlyRate  *int32  `json:"hourlyRate" swag-validate:"optional"`
	Cpu    *string `json:"cpu" swag-validate:"optional"`
	Memory *string `json:"memory" swag-validate:"optional"`
}

type DeleteMachinePlanArgs struct {
	Id string
}
