package service

import (
	"context"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/loads/fmts"
	log "github.com/sirupsen/logrus"
)

func init() {
	loads.AddLoader(fmts.YAMLMatcher, fmts.YAMLDoc)
}

// Watcher watches for new services via cluster APIs. Kubernetes clusters work via k8s API, Coco via etcd.
type Watcher interface {
	Watch(ctx context.Context, registry *Registry) error
}

// Registry is an in memory collection of services found by the watcher.
type Registry struct {
	Services map[string]Service `json:"services"`
}

// Service is information gleaned from the /__api of a service
type Service struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	APIEndpoint string `json:"api"`
}

// NewRegistry creates a new registry
func NewRegistry() *Registry {
	return &Registry{Services: make(map[string]Service, 0)}
}

// RegisterService registers a new service by decoding the OpenAPI data from the provided /__api endpoint
func (r *Registry) RegisterService(name string, endpoint string) {
	swagger, err := loads.Spec(endpoint)
	if err != nil {
		log.WithField("service", name).WithField("endpoint", endpoint).WithError(err).Info("Failed to parse/locate OpenAPI spec for service.")
		return
	}

	r.Services[name] = Service{name, swagger.Spec().Info.Title, swagger.Spec().Info.Description, endpoint}
}
