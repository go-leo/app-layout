# app-layout
app-layout 是一个应用服务的分层架构，主要分为：
- cmd: 程序启动命令
- api: 接口定义
- ui: 用户界面(可以简单理解为传统MVC的controller层)
- service: 业务逻辑层，即use case,用来编排domain层的逻辑
- domain: 领域模型层，领域模型为充血模型，有自己的方法，但方法逻辑都是"原子性"的，需要通过service层来编排
- infra: 基础设施层，比如数据库、缓存、消息队列、RPC等
```go
├── api
│   └── v1
│       ├── app.pb.go
│       ├── app.pb.validate.go
│       ├── app.proto
│       ├── app_grpc.pb.go
│       ├── app_leo.core.pb.go
│       ├── app_leo.grpc.pb.go
│       └── app_leo.http.pb.go
├── cmd
│   ├── grpc.go
│   ├── http.go
│   ├── main.go
│   ├── wire.go
│   └── wire_gen.go
├── domain
│   └── wire.go
├── go.mod
├── go.sum
├── infra
│   └── wire.go
├── protoc.sh
├── service
│   └── wire.go
└── ui
    ├── user_service.go
    └── wire.go
```


