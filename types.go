package scarlet

type Map = map[string]any

type ScarletRouteHandler = func(ctx ScarletRequestContext) interface{}
type ScarletRoutes = map[string]map[string][]ScarletRouteHandler
