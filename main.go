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
	app.Name = "api-hub"

	app.Action = func(ctx *cli.Context) error {
		// watcher, err := NewEtcdWatcher([]string{"http://localhost:2379"})
		// if err != nil {
		// return err
		// }

		registry := NewAppRegistry()

		// watcherCallback := func(key, val string) {
		// 	if appNameRegex.MatchString(key) {
		// 		app := appNameRegex.FindStringSubmatch(key)[1]
		// 		log.WithField("app", app).WithField("api", val).Info("Found new app")
		// 		registry.RegisterApp(app, val)
		// 	}
		// }

		// watcher.Read("/ft/services", watcherCallback)

		// go watcher.Watch(context.Background(), "/ft/services", watcherCallback)

		k8s := NewK8sWatcher()
		k8s.Watch(context.Background())

		return serve(registry)
	}

	app.Run(os.Args)
}

func serve(registry *AppRegistry) error {
	r := vestigo.NewRouter()

	r.Get("/apps", appsHandler(registry))

	box := ui.UI()
	dist := http.FileServer(box.HTTPBox())
	r.Get("/*", dist.ServeHTTP)

	http.Handle("/", r)

	return http.ListenAndServe(":8080", nil)
}
