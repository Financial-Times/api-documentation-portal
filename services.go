package main

import (
	"github.com/go-openapi/loads"
	"github.com/go-openapi/loads/fmts"
	log "github.com/sirupsen/logrus"
)

func init() {
	loads.AddLoader(fmts.YAMLMatcher, fmts.YAMLDoc)
}

type ServiceRegistry struct {
	Services map[string]Service `json:"services"`
}

type Service struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	APIEndpoint string `json:"api"`
}

func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{Services: make(map[string]Service, 0)}
}

func (s *ServiceRegistry) RegisterService(name string, endpoint string) {
	swagger, err := loads.Spec(endpoint)
	if err != nil {
		log.WithField("service", name).WithField("endpoint", endpoint).WithError(err).Info("Failed to parse/locate OpenAPI spec for service.")
		return
	}

	s.Services[name] = Service{name, swagger.Spec().Info.Title, swagger.Spec().Info.Description, endpoint}
}
