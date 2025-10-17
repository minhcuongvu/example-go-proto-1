# proto

```
protoc --version
which protoc-gen-go
which protoc-gen-go-grpc

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative calculator.proto

go mod init github.com/adrian/golang-grpc-1
go mod tidy
```
