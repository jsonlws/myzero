package lib

/* 业务code码定义 */
const (
	Success     = 200
	Fail        = 400
	ParamsError = 1002
)

/* 定义全局返回参数格式 */

//成功
type DoSuccessResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//失败
type DoErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
