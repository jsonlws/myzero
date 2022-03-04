# !/bin/bash

source ./server_define.sh


for loop in ${server_array[@]}
do 
   nohup ./../bin/${loop}_rpc_server -f ./../config/${loop}_rpc.yaml >./../logs/${loop}_rpc.log 2>&1 &
   echo "${loop} rpc服务启动" 
   sleep 2
   nohup ./../bin/${loop}_api_server -f ./../config/${loop}_api.yaml >./../logs/${loop}_api.log 2>&1 &
   echo "${loop} api服务启动" 
done



