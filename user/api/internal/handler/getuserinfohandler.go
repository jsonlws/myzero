package handler

import (
	"net/http"

	"lesjp/lib"
	"lesjp/user/api/internal/logic"
	"lesjp/user/api/internal/svc"
	"lesjp/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func getUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, lib.DoErrorResponse{
				Code: lib.ParamsError,
				Msg:  err.Error(),
			})
			return
		}

		l := logic.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo(&req)
		if err != nil {
			httpx.OkJson(w, lib.DoErrorResponse{
				Code: lib.ParamsError,
				Msg:  err.Error(),
			})
		} else {
			httpx.OkJson(w, lib.DoSuccessResponse{
				Code: lib.Success,
				Msg:  "获取数据成功",
				Data: resp,
			})
		}
	}
}
