# go-zero

## Usage
1. 确保自己拥有 `google api`
   可以放入项目根目录的 `third-party` 目录中，类似这样：
   ```shell
   $ ls ./third-party
   └── google
      ├── api
      │  ├── annotations.proto
      │  └── http.proto
      └── protobuf
         ├── compiler
         │  └── plugin.proto
         ├── any.proto
         ├── api.proto
         ├── descriptor.proto
         ├── duration.proto
         ├── empty.proto
         ├── field_mask.proto
         ├── source_context.proto
         ├── struct.proto
         ├── timestamp.proto
         ├── type.proto
         └── wrappers.proto
   ```

2. 安装依赖
   ```shell
   go install github.com/zeromicro/go-zero/tools/goctl@latest
   go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
   go install github.com/envoyproxy/protoc-gen-validate@latest
   ```

3. 生成一个服务
   1. 创建目录和 `proto`
      ```shell
      # 创建目录
      mkdir car && cd car 

      # 生成 proto 文件
      goctl rpc -o api/car/v1/orcarder.proto
      ```
      
      此时的目录结构为：
      ```shell
      ├── api
      │  └── car
      │     └── v1
      │        ├── car.proto
      ```

      `proto` 内容：
      ```protobuf
      syntax = "proto3";

      package order;
      option go_package="./car";

      message Request {
        string ping = 1;
      }

      message Response {
        string pong = 1;
      }

      service Car {
        rpc Ping(Request) returns(Response);
      }
      ```

   2. 生成 `grpc` 代码
      修改 `proto` 中 `go_package` 的值为 `./api/car/v1`，`./api/car/v1`指定生成路径：
      ```protobuf
      syntax = "proto3";

      package order;
      option go_package="./api/car/v1";

      message Request {
        string ping = 1;
      }

      message Response {
        string pong = 1;
      }

      service Car {
        rpc Ping(Request) returns(Response);
      }
      ```

      继续生成 `grpc` 代码，其中 `-m` 是生成多个对应 `proto` 中 `service` 的 `service` 目录
      ```shell
      goctl rpc protoc -m \
           --go_out=. \
           --go-grpc_out=. \
           --zrpc_out=. \
           api/car/v1/car.proto
      ```

   3. 创建 `pb` 和 `openapi` 文件：
      ```shell
      mkdir docs
      protoc --include_imports \
		--proto_path=. \
        --validate_out=paths=source_relative,lang=go:. \
		--openapi_out=paths=source_relative:./docs \
		--descriptor_set_out=./api/car/v1/car.pb \
		./api/car/v1/car.proto
      ```

      此时的目录结构为:
      ```shell
      ├── api
      │  └── car
      │     └── v1
      │        ├── car.pb # pb 文件，gateway 使用
      │        ├── car.pb.go
      │        ├── car.pb.validate.go  # validate
      │        ├── car.proto # proto 文件
      │        └── car_grpc.pb.go
      ├── client # grpc 自动生成，不要修改
      │  └── car
      │     └── car.go
      ├── docs # openapi
      │  └── openapi.yaml
      ├── etc # 配置文件
      │  └── car.yaml
      ├── internal
      │  ├── config # 配置结构体
      │  │  └── config.go
      │  ├── logic # 业务逻辑
      │  │  └── car
      │  │     └── pinglogic.go
      │  ├── server # grpc 自动生成，不要修改
      │  │  └── car # proto 中的 service 名称，如果 proto 中有多个 service 或者有多个 proto，则生成多个目录
      │  │     └── carserver.go
      │  └── svc # 服务上下文结构体
      │     └── servicecontext.go
      ├── car.go # main 文件，其中包括配置加载和服务启动
      ```
