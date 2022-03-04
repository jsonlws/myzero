开发语言
go  1.17
shell

使用前提要对微服务有所了解，shell编程

第三方服务
etcd  mysql redis

微服务框架
go-zero

官方文档 https://go-zero.dev/cn/

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
项目启动
进入script 目录

chmod +x *sh

server_define.sh  定义服务文件
build_all_server.sh  构建所有服务的脚步文件
start_all.sh 启动所有服务的脚步文件
stop_all.sh 关闭所有服务的脚本文件
/////////////////////////////////////////////////////////////////////////////////////////////////////////////
//特殊文件说明
create_api_and_rpc.sh 构建所有api和rpc模版的脚本文件【注意，已经好的服务切勿重新构建否则会覆盖已经写好的文件,只能在开发阶段使用】
需提前创建好目录
xxx
   ---rpc
      xxx.proto
   ---api
      xxx.api
目录名如上面结构所示
xxx代表的就是服务的名字      
需要进入脚本在server_array中定义服务的名字
完成上面操作就执行
 ./create_api_and_rpc.sh 
 就可以在目录中看自动生成的文件
 然后就可以开启微服务开发了
开发规范参照官方

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
config目录放的是服务的配置项
创建规则xxx_api.yaml   xxx_rpc.yaml 
配置内容直接复制对应服务etc中的内容

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
lib包存放项目开发公共开发包

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
为了api友好提示需要注意修改一下自动生成的文件
xxx/api/internal/handler/xxxhandler.go
//实例如下
func xxxHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, lib.DoErrorResponse{
				Code: lib.ParamsError,
				Msg:  err.Error(),
			})
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.OkJson(w, lib.DoErrorResponse{
				Code: lib.Fail,
				Msg:  err.Error(),
			})
		} else {
			httpx.OkJson(w, lib.DoSuccessResponse{
				Code: lib.Success,
				Msg:  "success",
				Data: resp,
			})
		}
	}
}
