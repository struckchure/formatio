package types

import (
	storm "github.com/Overal-X/formatio.storm"
	"pkg.formatio/lib"
)

type Repository struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
	Private  bool   `json:"private"`
	HTMLURL  string `json:"url"`
}

type ListRepositoriesArgs struct {
	UserId     string `json:"-" swaggerignore:"true"`
	PageNumber *int   `json:"pageNumber,omitempty" swag-validate:"optional" default:"1"`
	PageSize   *int   `json:"pageSize,omitempty" swag-validate:"optional" default:"20"`
}

type GithubUser struct {
	Id       int    `json:"id"`
	Username string `json:"login"`
	Type     string `json:"type"`
	FullName string `json:"name"`
	Email    string `json:"email"`
}

type ListGithubAccountConnectionsArgs struct {
	lib.BaseListFilterArgs

	UserId         *string `swaggerignore:"true"`
	InstallationId *int
}

type CreateGithubAccountConnectionArgs struct {
	UserId               string
	GithubId             string
	GithubInstallationId *int
	GithubUsername       string
	GithubEmail          string
}

type GetGithubAccountConnectionsArgs struct {
	Id     *string
	UserId *string
}

type UpdateGithubAccountConnectionArgs struct {
	Id                   string
	GithubId             *string
	GithubInstallationId *int
	GithubUsername       *string
	GithubEmail          *string
}

type DeleteGithubAccountConnectionsArgs struct {
	Id string
}

type Action struct {
	Name string `yaml:"name"`

	On struct {
		Push struct {
			Branches []string `yaml:"branches"`
		} `yaml:"push"`
	} `yaml:"on"`

	Jobs []storm.Job `yaml:"jobs"`
}

type ExecuteActionArgs struct {
	Action         Action
	DeploymentId   string
	DeploymentName string
}

type GetFileFromRepoArgs struct {
	RepoURL  string
	FilePath string
}
