package infrastructure

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRouter),
	fx.Provide(NewFBApp),
	fx.Provide(NewFBAuth),
)
