package sysbanner

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysBannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysBannerLogic {
	return &UpdateSysBannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysBannerLogic) UpdateSysBanner(req *types.SysBannerInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.UpdateSysBanner(l.ctx,
		&core.SysBannerInfo{
			Id:     req.Id,
			Status: req.Status,
			ShowAt: req.ShowAt,
			Url:    req.Url,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
