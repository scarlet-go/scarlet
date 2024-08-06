package scarlet

import (
	"github.com/fatih/color"
)

type Scarlet struct {
	config ScarletConfig
	routes ScarletRoutes
}

type ScarletConfig struct {
	Prefix string
	Name   string
}

type ScarletInterface interface {
	Get(route string, handlers ...ScarletRouteHandler) *Scarlet
	Post(route string, handlers ...ScarletRouteHandler) *Scarlet
	Put(route string, handlers ...ScarletRouteHandler) *Scarlet
	Patch(route string, handlers ...ScarletRouteHandler) *Scarlet
	Delete(route string, handlers ...ScarletRouteHandler) *Scarlet
	Use(handler *Scarlet) *Scarlet
	Listen(port string)
}

func New(config ScarletConfig) (instance *Scarlet) {

	app := Scarlet{}
	app.routes = make(ScarletRoutes)
	app.config = config

	return &app
}

func CreateScarletApplication(config ScarletConfig) ScarletInterface {
	color.Green("[Scarlet] LOG [ScarletFactory] Starting Scarlet Appplication...")

	return New(config)
}

// func checkInterface() ScarletInterface {
// 	return New(ConfigInstance{})
// }
