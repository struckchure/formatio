package tasks

import (
	"encoding/json"
	"log"

	"github.com/overal-x/rodelar-go-sdk"
	"pkg.formatio/lib"
	"pkg.formatio/types"
)

type DeploymentTasks struct {
	rmq           lib.RabbitMQ
	rodelarClient rodelar.IRodelarClient
}

func (t *DeploymentTasks) DeploymentNotificationTask() {
	t.rmq.Subscribe(lib.SubscribeArgs{
		Queue: types.DEPLOYMENT_NOTIFICATION_EVENT,
		Callback: func(body string) error {
			var dto map[string]any
			err := json.Unmarshal([]byte(body), &dto)
			if err != nil {
				log.Println(err)

				return err
			}

			// err = t.rodelarClient.Publish(rodelar.PublishArgs{
			// 	Event:   fmt.Sprintf("%s/%s", types.DEPLOYMENT_NOTIFICATION_EVENT, dto["machineId"]),
			// 	Message: dto,
			// })
			// if err != nil {
			// 	log.Println(err)

			// 	return err
			// }

			return nil
		},
	})
}

func (t *DeploymentTasks) DeploymentLogsTask() {
	t.rmq.Subscribe(lib.SubscribeArgs{
		Queue: types.DEPLOYMENT_LOG_EVENT_QUEUE,
		Callback: func(body string) error {
			// var payload db.DeploymentLogModel
			// err := json.Unmarshal([]byte(body), &payload)
			// if err != nil {
			// 	log.Println(err)

			// 	return err
			// }

			// err = t.rodelarClient.Publish(rodelar.PublishArgs{
			// 	Event:   fmt.Sprintf("%s/%s", types.DEPLOYMENT_LOG_STREAM_EVENT, payload.DeploymentID),
			// 	Message: body,
			// })
			// if err != nil {
			// 	log.Println(err)

			// 	return err
			// }

			return nil
		},
	})
}

func NewDeploymentTasks(
	rmq lib.RabbitMQ,
	rodelarClient rodelar.IRodelarClient,
) DeploymentTasks {
	return DeploymentTasks{
		rmq:           rmq,
		rodelarClient: rodelarClient,
	}
}
