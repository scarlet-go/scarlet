package scarlet

import (
	"net/http"

	"github.com/fatih/color"
)

func (s *Scarlet) Use(handler *Scarlet) *Scarlet {
	color.Green("[Scarlet] LOG [ModuleLoader] Loading module: %s", handler.config.Name)

	for route, methods := range handler.routes {
		for method := range methods {
			targetRoute := handler.config.Prefix + route

			if s.routes[targetRoute] == nil {
				s.routes[targetRoute] = make(map[string][]func(ctx http.Request) interface{})
			}

			s.routes[targetRoute][method] = handler.routes[route][method]
		}
	}

	return s
}
