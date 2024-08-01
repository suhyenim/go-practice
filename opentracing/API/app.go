package API

import (
	"go-code/API/network"
	"go-code/jaeger"
)

type App struct {
	Client *jaeger.Client
}

func NewApp(serviceName string) {
	a := &App{}

	a.Client = jaeger.NewClient(serviceName)
	n := network.NewNetwork(a.Client)
	n.Start()

	// TODO Network Start
}
