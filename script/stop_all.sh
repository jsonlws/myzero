# !/bin/bash

source ./server_define.sh

for loop in ${server_array[@]}
do 
   rpcPid=`ps aux|grep ${loop}_rpc_server|grep -v grep|awk '{print $2}'`
    if [ -n "${rpcPid}" ]
    then
       kill -TERM ${rpcPid} >/dev/null 2>&1
       echo "${loop} rpc服务已停止"
    else
       echo "${loop} rpc服务未启动,无需停止"
    fi
    sleep 1
    apiPid=`ps aux|grep ${loop}_api_server|grep -v grep|awk '{print $2}'`
    if [ -n "${apiPid}" ]
    then
       kill -TERM ${apiPid} >/dev/null 2>&1
       echo "${loop} api服务已停止"
    else
       echo "${loop} api服务未启动,无需停止"
    fi
done