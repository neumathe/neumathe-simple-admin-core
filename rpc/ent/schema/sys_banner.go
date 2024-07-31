package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type SysBanner struct {
	ent.Schema
}

func (SysBanner) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("show_at").Comment("Show At | 1: web| 2:miniprogram | 展示位置 1网页 2小程序"),
		field.String("url").Comment("Url | 图片链接")}
}
func (SysBanner) Edges() []ent.Edge {
	return nil
}

func (SysBanner) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "sys_banners"},
	}
}

func (SysBanner) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
	}
}
