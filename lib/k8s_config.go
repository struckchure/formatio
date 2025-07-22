package lib

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewK8SConfig(env Env) (*kubernetes.Clientset, *rest.Config) {
	var kubeConfigPath string = "./.kube/config"
	WriteTextToFile(kubeConfigPath, env.K8S_CLUSTER_CONFIG)

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Panicln("k8s-error[config]: ", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panicln("k8s-error[clientset]: ", err)
	}

	return clientset, config
}
