// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"fmt"

	"github.com/protobom/protobom/pkg/sbom"

	"github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/annotation"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

func (backend *Backend) createAnnotations(data ...*ent.Annotation) error {
	tx, err := backend.txClient()
	if err != nil {
		return err
	}

	builders := []*ent.AnnotationCreate{}

	for idx := range data {
		builder := tx.Annotation.Create().
			SetDocumentID(data[idx].DocumentID).
			SetName(data[idx].Name).
			SetValue(data[idx].Value).
			SetIsUnique(data[idx].IsUnique)

		builders = append(builders, builder)
	}

	err = tx.Annotation.CreateBulk(builders...).
		OnConflict().
		UpdateNewValues().
		Exec(backend.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return rollback(tx, fmt.Errorf("creating annotations: %w", err))
	}

	if err := tx.Commit(); err != nil {
		return rollback(tx, err)
	}

	return nil
}

// AddAnnotations applies multiple named annotation values to a single document.
func (backend *Backend) AddAnnotations(documentID, name string, values ...string) error {
	data := ent.Annotations{}
	for _, value := range values {
		data = append(data, &ent.Annotation{
			DocumentID: documentID,
			Name:       name,
			Value:      value,
		})
	}

	return backend.createAnnotations(data...)
}

// AddAnnotationToDocuments applies a single named annotation value to multiple documents.
func (backend *Backend) AddAnnotationToDocuments(name, value string, documentIDs ...string) error {
	data := ent.Annotations{}
	for _, documentID := range documentIDs {
		data = append(data, &ent.Annotation{
			DocumentID: documentID,
			Name:       name,
			Value:      value,
		})
	}

	return backend.createAnnotations(data...)
}

// ClearAnnotations removes all annotations from the specified documents.
func (backend *Backend) ClearAnnotations(documentIDs ...string) error {
	if len(documentIDs) == 0 {
		return nil
	}

	tx, err := backend.txClient()
	if err != nil {
		return err
	}

	_, err = tx.Annotation.Delete().Where(annotation.DocumentIDIn(documentIDs...)).Exec(backend.ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("clearing annotations: %w", err))
	}

	if err := tx.Commit(); err != nil {
		return rollback(tx, err)
	}

	return nil
}

// GetDocumentAnnotations gets all annotations for the specified
// document, limited to a set of annotation names if specified.
func (backend *Backend) GetDocumentAnnotations(documentID string, names ...string) (ent.Annotations, error) {
	if backend.client == nil {
		return nil, errUninitializedClient
	}

	predicates := []predicate.Annotation{
		annotation.HasDocumentWith(document.IDEQ(documentID)),
	}

	if len(names) > 0 {
		predicates = append(predicates, annotation.NameIn(names...))
	}

	annotations, err := backend.client.Annotation.Query().Where(predicates...).All(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying annotations: %w", err)
	}

	return annotations, nil
}

// GetDocumentsByAnnotation gets all documents having the specified named
// annotation, limited to a set of annotation values if specified.
func (backend *Backend) GetDocumentsByAnnotation(name string, values ...string) ([]*sbom.Document, error) {
	if backend.client == nil {
		return nil, errUninitializedClient
	}

	predicates := []predicate.Document{
		document.HasAnnotationsWith(annotation.NameEQ(name)),
	}

	if len(values) > 0 {
		predicates = append(predicates, document.HasAnnotationsWith(annotation.ValueIn(values...)))
	}

	ids, err := backend.client.Document.Query().Where(predicates...).IDs(backend.ctx)
	if err != nil {
		return nil, fmt.Errorf("querying documents table: %w", err)
	}

	return backend.GetDocumentsByID(ids...)
}

// GetDocumentUniqueAnnotation gets the value for a unique annotation.
func (backend *Backend) GetDocumentUniqueAnnotation(documentID, name string) (string, error) {
	if backend.client == nil {
		return "", errUninitializedClient
	}

	result, err := backend.client.Annotation.Query().
		Where(
			annotation.HasDocumentWith(document.IDEQ(documentID)),
			annotation.NameEQ(name),
			annotation.IsUniqueEQ(true),
		).
		Only(backend.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return "", nil
		}

		return "", fmt.Errorf("retrieving unique annotation for document: %w", err)
	}

	return result.Value, nil
}

// RemoveAnnotations removes all annotations with the specified name from
// the document, limited to a set of annotation values if specified.
func (backend *Backend) RemoveAnnotations(documentID, name string, values ...string) error {
	tx, err := backend.txClient()
	if err != nil {
		return err
	}

	predicates := []predicate.Annotation{
		annotation.DocumentIDEQ(documentID),
		annotation.NameEQ(name),
	}

	if len(values) > 0 {
		predicates = append(predicates, annotation.ValueIn(values...))
	}

	if _, err := tx.Annotation.Delete().Where(predicates...).Exec(backend.ctx); err != nil {
		return rollback(tx, fmt.Errorf("removing annotations: %w", err))
	}

	if err := tx.Commit(); err != nil {
		return rollback(tx, err)
	}

	return nil
}

// SetAnnotations explicitly sets the named annotations for the specified document.
func (backend *Backend) SetAnnotations(documentID, name string, values ...string) error {
	if err := backend.ClearAnnotations(documentID); err != nil {
		return err
	}

	return backend.AddAnnotations(documentID, name, values...)
}

// SetUniqueAnnotation sets a named annotation value that is unique to the specified document.
func (backend *Backend) SetUniqueAnnotation(documentID, name, value string) error {
	return backend.createAnnotations(&ent.Annotation{
		DocumentID: documentID,
		Name:       name,
		Value:      value,
		IsUnique:   true,
	})
}