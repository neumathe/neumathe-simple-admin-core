package user

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

type UpdateSysUserConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSysUserConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysUserConfigLogic {
	return &UpdateSysUserConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSysUserConfigLogic) UpdateSysUserConfig(in *core.SysUserConfigInfo) (*core.BaseResp, error) {
	query:= l.svcCtx.DB.SysUserConfig.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
			SetNotNilAnonymous(in.Anonymous).
			SetNotNilShowAnswer(in.ShowAnswer).
			SetNotNilShowAnalysis(in.ShowAnalysis).
			SetNotNilShowComment(in.ShowComment).
			SetNotNilStuid(in.Stuid).
			SetNotNilEmailIsCheck(in.EmailIsCheck).
			SetNotNilPhoneIsCheck(in.PhoneIsCheck)

	if in.Institute != nil {
		query.SetNotNilInstitute(pointy.GetPointer(int16(*in.Institute)))
	}

	 err := query.Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.UpdateSuccess }, nil
}
