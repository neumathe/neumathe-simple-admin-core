package sysbanner

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/sysbanner"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /sys_banner/list sysbanner GetSysBannerList
//
// Get sys banner list | 获取SysBanner列表
//
// Get sys banner list | 获取SysBanner列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SysBannerListReq
//
// Responses:
//  200: SysBannerListResp

func GetSysBannerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysBannerListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysbanner.NewGetSysBannerListLogic(r.Context(), svcCtx)
		resp, err := l.GetSysBannerList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
