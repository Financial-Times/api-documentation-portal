package coco

import (
	"context"
	"net/http"
	"time"

	etcdClient "github.com/coreos/etcd/client"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/proxy"
)

// Watcher see Watch func for details
type Watcher interface {
	Watch(ctx context.Context, key string, callback func(key string, val string))
	Read(key string, callback func(key string, val string)) error
}

type etcdWatcher struct {
	api etcdClient.KeysAPI
}

// NewEtcdWatcher returns a new etcd watcher
func NewEtcdWatcher(endpointsList []string) (Watcher, error) {
	transport := &http.Transport{
		Dial: proxy.Direct.Dial,
		ResponseHeaderTimeout: 10 * time.Second,
		MaxIdleConnsPerHost:   100,
	}

	etcdCfg := etcdClient.Config{
		Endpoints:               endpointsList,
		Transport:               transport,
		HeaderTimeoutPerRequest: 10 * time.Second,
	}

	client, err := etcdClient.New(etcdCfg)
	if err != nil {
		log.Printf("Cannot load etcd configuration: [%v]", err)
		return nil, err
	}

	api := etcdClient.NewKeysAPI(client)
	return &etcdWatcher{api}, nil
}

func (e *etcdWatcher) Read(key string, callback func(key string, val string)) error {
	resp, err := e.api.Get(context.Background(), key, &etcdClient.GetOptions{Recursive: true})
	if err != nil {
		return err
	}

	if resp != nil {
		processNodeAndChildren(resp.Node, callback)
	}

	return nil
}

// Watch starts an etcd watch on a given key, and triggers the callback when found
func (e *etcdWatcher) Watch(ctx context.Context, key string, callback func(key string, val string)) {
	watcher := e.api.Watcher(key, &etcdClient.WatcherOptions{AfterIndex: 0, Recursive: true})

	for {
		if ctx.Err() != nil {
			log.WithField("key", key).Info("Etcd watcher cancelled.")
			break
		}

		resp, err := watcher.Next(ctx)
		if err != nil {
			log.WithError(err).WithField("key", key).Info("Error occurred while waiting for change in etcd. Sleeping 10s")
			time.Sleep(10 * time.Second)
			continue
		}

		if resp != nil {
			processNodeAndChildren(resp.Node, callback)
		}
	}
}

func processNodeAndChildren(node *etcdClient.Node, callback func(key string, val string)) {
	if node == nil {
		return
	}

	runCallback(node, callback)

	if node.Nodes == nil {
		return
	}

	for _, child := range node.Nodes {
		processNodeAndChildren(child, callback)
	}
}

func runCallback(node *etcdClient.Node, callback func(key string, val string)) {
	defer func() {
		if r := recover(); r != nil {
			log.WithField("panic", r).Error("Watcher callback panicked! This should not happen, and indicates there is a bug.")
		}
	}()

	callback(node.Key, node.Value)
}
