package scarlet

import "net/http"

func (s *Scarlet) Get(route string, handlers ...func(ctx http.Request) interface{}) *Scarlet {
	return processMethod(s, route, "GET", handlers)
}

func (s *Scarlet) Post(route string, handlers ...func(ctx http.Request) interface{}) *Scarlet {
	return processMethod(s, route, "POST", handlers)

}

func (s *Scarlet) Put(route string, handlers ...func(ctx http.Request) interface{}) *Scarlet {
	return processMethod(s, route, "PUT", handlers)
}

func (s *Scarlet) Patch(route string, handlers ...func(ctx http.Request) interface{}) *Scarlet {
	return processMethod(s, route, "PATCH", handlers)

}

func (s *Scarlet) Delete(route string, handlers ...func(ctx http.Request) interface{}) *Scarlet {
	return processMethod(s, route, "DELETE", handlers)

}

func processMethod(s *Scarlet, route string, method string, handlers []func(ctx http.Request) interface{}) *Scarlet {
	if s.routes[route] == nil {
		s.routes[route] = make(map[string][]func(ctx http.Request) interface{})
	}

	s.routes[route][method] = handlers

	if len(handlers) == 0 {
		println("No handlers implemented for route " + route)
	}

	return s
}
