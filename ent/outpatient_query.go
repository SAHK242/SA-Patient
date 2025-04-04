// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"patient/ent/outpatient"
	"patient/ent/patient"
	"patient/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OutpatientQuery is the builder for querying Outpatient entities.
type OutpatientQuery struct {
	config
	ctx         *QueryContext
	order       []outpatient.OrderOption
	inters      []Interceptor
	predicates  []predicate.Outpatient
	withPatient *PatientQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutpatientQuery builder.
func (oq *OutpatientQuery) Where(ps ...predicate.Outpatient) *OutpatientQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit the number of records to be returned by this query.
func (oq *OutpatientQuery) Limit(limit int) *OutpatientQuery {
	oq.ctx.Limit = &limit
	return oq
}

// Offset to start from.
func (oq *OutpatientQuery) Offset(offset int) *OutpatientQuery {
	oq.ctx.Offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *OutpatientQuery) Unique(unique bool) *OutpatientQuery {
	oq.ctx.Unique = &unique
	return oq
}

// Order specifies how the records should be ordered.
func (oq *OutpatientQuery) Order(o ...outpatient.OrderOption) *OutpatientQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// QueryPatient chains the current query on the "patient" edge.
func (oq *OutpatientQuery) QueryPatient() *PatientQuery {
	query := (&PatientClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outpatient.Table, outpatient.FieldID, selector),
			sqlgraph.To(patient.Table, patient.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outpatient.PatientTable, outpatient.PatientColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Outpatient entity from the query.
// Returns a *NotFoundError when no Outpatient was found.
func (oq *OutpatientQuery) First(ctx context.Context) (*Outpatient, error) {
	nodes, err := oq.Limit(1).All(setContextOp(ctx, oq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outpatient.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *OutpatientQuery) FirstX(ctx context.Context) *Outpatient {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Outpatient ID from the query.
// Returns a *NotFoundError when no Outpatient ID was found.
func (oq *OutpatientQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = oq.Limit(1).IDs(setContextOp(ctx, oq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outpatient.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *OutpatientQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Outpatient entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Outpatient entity is found.
// Returns a *NotFoundError when no Outpatient entities are found.
func (oq *OutpatientQuery) Only(ctx context.Context) (*Outpatient, error) {
	nodes, err := oq.Limit(2).All(setContextOp(ctx, oq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outpatient.Label}
	default:
		return nil, &NotSingularError{outpatient.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *OutpatientQuery) OnlyX(ctx context.Context) *Outpatient {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Outpatient ID in the query.
// Returns a *NotSingularError when more than one Outpatient ID is found.
// Returns a *NotFoundError when no entities are found.
func (oq *OutpatientQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = oq.Limit(2).IDs(setContextOp(ctx, oq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outpatient.Label}
	default:
		err = &NotSingularError{outpatient.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *OutpatientQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Outpatients.
func (oq *OutpatientQuery) All(ctx context.Context) ([]*Outpatient, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryAll)
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Outpatient, *OutpatientQuery]()
	return withInterceptors[[]*Outpatient](ctx, oq, qr, oq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oq *OutpatientQuery) AllX(ctx context.Context) []*Outpatient {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Outpatient IDs.
func (oq *OutpatientQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if oq.ctx.Unique == nil && oq.path != nil {
		oq.Unique(true)
	}
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryIDs)
	if err = oq.Select(outpatient.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *OutpatientQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *OutpatientQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryCount)
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oq, querierCount[*OutpatientQuery](), oq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oq *OutpatientQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *OutpatientQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryExist)
	switch _, err := oq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *OutpatientQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutpatientQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *OutpatientQuery) Clone() *OutpatientQuery {
	if oq == nil {
		return nil
	}
	return &OutpatientQuery{
		config:      oq.config,
		ctx:         oq.ctx.Clone(),
		order:       append([]outpatient.OrderOption{}, oq.order...),
		inters:      append([]Interceptor{}, oq.inters...),
		predicates:  append([]predicate.Outpatient{}, oq.predicates...),
		withPatient: oq.withPatient.Clone(),
		// clone intermediate query.
		sql:  oq.sql.Clone(),
		path: oq.path,
	}
}

// WithPatient tells the query-builder to eager-load the nodes that are connected to
// the "patient" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OutpatientQuery) WithPatient(opts ...func(*PatientQuery)) *OutpatientQuery {
	query := (&PatientClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withPatient = query
	return oq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PatientID uuid.UUID `json:"patient_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Outpatient.Query().
//		GroupBy(outpatient.FieldPatientID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oq *OutpatientQuery) GroupBy(field string, fields ...string) *OutpatientGroupBy {
	oq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OutpatientGroupBy{build: oq}
	grbuild.flds = &oq.ctx.Fields
	grbuild.label = outpatient.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PatientID uuid.UUID `json:"patient_id,omitempty"`
//	}
//
//	client.Outpatient.Query().
//		Select(outpatient.FieldPatientID).
//		Scan(ctx, &v)
func (oq *OutpatientQuery) Select(fields ...string) *OutpatientSelect {
	oq.ctx.Fields = append(oq.ctx.Fields, fields...)
	sbuild := &OutpatientSelect{OutpatientQuery: oq}
	sbuild.label = outpatient.Label
	sbuild.flds, sbuild.scan = &oq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OutpatientSelect configured with the given aggregations.
func (oq *OutpatientQuery) Aggregate(fns ...AggregateFunc) *OutpatientSelect {
	return oq.Select().Aggregate(fns...)
}

func (oq *OutpatientQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oq); err != nil {
				return err
			}
		}
	}
	for _, f := range oq.ctx.Fields {
		if !outpatient.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *OutpatientQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Outpatient, error) {
	var (
		nodes       = []*Outpatient{}
		_spec       = oq.querySpec()
		loadedTypes = [1]bool{
			oq.withPatient != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Outpatient).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Outpatient{config: oq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oq.withPatient; query != nil {
		if err := oq.loadPatient(ctx, query, nodes, nil,
			func(n *Outpatient, e *Patient) { n.Edges.Patient = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oq *OutpatientQuery) loadPatient(ctx context.Context, query *PatientQuery, nodes []*Outpatient, init func(*Outpatient), assign func(*Outpatient, *Patient)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Outpatient)
	for i := range nodes {
		fk := nodes[i].PatientID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(patient.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "patient_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (oq *OutpatientQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	_spec.Node.Columns = oq.ctx.Fields
	if len(oq.ctx.Fields) > 0 {
		_spec.Unique = oq.ctx.Unique != nil && *oq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *OutpatientQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(outpatient.Table, outpatient.Columns, sqlgraph.NewFieldSpec(outpatient.FieldID, field.TypeUUID))
	_spec.From = oq.sql
	if unique := oq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oq.path != nil {
		_spec.Unique = true
	}
	if fields := oq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outpatient.FieldID)
		for i := range fields {
			if fields[i] != outpatient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if oq.withPatient != nil {
			_spec.Node.AddColumnOnce(outpatient.FieldPatientID)
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *OutpatientQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(outpatient.Table)
	columns := oq.ctx.Fields
	if len(columns) == 0 {
		columns = outpatient.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oq.ctx.Unique != nil && *oq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OutpatientGroupBy is the group-by builder for Outpatient entities.
type OutpatientGroupBy struct {
	selector
	build *OutpatientQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *OutpatientGroupBy) Aggregate(fns ...AggregateFunc) *OutpatientGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the selector query and scans the result into the given value.
func (ogb *OutpatientGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ogb.build.ctx, ent.OpQueryGroupBy)
	if err := ogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OutpatientQuery, *OutpatientGroupBy](ctx, ogb.build, ogb, ogb.build.inters, v)
}

func (ogb *OutpatientGroupBy) sqlScan(ctx context.Context, root *OutpatientQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ogb.flds)+len(ogb.fns))
		for _, f := range *ogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OutpatientSelect is the builder for selecting fields of Outpatient entities.
type OutpatientSelect struct {
	*OutpatientQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (os *OutpatientSelect) Aggregate(fns ...AggregateFunc) *OutpatientSelect {
	os.fns = append(os.fns, fns...)
	return os
}

// Scan applies the selector query and scans the result into the given value.
func (os *OutpatientSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, os.ctx, ent.OpQuerySelect)
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OutpatientQuery, *OutpatientSelect](ctx, os.OutpatientQuery, os, os.inters, v)
}

func (os *OutpatientSelect) sqlScan(ctx context.Context, root *OutpatientQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(os.fns))
	for _, fn := range os.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*os.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
