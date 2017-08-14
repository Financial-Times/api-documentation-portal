package main

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type k8sWatcher struct {
	k8s kubernetes.Interface
}

func NewLocalK8sWatcher(kubeconfig string) *k8sWatcher {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	return newK8sWatcherForConfig(config)
}

func NewK8sWatcher() *k8sWatcher {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	return newK8sWatcherForConfig(config)
}

func newK8sWatcherForConfig(config *rest.Config) *k8sWatcher {
	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(fmt.Sprintf("Failed to create k8s client, error was: %v", err.Error()))
	}

	k8sService := &k8sWatcher{
		k8s: k8sClient,
	}

	return k8sService
}

func (k *k8sWatcher) Watch(ctx context.Context, registry *ServiceRegistry) error {
	log.Info("Watching k8s")
	watcher, err := k.k8s.CoreV1().Services("default").Watch(meta_v1.ListOptions{LabelSelector: "hasHealthcheck=true"})
	if err != nil {
		return err
	}

	results := watcher.ResultChan()
	for service := range results {
		k8sService := service.Object.(*v1.Service)
		log.Info(k8sService.Name)
		// TODO: get the port properly

		registry.RegisterService(k8sService.Name, "http://"+k8sService.Name+":8080/__api")
	}

	return nil
}
