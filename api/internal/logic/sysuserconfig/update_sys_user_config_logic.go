package sysuserconfig

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysUserConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysUserConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysUserConfigLogic {
	return &UpdateSysUserConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysUserConfigLogic) UpdateSysUserConfig(req *types.SysUserConfigInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.UpdateSysUserConfig(l.ctx,
		&core.SysUserConfigInfo{
			Id:           req.Id,
			Anonymous:    req.Anonymous,
			ShowAnswer:   req.ShowAnswer,
			ShowAnalysis: req.ShowAnalysis,
			ShowComment:  req.ShowComment,
			Stuid:        req.Stuid,
			Institute:    req.Institute,
			EmailIsCheck: req.EmailIsCheck,
			PhoneIsCheck: req.PhoneIsCheck,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
