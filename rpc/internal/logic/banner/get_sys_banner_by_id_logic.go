package banner

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysBannerByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSysBannerByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysBannerByIdLogic {
	return &GetSysBannerByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSysBannerByIdLogic) GetSysBannerById(in *core.UUIDReq) (*core.SysBannerInfo, error) {
	result, err := l.svcCtx.DB.SysBanner.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.SysBannerInfo{
		Id:          pointy.GetPointer(result.ID.String()),
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Status:	pointy.GetPointer(uint32(result.Status)),
		ShowAt:	pointy.GetPointer(int32(result.ShowAt)),
		Url:	&result.URL,
	}, nil
}

