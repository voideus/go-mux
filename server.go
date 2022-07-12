package main

import (
	"net/http"

	"github.com/voideus/golang-mux-rest/controller"
	router "github.com/voideus/golang-mux-rest/http"
	"github.com/voideus/golang-mux-rest/lib"
	"github.com/voideus/golang-mux-rest/repository"
	"github.com/voideus/golang-mux-rest/service"
	"go.uber.org/fx"
)


func main() {	
	fx.New(fx.Options(
		router.Module,
		lib.Module,
		controller.Module,
		service.Module,
		repository.Module,
		fx.Provide(http.NewServeMux),
	),
		fx.Invoke(startApp),
	).Run()
}

func startApp(router router.MuxHandler, mux *http.ServeMux) {
	router.SetupRoutes()
	go http.ListenAndServe(":5000", mux)
}
