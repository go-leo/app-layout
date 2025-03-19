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
	"github.com/gorilla/mux"
)

var Provider = wire.NewSet(
	domain.Provider,
	infra.Provider,
	service.Provider,
	ui.Provider,
	muxMiddlewares,
	newRouter,
	NewHttpServer,
	NewGrpcServer,
)

func InitGrpcServer(ctx context.Context) (*grpcserverx.Server, error) {
	wire.Build(Provider)
	return nil, nil
}

func InitHttpServer(ctx context.Context) (*httpserverx.Server, error) {
	wire.Build(Provider)
	return nil, nil
}

func NewGrpcServer() *grpcserverx.Server {
	return grpcserverx.NewServer(grpcServerOptions()...)
}

func NewHttpServer(router *mux.Router) *httpserverx.Server {
	return httpserverx.NewServer(router, httpServerOptions()...)
}
