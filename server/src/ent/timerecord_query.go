// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"oruka/ent/predicate"
	"oruka/ent/timerecord"
	"oruka/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TimeRecordQuery is the builder for querying TimeRecord entities.
type TimeRecordQuery struct {
	config
	ctx            *QueryContext
	order          []OrderFunc
	inters         []Interceptor
	predicates     []predicate.TimeRecord
	withTimeKeeper *UserQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TimeRecordQuery builder.
func (trq *TimeRecordQuery) Where(ps ...predicate.TimeRecord) *TimeRecordQuery {
	trq.predicates = append(trq.predicates, ps...)
	return trq
}

// Limit the number of records to be returned by this query.
func (trq *TimeRecordQuery) Limit(limit int) *TimeRecordQuery {
	trq.ctx.Limit = &limit
	return trq
}

// Offset to start from.
func (trq *TimeRecordQuery) Offset(offset int) *TimeRecordQuery {
	trq.ctx.Offset = &offset
	return trq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (trq *TimeRecordQuery) Unique(unique bool) *TimeRecordQuery {
	trq.ctx.Unique = &unique
	return trq
}

// Order specifies how the records should be ordered.
func (trq *TimeRecordQuery) Order(o ...OrderFunc) *TimeRecordQuery {
	trq.order = append(trq.order, o...)
	return trq
}

// QueryTimeKeeper chains the current query on the "timeKeeper" edge.
func (trq *TimeRecordQuery) QueryTimeKeeper() *UserQuery {
	query := (&UserClient{config: trq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := trq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := trq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(timerecord.Table, timerecord.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, timerecord.TimeKeeperTable, timerecord.TimeKeeperColumn),
		)
		fromU = sqlgraph.SetNeighbors(trq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TimeRecord entity from the query.
// Returns a *NotFoundError when no TimeRecord was found.
func (trq *TimeRecordQuery) First(ctx context.Context) (*TimeRecord, error) {
	nodes, err := trq.Limit(1).All(setContextOp(ctx, trq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{timerecord.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (trq *TimeRecordQuery) FirstX(ctx context.Context) *TimeRecord {
	node, err := trq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TimeRecord ID from the query.
// Returns a *NotFoundError when no TimeRecord ID was found.
func (trq *TimeRecordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = trq.Limit(1).IDs(setContextOp(ctx, trq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{timerecord.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (trq *TimeRecordQuery) FirstIDX(ctx context.Context) int {
	id, err := trq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TimeRecord entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TimeRecord entity is found.
// Returns a *NotFoundError when no TimeRecord entities are found.
func (trq *TimeRecordQuery) Only(ctx context.Context) (*TimeRecord, error) {
	nodes, err := trq.Limit(2).All(setContextOp(ctx, trq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{timerecord.Label}
	default:
		return nil, &NotSingularError{timerecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (trq *TimeRecordQuery) OnlyX(ctx context.Context) *TimeRecord {
	node, err := trq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TimeRecord ID in the query.
// Returns a *NotSingularError when more than one TimeRecord ID is found.
// Returns a *NotFoundError when no entities are found.
func (trq *TimeRecordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = trq.Limit(2).IDs(setContextOp(ctx, trq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{timerecord.Label}
	default:
		err = &NotSingularError{timerecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (trq *TimeRecordQuery) OnlyIDX(ctx context.Context) int {
	id, err := trq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TimeRecords.
func (trq *TimeRecordQuery) All(ctx context.Context) ([]*TimeRecord, error) {
	ctx = setContextOp(ctx, trq.ctx, "All")
	if err := trq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TimeRecord, *TimeRecordQuery]()
	return withInterceptors[[]*TimeRecord](ctx, trq, qr, trq.inters)
}

// AllX is like All, but panics if an error occurs.
func (trq *TimeRecordQuery) AllX(ctx context.Context) []*TimeRecord {
	nodes, err := trq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TimeRecord IDs.
func (trq *TimeRecordQuery) IDs(ctx context.Context) (ids []int, err error) {
	if trq.ctx.Unique == nil && trq.path != nil {
		trq.Unique(true)
	}
	ctx = setContextOp(ctx, trq.ctx, "IDs")
	if err = trq.Select(timerecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (trq *TimeRecordQuery) IDsX(ctx context.Context) []int {
	ids, err := trq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (trq *TimeRecordQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, trq.ctx, "Count")
	if err := trq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, trq, querierCount[*TimeRecordQuery](), trq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (trq *TimeRecordQuery) CountX(ctx context.Context) int {
	count, err := trq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (trq *TimeRecordQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, trq.ctx, "Exist")
	switch _, err := trq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (trq *TimeRecordQuery) ExistX(ctx context.Context) bool {
	exist, err := trq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TimeRecordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (trq *TimeRecordQuery) Clone() *TimeRecordQuery {
	if trq == nil {
		return nil
	}
	return &TimeRecordQuery{
		config:         trq.config,
		ctx:            trq.ctx.Clone(),
		order:          append([]OrderFunc{}, trq.order...),
		inters:         append([]Interceptor{}, trq.inters...),
		predicates:     append([]predicate.TimeRecord{}, trq.predicates...),
		withTimeKeeper: trq.withTimeKeeper.Clone(),
		// clone intermediate query.
		sql:  trq.sql.Clone(),
		path: trq.path,
	}
}

// WithTimeKeeper tells the query-builder to eager-load the nodes that are connected to
// the "timeKeeper" edge. The optional arguments are used to configure the query builder of the edge.
func (trq *TimeRecordQuery) WithTimeKeeper(opts ...func(*UserQuery)) *TimeRecordQuery {
	query := (&UserClient{config: trq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	trq.withTimeKeeper = query
	return trq
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
//	client.TimeRecord.Query().
//		GroupBy(timerecord.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (trq *TimeRecordQuery) GroupBy(field string, fields ...string) *TimeRecordGroupBy {
	trq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TimeRecordGroupBy{build: trq}
	grbuild.flds = &trq.ctx.Fields
	grbuild.label = timerecord.Label
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
//	client.TimeRecord.Query().
//		Select(timerecord.FieldCreatedAt).
//		Scan(ctx, &v)
func (trq *TimeRecordQuery) Select(fields ...string) *TimeRecordSelect {
	trq.ctx.Fields = append(trq.ctx.Fields, fields...)
	sbuild := &TimeRecordSelect{TimeRecordQuery: trq}
	sbuild.label = timerecord.Label
	sbuild.flds, sbuild.scan = &trq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TimeRecordSelect configured with the given aggregations.
func (trq *TimeRecordQuery) Aggregate(fns ...AggregateFunc) *TimeRecordSelect {
	return trq.Select().Aggregate(fns...)
}

func (trq *TimeRecordQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range trq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, trq); err != nil {
				return err
			}
		}
	}
	for _, f := range trq.ctx.Fields {
		if !timerecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if trq.path != nil {
		prev, err := trq.path(ctx)
		if err != nil {
			return err
		}
		trq.sql = prev
	}
	return nil
}

func (trq *TimeRecordQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TimeRecord, error) {
	var (
		nodes       = []*TimeRecord{}
		withFKs     = trq.withFKs
		_spec       = trq.querySpec()
		loadedTypes = [1]bool{
			trq.withTimeKeeper != nil,
		}
	)
	if trq.withTimeKeeper != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, timerecord.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TimeRecord).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TimeRecord{config: trq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, trq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := trq.withTimeKeeper; query != nil {
		if err := trq.loadTimeKeeper(ctx, query, nodes, nil,
			func(n *TimeRecord, e *User) { n.Edges.TimeKeeper = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (trq *TimeRecordQuery) loadTimeKeeper(ctx context.Context, query *UserQuery, nodes []*TimeRecord, init func(*TimeRecord), assign func(*TimeRecord, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TimeRecord)
	for i := range nodes {
		if nodes[i].user_time_records == nil {
			continue
		}
		fk := *nodes[i].user_time_records
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_time_records" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (trq *TimeRecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := trq.querySpec()
	_spec.Node.Columns = trq.ctx.Fields
	if len(trq.ctx.Fields) > 0 {
		_spec.Unique = trq.ctx.Unique != nil && *trq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, trq.driver, _spec)
}

func (trq *TimeRecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(timerecord.Table, timerecord.Columns, sqlgraph.NewFieldSpec(timerecord.FieldID, field.TypeInt))
	_spec.From = trq.sql
	if unique := trq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if trq.path != nil {
		_spec.Unique = true
	}
	if fields := trq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, timerecord.FieldID)
		for i := range fields {
			if fields[i] != timerecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := trq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := trq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := trq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := trq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (trq *TimeRecordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(trq.driver.Dialect())
	t1 := builder.Table(timerecord.Table)
	columns := trq.ctx.Fields
	if len(columns) == 0 {
		columns = timerecord.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if trq.sql != nil {
		selector = trq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if trq.ctx.Unique != nil && *trq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range trq.predicates {
		p(selector)
	}
	for _, p := range trq.order {
		p(selector)
	}
	if offset := trq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := trq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TimeRecordGroupBy is the group-by builder for TimeRecord entities.
type TimeRecordGroupBy struct {
	selector
	build *TimeRecordQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (trgb *TimeRecordGroupBy) Aggregate(fns ...AggregateFunc) *TimeRecordGroupBy {
	trgb.fns = append(trgb.fns, fns...)
	return trgb
}

// Scan applies the selector query and scans the result into the given value.
func (trgb *TimeRecordGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, trgb.build.ctx, "GroupBy")
	if err := trgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TimeRecordQuery, *TimeRecordGroupBy](ctx, trgb.build, trgb, trgb.build.inters, v)
}

func (trgb *TimeRecordGroupBy) sqlScan(ctx context.Context, root *TimeRecordQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(trgb.fns))
	for _, fn := range trgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*trgb.flds)+len(trgb.fns))
		for _, f := range *trgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*trgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := trgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TimeRecordSelect is the builder for selecting fields of TimeRecord entities.
type TimeRecordSelect struct {
	*TimeRecordQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (trs *TimeRecordSelect) Aggregate(fns ...AggregateFunc) *TimeRecordSelect {
	trs.fns = append(trs.fns, fns...)
	return trs
}

// Scan applies the selector query and scans the result into the given value.
func (trs *TimeRecordSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, trs.ctx, "Select")
	if err := trs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TimeRecordQuery, *TimeRecordSelect](ctx, trs.TimeRecordQuery, trs, trs.inters, v)
}

func (trs *TimeRecordSelect) sqlScan(ctx context.Context, root *TimeRecordQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(trs.fns))
	for _, fn := range trs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*trs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := trs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}