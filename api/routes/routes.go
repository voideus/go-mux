package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewPostRoutes),
	fx.Provide(NewRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(PostRoutes *PostRoutes) Routes {
	return Routes{
		PostRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
