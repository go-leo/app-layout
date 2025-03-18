package main

import (
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/healthx"
	"github.com/go-leo/leo/v3/pprofx"
	"github.com/go-leo/leo/v3/prometheusx"
	"github.com/go-leo/leo/v3/serverx/actuator"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "HTTP server application",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		httpSrv, err := InitHttpServer(ctx)
		if err != nil {
			panic(err)
		}

		actuatorRouter := mux.NewRouter()
		actuatorRouter = pprofx.Append(actuatorRouter)
		actuatorRouter = prometheusx.Append(actuatorRouter)
		actuatorRouter = healthx.Append(actuatorRouter)
		actuatorSrv := actuator.NewServer(16060, actuatorRouter)
		app := leo.NewApp(leo.Runner(httpSrv, actuatorSrv))
		if err := app.Run(ctx); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}

func newRouter(mwf ...mux.MiddlewareFunc) *mux.Router {
	router := mux.NewRouter()
	// add middleware
	router.Use(mwf...)
	return router
}

func muxMiddlewares() []mux.MiddlewareFunc {
	return []mux.MiddlewareFunc{}
}

func httpServerOptions() []httpserverx.Option {
	return []httpserverx.Option{}
}
