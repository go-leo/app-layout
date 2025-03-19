package main

import (
	"context"
	"github.com/go-leo/app-layout/api/v1"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/healthx"
	"github.com/go-leo/leo/v3/pprofx"
	"github.com/go-leo/leo/v3/prometheusx"
	"github.com/go-leo/leo/v3/serverx/actuator"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/go-leo/leo/v3/transportx/httptransportx"
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
		actuatorSrv := actuator.NewServer(6060, actuatorRouter)
		app := leo.NewApp(leo.Runner(httpSrv, actuatorSrv))
		if err := app.Run(ctx); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}

func newRouter(
	ctx context.Context,
	userService api.UserService,
	mwf []mux.MiddlewareFunc,
	httpOptions []httptransportx.ServerOption,
) *mux.Router {
	router := mux.NewRouter()
	// 添加gorilla/mux中间件
	router.Use(mwf...)
	// 添加路由
	api.AppendUserHttpServerRoutes(router, userService, httpOptions...)
	return router
}

// muxMiddlewares 设置mux中间件
func muxMiddlewares(ctx context.Context) []mux.MiddlewareFunc {
	return []mux.MiddlewareFunc{}
}

// httpTransportOptions 设置Http Transport可选项
func httpTransportOptions(ctx context.Context) []httptransportx.ServerOption {
	return []httptransportx.ServerOption{}
}

// httpServerOptions 设置Http Server可选项
func httpServerOptions(ctx context.Context) []httpserverx.Option {
	return []httpserverx.Option{
		httpserverx.Port(8080),
	}
}

// httpServer 创建Http Server
func httpServer(ctx context.Context, router *mux.Router, httpServerOptions []httpserverx.Option) *httpserverx.Server {
	return httpserverx.NewServer(router, httpServerOptions...)
}
