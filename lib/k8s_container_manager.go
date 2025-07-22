package lib

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/client-go/util/retry"
)

type CustomReader struct {
	Underlying io.Reader
	Handler    LogHandler
	LogLevel   string
}

func (cr CustomReader) Read(p []byte) (n int, err error) {
	n, err = cr.Underlying.Read(p)
	if n > 0 {
		data := p[:n]

		if cr.Handler != nil {
			cr.Handler(string(data), cr.LogLevel)
		}
	}

	return n, err
}

type CustomWriter struct {
	Underlying io.Writer
	Handler    LogHandler
	LogLevel   string
}

func (cw CustomWriter) Write(p []byte) (n int, err error) {
	if cw.Handler != nil {
		cw.Handler(string(p), cw.LogLevel)
	}

	n, err = cw.Underlying.Write([]byte(""))
	return n, err
}

type IContainerManager interface {
	ListContainers() ([]appsv1.Deployment, error)
	CreateContainer(CreateContainerArgs) (metav1.Object, error)
	GetContainer(GetContainerArgs) (*appsv1.Deployment, error)
	UpdateContainer(UpdateContainerArgs) error
	DeleteContainer(DeleteContainerArgs) error

	ListDeploymentPods(string) (*apiv1.PodList, error)
	GetDeploymentByPodName(string) (*appsv1.Deployment, error)
	ExecuteCommandInContainer(ExecuteCommandInContainerArgs) error
}

type ContainerManager struct {
	config            *rest.Config
	restClient        rest.Interface
	deploymentsClient v1.DeploymentInterface
	podsClient        corev1.PodInterface
	replicaSetClient  v1.ReplicaSetInterface

	networkManager INetworkManager
}

func (c *ContainerManager) ListContainers() ([]appsv1.Deployment, error) {
	list, err := c.deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return []appsv1.Deployment{}, nil
	}

	return list.Items, nil
}

func (c *ContainerManager) CreateContainer(args CreateContainerArgs) (metav1.Object, error) {
	deploymentName := RandomString(10)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: Slugify(fmt.Sprintf("%s-%s-deployment", args.Name, deploymentName)),
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: Int32Ptr(args.Replicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: args.Labels,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: args.Labels,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  Slugify(args.Name),
							Image: args.Image,
							Ports: args.Ports,
							Resources: apiv1.ResourceRequirements{
								Requests: apiv1.ResourceList{
									apiv1.ResourceMemory: resource.MustParse(args.Memory),
								},
								Limits: apiv1.ResourceList{
									apiv1.ResourceMemory: resource.MustParse(args.Memory),
								},
							},
							ImagePullPolicy: apiv1.PullAlways,
						},
					},
				},
			},
		},
	}

	result, err := c.deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return &metav1.ObjectMeta{}, err
	}

	return result.GetObjectMeta(), nil
}

func (c *ContainerManager) GetContainer(args GetContainerArgs) (*appsv1.Deployment, error) {
	result, err := c.deploymentsClient.Get(context.TODO(), args.DeploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return result, err
}

func (c *ContainerManager) UpdateContainer(args UpdateContainerArgs) error {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		result, err := c.GetContainer(GetContainerArgs{DeploymentName: args.DeploymentName})
		if err != nil {
			return err
		}

		var replicas int32 = *result.Spec.Replicas
		var containerImage string = result.Spec.Template.Spec.Containers[0].Image
		var containerPorts []apiv1.ContainerPort

		if args.Replicas > 0 {
			replicas = args.Replicas
		}

		if len(args.Image) > 0 {
			containerImage = args.Image
		}

		if len(args.Ports) > 0 {
			for _, port := range args.Ports {
				containerPorts = append(containerPorts, c.networkManager.convertToContainerPort(port))
			}
		} else {
			containerPorts = result.Spec.Template.Spec.Containers[0].Ports
		}

		result.Spec.Replicas = &replicas
		result.Spec.Template.Spec.Containers[0].Image = containerImage
		result.Spec.Template.Spec.Containers[0].Ports = containerPorts

		// TODO: resources should be updated if present
		// result.Spec.Template.Spec.Containers[0].Resources.Requests = apiv1.ResourceList{
		// 	apiv1.ResourceMemory: resource.MustParse(args.Memory),
		// }
		// result.Spec.Template.Spec.Containers[0].Resources.Limits = apiv1.ResourceList{
		// 	apiv1.ResourceMemory: resource.MustParse(args.Memory),
		// }

		_, updateErr := c.deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})

		return updateErr
	})

	if retryErr != nil {
		log.Println(retryErr)

		return retryErr
	}

	return nil
}

func (c *ContainerManager) DeleteContainer(args DeleteContainerArgs) error {
	deletePolicy := metav1.DeletePropagationForeground
	if err := c.deploymentsClient.Delete(context.TODO(), args.DeploymentName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *ContainerManager) ListDeploymentPods(deploymentName string) (*apiv1.PodList, error) {
	// Get the Deployment object
	deployment, err := c.deploymentsClient.Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// Get the selector labels from the Deployment
	selectorLabels := deployment.Spec.Selector.MatchLabels

	// Create a label selector string from the labels
	labelSelector := metav1.FormatLabelSelector(&metav1.LabelSelector{MatchLabels: selectorLabels})

	// List Pods using the label selector
	pods, err := c.podsClient.List(context.TODO(), metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		return nil, err
	}

	return pods, nil
}

func (c *ContainerManager) GetDeploymentByPodName(podName string) (*appsv1.Deployment, error) {
	// Step 1: Get the Pod by its name and namespace
	pod, err := c.podsClient.Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pod: %v", err)
	}

	// Step 2: Check the Pod's OwnerReferences
	for _, owner := range pod.OwnerReferences {
		if owner.Kind == "ReplicaSet" {
			// Step 3: Get the ReplicaSet
			rs, err := c.replicaSetClient.Get(context.TODO(), owner.Name, metav1.GetOptions{})
			if err != nil {
				return nil, fmt.Errorf("failed to get replicaset: %v", err)
			}

			// Step 4: Check the ReplicaSet's OwnerReferences for a Deployment
			for _, rsOwner := range rs.OwnerReferences {
				if rsOwner.Kind == "Deployment" {
					return c.deploymentsClient.Get(context.TODO(), rsOwner.Name, metav1.GetOptions{})
				}
			}
		}
	}

	return nil, fmt.Errorf("no deployment found for pod: %s", podName)
}

func (c *ContainerManager) ExecuteCommandInContainer(args ExecuteCommandInContainerArgs) error {
	podList, err := c.ListDeploymentPods(args.DeploymentName)
	if err != nil {
		return err
	}

	if len(podList.Items) == 0 {
		return errors.New("no pods found for the deployment")
	}

	scheme := runtime.NewScheme()
	if err := apiv1.AddToScheme(scheme); err != nil {
		return err
	}

	parameterCodec := runtime.NewParameterCodec(scheme)

	for _, pod := range podList.Items {
		podName := pod.GetName()
		podContainerName := pod.Spec.Containers[0].Name

		req := c.restClient.
			Post().
			Name(podName).
			Namespace(apiv1.NamespaceDefault).
			Resource("pods").
			SubResource("exec").
			VersionedParams(&apiv1.PodExecOptions{
				Container: podContainerName,
				Command:   args.Command,
				Stdin:     true,
				Stdout:    true,
				Stderr:    true,
			}, parameterCodec)

		exec, err := remotecommand.NewSPDYExecutor(c.config, "POST", req.URL())
		if err != nil {
			return err
		}

		customStdin := CustomReader{Underlying: os.Stdin, Handler: args.LogHandler, LogLevel: "INFO"}
		customStdout := CustomWriter{Underlying: os.Stdout, Handler: args.LogHandler, LogLevel: "INFO"}
		customStderr := CustomWriter{Underlying: os.Stderr, Handler: args.LogHandler, LogLevel: "ERROR"}

		err = exec.StreamWithContext(
			context.TODO(),
			remotecommand.StreamOptions{
				Stdin:  &customStdin,
				Stdout: &customStdout,
				Stderr: &customStderr,
				Tty:    false,
			})

		if err != nil {
			return err
		}
	}

	return nil
}

func NewContainerManager(
	clientset *kubernetes.Clientset,
	config *rest.Config,

	networkManager INetworkManager,
) IContainerManager {
	restClient := clientset.CoreV1().RESTClient()
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	podsClient := clientset.CoreV1().Pods(apiv1.NamespaceDefault)
	replicaSetClient := clientset.AppsV1().ReplicaSets(apiv1.NamespaceDefault)

	return &ContainerManager{
		config:            config,
		restClient:        restClient,
		deploymentsClient: deploymentsClient,
		podsClient:        podsClient,
		networkManager:    networkManager,
		replicaSetClient:  replicaSetClient,
	}
}
