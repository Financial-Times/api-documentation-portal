package resources

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Financial-Times/api-documentation-portal/service"
	"github.com/husobee/vestigo"
	log "github.com/sirupsen/logrus"
)

const defaultProtocol = "https://"
const defaultServiceName = "/__api-documentation-portal"
const defaultAPIEndpoint = "/__api"

type services struct {
	Services []service.Service `json:"services"`
}

// Services returns the list of registered apps
func Services(registry *service.Registry) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		enc := json.NewEncoder(w)

		a := services{Services: make([]service.Service, 0)}
		for _, v := range registry.Services {
			v.APIEndpoint = defaultProtocol + r.Host + defaultServiceName + r.URL.Path + "/" + v.Name + defaultAPIEndpoint
			a.Services = append(a.Services, v)
		}

		enc.Encode(a)
	}
}

// ProxyAPI proxies /__api requests to the app
func ProxyAPI(registry *service.Registry) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		serviceName := vestigo.Param(r, "service")
		w.Header().Add("Content-Type", "application/json")

		service := service.Service{}
		for _, v := range registry.Services {
			if v.Name == serviceName {
				service = v
				break
			}
		}

		req, _ := http.NewRequest("GET", service.APIEndpoint, nil)
		req.Header.Add("X-Original-Request-URL", defaultProtocol+r.Host+"/__"+service.Name+defaultAPIEndpoint)
		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		_, err = io.Copy(w, resp.Body)
		if err != nil {
			log.WithError(err).WithField("service", serviceName).Error("Failed to proxy /__api endpoint for service")
		}
	}
}

// GTG is a no-op gtg for k8s purposes
func GTG() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}
}
