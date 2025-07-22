package lib

import (
	"log"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

type IInformer interface {
	Deployments(string, func(InformerEventType, ...*appsv1.Deployment) error) error
	Pods(string, func(InformerEventType, ...*apiv1.Pod) error) error
	Services(string, func(InformerEventType, ...*apiv1.Service) error) error

	Start(chan struct{})
	Stop(chan struct{})
}

type Informer struct {
	informerFactory    informers.SharedInformerFactory
	deploymentInformer cache.SharedIndexInformer
	podInformer        cache.SharedIndexInformer
	serviceInformer    cache.SharedIndexInformer
}

type InformerEventType int

const (
	InformerEventCreated InformerEventType = iota + 1
	InformerEventUpdated
	InformerEventDeleted
)

func (i *Informer) Deployments(namespace string, callback func(InformerEventType, ...*appsv1.Deployment) error) error {
	_, err := i.deploymentInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			deployment := obj.(*appsv1.Deployment)

			if deployment.Namespace == namespace {
				err := callback(InformerEventCreated, deployment)
				if err != nil {
					log.Println("Informer: ", err)
					return
				}
			}
		},
		DeleteFunc: func(obj interface{}) {
			deployment := obj.(*appsv1.Deployment)

			if deployment.Namespace == namespace {
				err := callback(InformerEventDeleted, deployment)
				if err != nil {
					log.Println("Informer: ", err)
					return
				}
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldDeployment := oldObj.(*appsv1.Deployment)
			newDeployment := newObj.(*appsv1.Deployment)

			// weird check, but it wouldn't hurt
			if oldDeployment.Namespace == namespace && newDeployment.Namespace == namespace {
				err := callback(InformerEventUpdated, oldDeployment, newDeployment)
				if err != nil {
					log.Println("Informer: ", err)
					return
				}
			}
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (i *Informer) Pods(namespace string, callback func(InformerEventType, ...*apiv1.Pod) error) error {
	_, err := i.podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod := obj.(*apiv1.Pod)

			if pod.Namespace == namespace {
				err := callback(InformerEventCreated, pod)
				if err != nil {
					log.Println("Informer: ", err)
					return
				}
			}
		},
		DeleteFunc: func(obj interface{}) {
			pod := obj.(*apiv1.Pod)

			if pod.Namespace == namespace {
				err := callback(InformerEventDeleted, pod)
				if err != nil {
					log.Println("Informer: ", err)
					return
				}
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldPod := oldObj.(*apiv1.Pod)
			newPod := newObj.(*apiv1.Pod)

			// weird check, but it wouldn't hurt
			if oldPod.Namespace == namespace && newPod.Namespace == namespace {
				err := callback(InformerEventUpdated, oldPod, newPod)
				if err != nil {
					log.Println("Informer: ", err)
					return
				}
			}
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (i *Informer) Services(namespace string, callback func(InformerEventType, ...*apiv1.Service) error) error {
	_, err := i.serviceInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			service := obj.(*apiv1.Service)

			if service.Namespace == namespace {
				err := callback(InformerEventCreated, service)
				if err != nil {
					log.Println("Informer: ", err)
					return
				}
			}
		},
		DeleteFunc: func(obj interface{}) {
			pod := obj.(*apiv1.Service)

			if pod.Namespace == namespace {
				err := callback(InformerEventDeleted, pod)
				if err != nil {
					log.Println("Informer: ", err)
					return
				}
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldService := oldObj.(*apiv1.Service)
			newService := newObj.(*apiv1.Service)

			// weird check, but it wouldn't hurt
			if oldService.Namespace == namespace && newService.Namespace == namespace {
				err := callback(InformerEventUpdated, oldService, newService)
				if err != nil {
					log.Println("Informer: ", err)
					return
				}
			}
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (i *Informer) Start(stopCh chan struct{}) {
	i.informerFactory.Start(stopCh)

	// Wait for the cache to sync
	if !cache.WaitForCacheSync(stopCh, i.podInformer.HasSynced) {
		log.Println("PodInformer: Timed out waiting for caches to sync")
		return
	}

	if !cache.WaitForCacheSync(stopCh, i.serviceInformer.HasSynced) {
		log.Println("ServiceInformer: Timed out waiting for caches to sync")
		return
	}

	// Keep the application running
	<-stopCh
}

func (i *Informer) Stop(stopCh chan struct{}) {
	close(stopCh)
}

func NewInformer(clientset *kubernetes.Clientset, config *rest.Config) IInformer {
	informerFactory := informers.NewSharedInformerFactory(clientset, 0)
	deploymentInformer := informerFactory.Apps().V1().Deployments().Informer()
	podInformer := informerFactory.Core().V1().Pods().Informer()
	serviceInformer := informerFactory.Core().V1().Services().Informer()

	return &Informer{
		informerFactory:    informerFactory,
		deploymentInformer: deploymentInformer,
		podInformer:        podInformer,
		serviceInformer:    serviceInformer,
	}
}
