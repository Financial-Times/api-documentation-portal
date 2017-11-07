package service

import (
	"context"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/loads/fmts"
	"github.com/go-openapi/spec"
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
	Name        string   `json:"name"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Paths       []string `json:"paths"`
	APIEndpoint string   `json:"api"`
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

	spec := swagger.Spec()
	paths, tags := getPathsAndTags(spec)

	r.Services[name] = Service{
		Name:        name,
		Title:       spec.Info.Title,
		Description: spec.Info.Description,
		Tags:        tags,
		Paths:       paths,
		APIEndpoint: endpoint,
	}
}

func getPathsAndTags(spec *spec.Swagger) ([]string, []string) {
	tagSet := make(map[string]struct{})
	pathSet := make(map[string]struct{})

	for p, path := range spec.Paths.Paths {
		pathSet[p] = struct{}{}

		if path.Get != nil {
			appendToMap(tagSet, path.Get.Tags...)
		}

		if path.Post != nil {
			appendToMap(tagSet, path.Post.Tags...)
		}

		if path.Put != nil {
			appendToMap(tagSet, path.Post.Tags...)
		}

		if path.Delete != nil {
			appendToMap(tagSet, path.Post.Tags...)
		}
	}

	return keysToArray(pathSet), keysToArray(tagSet)
}

func keysToArray(m map[string]struct{}) []string {
	r := make([]string, 0)
	for k := range m {
		r = append(r, k)
	}
	return r
}

func appendToMap(m map[string]struct{}, data ...string) {
	for _, v := range data {
		m[v] = struct{}{}
	}
}
