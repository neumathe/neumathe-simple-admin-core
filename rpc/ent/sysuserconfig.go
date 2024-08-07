// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-core/rpc/ent/sysuserconfig"
)

// SysUserConfig is the model entity for the SysUserConfig schema.
type SysUserConfig struct {
	config `json:"-"`
	// ID of the ent.
	// UUID
	ID uuid.UUID `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Delete Time | 删除日期
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Anonymous | 匿名 0:不开启 1:开启
	Anonymous bool `json:"anonymous,omitempty"`
	// Show Answer | 默认显示答案 0:不开启 1:开启
	ShowAnswer bool `json:"show_answer,omitempty"`
	// Show Analysis | 默认显示解析 0:不开启 1:开启
	ShowAnalysis bool `json:"show_analysis,omitempty"`
	// Show Comment | 默认显示评论 0:不开启 1:开启
	ShowComment bool `json:"show_comment,omitempty"`
	// Stuid | 学号
	Stuid int32 `json:"stuid,omitempty"`
	// institute | 学院
	Institute int16 `json:"institute,omitempty"`
	// Email Is Check | 邮箱是否通过验证 | 0 未通过 1 通过
	EmailIsCheck bool `json:"email_is_check,omitempty"`
	// Phone Is Check | 手机号是否验证通过 | 0 未通过 1 通过
	PhoneIsCheck bool `json:"phone_is_check,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysUserConfig) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysuserconfig.FieldAnonymous, sysuserconfig.FieldShowAnswer, sysuserconfig.FieldShowAnalysis, sysuserconfig.FieldShowComment, sysuserconfig.FieldEmailIsCheck, sysuserconfig.FieldPhoneIsCheck:
			values[i] = new(sql.NullBool)
		case sysuserconfig.FieldStuid, sysuserconfig.FieldInstitute:
			values[i] = new(sql.NullInt64)
		case sysuserconfig.FieldCreatedAt, sysuserconfig.FieldUpdatedAt, sysuserconfig.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case sysuserconfig.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysUserConfig fields.
func (suc *SysUserConfig) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysuserconfig.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				suc.ID = *value
			}
		case sysuserconfig.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				suc.CreatedAt = value.Time
			}
		case sysuserconfig.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				suc.UpdatedAt = value.Time
			}
		case sysuserconfig.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				suc.DeletedAt = value.Time
			}
		case sysuserconfig.FieldAnonymous:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field anonymous", values[i])
			} else if value.Valid {
				suc.Anonymous = value.Bool
			}
		case sysuserconfig.FieldShowAnswer:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field show_answer", values[i])
			} else if value.Valid {
				suc.ShowAnswer = value.Bool
			}
		case sysuserconfig.FieldShowAnalysis:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field show_analysis", values[i])
			} else if value.Valid {
				suc.ShowAnalysis = value.Bool
			}
		case sysuserconfig.FieldShowComment:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field show_comment", values[i])
			} else if value.Valid {
				suc.ShowComment = value.Bool
			}
		case sysuserconfig.FieldStuid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field stuid", values[i])
			} else if value.Valid {
				suc.Stuid = int32(value.Int64)
			}
		case sysuserconfig.FieldInstitute:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field institute", values[i])
			} else if value.Valid {
				suc.Institute = int16(value.Int64)
			}
		case sysuserconfig.FieldEmailIsCheck:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field email_is_check", values[i])
			} else if value.Valid {
				suc.EmailIsCheck = value.Bool
			}
		case sysuserconfig.FieldPhoneIsCheck:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field phone_is_check", values[i])
			} else if value.Valid {
				suc.PhoneIsCheck = value.Bool
			}
		default:
			suc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SysUserConfig.
// This includes values selected through modifiers, order, etc.
func (suc *SysUserConfig) Value(name string) (ent.Value, error) {
	return suc.selectValues.Get(name)
}

// Update returns a builder for updating this SysUserConfig.
// Note that you need to call SysUserConfig.Unwrap() before calling this method if this SysUserConfig
// was returned from a transaction, and the transaction was committed or rolled back.
func (suc *SysUserConfig) Update() *SysUserConfigUpdateOne {
	return NewSysUserConfigClient(suc.config).UpdateOne(suc)
}

// Unwrap unwraps the SysUserConfig entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (suc *SysUserConfig) Unwrap() *SysUserConfig {
	_tx, ok := suc.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysUserConfig is not a transactional entity")
	}
	suc.config.driver = _tx.drv
	return suc
}

// String implements the fmt.Stringer.
func (suc *SysUserConfig) String() string {
	var builder strings.Builder
	builder.WriteString("SysUserConfig(")
	builder.WriteString(fmt.Sprintf("id=%v, ", suc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(suc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(suc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(suc.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("anonymous=")
	builder.WriteString(fmt.Sprintf("%v", suc.Anonymous))
	builder.WriteString(", ")
	builder.WriteString("show_answer=")
	builder.WriteString(fmt.Sprintf("%v", suc.ShowAnswer))
	builder.WriteString(", ")
	builder.WriteString("show_analysis=")
	builder.WriteString(fmt.Sprintf("%v", suc.ShowAnalysis))
	builder.WriteString(", ")
	builder.WriteString("show_comment=")
	builder.WriteString(fmt.Sprintf("%v", suc.ShowComment))
	builder.WriteString(", ")
	builder.WriteString("stuid=")
	builder.WriteString(fmt.Sprintf("%v", suc.Stuid))
	builder.WriteString(", ")
	builder.WriteString("institute=")
	builder.WriteString(fmt.Sprintf("%v", suc.Institute))
	builder.WriteString(", ")
	builder.WriteString("email_is_check=")
	builder.WriteString(fmt.Sprintf("%v", suc.EmailIsCheck))
	builder.WriteString(", ")
	builder.WriteString("phone_is_check=")
	builder.WriteString(fmt.Sprintf("%v", suc.PhoneIsCheck))
	builder.WriteByte(')')
	return builder.String()
}

// SysUserConfigs is a parsable slice of SysUserConfig.
type SysUserConfigs []*SysUserConfig
