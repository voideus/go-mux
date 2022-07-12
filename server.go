package main

import (
	"github.com/voideus/golang-mux-rest/bundlefx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		bundlefx.Module,
	).Run()
}
