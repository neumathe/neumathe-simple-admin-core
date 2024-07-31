package user

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSysUserConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSysUserConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSysUserConfigLogic {
	return &CreateSysUserConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSysUserConfigLogic) CreateSysUserConfig(in *core.SysUserConfigInfo) (*core.BaseUUIDResp, error) {
    query := l.svcCtx.DB.SysUserConfig.Create().
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

	result, err := query.Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess }, nil
}
