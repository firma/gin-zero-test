// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"miya/api/internal/app"
	"miya/api/internal/service"
)

// Injectors from wire.go:

// CreateApp init application.
func CreateApp() (*app.Application, func(), error) {
	rpcService := service.NewServiceContext()
	application, err := app.NewApp(rpcService)
	if err != nil {
		return nil, nil, err
	}
	return application, func() {
	}, nil
}
