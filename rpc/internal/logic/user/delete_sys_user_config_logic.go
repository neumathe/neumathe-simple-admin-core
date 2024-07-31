package user

import (
	"context"

    "github.com/suyuan32/simple-admin-core/rpc/ent/sysuserconfig"
    "github.com/suyuan32/simple-admin-core/rpc/internal/svc"
    "github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
    "github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/suyuan32/simple-admin-common/utils/uuidx"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysUserConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSysUserConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysUserConfigLogic {
	return &DeleteSysUserConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSysUserConfigLogic) DeleteSysUserConfig(in *core.UUIDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.SysUserConfig.Delete().Where(sysuserconfig.IDIn(uuidx.ParseUUIDSlice(in.Ids)...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
