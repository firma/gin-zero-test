//go:build wireinject

package main

import (
	"github.com/google/wire"
	"miya/api/internal"
	"miya/api/internal/app"
)

// CreateApp init application.
func CreateApp() (*app.Application, func(), error) {
	panic(wire.Build(internal.ProviderSet))
}
