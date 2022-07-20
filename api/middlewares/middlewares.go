package middlewares

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAuthMiddleware),
)

// type IMiddleware interface {
// 	Setup()
// }

// type Middlewares []IMiddleware

// func NewMiddlewares() Middlewares {
// 	return Middlewares{}
// }
