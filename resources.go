package main

import (
	"encoding/json"
	"net/http"
)

type apps struct {
	Apps []App `json:"apps"`
}

func appsHandler(registry *AppRegistry) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		enc := json.NewEncoder(w)

		a := apps{}
		for _, v := range registry.Apps {
			a.Apps = append(a.Apps, v)
		}

		enc.Encode(a)
	}
}
