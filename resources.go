package main

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type services struct {
	Services []Service `json:"services"`
}

func servicesHandler(registry *ServiceRegistry) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.Host, r.URL.String())

		w.Header().Add("Content-Type", "application/json")

		enc := json.NewEncoder(w)

		a := services{Services: make([]Service, 0)}
		for _, v := range registry.Services {
			a.Services = append(a.Services, v)
		}

		enc.Encode(a)
	}
}

// func proxyAPI(registry *ServiceRegistry) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		service := vestigo.Param(r, "service")
// 		w.Header().Add("Content-Type", "application/json")
//
// 		a := services{Services: make([]Service, 0)}
// 		for _, v := range registry.Services {
// 			a.Services = append(a.Services, v)
// 		}
//
// 	}
// }

func gtg() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}
}
