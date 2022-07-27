//go:build wireinject

package main

import (
	"github.com/google/wire"
	"miya/apix/internal"
	"miya/apix/internal/app"
)

// CreateApp init application.
func CreateApp() (*app.Application, func(), error) {
	panic(wire.Build(internal.ProviderSet))
}
