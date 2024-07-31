package sysbanner

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysBannerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysBannerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysBannerListLogic {
	return &GetSysBannerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysBannerListLogic) GetSysBannerList(req *types.SysBannerListReq) (resp *types.SysBannerListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetSysBannerList(l.ctx,
		&core.SysBannerListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Url:      req.Url,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.SysBannerListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.SysBannerInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Status: v.Status,
				ShowAt: v.ShowAt,
				Url:    v.Url,
			})
	}
	return resp, nil
}
