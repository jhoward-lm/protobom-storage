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
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/documenttype"
	"github.com/protobom/storage/internal/backends/ent/metadata"
	"github.com/protobom/storage/internal/backends/ent/person"
	"github.com/protobom/storage/internal/backends/ent/predicate"
	"github.com/protobom/storage/internal/backends/ent/tool"
)

// MetadataUpdate is the builder for updating Metadata entities.
type MetadataUpdate struct {
	config
	hooks    []Hook
	mutation *MetadataMutation
}

// Where appends a list predicates to the MetadataUpdate builder.
func (mu *MetadataUpdate) Where(ps ...predicate.Metadata) *MetadataUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetVersion sets the "version" field.
func (mu *MetadataUpdate) SetVersion(s string) *MetadataUpdate {
	mu.mutation.SetVersion(s)
	return mu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (mu *MetadataUpdate) SetNillableVersion(s *string) *MetadataUpdate {
	if s != nil {
		mu.SetVersion(*s)
	}
	return mu
}

// SetName sets the "name" field.
func (mu *MetadataUpdate) SetName(s string) *MetadataUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mu *MetadataUpdate) SetNillableName(s *string) *MetadataUpdate {
	if s != nil {
		mu.SetName(*s)
	}
	return mu
}

// SetDate sets the "date" field.
func (mu *MetadataUpdate) SetDate(t time.Time) *MetadataUpdate {
	mu.mutation.SetDate(t)
	return mu
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (mu *MetadataUpdate) SetNillableDate(t *time.Time) *MetadataUpdate {
	if t != nil {
		mu.SetDate(*t)
	}
	return mu
}

// SetComment sets the "comment" field.
func (mu *MetadataUpdate) SetComment(s string) *MetadataUpdate {
	mu.mutation.SetComment(s)
	return mu
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (mu *MetadataUpdate) SetNillableComment(s *string) *MetadataUpdate {
	if s != nil {
		mu.SetComment(*s)
	}
	return mu
}

// AddToolIDs adds the "tools" edge to the Tool entity by IDs.
func (mu *MetadataUpdate) AddToolIDs(ids ...int) *MetadataUpdate {
	mu.mutation.AddToolIDs(ids...)
	return mu
}

// AddTools adds the "tools" edges to the Tool entity.
func (mu *MetadataUpdate) AddTools(t ...*Tool) *MetadataUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return mu.AddToolIDs(ids...)
}

// AddAuthorIDs adds the "authors" edge to the Person entity by IDs.
func (mu *MetadataUpdate) AddAuthorIDs(ids ...int) *MetadataUpdate {
	mu.mutation.AddAuthorIDs(ids...)
	return mu
}

// AddAuthors adds the "authors" edges to the Person entity.
func (mu *MetadataUpdate) AddAuthors(p ...*Person) *MetadataUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.AddAuthorIDs(ids...)
}

// AddDocumentTypeIDs adds the "document_types" edge to the DocumentType entity by IDs.
func (mu *MetadataUpdate) AddDocumentTypeIDs(ids ...int) *MetadataUpdate {
	mu.mutation.AddDocumentTypeIDs(ids...)
	return mu
}

// AddDocumentTypes adds the "document_types" edges to the DocumentType entity.
func (mu *MetadataUpdate) AddDocumentTypes(d ...*DocumentType) *MetadataUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mu.AddDocumentTypeIDs(ids...)
}

// SetDocumentID sets the "document" edge to the Document entity by ID.
func (mu *MetadataUpdate) SetDocumentID(id string) *MetadataUpdate {
	mu.mutation.SetDocumentID(id)
	return mu
}

// SetNillableDocumentID sets the "document" edge to the Document entity by ID if the given value is not nil.
func (mu *MetadataUpdate) SetNillableDocumentID(id *string) *MetadataUpdate {
	if id != nil {
		mu = mu.SetDocumentID(*id)
	}
	return mu
}

// SetDocument sets the "document" edge to the Document entity.
func (mu *MetadataUpdate) SetDocument(d *Document) *MetadataUpdate {
	return mu.SetDocumentID(d.ID)
}

// Mutation returns the MetadataMutation object of the builder.
func (mu *MetadataUpdate) Mutation() *MetadataMutation {
	return mu.mutation
}

// ClearTools clears all "tools" edges to the Tool entity.
func (mu *MetadataUpdate) ClearTools() *MetadataUpdate {
	mu.mutation.ClearTools()
	return mu
}

// RemoveToolIDs removes the "tools" edge to Tool entities by IDs.
func (mu *MetadataUpdate) RemoveToolIDs(ids ...int) *MetadataUpdate {
	mu.mutation.RemoveToolIDs(ids...)
	return mu
}

// RemoveTools removes "tools" edges to Tool entities.
func (mu *MetadataUpdate) RemoveTools(t ...*Tool) *MetadataUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return mu.RemoveToolIDs(ids...)
}

// ClearAuthors clears all "authors" edges to the Person entity.
func (mu *MetadataUpdate) ClearAuthors() *MetadataUpdate {
	mu.mutation.ClearAuthors()
	return mu
}

// RemoveAuthorIDs removes the "authors" edge to Person entities by IDs.
func (mu *MetadataUpdate) RemoveAuthorIDs(ids ...int) *MetadataUpdate {
	mu.mutation.RemoveAuthorIDs(ids...)
	return mu
}

// RemoveAuthors removes "authors" edges to Person entities.
func (mu *MetadataUpdate) RemoveAuthors(p ...*Person) *MetadataUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.RemoveAuthorIDs(ids...)
}

// ClearDocumentTypes clears all "document_types" edges to the DocumentType entity.
func (mu *MetadataUpdate) ClearDocumentTypes() *MetadataUpdate {
	mu.mutation.ClearDocumentTypes()
	return mu
}

// RemoveDocumentTypeIDs removes the "document_types" edge to DocumentType entities by IDs.
func (mu *MetadataUpdate) RemoveDocumentTypeIDs(ids ...int) *MetadataUpdate {
	mu.mutation.RemoveDocumentTypeIDs(ids...)
	return mu
}

// RemoveDocumentTypes removes "document_types" edges to DocumentType entities.
func (mu *MetadataUpdate) RemoveDocumentTypes(d ...*DocumentType) *MetadataUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mu.RemoveDocumentTypeIDs(ids...)
}

// ClearDocument clears the "document" edge to the Document entity.
func (mu *MetadataUpdate) ClearDocument() *MetadataUpdate {
	mu.mutation.ClearDocument()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MetadataUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MetadataUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MetadataUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MetadataUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MetadataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(metadata.Table, metadata.Columns, sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Version(); ok {
		_spec.SetField(metadata.FieldVersion, field.TypeString, value)
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.SetField(metadata.FieldName, field.TypeString, value)
	}
	if value, ok := mu.mutation.Date(); ok {
		_spec.SetField(metadata.FieldDate, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Comment(); ok {
		_spec.SetField(metadata.FieldComment, field.TypeString, value)
	}
	if mu.mutation.ToolsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.ToolsTable,
			Columns: []string{metadata.ToolsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedToolsIDs(); len(nodes) > 0 && !mu.mutation.ToolsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.ToolsTable,
			Columns: []string{metadata.ToolsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ToolsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.ToolsTable,
			Columns: []string{metadata.ToolsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.AuthorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.AuthorsTable,
			Columns: []string{metadata.AuthorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedAuthorsIDs(); len(nodes) > 0 && !mu.mutation.AuthorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.AuthorsTable,
			Columns: []string{metadata.AuthorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.AuthorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.AuthorsTable,
			Columns: []string{metadata.AuthorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.DocumentTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.DocumentTypesTable,
			Columns: []string{metadata.DocumentTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(documenttype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedDocumentTypesIDs(); len(nodes) > 0 && !mu.mutation.DocumentTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.DocumentTypesTable,
			Columns: []string{metadata.DocumentTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(documenttype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.DocumentTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.DocumentTypesTable,
			Columns: []string{metadata.DocumentTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(documenttype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.DocumentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   metadata.DocumentTable,
			Columns: []string{metadata.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(document.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.DocumentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   metadata.DocumentTable,
			Columns: []string{metadata.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(document.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MetadataUpdateOne is the builder for updating a single Metadata entity.
type MetadataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MetadataMutation
}

// SetVersion sets the "version" field.
func (muo *MetadataUpdateOne) SetVersion(s string) *MetadataUpdateOne {
	muo.mutation.SetVersion(s)
	return muo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (muo *MetadataUpdateOne) SetNillableVersion(s *string) *MetadataUpdateOne {
	if s != nil {
		muo.SetVersion(*s)
	}
	return muo
}

// SetName sets the "name" field.
func (muo *MetadataUpdateOne) SetName(s string) *MetadataUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (muo *MetadataUpdateOne) SetNillableName(s *string) *MetadataUpdateOne {
	if s != nil {
		muo.SetName(*s)
	}
	return muo
}

// SetDate sets the "date" field.
func (muo *MetadataUpdateOne) SetDate(t time.Time) *MetadataUpdateOne {
	muo.mutation.SetDate(t)
	return muo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (muo *MetadataUpdateOne) SetNillableDate(t *time.Time) *MetadataUpdateOne {
	if t != nil {
		muo.SetDate(*t)
	}
	return muo
}

// SetComment sets the "comment" field.
func (muo *MetadataUpdateOne) SetComment(s string) *MetadataUpdateOne {
	muo.mutation.SetComment(s)
	return muo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (muo *MetadataUpdateOne) SetNillableComment(s *string) *MetadataUpdateOne {
	if s != nil {
		muo.SetComment(*s)
	}
	return muo
}

// AddToolIDs adds the "tools" edge to the Tool entity by IDs.
func (muo *MetadataUpdateOne) AddToolIDs(ids ...int) *MetadataUpdateOne {
	muo.mutation.AddToolIDs(ids...)
	return muo
}

// AddTools adds the "tools" edges to the Tool entity.
func (muo *MetadataUpdateOne) AddTools(t ...*Tool) *MetadataUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return muo.AddToolIDs(ids...)
}

// AddAuthorIDs adds the "authors" edge to the Person entity by IDs.
func (muo *MetadataUpdateOne) AddAuthorIDs(ids ...int) *MetadataUpdateOne {
	muo.mutation.AddAuthorIDs(ids...)
	return muo
}

// AddAuthors adds the "authors" edges to the Person entity.
func (muo *MetadataUpdateOne) AddAuthors(p ...*Person) *MetadataUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.AddAuthorIDs(ids...)
}

// AddDocumentTypeIDs adds the "document_types" edge to the DocumentType entity by IDs.
func (muo *MetadataUpdateOne) AddDocumentTypeIDs(ids ...int) *MetadataUpdateOne {
	muo.mutation.AddDocumentTypeIDs(ids...)
	return muo
}

// AddDocumentTypes adds the "document_types" edges to the DocumentType entity.
func (muo *MetadataUpdateOne) AddDocumentTypes(d ...*DocumentType) *MetadataUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return muo.AddDocumentTypeIDs(ids...)
}

// SetDocumentID sets the "document" edge to the Document entity by ID.
func (muo *MetadataUpdateOne) SetDocumentID(id string) *MetadataUpdateOne {
	muo.mutation.SetDocumentID(id)
	return muo
}

// SetNillableDocumentID sets the "document" edge to the Document entity by ID if the given value is not nil.
func (muo *MetadataUpdateOne) SetNillableDocumentID(id *string) *MetadataUpdateOne {
	if id != nil {
		muo = muo.SetDocumentID(*id)
	}
	return muo
}

// SetDocument sets the "document" edge to the Document entity.
func (muo *MetadataUpdateOne) SetDocument(d *Document) *MetadataUpdateOne {
	return muo.SetDocumentID(d.ID)
}

// Mutation returns the MetadataMutation object of the builder.
func (muo *MetadataUpdateOne) Mutation() *MetadataMutation {
	return muo.mutation
}

// ClearTools clears all "tools" edges to the Tool entity.
func (muo *MetadataUpdateOne) ClearTools() *MetadataUpdateOne {
	muo.mutation.ClearTools()
	return muo
}

// RemoveToolIDs removes the "tools" edge to Tool entities by IDs.
func (muo *MetadataUpdateOne) RemoveToolIDs(ids ...int) *MetadataUpdateOne {
	muo.mutation.RemoveToolIDs(ids...)
	return muo
}

// RemoveTools removes "tools" edges to Tool entities.
func (muo *MetadataUpdateOne) RemoveTools(t ...*Tool) *MetadataUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return muo.RemoveToolIDs(ids...)
}

// ClearAuthors clears all "authors" edges to the Person entity.
func (muo *MetadataUpdateOne) ClearAuthors() *MetadataUpdateOne {
	muo.mutation.ClearAuthors()
	return muo
}

// RemoveAuthorIDs removes the "authors" edge to Person entities by IDs.
func (muo *MetadataUpdateOne) RemoveAuthorIDs(ids ...int) *MetadataUpdateOne {
	muo.mutation.RemoveAuthorIDs(ids...)
	return muo
}

// RemoveAuthors removes "authors" edges to Person entities.
func (muo *MetadataUpdateOne) RemoveAuthors(p ...*Person) *MetadataUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.RemoveAuthorIDs(ids...)
}

// ClearDocumentTypes clears all "document_types" edges to the DocumentType entity.
func (muo *MetadataUpdateOne) ClearDocumentTypes() *MetadataUpdateOne {
	muo.mutation.ClearDocumentTypes()
	return muo
}

// RemoveDocumentTypeIDs removes the "document_types" edge to DocumentType entities by IDs.
func (muo *MetadataUpdateOne) RemoveDocumentTypeIDs(ids ...int) *MetadataUpdateOne {
	muo.mutation.RemoveDocumentTypeIDs(ids...)
	return muo
}

// RemoveDocumentTypes removes "document_types" edges to DocumentType entities.
func (muo *MetadataUpdateOne) RemoveDocumentTypes(d ...*DocumentType) *MetadataUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return muo.RemoveDocumentTypeIDs(ids...)
}

// ClearDocument clears the "document" edge to the Document entity.
func (muo *MetadataUpdateOne) ClearDocument() *MetadataUpdateOne {
	muo.mutation.ClearDocument()
	return muo
}

// Where appends a list predicates to the MetadataUpdate builder.
func (muo *MetadataUpdateOne) Where(ps ...predicate.Metadata) *MetadataUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MetadataUpdateOne) Select(field string, fields ...string) *MetadataUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Metadata entity.
func (muo *MetadataUpdateOne) Save(ctx context.Context) (*Metadata, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MetadataUpdateOne) SaveX(ctx context.Context) *Metadata {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MetadataUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MetadataUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MetadataUpdateOne) sqlSave(ctx context.Context) (_node *Metadata, err error) {
	_spec := sqlgraph.NewUpdateSpec(metadata.Table, metadata.Columns, sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Metadata.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, metadata.FieldID)
		for _, f := range fields {
			if !metadata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != metadata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Version(); ok {
		_spec.SetField(metadata.FieldVersion, field.TypeString, value)
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.SetField(metadata.FieldName, field.TypeString, value)
	}
	if value, ok := muo.mutation.Date(); ok {
		_spec.SetField(metadata.FieldDate, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Comment(); ok {
		_spec.SetField(metadata.FieldComment, field.TypeString, value)
	}
	if muo.mutation.ToolsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.ToolsTable,
			Columns: []string{metadata.ToolsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedToolsIDs(); len(nodes) > 0 && !muo.mutation.ToolsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.ToolsTable,
			Columns: []string{metadata.ToolsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ToolsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.ToolsTable,
			Columns: []string{metadata.ToolsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.AuthorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.AuthorsTable,
			Columns: []string{metadata.AuthorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedAuthorsIDs(); len(nodes) > 0 && !muo.mutation.AuthorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.AuthorsTable,
			Columns: []string{metadata.AuthorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.AuthorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.AuthorsTable,
			Columns: []string{metadata.AuthorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.DocumentTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.DocumentTypesTable,
			Columns: []string{metadata.DocumentTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(documenttype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedDocumentTypesIDs(); len(nodes) > 0 && !muo.mutation.DocumentTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.DocumentTypesTable,
			Columns: []string{metadata.DocumentTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(documenttype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.DocumentTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.DocumentTypesTable,
			Columns: []string{metadata.DocumentTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(documenttype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.DocumentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   metadata.DocumentTable,
			Columns: []string{metadata.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(document.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.DocumentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   metadata.DocumentTable,
			Columns: []string{metadata.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(document.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Metadata{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
