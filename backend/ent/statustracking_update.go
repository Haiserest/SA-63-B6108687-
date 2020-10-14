// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/Haiserest/app/ent/predicate"
	"github.com/Haiserest/app/ent/statustracking"
	"github.com/Haiserest/app/ent/treatment"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// StatusTrackingUpdate is the builder for updating StatusTracking entities.
type StatusTrackingUpdate struct {
	config
	hooks      []Hook
	mutation   *StatusTrackingMutation
	predicates []predicate.StatusTracking
}

// Where adds a new predicate for the builder.
func (stu *StatusTrackingUpdate) Where(ps ...predicate.StatusTracking) *StatusTrackingUpdate {
	stu.predicates = append(stu.predicates, ps...)
	return stu
}

// SetStatus sets the status field.
func (stu *StatusTrackingUpdate) SetStatus(s string) *StatusTrackingUpdate {
	stu.mutation.SetStatus(s)
	return stu
}

// AddStatusIDIDs adds the status_id edge to Treatment by ids.
func (stu *StatusTrackingUpdate) AddStatusIDIDs(ids ...int) *StatusTrackingUpdate {
	stu.mutation.AddStatusIDIDs(ids...)
	return stu
}

// AddStatusID adds the status_id edges to Treatment.
func (stu *StatusTrackingUpdate) AddStatusID(t ...*Treatment) *StatusTrackingUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return stu.AddStatusIDIDs(ids...)
}

// Mutation returns the StatusTrackingMutation object of the builder.
func (stu *StatusTrackingUpdate) Mutation() *StatusTrackingMutation {
	return stu.mutation
}

// RemoveStatusIDIDs removes the status_id edge to Treatment by ids.
func (stu *StatusTrackingUpdate) RemoveStatusIDIDs(ids ...int) *StatusTrackingUpdate {
	stu.mutation.RemoveStatusIDIDs(ids...)
	return stu
}

// RemoveStatusID removes status_id edges to Treatment.
func (stu *StatusTrackingUpdate) RemoveStatusID(t ...*Treatment) *StatusTrackingUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return stu.RemoveStatusIDIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (stu *StatusTrackingUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := stu.mutation.Status(); ok {
		if err := statustracking.StatusValidator(v); err != nil {
			return 0, &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(stu.hooks) == 0 {
		affected, err = stu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatusTrackingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			stu.mutation = mutation
			affected, err = stu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(stu.hooks) - 1; i >= 0; i-- {
			mut = stu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, stu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (stu *StatusTrackingUpdate) SaveX(ctx context.Context) int {
	affected, err := stu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (stu *StatusTrackingUpdate) Exec(ctx context.Context) error {
	_, err := stu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stu *StatusTrackingUpdate) ExecX(ctx context.Context) {
	if err := stu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (stu *StatusTrackingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statustracking.Table,
			Columns: statustracking.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statustracking.FieldID,
			},
		},
	}
	if ps := stu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: statustracking.FieldStatus,
		})
	}
	if nodes := stu.mutation.RemovedStatusIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statustracking.StatusIDTable,
			Columns: []string{statustracking.StatusIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: treatment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.StatusIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statustracking.StatusIDTable,
			Columns: []string{statustracking.StatusIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: treatment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, stu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{statustracking.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// StatusTrackingUpdateOne is the builder for updating a single StatusTracking entity.
type StatusTrackingUpdateOne struct {
	config
	hooks    []Hook
	mutation *StatusTrackingMutation
}

// SetStatus sets the status field.
func (stuo *StatusTrackingUpdateOne) SetStatus(s string) *StatusTrackingUpdateOne {
	stuo.mutation.SetStatus(s)
	return stuo
}

// AddStatusIDIDs adds the status_id edge to Treatment by ids.
func (stuo *StatusTrackingUpdateOne) AddStatusIDIDs(ids ...int) *StatusTrackingUpdateOne {
	stuo.mutation.AddStatusIDIDs(ids...)
	return stuo
}

// AddStatusID adds the status_id edges to Treatment.
func (stuo *StatusTrackingUpdateOne) AddStatusID(t ...*Treatment) *StatusTrackingUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return stuo.AddStatusIDIDs(ids...)
}

// Mutation returns the StatusTrackingMutation object of the builder.
func (stuo *StatusTrackingUpdateOne) Mutation() *StatusTrackingMutation {
	return stuo.mutation
}

// RemoveStatusIDIDs removes the status_id edge to Treatment by ids.
func (stuo *StatusTrackingUpdateOne) RemoveStatusIDIDs(ids ...int) *StatusTrackingUpdateOne {
	stuo.mutation.RemoveStatusIDIDs(ids...)
	return stuo
}

// RemoveStatusID removes status_id edges to Treatment.
func (stuo *StatusTrackingUpdateOne) RemoveStatusID(t ...*Treatment) *StatusTrackingUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return stuo.RemoveStatusIDIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (stuo *StatusTrackingUpdateOne) Save(ctx context.Context) (*StatusTracking, error) {
	if v, ok := stuo.mutation.Status(); ok {
		if err := statustracking.StatusValidator(v); err != nil {
			return nil, &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}

	var (
		err  error
		node *StatusTracking
	)
	if len(stuo.hooks) == 0 {
		node, err = stuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatusTrackingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			stuo.mutation = mutation
			node, err = stuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(stuo.hooks) - 1; i >= 0; i-- {
			mut = stuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, stuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (stuo *StatusTrackingUpdateOne) SaveX(ctx context.Context) *StatusTracking {
	st, err := stuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return st
}

// Exec executes the query on the entity.
func (stuo *StatusTrackingUpdateOne) Exec(ctx context.Context) error {
	_, err := stuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stuo *StatusTrackingUpdateOne) ExecX(ctx context.Context) {
	if err := stuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (stuo *StatusTrackingUpdateOne) sqlSave(ctx context.Context) (st *StatusTracking, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statustracking.Table,
			Columns: statustracking.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statustracking.FieldID,
			},
		},
	}
	id, ok := stuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing StatusTracking.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := stuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: statustracking.FieldStatus,
		})
	}
	if nodes := stuo.mutation.RemovedStatusIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statustracking.StatusIDTable,
			Columns: []string{statustracking.StatusIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: treatment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.StatusIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statustracking.StatusIDTable,
			Columns: []string{statustracking.StatusIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: treatment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	st = &StatusTracking{config: stuo.config}
	_spec.Assign = st.assignValues
	_spec.ScanValues = st.scanValues()
	if err = sqlgraph.UpdateNode(ctx, stuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{statustracking.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return st, nil
}