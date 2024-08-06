package scarlet

type Map = map[string]any

type ScarletRoutes = map[string]map[string][]func(ctx ScarletContext) interface{}
