package tasks

import (
	"context"

	"go.uber.org/fx"
	"pkg.formatio/lib"
)

func NewRootTasks(
	lc fx.Lifecycle,

	machineTasks MachineTasks,
	deploymentTasks DeploymentTasks,
	githubTasks GithubTasks,
	billingTasks BillingTasks,

	k8sInformer lib.IInformer,
) {
	k8sInformerChannel := make(chan struct{})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go machineTasks.CreateMachineTask()
			go machineTasks.UpdateMachineTask()
			go machineTasks.DeleteMachineTask()
			go machineTasks.RedeployOnMachineUpdateTask()

			// go deploymentTasks.DeploymentNotificationTask()
			// go deploymentTasks.DeploymentLogsTask()

			go githubTasks.DeployRepoTask()

			go billingTasks.ScheduleMachineInvoicesTask()
			go billingTasks.ProcessInvoiceTask()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			k8sInformer.Stop(k8sInformerChannel)

			return nil
		},
	})
}
