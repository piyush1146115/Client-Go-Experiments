package api

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func int32Ptr(i int32) *int32 {
	return &i
}

func int64Ptr(i int64) *int64 {
	return &i
}

func strPtr(s string) *string {
	return &s
}

var deploymentName string

func SetDeploymentName(name string) {
	deploymentName = name
}

var replicas int32

func SetReplicas(rep int32) {
	replicas = rep
}

func CreateClientSet() kubernetes.Interface {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		panic(err)
	}

	var clientset kubernetes.Interface
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}
