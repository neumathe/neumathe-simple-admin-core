package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
	mixins2 "github.com/suyuan32/simple-admin-core/rpc/ent/schema/mixins"
)

type SysUserConfig struct {
	ent.Schema
}

func (SysUserConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("anonymous").Comment("Anonymous | 匿名 0:不开启 1:开启"),
		field.Bool("show_answer").Comment("Show Answer | 默认显示答案 0:不开启 1:开启"),
		field.Bool("show_analysis").Comment("Show Analysis | 默认显示解析 0:不开启 1:开启"),
		field.Bool("show_comment").Comment("Show Comment | 默认显示评论 0:不开启 1:开启"),
		field.Int32("stuid").Optional().Comment("Stuid | 学号"),
		field.Int16("institute").Optional().Comment("institute | 学院"),
		field.Bool("email_is_check").Comment("Email Is Check | 邮箱是否通过验证 | 0 未通过 1 通过"),
		field.Bool("phone_is_check").Comment("Phone Is Check | 手机号是否验证通过 | 0 未通过 1 通过")}
}
func (SysUserConfig) Edges() []ent.Edge {
	return nil
}

func (SysUserConfig) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "sys_user_configs"},
	}
}

func (SysUserConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins2.SoftDeleteMixin{},
	}
}
