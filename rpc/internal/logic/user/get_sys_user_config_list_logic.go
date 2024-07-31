package user

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent/sysuserconfig"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetSysUserConfigListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSysUserConfigListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysUserConfigListLogic {
	return &GetSysUserConfigListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSysUserConfigListLogic) GetSysUserConfigList(in *core.SysUserConfigListReq) (*core.SysUserConfigListResp, error) {
	var predicates []predicate.SysUserConfig
	result, err := l.svcCtx.DB.SysUserConfig.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.SysUserConfigListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.SysUserConfigInfo{
			Id:          pointy.GetPointer(v.ID.String()),
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Anonymous:	&v.Anonymous,
			ShowAnswer:	&v.ShowAnswer,
			ShowAnalysis:	&v.ShowAnalysis,
			ShowComment:	&v.ShowComment,
			Stuid:	&v.Stuid,
			Institute:	pointy.GetPointer(int32(v.Institute)),
			EmailIsCheck:	&v.EmailIsCheck,
			PhoneIsCheck:	&v.PhoneIsCheck,
		})
	}

	return resp, nil
}
