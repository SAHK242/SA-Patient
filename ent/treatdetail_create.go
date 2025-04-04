// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"patient/ent/treatdetail"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TreatDetailCreate is the builder for creating a TreatDetail entity.
type TreatDetailCreate struct {
	config
	mutation *TreatDetailMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// Mutation returns the TreatDetailMutation object of the builder.
func (tdc *TreatDetailCreate) Mutation() *TreatDetailMutation {
	return tdc.mutation
}

// Save creates the TreatDetail in the database.
func (tdc *TreatDetailCreate) Save(ctx context.Context) (*TreatDetail, error) {
	return withHooks(ctx, tdc.sqlSave, tdc.mutation, tdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tdc *TreatDetailCreate) SaveX(ctx context.Context) *TreatDetail {
	v, err := tdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tdc *TreatDetailCreate) Exec(ctx context.Context) error {
	_, err := tdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdc *TreatDetailCreate) ExecX(ctx context.Context) {
	if err := tdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tdc *TreatDetailCreate) check() error {
	return nil
}

func (tdc *TreatDetailCreate) sqlSave(ctx context.Context) (*TreatDetail, error) {
	if err := tdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tdc.mutation.id = &_node.ID
	tdc.mutation.done = true
	return _node, nil
}

func (tdc *TreatDetailCreate) createSpec() (*TreatDetail, *sqlgraph.CreateSpec) {
	var (
		_node = &TreatDetail{config: tdc.config}
		_spec = sqlgraph.NewCreateSpec(treatdetail.Table, sqlgraph.NewFieldSpec(treatdetail.FieldID, field.TypeInt))
	)
	_spec.OnConflict = tdc.conflict
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TreatDetail.Create().
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (tdc *TreatDetailCreate) OnConflict(opts ...sql.ConflictOption) *TreatDetailUpsertOne {
	tdc.conflict = opts
	return &TreatDetailUpsertOne{
		create: tdc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TreatDetail.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tdc *TreatDetailCreate) OnConflictColumns(columns ...string) *TreatDetailUpsertOne {
	tdc.conflict = append(tdc.conflict, sql.ConflictColumns(columns...))
	return &TreatDetailUpsertOne{
		create: tdc,
	}
}

type (
	// TreatDetailUpsertOne is the builder for "upsert"-ing
	//  one TreatDetail node.
	TreatDetailUpsertOne struct {
		create *TreatDetailCreate
	}

	// TreatDetailUpsert is the "OnConflict" setter.
	TreatDetailUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.TreatDetail.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TreatDetailUpsertOne) UpdateNewValues() *TreatDetailUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TreatDetail.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TreatDetailUpsertOne) Ignore() *TreatDetailUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TreatDetailUpsertOne) DoNothing() *TreatDetailUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TreatDetailCreate.OnConflict
// documentation for more info.
func (u *TreatDetailUpsertOne) Update(set func(*TreatDetailUpsert)) *TreatDetailUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TreatDetailUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *TreatDetailUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TreatDetailCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TreatDetailUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TreatDetailUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TreatDetailUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TreatDetailCreateBulk is the builder for creating many TreatDetail entities in bulk.
type TreatDetailCreateBulk struct {
	config
	err      error
	builders []*TreatDetailCreate
	conflict []sql.ConflictOption
}

// Save creates the TreatDetail entities in the database.
func (tdcb *TreatDetailCreateBulk) Save(ctx context.Context) ([]*TreatDetail, error) {
	if tdcb.err != nil {
		return nil, tdcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tdcb.builders))
	nodes := make([]*TreatDetail, len(tdcb.builders))
	mutators := make([]Mutator, len(tdcb.builders))
	for i := range tdcb.builders {
		func(i int, root context.Context) {
			builder := tdcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TreatDetailMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tdcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tdcb *TreatDetailCreateBulk) SaveX(ctx context.Context) []*TreatDetail {
	v, err := tdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tdcb *TreatDetailCreateBulk) Exec(ctx context.Context) error {
	_, err := tdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdcb *TreatDetailCreateBulk) ExecX(ctx context.Context) {
	if err := tdcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TreatDetail.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (tdcb *TreatDetailCreateBulk) OnConflict(opts ...sql.ConflictOption) *TreatDetailUpsertBulk {
	tdcb.conflict = opts
	return &TreatDetailUpsertBulk{
		create: tdcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TreatDetail.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tdcb *TreatDetailCreateBulk) OnConflictColumns(columns ...string) *TreatDetailUpsertBulk {
	tdcb.conflict = append(tdcb.conflict, sql.ConflictColumns(columns...))
	return &TreatDetailUpsertBulk{
		create: tdcb,
	}
}

// TreatDetailUpsertBulk is the builder for "upsert"-ing
// a bulk of TreatDetail nodes.
type TreatDetailUpsertBulk struct {
	create *TreatDetailCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TreatDetail.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TreatDetailUpsertBulk) UpdateNewValues() *TreatDetailUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TreatDetail.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TreatDetailUpsertBulk) Ignore() *TreatDetailUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TreatDetailUpsertBulk) DoNothing() *TreatDetailUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TreatDetailCreateBulk.OnConflict
// documentation for more info.
func (u *TreatDetailUpsertBulk) Update(set func(*TreatDetailUpsert)) *TreatDetailUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TreatDetailUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *TreatDetailUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TreatDetailCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TreatDetailCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TreatDetailUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
