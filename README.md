# go-proto-order
a test project by protobuf


# 若修改了proto文件，需要重新执行protoc命令
```protobuf
protoc -I ./pb \
--go_out ./ecommerce --go_opt paths=source_relative \
--go-grpc_out ecommerce/ --go-grpc_out --go_opt paths=source_relative \
order.proto
```


