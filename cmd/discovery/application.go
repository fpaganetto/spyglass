package main

import (
	"fmt"
	"net/http"

	"k8s.io/client-go/kubernetes"
)

type application struct {
	config   config
	server   *http.Server
	k8client *kubernetes.Clientset
}

func (app *application) start() error {
	app.server = &http.Server{
		Addr: fmt.Sprintf(":%d", app.config.port),
	}

	http.HandleFunc("/", discovery)

	err := app.server.ListenAndServe()
	return err
}
