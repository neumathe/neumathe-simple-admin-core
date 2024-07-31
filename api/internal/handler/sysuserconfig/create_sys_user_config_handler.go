package sysuserconfig

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/sysuserconfig"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /sys_user_config/create sysuserconfig CreateSysUserConfig
//
// Create sys user config information | 创建SysUserConfig
//
// Create sys user config information | 创建SysUserConfig
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SysUserConfigInfo
//
// Responses:
//  200: BaseMsgResp

func CreateSysUserConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysUserConfigInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysuserconfig.NewCreateSysUserConfigLogic(r.Context(), svcCtx)
		resp, err := l.CreateSysUserConfig(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
