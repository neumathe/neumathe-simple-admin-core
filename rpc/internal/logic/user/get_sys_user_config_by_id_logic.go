package user

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysUserConfigByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSysUserConfigByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysUserConfigByIdLogic {
	return &GetSysUserConfigByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSysUserConfigByIdLogic) GetSysUserConfigById(in *core.UUIDReq) (*core.SysUserConfigInfo, error) {
	result, err := l.svcCtx.DB.SysUserConfig.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.SysUserConfigInfo{
		Id:          pointy.GetPointer(result.ID.String()),
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Anonymous:	&result.Anonymous,
		ShowAnswer:	&result.ShowAnswer,
		ShowAnalysis:	&result.ShowAnalysis,
		ShowComment:	&result.ShowComment,
		Stuid:	&result.Stuid,
		Institute:	pointy.GetPointer(int32(result.Institute)),
		EmailIsCheck:	&result.EmailIsCheck,
		PhoneIsCheck:	&result.PhoneIsCheck,
	}, nil
}

