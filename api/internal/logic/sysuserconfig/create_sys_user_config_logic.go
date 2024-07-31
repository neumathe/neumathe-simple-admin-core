package sysuserconfig

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSysUserConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSysUserConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSysUserConfigLogic {
	return &CreateSysUserConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSysUserConfigLogic) CreateSysUserConfig(req *types.SysUserConfigInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.CreateSysUserConfig(l.ctx,
		&core.SysUserConfigInfo{
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
