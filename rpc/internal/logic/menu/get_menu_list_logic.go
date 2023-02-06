package menu

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/menu"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuListLogic) GetMenuList(in *core.PageInfoReq) (resp *core.MenuInfoList, err error) {
	menus, err := l.svcCtx.DB.Menu.Query().Page(l.ctx, in.Page, in.PageSize, func(pager *ent.MenuPager) {
		pager.Order = ent.Asc(menu.FieldSort)
	})
	if err != nil {
		logx.Error(err.Error())
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	resp = &core.MenuInfoList{}
	for _, v := range menus.List {
		resp.Data = append(resp.Data, &core.MenuInfo{
			Id:        v.ID,
			CreatedAt: v.CreatedAt.UnixMilli(),
			UpdatedAt: v.UpdatedAt.UnixMilli(),
			MenuType:  v.MenuType,
			Level:     v.MenuLevel,
			ParentId:  v.ParentID,
			Path:      v.Path,
			Name:      v.Name,
			Redirect:  v.Redirect,
			Component: v.Component,
			Sort:      v.Sort,
			Meta: &core.Meta{
				Title:              v.Title,
				Icon:               v.Icon,
				HideMenu:           v.HideMenu,
				HideBreadcrumb:     v.HideBreadcrumb,
				CurrentActiveMenu:  v.CurrentActiveMenu,
				IgnoreKeepAlive:    v.IgnoreKeepAlive,
				HideTab:            v.HideTab,
				FrameSrc:           v.FrameSrc,
				CarryParam:         v.CarryParam,
				HideChildrenInMenu: v.HideChildrenInMenu,
				Affix:              v.Affix,
				DynamicLevel:       v.DynamicLevel,
				RealPath:           v.RealPath,
			},
		})
	}
	resp.Total = menus.PageDetails.Total
	return resp, nil
}