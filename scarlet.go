package scarlet

import (
	"github.com/fatih/color"
)

type Scarlet struct {
	config ConfigInstance
	routes ScarletRoutes
}

type ConfigInstance struct {
	Prefix string
	Name   string
}

type ScarletInterface interface {
	Get(route string, handlers ...func(ctx ScarletContext) interface{}) *Scarlet
	Post(route string, handlers ...func(ctx ScarletContext) interface{}) *Scarlet
	Put(route string, handlers ...func(ctx ScarletContext) interface{}) *Scarlet
	Patch(route string, handlers ...func(ctx ScarletContext) interface{}) *Scarlet
	Delete(route string, handlers ...func(ctx ScarletContext) interface{}) *Scarlet
	Use(handler *Scarlet) *Scarlet
	Listen(port string)
}

func New(config ConfigInstance) (instance *Scarlet) {

	app := Scarlet{}
	app.routes = make(ScarletRoutes)
	app.config = config

	return &app
}

func CreateScarletApplication(config ConfigInstance) ScarletInterface {
	color.Green("[Scarlet] LOG [ScarletFactory] Starting Scarlet Appplication...")

	return New(config)
}

// func checkInterface() ScarletInterface {
// 	return New(ConfigInstance{})
// }
