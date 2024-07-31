package sysuserconfig

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/sysuserconfig"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /sys_user_config/list sysuserconfig GetSysUserConfigList
//
// Get sys user config list | 获取SysUserConfig列表
//
// Get sys user config list | 获取SysUserConfig列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SysUserConfigListReq
//
// Responses:
//  200: SysUserConfigListResp

func GetSysUserConfigListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysUserConfigListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysuserconfig.NewGetSysUserConfigListLogic(r.Context(), svcCtx)
		resp, err := l.GetSysUserConfigList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
