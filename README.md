# proto

```
protoc --version
which protoc-gen-go
which protoc-gen-go-grpc

# Output generated .pb.go and _grpc.pb.go files to ./pb_gen/
mkdir -p pb_go
protoc --go_out=pb_gen --go_opt=paths=source_relative \
       --go-grpc_out=pb_gen --go-grpc_opt=paths=source_relative \
       calculator.proto

go mod init github.com/adrian/golang-grpc-1
go mod tidy
```
