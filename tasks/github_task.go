package tasks

import (
	"encoding/json"
	"log"

	"pkg.formatio/lib"
	"pkg.formatio/services"
	"pkg.formatio/types"
)

type GithubTasks struct {
	rmq           lib.RabbitMQ
	githubService services.IGithubService
}

func (t *GithubTasks) DeployRepoTask() {
	t.rmq.SubscribeWithWorkers(2, lib.SubscribeArgs{
		Queue: types.DEPLOYMENT_DEPLOY_REPO_QUEUE,
		Callback: func(body string) error {
			var payload types.DeployRepoArgs
			err := json.Unmarshal([]byte(body), &payload)
			if err != nil {
				return err
			}

			err = t.githubService.DeployRepoHandler(payload)
			if err != nil {
				log.Println(err)

				return err
			}

			return nil
		},
	})
}

func NewGithubTasks(
	rmq lib.RabbitMQ,
	githubService services.IGithubService,
) GithubTasks {
	return GithubTasks{
		rmq:           rmq,
		githubService: githubService,
	}
}
