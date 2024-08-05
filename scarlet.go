package scarlet

import (
	"net/http"

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
	Get(route string, handlers ...func(ctx http.Request) interface{}) *Scarlet
	Post(route string, handlers ...func(ctx http.Request) interface{}) *Scarlet
	Put(route string, handlers ...func(ctx http.Request) interface{}) *Scarlet
	Patch(route string, handlers ...func(ctx http.Request) interface{}) *Scarlet
	Delete(route string, handlers ...func(ctx http.Request) interface{}) *Scarlet
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
