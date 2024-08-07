// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/ent/sysuserconfig"
)

// SysUserConfigDelete is the builder for deleting a SysUserConfig entity.
type SysUserConfigDelete struct {
	config
	hooks    []Hook
	mutation *SysUserConfigMutation
}

// Where appends a list predicates to the SysUserConfigDelete builder.
func (sucd *SysUserConfigDelete) Where(ps ...predicate.SysUserConfig) *SysUserConfigDelete {
	sucd.mutation.Where(ps...)
	return sucd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sucd *SysUserConfigDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sucd.sqlExec, sucd.mutation, sucd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sucd *SysUserConfigDelete) ExecX(ctx context.Context) int {
	n, err := sucd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sucd *SysUserConfigDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sysuserconfig.Table, sqlgraph.NewFieldSpec(sysuserconfig.FieldID, field.TypeUUID))
	if ps := sucd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sucd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sucd.mutation.done = true
	return affected, err
}

// SysUserConfigDeleteOne is the builder for deleting a single SysUserConfig entity.
type SysUserConfigDeleteOne struct {
	sucd *SysUserConfigDelete
}

// Where appends a list predicates to the SysUserConfigDelete builder.
func (sucdo *SysUserConfigDeleteOne) Where(ps ...predicate.SysUserConfig) *SysUserConfigDeleteOne {
	sucdo.sucd.mutation.Where(ps...)
	return sucdo
}

// Exec executes the deletion query.
func (sucdo *SysUserConfigDeleteOne) Exec(ctx context.Context) error {
	n, err := sucdo.sucd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sysuserconfig.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sucdo *SysUserConfigDeleteOne) ExecX(ctx context.Context) {
	if err := sucdo.Exec(ctx); err != nil {
		panic(err)
	}
}
