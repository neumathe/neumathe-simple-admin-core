package sysuserconfig

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysUserConfigByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysUserConfigByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysUserConfigByIdLogic {
	return &GetSysUserConfigByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysUserConfigByIdLogic) GetSysUserConfigById(req *types.UUIDReq) (resp *types.SysUserConfigInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetSysUserConfigById(l.ctx, &core.UUIDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.SysUserConfigInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.SysUserConfigInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Anonymous:    data.Anonymous,
			ShowAnswer:   data.ShowAnswer,
			ShowAnalysis: data.ShowAnalysis,
			ShowComment:  data.ShowComment,
			Stuid:        data.Stuid,
			Institute:    data.Institute,
			EmailIsCheck: data.EmailIsCheck,
			PhoneIsCheck: data.PhoneIsCheck,
		},
	}, nil
}
