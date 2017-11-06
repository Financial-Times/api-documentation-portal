package main

import (
	"context"
	"net/http"
	"os"
	"regexp"

	"github.com/Financial-Times/api-documentation-portal/k8s"
	"github.com/Financial-Times/api-documentation-portal/resources"
	"github.com/Financial-Times/api-documentation-portal/service"
	"github.com/Financial-Times/api-documentation-portal/ui"
	"github.com/husobee/vestigo"
	"github.com/urfave/cli"
)

var appNameRegex = regexp.MustCompile(`/ft/services/(.*)/api`)

func main() {
	app := cli.NewApp()
	app.Name = "api-documentation-portal"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "external-cluster",
			Usage:  "Use to connect to an external kubernetes cluster using the locally configured KUBECONFIG variable.",
			EnvVar: "EXTERNAL_CLUSTER",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		registry := service.NewRegistry()

		useExternalCluster := ctx.Bool("external-cluster")
		var watcher service.Watcher
		if useExternalCluster {
			watcher = k8s.NewLocalWatcher(os.Getenv("KUBECONFIG"))
		} else {
			watcher = k8s.NewWatcher()
		}

		go watcher.Watch(context.Background(), registry)

		return serve(registry)
	}

	app.Run(os.Args)
}

func serve(registry *service.Registry) error {
	r := vestigo.NewRouter()

	r.Get("/__gtg", resources.GTG())
	r.Get("/services", resources.Services(registry))
	r.Get("/services/:service/__api", resources.ProxyAPI(registry))

	box := ui.UI()
	dist := http.FileServer(box.HTTPBox())
	r.Get("/*", dist.ServeHTTP)

	http.Handle("/", r)

	return http.ListenAndServe(":8080", nil)
}
