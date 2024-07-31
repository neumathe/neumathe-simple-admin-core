package sysuserconfig

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/sysuserconfig"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /sys_user_config/update sysuserconfig UpdateSysUserConfig
//
// Update sys user config information | 更新SysUserConfig
//
// Update sys user config information | 更新SysUserConfig
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SysUserConfigInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateSysUserConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysUserConfigInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysuserconfig.NewUpdateSysUserConfigLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSysUserConfig(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
