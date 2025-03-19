//go:build wireinject

package main

import (
	"context"
	"github.com/go-leo/app-layout/domain"
	"github.com/go-leo/app-layout/infra"
	"github.com/go-leo/app-layout/service"
	"github.com/go-leo/app-layout/ui"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	domain.Provider,
	infra.Provider,
	service.Provider,
	ui.Provider,
	muxMiddlewares,
	newRouter,
	httpTransportOptions,
	httpServerOptions,
	httpServer,
	grpcServerOptions,
	grpcTransportOptions,
	grpcServerxOptions,
	grpcServer,
)

func InitGrpcServer(ctx context.Context) (*grpcserverx.Server, error) {
	wire.Build(Provider)
	return nil, nil
}

func InitHttpServer(ctx context.Context) (*httpserverx.Server, error) {
	wire.Build(Provider)
	return nil, nil
}
