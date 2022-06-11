## Sample gRPC client and server

## System Requirements

#### Building the App
[Go](https://go.dev/doc/install)

#### Compile Protobuf
[protoc](https://grpc.io/docs/protoc-installation/)

## Example Usage
### Run Server
 ```shell
 go build  -o ./bin/server ./cmd/server
 ./bin/server
 ```
### Run Client
 ```shell
%  go build  -o ./bin/client ./cmd/client
%  ./bin/client -method add -a 5 -b 6
2022/06/11 16:55:51 result: 11
%  ./bin/client -method sub -a 5 -b 6
2022/06/11 16:56:05 result: -1
%  ./bin/client -method mul -a 2 -b 6
2022/06/11 16:56:16 result: 12
%  ./bin/client -method mul -a 22 -b 3
2022/06/11 16:56:32 result: 66
%  ./bin/client -method mul -a 2 -b 0 
2022/06/11 16:56:39 result: 0
%  ./bin/client -method div -a 2 -b 0
2022/06/11 16:56:53 rpc error: code = InvalidArgument desc = cannot divide by zero
%  ./bin/client -method div -a 20 -b 3
2022/06/11 16:57:00 result: 6
  ```

### Compile proto file
```shell 
  protoc --go_out=. \
    --go-grpc_out=. \
    api/calculator.proto
   
```
