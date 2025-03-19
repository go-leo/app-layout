package main

import (
	"context"
	"github.com/go-leo/app-layout/api/v1"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/healthx"
	"github.com/go-leo/leo/v3/pprofx"
	"github.com/go-leo/leo/v3/prometheusx"
	"github.com/go-leo/leo/v3/serverx/actuator"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "gRPC server application",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		grpcSrv, err := InitGrpcServer(ctx)
		if err != nil {
			panic(err)
		}

		actuatorRouter := mux.NewRouter()
		actuatorRouter = pprofx.Append(actuatorRouter)
		actuatorRouter = prometheusx.Append(actuatorRouter)
		actuatorRouter = healthx.Append(actuatorRouter)
		actuatorSrv := actuator.NewServer(7070, actuatorRouter)
		app := leo.NewApp(leo.Runner(grpcSrv, actuatorSrv))
		if err := app.Run(ctx); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}

// grpcServerOptions 设置原生 gRPC Server 可选项
func grpcServerOptions() []grpc.ServerOption {
	return []grpc.ServerOption{}
}

// httpTransportOptions 设置Http Transport可选项
func grpcTransportOptions(ctx context.Context) []grpctransportx.ServerOption {
	return []grpctransportx.ServerOption{}
}

// grpcServerxOptions 设置 leo gRPC Server 可选项
func grpcServerxOptions(ctx context.Context, grpcServerOptions []grpc.ServerOption) []grpcserverx.Option {
	return []grpcserverx.Option{
		grpcserverx.Port(9090),
		grpcserverx.ServerOptions(grpcServerOptions...),
	}
}

// grpcServer 创建gRPC Server
func grpcServer(
	ctx context.Context,
	userService api.UserService,
	grpcTransportOptions []grpctransportx.ServerOption,
	grpcServerOptions []grpcserverx.Option,
) *grpcserverx.Server {
	// create gRPC server
	srv := grpcserverx.NewServer(grpcServerOptions...)
	// register grpc service
	api.RegisterUserServer(srv, api.NewUserGrpcServer(userService, grpcTransportOptions...))
	return srv
}
