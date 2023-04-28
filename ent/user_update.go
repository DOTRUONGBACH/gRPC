// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ent-grpc-example/ent/category"
	"ent-grpc-example/ent/predicate"
	"ent-grpc-example/ent/user"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetEmailAddress sets the "email_address" field.
func (uu *UserUpdate) SetEmailAddress(s string) *UserUpdate {
	uu.mutation.SetEmailAddress(s)
	return uu
}

// AddAdministeredIDs adds the "administered" edge to the Category entity by IDs.
func (uu *UserUpdate) AddAdministeredIDs(ids ...int) *UserUpdate {
	uu.mutation.AddAdministeredIDs(ids...)
	return uu
}

// AddAdministered adds the "administered" edges to the Category entity.
func (uu *UserUpdate) AddAdministered(c ...*Category) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddAdministeredIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearAdministered clears all "administered" edges to the Category entity.
func (uu *UserUpdate) ClearAdministered() *UserUpdate {
	uu.mutation.ClearAdministered()
	return uu
}

// RemoveAdministeredIDs removes the "administered" edge to Category entities by IDs.
func (uu *UserUpdate) RemoveAdministeredIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveAdministeredIDs(ids...)
	return uu
}

// RemoveAdministered removes "administered" edges to Category entities.
func (uu *UserUpdate) RemoveAdministered(c ...*Category) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveAdministeredIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.EmailAddress(); ok {
		_spec.SetField(user.FieldEmailAddress, field.TypeString, value)
	}
	if uu.mutation.AdministeredCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.AdministeredTable,
			Columns: []string{user.AdministeredColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedAdministeredIDs(); len(nodes) > 0 && !uu.mutation.AdministeredCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.AdministeredTable,
			Columns: []string{user.AdministeredColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AdministeredIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.AdministeredTable,
			Columns: []string{user.AdministeredColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetEmailAddress sets the "email_address" field.
func (uuo *UserUpdateOne) SetEmailAddress(s string) *UserUpdateOne {
	uuo.mutation.SetEmailAddress(s)
	return uuo
}

// AddAdministeredIDs adds the "administered" edge to the Category entity by IDs.
func (uuo *UserUpdateOne) AddAdministeredIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddAdministeredIDs(ids...)
	return uuo
}

// AddAdministered adds the "administered" edges to the Category entity.
func (uuo *UserUpdateOne) AddAdministered(c ...*Category) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddAdministeredIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearAdministered clears all "administered" edges to the Category entity.
func (uuo *UserUpdateOne) ClearAdministered() *UserUpdateOne {
	uuo.mutation.ClearAdministered()
	return uuo
}

// RemoveAdministeredIDs removes the "administered" edge to Category entities by IDs.
func (uuo *UserUpdateOne) RemoveAdministeredIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveAdministeredIDs(ids...)
	return uuo
}

// RemoveAdministered removes "administered" edges to Category entities.
func (uuo *UserUpdateOne) RemoveAdministered(c ...*Category) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveAdministeredIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.EmailAddress(); ok {
		_spec.SetField(user.FieldEmailAddress, field.TypeString, value)
	}
	if uuo.mutation.AdministeredCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.AdministeredTable,
			Columns: []string{user.AdministeredColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedAdministeredIDs(); len(nodes) > 0 && !uuo.mutation.AdministeredCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.AdministeredTable,
			Columns: []string{user.AdministeredColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AdministeredIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.AdministeredTable,
			Columns: []string{user.AdministeredColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
