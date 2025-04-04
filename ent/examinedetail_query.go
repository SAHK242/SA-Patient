// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"patient/ent/examinedetail"
	"patient/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamineDetailQuery is the builder for querying ExamineDetail entities.
type ExamineDetailQuery struct {
	config
	ctx        *QueryContext
	order      []examinedetail.OrderOption
	inters     []Interceptor
	predicates []predicate.ExamineDetail
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExamineDetailQuery builder.
func (edq *ExamineDetailQuery) Where(ps ...predicate.ExamineDetail) *ExamineDetailQuery {
	edq.predicates = append(edq.predicates, ps...)
	return edq
}

// Limit the number of records to be returned by this query.
func (edq *ExamineDetailQuery) Limit(limit int) *ExamineDetailQuery {
	edq.ctx.Limit = &limit
	return edq
}

// Offset to start from.
func (edq *ExamineDetailQuery) Offset(offset int) *ExamineDetailQuery {
	edq.ctx.Offset = &offset
	return edq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (edq *ExamineDetailQuery) Unique(unique bool) *ExamineDetailQuery {
	edq.ctx.Unique = &unique
	return edq
}

// Order specifies how the records should be ordered.
func (edq *ExamineDetailQuery) Order(o ...examinedetail.OrderOption) *ExamineDetailQuery {
	edq.order = append(edq.order, o...)
	return edq
}

// First returns the first ExamineDetail entity from the query.
// Returns a *NotFoundError when no ExamineDetail was found.
func (edq *ExamineDetailQuery) First(ctx context.Context) (*ExamineDetail, error) {
	nodes, err := edq.Limit(1).All(setContextOp(ctx, edq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{examinedetail.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (edq *ExamineDetailQuery) FirstX(ctx context.Context) *ExamineDetail {
	node, err := edq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ExamineDetail ID from the query.
// Returns a *NotFoundError when no ExamineDetail ID was found.
func (edq *ExamineDetailQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = edq.Limit(1).IDs(setContextOp(ctx, edq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{examinedetail.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (edq *ExamineDetailQuery) FirstIDX(ctx context.Context) int {
	id, err := edq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ExamineDetail entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ExamineDetail entity is found.
// Returns a *NotFoundError when no ExamineDetail entities are found.
func (edq *ExamineDetailQuery) Only(ctx context.Context) (*ExamineDetail, error) {
	nodes, err := edq.Limit(2).All(setContextOp(ctx, edq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{examinedetail.Label}
	default:
		return nil, &NotSingularError{examinedetail.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (edq *ExamineDetailQuery) OnlyX(ctx context.Context) *ExamineDetail {
	node, err := edq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ExamineDetail ID in the query.
// Returns a *NotSingularError when more than one ExamineDetail ID is found.
// Returns a *NotFoundError when no entities are found.
func (edq *ExamineDetailQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = edq.Limit(2).IDs(setContextOp(ctx, edq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{examinedetail.Label}
	default:
		err = &NotSingularError{examinedetail.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (edq *ExamineDetailQuery) OnlyIDX(ctx context.Context) int {
	id, err := edq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ExamineDetails.
func (edq *ExamineDetailQuery) All(ctx context.Context) ([]*ExamineDetail, error) {
	ctx = setContextOp(ctx, edq.ctx, ent.OpQueryAll)
	if err := edq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ExamineDetail, *ExamineDetailQuery]()
	return withInterceptors[[]*ExamineDetail](ctx, edq, qr, edq.inters)
}

// AllX is like All, but panics if an error occurs.
func (edq *ExamineDetailQuery) AllX(ctx context.Context) []*ExamineDetail {
	nodes, err := edq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ExamineDetail IDs.
func (edq *ExamineDetailQuery) IDs(ctx context.Context) (ids []int, err error) {
	if edq.ctx.Unique == nil && edq.path != nil {
		edq.Unique(true)
	}
	ctx = setContextOp(ctx, edq.ctx, ent.OpQueryIDs)
	if err = edq.Select(examinedetail.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (edq *ExamineDetailQuery) IDsX(ctx context.Context) []int {
	ids, err := edq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (edq *ExamineDetailQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, edq.ctx, ent.OpQueryCount)
	if err := edq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, edq, querierCount[*ExamineDetailQuery](), edq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (edq *ExamineDetailQuery) CountX(ctx context.Context) int {
	count, err := edq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (edq *ExamineDetailQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, edq.ctx, ent.OpQueryExist)
	switch _, err := edq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (edq *ExamineDetailQuery) ExistX(ctx context.Context) bool {
	exist, err := edq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExamineDetailQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (edq *ExamineDetailQuery) Clone() *ExamineDetailQuery {
	if edq == nil {
		return nil
	}
	return &ExamineDetailQuery{
		config:     edq.config,
		ctx:        edq.ctx.Clone(),
		order:      append([]examinedetail.OrderOption{}, edq.order...),
		inters:     append([]Interceptor{}, edq.inters...),
		predicates: append([]predicate.ExamineDetail{}, edq.predicates...),
		// clone intermediate query.
		sql:  edq.sql.Clone(),
		path: edq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (edq *ExamineDetailQuery) GroupBy(field string, fields ...string) *ExamineDetailGroupBy {
	edq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ExamineDetailGroupBy{build: edq}
	grbuild.flds = &edq.ctx.Fields
	grbuild.label = examinedetail.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (edq *ExamineDetailQuery) Select(fields ...string) *ExamineDetailSelect {
	edq.ctx.Fields = append(edq.ctx.Fields, fields...)
	sbuild := &ExamineDetailSelect{ExamineDetailQuery: edq}
	sbuild.label = examinedetail.Label
	sbuild.flds, sbuild.scan = &edq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ExamineDetailSelect configured with the given aggregations.
func (edq *ExamineDetailQuery) Aggregate(fns ...AggregateFunc) *ExamineDetailSelect {
	return edq.Select().Aggregate(fns...)
}

func (edq *ExamineDetailQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range edq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, edq); err != nil {
				return err
			}
		}
	}
	for _, f := range edq.ctx.Fields {
		if !examinedetail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if edq.path != nil {
		prev, err := edq.path(ctx)
		if err != nil {
			return err
		}
		edq.sql = prev
	}
	return nil
}

func (edq *ExamineDetailQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ExamineDetail, error) {
	var (
		nodes = []*ExamineDetail{}
		_spec = edq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ExamineDetail).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ExamineDetail{config: edq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, edq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (edq *ExamineDetailQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := edq.querySpec()
	_spec.Node.Columns = edq.ctx.Fields
	if len(edq.ctx.Fields) > 0 {
		_spec.Unique = edq.ctx.Unique != nil && *edq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, edq.driver, _spec)
}

func (edq *ExamineDetailQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(examinedetail.Table, examinedetail.Columns, sqlgraph.NewFieldSpec(examinedetail.FieldID, field.TypeInt))
	_spec.From = edq.sql
	if unique := edq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if edq.path != nil {
		_spec.Unique = true
	}
	if fields := edq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, examinedetail.FieldID)
		for i := range fields {
			if fields[i] != examinedetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := edq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := edq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := edq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := edq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (edq *ExamineDetailQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(edq.driver.Dialect())
	t1 := builder.Table(examinedetail.Table)
	columns := edq.ctx.Fields
	if len(columns) == 0 {
		columns = examinedetail.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if edq.sql != nil {
		selector = edq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if edq.ctx.Unique != nil && *edq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range edq.predicates {
		p(selector)
	}
	for _, p := range edq.order {
		p(selector)
	}
	if offset := edq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := edq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ExamineDetailGroupBy is the group-by builder for ExamineDetail entities.
type ExamineDetailGroupBy struct {
	selector
	build *ExamineDetailQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (edgb *ExamineDetailGroupBy) Aggregate(fns ...AggregateFunc) *ExamineDetailGroupBy {
	edgb.fns = append(edgb.fns, fns...)
	return edgb
}

// Scan applies the selector query and scans the result into the given value.
func (edgb *ExamineDetailGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, edgb.build.ctx, ent.OpQueryGroupBy)
	if err := edgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExamineDetailQuery, *ExamineDetailGroupBy](ctx, edgb.build, edgb, edgb.build.inters, v)
}

func (edgb *ExamineDetailGroupBy) sqlScan(ctx context.Context, root *ExamineDetailQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(edgb.fns))
	for _, fn := range edgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*edgb.flds)+len(edgb.fns))
		for _, f := range *edgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*edgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := edgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ExamineDetailSelect is the builder for selecting fields of ExamineDetail entities.
type ExamineDetailSelect struct {
	*ExamineDetailQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (eds *ExamineDetailSelect) Aggregate(fns ...AggregateFunc) *ExamineDetailSelect {
	eds.fns = append(eds.fns, fns...)
	return eds
}

// Scan applies the selector query and scans the result into the given value.
func (eds *ExamineDetailSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eds.ctx, ent.OpQuerySelect)
	if err := eds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExamineDetailQuery, *ExamineDetailSelect](ctx, eds.ExamineDetailQuery, eds, eds.inters, v)
}

func (eds *ExamineDetailSelect) sqlScan(ctx context.Context, root *ExamineDetailQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(eds.fns))
	for _, fn := range eds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*eds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
