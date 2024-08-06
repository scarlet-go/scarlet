package scarlet

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/fatih/color"
)

func createRoutesTree(routes ScarletRoutes, config ConfigInstance) {
	for route, methods := range routes {
		targetRoute := config.Prefix + route

		for method := range methods {
			color.Green("[Scarlet] LOG [RouterExplorer] Mapped {%s, %s} route", targetRoute, method)
		}

		http.HandleFunc(targetRoute, func(w http.ResponseWriter, r *http.Request) {
			if _, ok := methods[r.Method]; !ok {
				methodNotAllowed(w)
				return
			}

			for method := range methods {
				if method != r.Method {
					continue
				}

				ctx := &ScarletContext{Request: *r, inherited: make(map[string]interface{})}

				methodHandlers := methods[method]
				routerHandlerIndex := len(methodHandlers) - 1

				var routeHandler func(ctx ScarletContext) interface{}

				for i := 0; i < len(methodHandlers); i++ {
					if i == routerHandlerIndex {
						routeHandler = methodHandlers[i]
						break
					}

					handler := methodHandlers[i](*ctx)

					switch v := handler.(type) {
					case ScarletError:
						statusCode := v.StatusCode
						message := v.Message

						http.Error(w, message, statusCode)
						return
					}
				}

				if routeHandler != nil {
					handler := routeHandler(*ctx)

					switch v := handler.(type) {
					case string:
						io.WriteString(w, v)
					case Map:
						data, err := json.Marshal(v)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
						io.WriteString(w, string(data))
					case ScarletError:
						statusCode := v.StatusCode
						message := v.Message

						http.Error(w, message, statusCode)
					}
				}

			}

			// methodHandlers := methods[r.Method]
			// routerHandlerIndex := len(methodHandlers) - 1

			// var routeHandler func(ctx ScarletContext) interface{}

			// for i := 0; i < len(methodHandlers); i++ {
			// 	if i == routerHandlerIndex {
			// 		routeHandler = methodHandlers[i]
			// 		break
			// 	}

			// 	handler := methodHandlers[i](ScarletContext{Request: *r, inherited: make(map[string]interface{})})

			// 	switch v := handler.(type) {
			// 	case ScarletError:
			// 		statusCode := v.StatusCode
			// 		message := v.Message

			// 		http.Error(w, message, statusCode)
			// 		return
			// 	}
			// }

			// if routeHandler != nil {
			// 	handler := routeHandler(ScarletContext{Request: *r, inherited: make(map[string]interface{})})

			// 	switch v := handler.(type) {
			// 	case string:
			// 		io.WriteString(w, v)
			// 	case Map:
			// 		data, err := json.Marshal(v)
			// 		if err != nil {
			// 			http.Error(w, err.Error(), http.StatusInternalServerError)
			// 			return
			// 		}
			// 		io.WriteString(w, string(data))
			// 	case ScarletError:
			// 		statusCode := v.StatusCode
			// 		message := v.Message

			// 		http.Error(w, message, statusCode)
			// 	}
			// }
		})
	}
}

func methodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
