//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"dataproducer/internal/admin"
	"dataproducer/internal/admin/service"
	"dataproducer/internal/conf"
	"dataproducer/internal/data"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(admin.ProviderSet, service.ProviderSet, data.AdminProviderSet, newApp))
}
