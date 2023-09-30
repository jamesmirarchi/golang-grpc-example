# grpc-go


### install pkgs
```shell
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install github.com/bufbuild/buf/cmd/buf@latest
```

### Setup env
```shell
[ -n "$(go env GOBIN)" ] && export PATH="$(go env GOBIN):${PATH}"
[ -n "$(go env GOPATH)" ] && export PATH="$(go env GOPATH)/bin:${PATH}"
```

### Build proto

```shell
─$ protoc event/v1/event.proto \
  --go_out=. \
  --go_opt=paths=source_relative \
  --proto_path=. \
--go-grpc_out=. --go-grpc_opt=paths=source_relative
```

## Server 
```shell
go run cmd/server/main.go 
```

## Client commands for testing

```shell
grpcurl -plaintext localhost:50001 event.v1.EventService/Retrieve

grpcurl -plaintext localhost:50001 list
grpcurl -plaintext localhost:50001 describe event.v1.EventService
grpcurl -plaintext localhost:50001 list event.v1.EventService
```


```shell

grpcurl -plaintext \
  -d '{ "id":"1", "Pilot":"Viper","ac":"F18","last_update":"9-29-23 21:29:23","squadron":"abc",fuel_status": 0.95 }' \
  localhost:50001 event.v1EventService/Insert
```