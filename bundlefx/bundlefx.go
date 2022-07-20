package bundlefx

import (
	"context"
	"fmt"

	"github.com/voideus/golang-mux-rest/api/middlewares"
	"github.com/voideus/golang-mux-rest/api/routes"
	"github.com/voideus/golang-mux-rest/controller"
	"github.com/voideus/golang-mux-rest/infrastructure"
	"github.com/voideus/golang-mux-rest/lib"
	"github.com/voideus/golang-mux-rest/repository"
	"github.com/voideus/golang-mux-rest/service"
	"go.uber.org/fx"
)

func registerHooks(
	lifecycle fx.Lifecycle,
	router infrastructure.Router, env *lib.Env,
	route routes.Routes,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				// router.SetupRoutes()
				// go http.ListenAndServe(env.ServerPort, gin)
				route.Setup()
				// if err := go router.Run(":5000"); err != nil {
				// 	fmt.Printf("error running server: %s", err)
				// 	return err
				// }
				go router.Run(":5000")
				fmt.Printf("Server up and running on port: %v", env.ServerPort)
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
	routes.Module,
	lib.Module,
	controller.Module,
	service.Module,
	repository.Module,
	middlewares.Module,
	infrastructure.Module,
	// fx.Provide(http.NewServeMux),
	// fx.Provide(gin.Default()),
	fx.Invoke(registerHooks),
)
