package tasks

import (
	"encoding/json"
	"errors"
	"log"

	v1 "k8s.io/api/core/v1"
	"pkg.formatio/dao"
	"pkg.formatio/lib"
	"pkg.formatio/services"
	"pkg.formatio/types"
)

type MachineTasks struct {
	rmq                 lib.RabbitMQ
	machineService      services.IMachineService
	k8sInformer         lib.IInformer
	k8sInformerChannel  chan struct{}
	k8sContainerManager lib.IContainerManager
	deploymentService   *services.DeploymentService
	machineDao          dao.IMachineDao
	repoConnectionDao   dao.IRepoConnectionDao
}

func (r *MachineTasks) CreateMachineTask() {
	r.rmq.Subscribe(lib.SubscribeArgs{
		Queue: services.CREATE_MACHINE_QUEUE,
		Callback: func(body string) error {
			var payload types.CreateMachineEventHandlerArgs
			if err := json.Unmarshal([]byte(body), &payload); err != nil {
				// TODO: publish error back to exchange, then websocket
				log.Println(err)

				return err
			}

			err := r.machineService.CreateMachineEventHandler(payload)
			if err != nil {
				log.Println(err)

				return err
			}

			return nil
		},
	})
}

func (r *MachineTasks) UpdateMachineTask() {
	r.rmq.Subscribe(lib.SubscribeArgs{
		Queue: services.UPDATE_MACHINE_QUEUE,
		Callback: func(body string) error {
			var payload types.UpdateMachineArgs
			if err := json.Unmarshal([]byte(body), &payload); err != nil {
				// TODO: publish error back to exchange, then websocket
				log.Println(err)

				return err
			}

			err := r.machineService.UpdateMachineEventHandler(payload)
			if err != nil {
				log.Println(err)

				return err
			}

			return nil
		},
	})
}

func (r *MachineTasks) DeleteMachineTask() {
	r.rmq.Subscribe(lib.SubscribeArgs{
		Queue: services.DELETE_MACHINE_QUEUE,
		Callback: func(body string) error {
			var payload struct {
				Id        string
				MachineId string
			}
			if err := json.Unmarshal([]byte(body), &payload); err != nil {
				// TODO: publish error back to exchange, then websocket
				log.Println(err)

				return err
			}

			err := r.machineService.DeleteMachineEventHandler(types.DeleteMachineEventHandlerArgs{
				Id:          payload.Id,
				ContainerId: &payload.MachineId,
			})
			if err != nil {
				log.Println(err)

				return err
			}

			return nil
		},
	})
}

func (r *MachineTasks) RedeployOnMachineUpdateTask() {
	r.k8sInformer.Pods(v1.NamespaceDefault, func(iet lib.InformerEventType, p ...*v1.Pod) error {
		if iet == lib.InformerEventCreated {
			pod := p[0]
			if pod.Status.Phase != v1.PodRunning {
				return nil
			}

			deployment, err := r.k8sContainerManager.GetDeploymentByPodName(pod.Name)
			if err != nil {
				return err
			}

			machine, err := r.machineDao.GetMachine(types.GetMachineArgs{ContainerId: &deployment.Name})
			if err != nil {
				return err
			}

			repoConnections, err := r.repoConnectionDao.ListRepoConnections(types.ListRepoConnectionArgs{
				MachineId: &machine.ID,
			})
			if err != nil {
				return err
			}

			if len(repoConnections) == 0 {
				return errors.New("no repo connection found")
			}

			deployments, err := r.deploymentService.ListDeployments(types.ListDeploymentArgs{
				MachineId: &machine.ID,
			})
			if err != nil {
				return err
			}

			if len(deployments) == 0 {
				return nil
			}

			err = r.deploymentService.DeployRepo(services.DeployRepoArgs{
				UserId:       machine.OwnerID,
				ConnectionId: repoConnections[0].ID,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})

	r.k8sInformer.Start(r.k8sInformerChannel)
}

func NewMachineTasks(
	rmq lib.RabbitMQ,
	machineService services.IMachineService,
	k8sInformer lib.IInformer,
	k8sContainerManager lib.IContainerManager,
	deploymentService *services.DeploymentService,
	machineDao dao.IMachineDao,
	repoConnectionDao dao.IRepoConnectionDao,
) MachineTasks {
	return MachineTasks{
		rmq:                 rmq,
		machineService:      machineService,
		k8sInformer:         k8sInformer,
		k8sInformerChannel:  make(chan struct{}),
		k8sContainerManager: k8sContainerManager,
		deploymentService:   deploymentService,
		machineDao:          machineDao,
		repoConnectionDao:   repoConnectionDao,
	}
}
