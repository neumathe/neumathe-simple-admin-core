package banner

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent/sysbanner"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetSysBannerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSysBannerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysBannerListLogic {
	return &GetSysBannerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSysBannerListLogic) GetSysBannerList(in *core.SysBannerListReq) (*core.SysBannerListResp, error) {
	var predicates []predicate.SysBanner
	if in.Url != nil {
		predicates = append(predicates, sysbanner.URLContains(*in.Url))
	}
	result, err := l.svcCtx.DB.SysBanner.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.SysBannerListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.SysBannerInfo{
			Id:          pointy.GetPointer(v.ID.String()),
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:	pointy.GetPointer(uint32(v.Status)),
			ShowAt:	pointy.GetPointer(int32(v.ShowAt)),
			Url:	&v.URL,
		})
	}

	return resp, nil
}
