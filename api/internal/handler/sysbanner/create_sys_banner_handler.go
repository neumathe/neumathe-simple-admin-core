package sysbanner

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/sysbanner"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /sys_banner/create sysbanner CreateSysBanner
//
// Create sys banner information | 创建SysBanner
//
// Create sys banner information | 创建SysBanner
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SysBannerInfo
//
// Responses:
//  200: BaseMsgResp

func CreateSysBannerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysBannerInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysbanner.NewCreateSysBannerLogic(r.Context(), svcCtx)
		resp, err := l.CreateSysBanner(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
