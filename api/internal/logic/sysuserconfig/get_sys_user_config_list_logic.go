package sysuserconfig

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysUserConfigListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysUserConfigListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysUserConfigListLogic {
	return &GetSysUserConfigListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysUserConfigListLogic) GetSysUserConfigList(req *types.SysUserConfigListReq) (resp *types.SysUserConfigListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetSysUserConfigList(l.ctx,
		&core.SysUserConfigListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.SysUserConfigListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.SysUserConfigInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Anonymous:    v.Anonymous,
				ShowAnswer:   v.ShowAnswer,
				ShowAnalysis: v.ShowAnalysis,
				ShowComment:  v.ShowComment,
				Stuid:        v.Stuid,
				Institute:    v.Institute,
				EmailIsCheck: v.EmailIsCheck,
				PhoneIsCheck: v.PhoneIsCheck,
			})
	}
	return resp, nil
}
