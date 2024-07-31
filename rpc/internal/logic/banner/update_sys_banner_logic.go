package banner

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysBannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSysBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysBannerLogic {
	return &UpdateSysBannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSysBannerLogic) UpdateSysBanner(in *core.SysBannerInfo) (*core.BaseResp, error) {
	query:= l.svcCtx.DB.SysBanner.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
			SetNotNilURL(in.Url)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}
	if in.ShowAt != nil {
		query.SetNotNilShowAt(pointy.GetPointer(int8(*in.ShowAt)))
	}

	 err := query.Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.UpdateSuccess }, nil
}
