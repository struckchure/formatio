package lib

import (
	apiv1 "k8s.io/api/core/v1"
)

type LogHandler func(message string, logLevel string) error

type ListContainersArgs struct{}

type CreateContainerArgs struct {
	Replicas int32
	Labels   map[string]string

	Name   string
	Image  string
	CPU    string
	Memory string
	Ports  []apiv1.ContainerPort
}

type GetContainerArgs struct {
	DeploymentName string
}

type UpdateContainerArgs struct {
	DeploymentName string `json:"deploymentName,omitempty"`
	Replicas       int32  `json:"replicas,omitempty"`

	Image  string        `json:"image,omitempty"`
	CPU    string        `json:"cpu,omitempty"`
	Memory string        `json:"memory,omitempty"`
	Ports  []NetworkPort `json:"ports,omitempty"`
}

type DeleteContainerArgs struct {
	DeploymentName string
}

type ExecuteCommandInContainerArgs struct {
	DeploymentName string
	Command        []string

	LogHandler LogHandler
}

type NetworkPort struct {
	Protocol        string
	DestinationPort int
}

type ListNetworkArgs struct{}

type CreateNetworkArgs struct {
	Name     string
	Labels   map[string]string
	HostName string
	Port     NetworkPort
}

type CreateNetworkResult struct {
	ServiceID string `json:"serviceId"`
	IngressID string `json:"ingressId"`
	HostName  string `json:"hostName"`
}

type GetNetworkArgs struct{}

type UpdateNetworkArgs struct{}

type DeleteNetworkArgs struct {
	ServiceID string `json:"serviceId"`
	IngressID string `json:"ingressId"`
}
