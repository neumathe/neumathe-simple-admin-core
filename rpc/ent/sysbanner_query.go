// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/ent/sysbanner"
)

// SysBannerQuery is the builder for querying SysBanner entities.
type SysBannerQuery struct {
	config
	ctx        *QueryContext
	order      []sysbanner.OrderOption
	inters     []Interceptor
	predicates []predicate.SysBanner
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysBannerQuery builder.
func (sbq *SysBannerQuery) Where(ps ...predicate.SysBanner) *SysBannerQuery {
	sbq.predicates = append(sbq.predicates, ps...)
	return sbq
}

// Limit the number of records to be returned by this query.
func (sbq *SysBannerQuery) Limit(limit int) *SysBannerQuery {
	sbq.ctx.Limit = &limit
	return sbq
}

// Offset to start from.
func (sbq *SysBannerQuery) Offset(offset int) *SysBannerQuery {
	sbq.ctx.Offset = &offset
	return sbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sbq *SysBannerQuery) Unique(unique bool) *SysBannerQuery {
	sbq.ctx.Unique = &unique
	return sbq
}

// Order specifies how the records should be ordered.
func (sbq *SysBannerQuery) Order(o ...sysbanner.OrderOption) *SysBannerQuery {
	sbq.order = append(sbq.order, o...)
	return sbq
}

// First returns the first SysBanner entity from the query.
// Returns a *NotFoundError when no SysBanner was found.
func (sbq *SysBannerQuery) First(ctx context.Context) (*SysBanner, error) {
	nodes, err := sbq.Limit(1).All(setContextOp(ctx, sbq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysbanner.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sbq *SysBannerQuery) FirstX(ctx context.Context) *SysBanner {
	node, err := sbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysBanner ID from the query.
// Returns a *NotFoundError when no SysBanner ID was found.
func (sbq *SysBannerQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sbq.Limit(1).IDs(setContextOp(ctx, sbq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysbanner.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sbq *SysBannerQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := sbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysBanner entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysBanner entity is found.
// Returns a *NotFoundError when no SysBanner entities are found.
func (sbq *SysBannerQuery) Only(ctx context.Context) (*SysBanner, error) {
	nodes, err := sbq.Limit(2).All(setContextOp(ctx, sbq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysbanner.Label}
	default:
		return nil, &NotSingularError{sysbanner.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sbq *SysBannerQuery) OnlyX(ctx context.Context) *SysBanner {
	node, err := sbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysBanner ID in the query.
// Returns a *NotSingularError when more than one SysBanner ID is found.
// Returns a *NotFoundError when no entities are found.
func (sbq *SysBannerQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sbq.Limit(2).IDs(setContextOp(ctx, sbq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysbanner.Label}
	default:
		err = &NotSingularError{sysbanner.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sbq *SysBannerQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := sbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysBanners.
func (sbq *SysBannerQuery) All(ctx context.Context) ([]*SysBanner, error) {
	ctx = setContextOp(ctx, sbq.ctx, "All")
	if err := sbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SysBanner, *SysBannerQuery]()
	return withInterceptors[[]*SysBanner](ctx, sbq, qr, sbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sbq *SysBannerQuery) AllX(ctx context.Context) []*SysBanner {
	nodes, err := sbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysBanner IDs.
func (sbq *SysBannerQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if sbq.ctx.Unique == nil && sbq.path != nil {
		sbq.Unique(true)
	}
	ctx = setContextOp(ctx, sbq.ctx, "IDs")
	if err = sbq.Select(sysbanner.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sbq *SysBannerQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := sbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sbq *SysBannerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sbq.ctx, "Count")
	if err := sbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sbq, querierCount[*SysBannerQuery](), sbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sbq *SysBannerQuery) CountX(ctx context.Context) int {
	count, err := sbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sbq *SysBannerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sbq.ctx, "Exist")
	switch _, err := sbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sbq *SysBannerQuery) ExistX(ctx context.Context) bool {
	exist, err := sbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysBannerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sbq *SysBannerQuery) Clone() *SysBannerQuery {
	if sbq == nil {
		return nil
	}
	return &SysBannerQuery{
		config:     sbq.config,
		ctx:        sbq.ctx.Clone(),
		order:      append([]sysbanner.OrderOption{}, sbq.order...),
		inters:     append([]Interceptor{}, sbq.inters...),
		predicates: append([]predicate.SysBanner{}, sbq.predicates...),
		// clone intermediate query.
		sql:  sbq.sql.Clone(),
		path: sbq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SysBanner.Query().
//		GroupBy(sysbanner.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sbq *SysBannerQuery) GroupBy(field string, fields ...string) *SysBannerGroupBy {
	sbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SysBannerGroupBy{build: sbq}
	grbuild.flds = &sbq.ctx.Fields
	grbuild.label = sysbanner.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.SysBanner.Query().
//		Select(sysbanner.FieldCreatedAt).
//		Scan(ctx, &v)
func (sbq *SysBannerQuery) Select(fields ...string) *SysBannerSelect {
	sbq.ctx.Fields = append(sbq.ctx.Fields, fields...)
	sbuild := &SysBannerSelect{SysBannerQuery: sbq}
	sbuild.label = sysbanner.Label
	sbuild.flds, sbuild.scan = &sbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SysBannerSelect configured with the given aggregations.
func (sbq *SysBannerQuery) Aggregate(fns ...AggregateFunc) *SysBannerSelect {
	return sbq.Select().Aggregate(fns...)
}

func (sbq *SysBannerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sbq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sbq); err != nil {
				return err
			}
		}
	}
	for _, f := range sbq.ctx.Fields {
		if !sysbanner.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sbq.path != nil {
		prev, err := sbq.path(ctx)
		if err != nil {
			return err
		}
		sbq.sql = prev
	}
	return nil
}

func (sbq *SysBannerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysBanner, error) {
	var (
		nodes = []*SysBanner{}
		_spec = sbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysBanner).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysBanner{config: sbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(sbq.modifiers) > 0 {
		_spec.Modifiers = sbq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sbq *SysBannerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sbq.querySpec()
	if len(sbq.modifiers) > 0 {
		_spec.Modifiers = sbq.modifiers
	}
	_spec.Node.Columns = sbq.ctx.Fields
	if len(sbq.ctx.Fields) > 0 {
		_spec.Unique = sbq.ctx.Unique != nil && *sbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sbq.driver, _spec)
}

func (sbq *SysBannerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sysbanner.Table, sysbanner.Columns, sqlgraph.NewFieldSpec(sysbanner.FieldID, field.TypeUUID))
	_spec.From = sbq.sql
	if unique := sbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sbq.path != nil {
		_spec.Unique = true
	}
	if fields := sbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysbanner.FieldID)
		for i := range fields {
			if fields[i] != sysbanner.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sbq *SysBannerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sbq.driver.Dialect())
	t1 := builder.Table(sysbanner.Table)
	columns := sbq.ctx.Fields
	if len(columns) == 0 {
		columns = sysbanner.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sbq.sql != nil {
		selector = sbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sbq.ctx.Unique != nil && *sbq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range sbq.modifiers {
		m(selector)
	}
	for _, p := range sbq.predicates {
		p(selector)
	}
	for _, p := range sbq.order {
		p(selector)
	}
	if offset := sbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sbq *SysBannerQuery) Modify(modifiers ...func(s *sql.Selector)) *SysBannerSelect {
	sbq.modifiers = append(sbq.modifiers, modifiers...)
	return sbq.Select()
}

// SysBannerGroupBy is the group-by builder for SysBanner entities.
type SysBannerGroupBy struct {
	selector
	build *SysBannerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sbgb *SysBannerGroupBy) Aggregate(fns ...AggregateFunc) *SysBannerGroupBy {
	sbgb.fns = append(sbgb.fns, fns...)
	return sbgb
}

// Scan applies the selector query and scans the result into the given value.
func (sbgb *SysBannerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sbgb.build.ctx, "GroupBy")
	if err := sbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysBannerQuery, *SysBannerGroupBy](ctx, sbgb.build, sbgb, sbgb.build.inters, v)
}

func (sbgb *SysBannerGroupBy) sqlScan(ctx context.Context, root *SysBannerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sbgb.fns))
	for _, fn := range sbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sbgb.flds)+len(sbgb.fns))
		for _, f := range *sbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SysBannerSelect is the builder for selecting fields of SysBanner entities.
type SysBannerSelect struct {
	*SysBannerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sbs *SysBannerSelect) Aggregate(fns ...AggregateFunc) *SysBannerSelect {
	sbs.fns = append(sbs.fns, fns...)
	return sbs
}

// Scan applies the selector query and scans the result into the given value.
func (sbs *SysBannerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sbs.ctx, "Select")
	if err := sbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysBannerQuery, *SysBannerSelect](ctx, sbs.SysBannerQuery, sbs, sbs.inters, v)
}

func (sbs *SysBannerSelect) sqlScan(ctx context.Context, root *SysBannerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sbs.fns))
	for _, fn := range sbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sbs *SysBannerSelect) Modify(modifiers ...func(s *sql.Selector)) *SysBannerSelect {
	sbs.modifiers = append(sbs.modifiers, modifiers...)
	return sbs
}
