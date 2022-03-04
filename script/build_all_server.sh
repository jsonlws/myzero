# !/bin/bash

source ./server_define.sh


#CGO_ENABLED=0  GOOS=linux  GOARCH=amd64 编译linux服务下可执行文件添加此操作

for loop in ${server_array[@]}
do 
   go build -ldflags "-s -w" -o ./../bin/${loop}_rpc_server  ./../${loop}/rpc/${loop}.go
   go build -ldflags "-s -w" -o ./../bin/${loop}_api_server  ./../${loop}/api/${loop}.go
done