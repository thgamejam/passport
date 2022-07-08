//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"passport/internal/biz"
	"passport/internal/conf"
	"passport/internal/data"
	"passport/internal/server"
	"passport/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Passport, registry.Registrar, registry.Discovery, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
