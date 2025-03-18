package main

import (
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/healthx"
	"github.com/go-leo/leo/v3/pprofx"
	"github.com/go-leo/leo/v3/prometheusx"
	"github.com/go-leo/leo/v3/serverx/actuator"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
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
		actuatorSrv := actuator.NewServer(16060, actuatorRouter)
		app := leo.NewApp(leo.Runner(grpcSrv, actuatorSrv))
		if err := app.Run(ctx); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}

func grpcServerOptions() []grpcserverx.Option {
	return []grpcserverx.Option{}
}
