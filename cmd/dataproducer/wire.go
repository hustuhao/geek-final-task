//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"dataproducer/internal/biz"
	"dataproducer/internal/conf"
	"dataproducer/internal/data"
	"dataproducer/internal/dataproducer"
	"dataproducer/internal/dataproducer/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(dataproducer.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
