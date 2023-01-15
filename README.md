# go-httpserver

本仓库是使用 Golang 的 Gin 框架编写的一个供个人云原生学习的 HttpServer

该 HttpServer 默认运行在 8080 端口提供服务

本地构建测试：
```shell
go mod tidy

go build -o httpserver
```

本地运行调试：
```shell
go run main.go
```

该 HttpServer 共有两个测试的 URL，分别为 `/` 路径与应用健康检查 `/healthy` 路径

同时针对于 `/healthy` 路径 提供了 POST 请求来动态修改应用健康状态
```shell
curl localhost:8080

curl localhost:8080/healthy

curl -X POST -d '{"Status":"Fail"}' http://localhost:8080/healthy

curl -X POST -d '{"Status":"OK"}' http://localhost:8080/healthy
```
