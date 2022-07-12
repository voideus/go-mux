package router

import (
	"net/http"

	"go.uber.org/fx"
)

type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}

var Module = fx.Options(
	fx.Provide(NewMuxRouter),
)
