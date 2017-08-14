package main

import (
	"encoding/json"
	"net/http"
)

type services struct {
	Apps []App `json:"services"`
}

func appsHandler(registry *AppRegistry) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		enc := json.NewEncoder(w)

		a := services{Apps: make([]App, 0)}
		for _, v := range registry.Apps {
			a.Apps = append(a.Apps, v)
		}

		enc.Encode(a)
	}
}

func gtg() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}
}
