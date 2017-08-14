package main

import (
	"context"
	"net/http"
	"os"
	"regexp"

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
			EnvVar: "EXTERNAL_CLUSTER",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		// watcher, err := NewEtcdWatcher([]string{"http://localhost:2379"})
		// if err != nil {
		// return err
		// }

		registry := NewServiceRegistry()

		// watcherCallback := func(key, val string) {
		// 	if appNameRegex.MatchString(key) {
		// 		app := appNameRegex.FindStringSubmatch(key)[1]
		// 		log.WithField("app", app).WithField("api", val).Info("Found new app")
		// 		registry.RegisterApp(app, val)
		// 	}
		// }

		// watcher.Read("/ft/services", watcherCallback)

		// go watcher.Watch(context.Background(), "/ft/services", watcherCallback)

		useExternalCluster := ctx.Bool("external-cluster")
		var k8s *k8sWatcher
		if useExternalCluster {
			k8s = NewLocalK8sWatcher(os.Getenv("KUBECONFIG"))
		} else {
			k8s = NewK8sWatcher()
		}

		go k8s.Watch(context.Background(), registry)

		return serve(registry)
	}

	app.Run(os.Args)
}

func serve(registry *ServiceRegistry) error {
	r := vestigo.NewRouter()

	r.Get("/__gtg", gtg())
	r.Get("/services", servicesHandler(registry))

	box := ui.UI()
	dist := http.FileServer(box.HTTPBox())
	r.Get("/*", dist.ServeHTTP)

	http.Handle("/", r)

	return http.ListenAndServe(":8080", nil)
}
