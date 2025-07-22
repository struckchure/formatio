package types

import (
	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
)

const (
	DEPLOYMENT_DEPLOY_REPO_QUEUE = "deployment-deploy-repo-queue"
	DEPLOYMENT_LOG_EVENT_QUEUE   = "deployment-log-queue"

	DEPLOYMENT_NOTIFICATION_EVENT = "deployment-notification-event"
	DEPLOYMENT_LOG_STREAM_EVENT   = "deployment-log-stream-event"
)

type ListDeploymentArgs struct {
	lib.BaseListFilterArgs

	MachineId        *string `query:"machineId"`
	RepoConnectionId *string `query:"repoConnectionId" swag-validate:"optional"`
}

type CreateDeploymentArgs struct {
	MachineId        string
	RepoConnectionId string
	CommitHash       string
	CommitMessage    string
	Status           db.DeploymentStatus
	Actor            string
}

type GetDeploymentArgs struct {
	Id string
}

type UpdateDeploymentArgs struct {
	Id               string
	RepoConnectionId string
	CommitHash       *string
	CommitMessage    *string
	Status           *db.DeploymentStatus
	Actor            *string
}

type DeleteDeploymentArgs struct {
	Id string
}

type ListDeploymentLogArgs struct {
	lib.BaseListFilterArgs

	DeploymentId *string
	JobId        *string
	Message      *string
}

type CreateDeploymentLogArgs struct {
	DeploymentId string
	Message      string
	JobId        string
}

type GetDeploymentLogArgs struct {
	Id string
}

type UpdateDeploymentLogArgs struct {
	Id      string
	Message *string
	JobId   *string
}

type DeleteDeploymentLogArgs struct {
	Id string
}

// TODO: should be same as `CreateDeploymentLogArgs`
type DeploymentNotificationPayload struct {
	DeploymentId  string
	CommitHash    string
	CommitMessage string
}
