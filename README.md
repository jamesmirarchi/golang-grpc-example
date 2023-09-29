# grpc-go

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