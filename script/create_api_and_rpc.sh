# !/bin/bash

#执行生成api和rpc服务模版文件

server_array=(
    user
)

for loop in ${server_array[@]}
do 
   goctl api go -api  ./../${loop}/api/${loop}.api -dir ./../${loop}/api
   goctl rpc proto -src ./../${loop}/rpc/${loop}.proto -dir ./../${loop}/rpc
done

