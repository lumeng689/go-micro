# proto生成参考

##  安装插件

```bash
go install  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway 
go install  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger 
go install  github.com/golang/protobuf/protoc-gen-go
```

```bash
# 编译google.api
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:../../ google/api/*.proto

# 编译本地文件
protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=google/api:../../ demo/helloworld.proto

# 编译GW模块
protoc --grpc-gateway_out=logtostderr=true:../../ demo/helloworld.proto

```

```
protoc demo/helloworld.proto    --go_out=plugins=grpc:../../ --proto_path=./

```




### 3 protoc 安装

访问[https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)
以当前版本【3.7.1】为例

#### windows
下载 protoc-3.7.1-win64.zip，解压，将bin下的 protoc.exe 文件丢到一个可执行路径里就行了。

#### mac
下载 protoc-3.7.1-osx-x86_64.zip，然后解压文件
```bash
cd ~/Downloads/protoc-3.7.1-osx-x86_64
cp -R ./bin /usr/local/
cp -R ./includes /usr/local
```

### 4 protoc-gen-go 用法

安装
```bash
# 下载protoc文件 https://github.com/protocolbuffers/protobuf/releases

# 下载protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go
```


```bash
# 执行命令进行生成
protoc --go_out=. *.proto
# 生成grpc调用代码，得写上 plguins， ： 后面 是输出路径
protoc ./proto/helloworld.proto --go_out=plugins=grpc:./
```
