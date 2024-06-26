// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/protobom/storage/internal/backends/ent/externalreference"
	"github.com/protobom/storage/internal/backends/ent/hashesentry"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// HashesEntryUpdate is the builder for updating HashesEntry entities.
type HashesEntryUpdate struct {
	config
	hooks    []Hook
	mutation *HashesEntryMutation
}

// Where appends a list predicates to the HashesEntryUpdate builder.
func (heu *HashesEntryUpdate) Where(ps ...predicate.HashesEntry) *HashesEntryUpdate {
	heu.mutation.Where(ps...)
	return heu
}

// SetExternalReferenceID sets the "external_reference_id" field.
func (heu *HashesEntryUpdate) SetExternalReferenceID(i int) *HashesEntryUpdate {
	heu.mutation.SetExternalReferenceID(i)
	return heu
}

// SetNillableExternalReferenceID sets the "external_reference_id" field if the given value is not nil.
func (heu *HashesEntryUpdate) SetNillableExternalReferenceID(i *int) *HashesEntryUpdate {
	if i != nil {
		heu.SetExternalReferenceID(*i)
	}
	return heu
}

// ClearExternalReferenceID clears the value of the "external_reference_id" field.
func (heu *HashesEntryUpdate) ClearExternalReferenceID() *HashesEntryUpdate {
	heu.mutation.ClearExternalReferenceID()
	return heu
}

// SetNodeID sets the "node_id" field.
func (heu *HashesEntryUpdate) SetNodeID(s string) *HashesEntryUpdate {
	heu.mutation.SetNodeID(s)
	return heu
}

// SetNillableNodeID sets the "node_id" field if the given value is not nil.
func (heu *HashesEntryUpdate) SetNillableNodeID(s *string) *HashesEntryUpdate {
	if s != nil {
		heu.SetNodeID(*s)
	}
	return heu
}

// ClearNodeID clears the value of the "node_id" field.
func (heu *HashesEntryUpdate) ClearNodeID() *HashesEntryUpdate {
	heu.mutation.ClearNodeID()
	return heu
}

// SetHashAlgorithmType sets the "hash_algorithm_type" field.
func (heu *HashesEntryUpdate) SetHashAlgorithmType(hat hashesentry.HashAlgorithmType) *HashesEntryUpdate {
	heu.mutation.SetHashAlgorithmType(hat)
	return heu
}

// SetNillableHashAlgorithmType sets the "hash_algorithm_type" field if the given value is not nil.
func (heu *HashesEntryUpdate) SetNillableHashAlgorithmType(hat *hashesentry.HashAlgorithmType) *HashesEntryUpdate {
	if hat != nil {
		heu.SetHashAlgorithmType(*hat)
	}
	return heu
}

// SetHashData sets the "hash_data" field.
func (heu *HashesEntryUpdate) SetHashData(s string) *HashesEntryUpdate {
	heu.mutation.SetHashData(s)
	return heu
}

// SetNillableHashData sets the "hash_data" field if the given value is not nil.
func (heu *HashesEntryUpdate) SetNillableHashData(s *string) *HashesEntryUpdate {
	if s != nil {
		heu.SetHashData(*s)
	}
	return heu
}

// SetExternalReference sets the "external_reference" edge to the ExternalReference entity.
func (heu *HashesEntryUpdate) SetExternalReference(e *ExternalReference) *HashesEntryUpdate {
	return heu.SetExternalReferenceID(e.ID)
}

// SetNode sets the "node" edge to the Node entity.
func (heu *HashesEntryUpdate) SetNode(n *Node) *HashesEntryUpdate {
	return heu.SetNodeID(n.ID)
}

// Mutation returns the HashesEntryMutation object of the builder.
func (heu *HashesEntryUpdate) Mutation() *HashesEntryMutation {
	return heu.mutation
}

// ClearExternalReference clears the "external_reference" edge to the ExternalReference entity.
func (heu *HashesEntryUpdate) ClearExternalReference() *HashesEntryUpdate {
	heu.mutation.ClearExternalReference()
	return heu
}

// ClearNode clears the "node" edge to the Node entity.
func (heu *HashesEntryUpdate) ClearNode() *HashesEntryUpdate {
	heu.mutation.ClearNode()
	return heu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (heu *HashesEntryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, heu.sqlSave, heu.mutation, heu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (heu *HashesEntryUpdate) SaveX(ctx context.Context) int {
	affected, err := heu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (heu *HashesEntryUpdate) Exec(ctx context.Context) error {
	_, err := heu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (heu *HashesEntryUpdate) ExecX(ctx context.Context) {
	if err := heu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (heu *HashesEntryUpdate) check() error {
	if v, ok := heu.mutation.HashAlgorithmType(); ok {
		if err := hashesentry.HashAlgorithmTypeValidator(v); err != nil {
			return &ValidationError{Name: "hash_algorithm_type", err: fmt.Errorf(`ent: validator failed for field "HashesEntry.hash_algorithm_type": %w`, err)}
		}
	}
	return nil
}

func (heu *HashesEntryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := heu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(hashesentry.Table, hashesentry.Columns, sqlgraph.NewFieldSpec(hashesentry.FieldID, field.TypeInt))
	if ps := heu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := heu.mutation.HashAlgorithmType(); ok {
		_spec.SetField(hashesentry.FieldHashAlgorithmType, field.TypeEnum, value)
	}
	if value, ok := heu.mutation.HashData(); ok {
		_spec.SetField(hashesentry.FieldHashData, field.TypeString, value)
	}
	if heu.mutation.ExternalReferenceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hashesentry.ExternalReferenceTable,
			Columns: []string{hashesentry.ExternalReferenceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(externalreference.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heu.mutation.ExternalReferenceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hashesentry.ExternalReferenceTable,
			Columns: []string{hashesentry.ExternalReferenceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(externalreference.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if heu.mutation.NodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hashesentry.NodeTable,
			Columns: []string{hashesentry.NodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heu.mutation.NodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hashesentry.NodeTable,
			Columns: []string{hashesentry.NodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, heu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hashesentry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	heu.mutation.done = true
	return n, nil
}

// HashesEntryUpdateOne is the builder for updating a single HashesEntry entity.
type HashesEntryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HashesEntryMutation
}

// SetExternalReferenceID sets the "external_reference_id" field.
func (heuo *HashesEntryUpdateOne) SetExternalReferenceID(i int) *HashesEntryUpdateOne {
	heuo.mutation.SetExternalReferenceID(i)
	return heuo
}

// SetNillableExternalReferenceID sets the "external_reference_id" field if the given value is not nil.
func (heuo *HashesEntryUpdateOne) SetNillableExternalReferenceID(i *int) *HashesEntryUpdateOne {
	if i != nil {
		heuo.SetExternalReferenceID(*i)
	}
	return heuo
}

// ClearExternalReferenceID clears the value of the "external_reference_id" field.
func (heuo *HashesEntryUpdateOne) ClearExternalReferenceID() *HashesEntryUpdateOne {
	heuo.mutation.ClearExternalReferenceID()
	return heuo
}

// SetNodeID sets the "node_id" field.
func (heuo *HashesEntryUpdateOne) SetNodeID(s string) *HashesEntryUpdateOne {
	heuo.mutation.SetNodeID(s)
	return heuo
}

// SetNillableNodeID sets the "node_id" field if the given value is not nil.
func (heuo *HashesEntryUpdateOne) SetNillableNodeID(s *string) *HashesEntryUpdateOne {
	if s != nil {
		heuo.SetNodeID(*s)
	}
	return heuo
}

// ClearNodeID clears the value of the "node_id" field.
func (heuo *HashesEntryUpdateOne) ClearNodeID() *HashesEntryUpdateOne {
	heuo.mutation.ClearNodeID()
	return heuo
}

// SetHashAlgorithmType sets the "hash_algorithm_type" field.
func (heuo *HashesEntryUpdateOne) SetHashAlgorithmType(hat hashesentry.HashAlgorithmType) *HashesEntryUpdateOne {
	heuo.mutation.SetHashAlgorithmType(hat)
	return heuo
}

// SetNillableHashAlgorithmType sets the "hash_algorithm_type" field if the given value is not nil.
func (heuo *HashesEntryUpdateOne) SetNillableHashAlgorithmType(hat *hashesentry.HashAlgorithmType) *HashesEntryUpdateOne {
	if hat != nil {
		heuo.SetHashAlgorithmType(*hat)
	}
	return heuo
}

// SetHashData sets the "hash_data" field.
func (heuo *HashesEntryUpdateOne) SetHashData(s string) *HashesEntryUpdateOne {
	heuo.mutation.SetHashData(s)
	return heuo
}

// SetNillableHashData sets the "hash_data" field if the given value is not nil.
func (heuo *HashesEntryUpdateOne) SetNillableHashData(s *string) *HashesEntryUpdateOne {
	if s != nil {
		heuo.SetHashData(*s)
	}
	return heuo
}

// SetExternalReference sets the "external_reference" edge to the ExternalReference entity.
func (heuo *HashesEntryUpdateOne) SetExternalReference(e *ExternalReference) *HashesEntryUpdateOne {
	return heuo.SetExternalReferenceID(e.ID)
}

// SetNode sets the "node" edge to the Node entity.
func (heuo *HashesEntryUpdateOne) SetNode(n *Node) *HashesEntryUpdateOne {
	return heuo.SetNodeID(n.ID)
}

// Mutation returns the HashesEntryMutation object of the builder.
func (heuo *HashesEntryUpdateOne) Mutation() *HashesEntryMutation {
	return heuo.mutation
}

// ClearExternalReference clears the "external_reference" edge to the ExternalReference entity.
func (heuo *HashesEntryUpdateOne) ClearExternalReference() *HashesEntryUpdateOne {
	heuo.mutation.ClearExternalReference()
	return heuo
}

// ClearNode clears the "node" edge to the Node entity.
func (heuo *HashesEntryUpdateOne) ClearNode() *HashesEntryUpdateOne {
	heuo.mutation.ClearNode()
	return heuo
}

// Where appends a list predicates to the HashesEntryUpdate builder.
func (heuo *HashesEntryUpdateOne) Where(ps ...predicate.HashesEntry) *HashesEntryUpdateOne {
	heuo.mutation.Where(ps...)
	return heuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (heuo *HashesEntryUpdateOne) Select(field string, fields ...string) *HashesEntryUpdateOne {
	heuo.fields = append([]string{field}, fields...)
	return heuo
}

// Save executes the query and returns the updated HashesEntry entity.
func (heuo *HashesEntryUpdateOne) Save(ctx context.Context) (*HashesEntry, error) {
	return withHooks(ctx, heuo.sqlSave, heuo.mutation, heuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (heuo *HashesEntryUpdateOne) SaveX(ctx context.Context) *HashesEntry {
	node, err := heuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (heuo *HashesEntryUpdateOne) Exec(ctx context.Context) error {
	_, err := heuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (heuo *HashesEntryUpdateOne) ExecX(ctx context.Context) {
	if err := heuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (heuo *HashesEntryUpdateOne) check() error {
	if v, ok := heuo.mutation.HashAlgorithmType(); ok {
		if err := hashesentry.HashAlgorithmTypeValidator(v); err != nil {
			return &ValidationError{Name: "hash_algorithm_type", err: fmt.Errorf(`ent: validator failed for field "HashesEntry.hash_algorithm_type": %w`, err)}
		}
	}
	return nil
}

func (heuo *HashesEntryUpdateOne) sqlSave(ctx context.Context) (_node *HashesEntry, err error) {
	if err := heuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(hashesentry.Table, hashesentry.Columns, sqlgraph.NewFieldSpec(hashesentry.FieldID, field.TypeInt))
	id, ok := heuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HashesEntry.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := heuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hashesentry.FieldID)
		for _, f := range fields {
			if !hashesentry.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != hashesentry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := heuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := heuo.mutation.HashAlgorithmType(); ok {
		_spec.SetField(hashesentry.FieldHashAlgorithmType, field.TypeEnum, value)
	}
	if value, ok := heuo.mutation.HashData(); ok {
		_spec.SetField(hashesentry.FieldHashData, field.TypeString, value)
	}
	if heuo.mutation.ExternalReferenceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hashesentry.ExternalReferenceTable,
			Columns: []string{hashesentry.ExternalReferenceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(externalreference.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heuo.mutation.ExternalReferenceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hashesentry.ExternalReferenceTable,
			Columns: []string{hashesentry.ExternalReferenceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(externalreference.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if heuo.mutation.NodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hashesentry.NodeTable,
			Columns: []string{hashesentry.NodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heuo.mutation.NodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hashesentry.NodeTable,
			Columns: []string{hashesentry.NodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &HashesEntry{config: heuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, heuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hashesentry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	heuo.mutation.done = true
	return _node, nil
}
