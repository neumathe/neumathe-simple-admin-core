package banner

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSysBannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSysBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSysBannerLogic {
	return &CreateSysBannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSysBannerLogic) CreateSysBanner(in *core.SysBannerInfo) (*core.BaseUUIDResp, error) {
    query := l.svcCtx.DB.SysBanner.Create().
			SetNotNilURL(in.Url)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}
	if in.ShowAt != nil {
		query.SetNotNilShowAt(pointy.GetPointer(int8(*in.ShowAt)))
	}

	result, err := query.Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess }, nil
}
