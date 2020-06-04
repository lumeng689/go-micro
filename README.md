# go-micro

本项目基于gRPC+gRPC Gateway搭建，可以同时提供gRPC服务与REST服务。

文档参考自：

- [Grpc+Grpc Gateway实践一 介绍与环境安装](https://segmentfault.com/a/1190000013339403)
- [Grpc+Grpc Gateway实践二 有些复杂的Hello World](https://segmentfault.com/a/1190000013408485)
- [Grpc+Grpc Gateway实践三 Swagger了解一下](https://segmentfault.com/a/1190000013513469)

# 启动项目

```bash
# 安装包
go list -m -json all
```

* 启动项目

```bash
go build main.go  server -p=50053 --secure=true
```


