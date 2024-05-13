// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"fmt"

	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"

	"github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/documenttype"
)

type (
	DocumentTypeBackend struct {
		*Backend[*sbom.DocumentType]
		Options DocumentTypeBackendOptions
	}

	DocumentTypeBackendOptions struct {
		*BackendOptions[*sbom.DocumentType]
		MetadataID string
	}
)

var _ storage.StoreRetriever[*sbom.DocumentType] = (*DocumentTypeBackend)(nil)

func (backend *DocumentTypeBackend) Store(docType *sbom.DocumentType, opts *storage.StoreOptions) error {
	entOpts, ok := opts.BackendOptions.(DocumentTypeBackendOptions)
	if !ok || entOpts.MetadataID == "" {
		return fmt.Errorf("%w: %v", errInvalidEntOptions, opts.BackendOptions)
	}

	typeName := documenttype.Type(docType.Type.String())
	err := backend.client.DocumentType.Create().
		SetNillableName(docType.Name).
		SetNillableDescription(docType.Description).
		SetNillableType(&typeName).
		SetMetadataID(entOpts.MetadataID).
		Exec(backend.ctx)

	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("failed creating ent.DocumentType: %w", err)
	}

	return nil
}

func (backend *DocumentTypeBackend) Retrieve(_id string, _opts *storage.RetrieveOptions) (*sbom.DocumentType, error) {
	return nil, nil
}

func (backend *DocumentTypeBackend) WithMetadataID(id string) *DocumentTypeBackend {
	backend.Options.MetadataID = id

	return backend
}
