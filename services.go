package main

import (
	"github.com/go-openapi/loads"
	"github.com/go-openapi/loads/fmts"
)

func init() {
	loads.AddLoader(fmts.YAMLMatcher, fmts.YAMLDoc)
}

type AppRegistry struct {
	Apps map[string]App `json:"services"`
}

type App struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	APIEndpoint string `json:"api"`
}

func NewAppRegistry() *AppRegistry {
	return &AppRegistry{Apps: make(map[string]App, 0)}
}

func (s *AppRegistry) RegisterApp(app string, endpoint string) {
	swagger, err := loads.Spec(endpoint)
	if err != nil {
		return
	}

	s.Apps[app] = App{app, swagger.Spec().Info.Title, swagger.Spec().Info.Description, endpoint}
}
