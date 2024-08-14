package main

import (
	"github.com/alexohneander/image-updater/internal/k8s"
)

func main() {

	// get config
	// creates the in-cluster config
	// config, err := rest.InClusterConfig()
	// if err != nil {
	// 	panic(err.Error())
	// }

	config, err := k8s.GetLocalConfig()
	if err != nil {
		panic(err.Error())
	}

	// start watcher
	k8s.StartWatcher(config)
}
