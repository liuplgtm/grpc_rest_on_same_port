# grpc_rest_on_same_port
grpc and grpc gateway served by the same port

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
