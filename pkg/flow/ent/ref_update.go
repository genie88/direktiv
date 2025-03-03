// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/direktiv/direktiv/pkg/flow/ent/ref"
	"github.com/direktiv/direktiv/pkg/flow/ent/revision"
	"github.com/direktiv/direktiv/pkg/flow/ent/route"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// RefUpdate is the builder for updating Ref entities.
type RefUpdate struct {
	config
	hooks    []Hook
	mutation *RefMutation
}

// Where appends a list predicates to the RefUpdate builder.
func (ru *RefUpdate) Where(ps ...predicate.Ref) *RefUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (ru *RefUpdate) SetWorkflowID(id uuid.UUID) *RefUpdate {
	ru.mutation.SetWorkflowID(id)
	return ru
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (ru *RefUpdate) SetWorkflow(w *Workflow) *RefUpdate {
	return ru.SetWorkflowID(w.ID)
}

// SetRevisionID sets the "revision" edge to the Revision entity by ID.
func (ru *RefUpdate) SetRevisionID(id uuid.UUID) *RefUpdate {
	ru.mutation.SetRevisionID(id)
	return ru
}

// SetRevision sets the "revision" edge to the Revision entity.
func (ru *RefUpdate) SetRevision(r *Revision) *RefUpdate {
	return ru.SetRevisionID(r.ID)
}

// AddRouteIDs adds the "routes" edge to the Route entity by IDs.
func (ru *RefUpdate) AddRouteIDs(ids ...uuid.UUID) *RefUpdate {
	ru.mutation.AddRouteIDs(ids...)
	return ru
}

// AddRoutes adds the "routes" edges to the Route entity.
func (ru *RefUpdate) AddRoutes(r ...*Route) *RefUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddRouteIDs(ids...)
}

// Mutation returns the RefMutation object of the builder.
func (ru *RefUpdate) Mutation() *RefMutation {
	return ru.mutation
}

// ClearWorkflow clears the "workflow" edge to the Workflow entity.
func (ru *RefUpdate) ClearWorkflow() *RefUpdate {
	ru.mutation.ClearWorkflow()
	return ru
}

// ClearRevision clears the "revision" edge to the Revision entity.
func (ru *RefUpdate) ClearRevision() *RefUpdate {
	ru.mutation.ClearRevision()
	return ru
}

// ClearRoutes clears all "routes" edges to the Route entity.
func (ru *RefUpdate) ClearRoutes() *RefUpdate {
	ru.mutation.ClearRoutes()
	return ru
}

// RemoveRouteIDs removes the "routes" edge to Route entities by IDs.
func (ru *RefUpdate) RemoveRouteIDs(ids ...uuid.UUID) *RefUpdate {
	ru.mutation.RemoveRouteIDs(ids...)
	return ru
}

// RemoveRoutes removes "routes" edges to Route entities.
func (ru *RefUpdate) RemoveRoutes(r ...*Route) *RefUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveRouteIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RefUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RefMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RefUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RefUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RefUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RefUpdate) check() error {
	if _, ok := ru.mutation.WorkflowID(); ru.mutation.WorkflowCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"workflow\"")
	}
	if _, ok := ru.mutation.RevisionID(); ru.mutation.RevisionCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"revision\"")
	}
	return nil
}

func (ru *RefUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ref.Table,
			Columns: ref.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: ref.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ru.mutation.WorkflowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ref.WorkflowTable,
			Columns: []string{ref.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ref.WorkflowTable,
			Columns: []string{ref.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.RevisionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ref.RevisionTable,
			Columns: []string{ref.RevisionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: revision.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RevisionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ref.RevisionTable,
			Columns: []string{ref.RevisionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: revision.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.RoutesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ref.RoutesTable,
			Columns: []string{ref.RoutesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: route.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedRoutesIDs(); len(nodes) > 0 && !ru.mutation.RoutesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ref.RoutesTable,
			Columns: []string{ref.RoutesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: route.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RoutesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ref.RoutesTable,
			Columns: []string{ref.RoutesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: route.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ref.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RefUpdateOne is the builder for updating a single Ref entity.
type RefUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RefMutation
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (ruo *RefUpdateOne) SetWorkflowID(id uuid.UUID) *RefUpdateOne {
	ruo.mutation.SetWorkflowID(id)
	return ruo
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (ruo *RefUpdateOne) SetWorkflow(w *Workflow) *RefUpdateOne {
	return ruo.SetWorkflowID(w.ID)
}

// SetRevisionID sets the "revision" edge to the Revision entity by ID.
func (ruo *RefUpdateOne) SetRevisionID(id uuid.UUID) *RefUpdateOne {
	ruo.mutation.SetRevisionID(id)
	return ruo
}

// SetRevision sets the "revision" edge to the Revision entity.
func (ruo *RefUpdateOne) SetRevision(r *Revision) *RefUpdateOne {
	return ruo.SetRevisionID(r.ID)
}

// AddRouteIDs adds the "routes" edge to the Route entity by IDs.
func (ruo *RefUpdateOne) AddRouteIDs(ids ...uuid.UUID) *RefUpdateOne {
	ruo.mutation.AddRouteIDs(ids...)
	return ruo
}

// AddRoutes adds the "routes" edges to the Route entity.
func (ruo *RefUpdateOne) AddRoutes(r ...*Route) *RefUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddRouteIDs(ids...)
}

// Mutation returns the RefMutation object of the builder.
func (ruo *RefUpdateOne) Mutation() *RefMutation {
	return ruo.mutation
}

// ClearWorkflow clears the "workflow" edge to the Workflow entity.
func (ruo *RefUpdateOne) ClearWorkflow() *RefUpdateOne {
	ruo.mutation.ClearWorkflow()
	return ruo
}

// ClearRevision clears the "revision" edge to the Revision entity.
func (ruo *RefUpdateOne) ClearRevision() *RefUpdateOne {
	ruo.mutation.ClearRevision()
	return ruo
}

// ClearRoutes clears all "routes" edges to the Route entity.
func (ruo *RefUpdateOne) ClearRoutes() *RefUpdateOne {
	ruo.mutation.ClearRoutes()
	return ruo
}

// RemoveRouteIDs removes the "routes" edge to Route entities by IDs.
func (ruo *RefUpdateOne) RemoveRouteIDs(ids ...uuid.UUID) *RefUpdateOne {
	ruo.mutation.RemoveRouteIDs(ids...)
	return ruo
}

// RemoveRoutes removes "routes" edges to Route entities.
func (ruo *RefUpdateOne) RemoveRoutes(r ...*Route) *RefUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveRouteIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RefUpdateOne) Select(field string, fields ...string) *RefUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Ref entity.
func (ruo *RefUpdateOne) Save(ctx context.Context) (*Ref, error) {
	var (
		err  error
		node *Ref
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RefMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RefUpdateOne) SaveX(ctx context.Context) *Ref {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RefUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RefUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RefUpdateOne) check() error {
	if _, ok := ruo.mutation.WorkflowID(); ruo.mutation.WorkflowCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"workflow\"")
	}
	if _, ok := ruo.mutation.RevisionID(); ruo.mutation.RevisionCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"revision\"")
	}
	return nil
}

func (ruo *RefUpdateOne) sqlSave(ctx context.Context) (_node *Ref, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ref.Table,
			Columns: ref.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: ref.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Ref.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ref.FieldID)
		for _, f := range fields {
			if !ref.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ref.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ruo.mutation.WorkflowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ref.WorkflowTable,
			Columns: []string{ref.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ref.WorkflowTable,
			Columns: []string{ref.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.RevisionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ref.RevisionTable,
			Columns: []string{ref.RevisionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: revision.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RevisionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ref.RevisionTable,
			Columns: []string{ref.RevisionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: revision.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.RoutesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ref.RoutesTable,
			Columns: []string{ref.RoutesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: route.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedRoutesIDs(); len(nodes) > 0 && !ruo.mutation.RoutesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ref.RoutesTable,
			Columns: []string{ref.RoutesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: route.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RoutesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ref.RoutesTable,
			Columns: []string{ref.RoutesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: route.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Ref{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ref.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
