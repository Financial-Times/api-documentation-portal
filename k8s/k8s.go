package k8s

import (
	"context"
	"fmt"

	"github.com/Financial-Times/api-documentation-portal/service"

	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type k8sWatcher struct {
	k8s kubernetes.Interface
}

func NewLocalWatcher(kubeconfig string) service.Watcher {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	return newK8sWatcherForConfig(config)
}

func NewWatcher() service.Watcher {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	return newK8sWatcherForConfig(config)
}

func newK8sWatcherForConfig(config *rest.Config) service.Watcher {
	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(fmt.Sprintf("Failed to create k8s client, error was: %v", err.Error()))
	}

	k8sService := &k8sWatcher{
		k8s: k8sClient,
	}

	return k8sService
}

func (k *k8sWatcher) Watch(ctx context.Context, registry *service.Registry) error {
	watcher, err := k.k8s.CoreV1().Services("default").Watch(meta_v1.ListOptions{LabelSelector: "hasOpenAPI=true"})
	if err != nil {
		return err
	}

	results := watcher.ResultChan()
	for service := range results {
		k8sService := service.Object.(*v1.Service)
		// TODO: get the port properly
		registry.RegisterService(k8sService.Name, "http://"+k8sService.Name+":8080/__api")
	}

	return nil
}
