package lib

import (
	"context"
	"fmt"
	"log"
	"strings"

	apiv1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	networkingv1Typed "k8s.io/client-go/kubernetes/typed/networking/v1"
)

type INetworkManager interface {
	ListNetworks(ListNetworkArgs) ([]CreateNetworkResult, error)
	CreateNetwork(CreateNetworkArgs) (*CreateNetworkResult, error)
	GetNetwork(GetNetworkArgs) (*CreateNetworkResult, error)
	UpdateNetwork(UpdateNetworkArgs) (*CreateNetworkResult, error)
	DeleteNetwork(DeleteNetworkArgs) error

	convertToServicePort(NetworkPort) apiv1.ServicePort
	convertToContainerPort(NetworkPort) apiv1.ContainerPort
}

type NetworkManager struct {
	servicesClient v1.ServiceInterface
	ingressClient  networkingv1Typed.IngressInterface
}

// ListNetworks implements NetworkInterface.
func (*NetworkManager) ListNetworks(ListNetworkArgs) ([]CreateNetworkResult, error) {
	panic("unimplemented")
}

// CreateNetwork implements NetworkInterface.
func (n *NetworkManager) CreateNetwork(args CreateNetworkArgs) (*CreateNetworkResult, error) {
	serviceName := Slugify(fmt.Sprintf("%s-service-%s", args.Name, RandomString(6)))
	ingressName := Slugify(fmt.Sprintf("%s-ingress-%s", args.Name, RandomString(6)))

	ingressAnnotations := map[string]string{"cert-manager.io/cluster-issuer": "letsencrypt-nginx"}
	ingressClassName := "nginx"
	ingressPath := "/"
	ingressPathType := "Prefix"

	// create service for deployment
	service, err := n.servicesClient.Create(context.TODO(), &apiv1.Service{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:   serviceName,
			Labels: args.Labels,
		},
		Spec: apiv1.ServiceSpec{
			Ports:    []apiv1.ServicePort{n.convertToServicePort(args.Port)},
			Selector: args.Labels,
		},
	}, metav1.CreateOptions{})
	if err != nil {
		log.Println(err)

		return nil, err
	}

	ingress, err := n.ingressClient.Create(context.TODO(), &networkingv1.Ingress{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:        ingressName,
			Annotations: ingressAnnotations,
			Labels:      args.Labels,
		},
		Spec: networkingv1.IngressSpec{
			IngressClassName: &ingressClassName,
			TLS: []networkingv1.IngressTLS{
				{
					Hosts:      []string{args.HostName},
					SecretName: args.HostName,
				},
			},
			Rules: []networkingv1.IngressRule{
				{
					Host: args.HostName,
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Path:     ingressPath,
									PathType: (*networkingv1.PathType)(&ingressPathType),
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: serviceName,
											Port: networkingv1.ServiceBackendPort{
												Number: int32(args.Port.DestinationPort),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}, metav1.CreateOptions{})
	if err != nil {
		// TODO: delete service
		log.Println(err)

		return nil, err
	}

	return &CreateNetworkResult{
		ServiceID: service.Name,
		IngressID: ingress.Name,
		HostName:  ingress.Spec.Rules[0].Host,
	}, nil
}

// GetNetwork implements NetworkInterface.
func (*NetworkManager) GetNetwork(GetNetworkArgs) (*CreateNetworkResult, error) {
	panic("unimplemented")
}

// UpdateNetwork implements NetworkInterface.
func (*NetworkManager) UpdateNetwork(UpdateNetworkArgs) (*CreateNetworkResult, error) {
	panic("unimplemented")
}

// DeleteNetwork implements NetworkInterface.
func (n *NetworkManager) DeleteNetwork(args DeleteNetworkArgs) error {
	deletePolicy := metav1.DeletePropagationForeground
	err := n.servicesClient.Delete(context.TODO(), args.ServiceID, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	if err != nil {
		return err
	}

	err = n.ingressClient.Delete(context.TODO(), args.IngressID, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	if err != nil {
		return err
	}

	return nil
}

// convertToServicePort implements NetworkInterface.
func (*NetworkManager) convertToServicePort(port NetworkPort) apiv1.ServicePort {
	return apiv1.ServicePort{
		Name:       strings.ToLower(RandomString(6)),
		TargetPort: intstr.FromInt(port.DestinationPort),
		Port:       int32(port.DestinationPort),
		Protocol:   apiv1.Protocol(port.Protocol),
	}
}

// convertToServicePort implements NetworkInterface.
func (*NetworkManager) convertToContainerPort(port NetworkPort) apiv1.ContainerPort {
	return apiv1.ContainerPort{
		ContainerPort: int32(port.DestinationPort),
		Protocol:      apiv1.Protocol(port.Protocol),
	}
}

func NewNetworkManager(clientset *kubernetes.Clientset) INetworkManager {
	servicesClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)
	ingressClient := clientset.NetworkingV1().Ingresses(apiv1.NamespaceDefault)

	return &NetworkManager{
		servicesClient: servicesClient,
		ingressClient:  ingressClient,
	}
}
