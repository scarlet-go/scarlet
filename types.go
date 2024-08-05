package scarlet

import "net/http"

type Map = map[string]any

type ScarletRoutes = map[string]map[string][]func(ctx http.Request) interface{}
