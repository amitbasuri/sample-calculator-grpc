## Sample gRPC client and server


### Example Usage
### Run Server
 ```shell
 go build  -o ./bin/server ./cmd/server
 ./bin/server
 ```
### Run Client
 ```shell
%  go build  -o ./bin/client ./cmd/client
%  ./bin/client -method add -a 5 -b 6
2022/06/11 16:55:51 res: 11
%  ./bin/client -method sub -a 5 -b 6
2022/06/11 16:56:05 res: -1
%  ./bin/client -method mul -a 2 -b 6
2022/06/11 16:56:16 res: 12
%  ./bin/client -method mul -a 22 -b 3
2022/06/11 16:56:32 res: 66
%  ./bin/client -method mul -a 2 -b 0 
2022/06/11 16:56:39 res: 0
%  ./bin/client -method div -a 2 -b 0
2022/06/11 16:56:53 rpc error: code = InvalidArgument desc = cannot divide by zero
%  ./bin/client -method div -a 20 -b 3
2022/06/11 16:57:00 res: 6
  ```

### Compile proto file
```shell 
  protoc --go_out=. \
    --go-grpc_out=. \
    api/calculator.proto
   
```
