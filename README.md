# go-proto-order
a test project by protobuf


# 若修改了proto文件，需要重新执行protoc命令
```shell
protoc -I ./pb \
--go_out ./ecommerce --go_opt paths=source_relative \
--go-grpc_out ecommerce/ --go-grpc_opt paths=source_relative \
order.proto
```

# 使用
```shell
cd demo/server
## 开启服务端
go run main.go

cd demo/client
## 开启客户端
go run main.go
```
