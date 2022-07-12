package bundlefx

import (
	"context"
	"fmt"
	"net/http"

	"github.com/voideus/golang-mux-rest/controller"
	router "github.com/voideus/golang-mux-rest/http"
	"github.com/voideus/golang-mux-rest/lib"
	"github.com/voideus/golang-mux-rest/repository"
	"github.com/voideus/golang-mux-rest/service"
	"go.uber.org/fx"
)

func registerHooks(
	lifecycle fx.Lifecycle,
	router router.MuxHandler, mux *http.ServeMux,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				fmt.Println("Server up and running")
				router.SetupRoutes()
				go http.ListenAndServe(":5000", mux)
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}

var Module = fx.Options(
	router.Module,
	lib.Module,
	controller.Module,
	service.Module,
	repository.Module,
	fx.Provide(http.NewServeMux),
	fx.Invoke(registerHooks),
)