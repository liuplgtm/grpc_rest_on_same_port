# grpc_rest_on_same_port
grpc and grpc gateway served by the same port.

Internally, it relies on cmux to multiplex into two connections, one for grpc and the other for grpc_gateway (for REST)

first, some preparation.
```
$ go mod init nice
$ protoc -I/usr/local/include -I. \
  -I$GOPATH/src \  
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  hello/hello.proto  
$ protoc -I/usr/local/include -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:. \
    hello/hello.proto
```

terminal 1:
```
go run server/main.go
```

terminal 2:
```
go run client/main.go
```
Note: this is for grpc. I cannot get grpc_cli working.

terminal 2:
```
curl -X POST -i http://localhost:8080/v1/example/echo/abcde -d '{"num": "100"}'
```
{"id":"abcde","num":"100"}%
