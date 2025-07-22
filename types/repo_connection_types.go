package types

import "pkg.formatio/lib"

type ListRepoConnectionArgs struct {
	*lib.BaseListFilterArgs

	RepoId    *string `json:"repoId,omitempty" swag-validate:"optional" swaggerignore:"true"`
	OwnerId   *string `json:"ownerId,omitempty" swaggerignore:"true"`
	MachineId *string `json:"machineId,omitempty" swag-validate:"optional"`
}

type CreateRepoConnectionArgs struct {
	MachineId string `json:"machineId"`
	UserId    string `json:"-" swaggerignore:"true"`
	RepoId    string `json:"repoId"`
	RepoName  string `json:"repoName"`
}

type GetRepoConnectionArgs struct {
	Id string `json:"id"`
}

type UpdateRepoConnectionArgs struct {
	Id        string  `json:"id" swaggerignore:"true"`
	MachineId *string `json:"machineId" swag-validate:"optional"`
	RepoId    *string `json:"repoId" swag-validate:"optional"`
	RepoName  *string `json:"repoName" swag-validate:"optional"`
}

type DeleteRepoConnectionArgs struct {
	Id string `json:"id"`
}
