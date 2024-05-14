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
)

type MetadataBackend struct {
	*Backend[*sbom.Metadata]
}

var _ storage.StoreRetriever[*sbom.Metadata] = (*MetadataBackend)(nil)

func (backend *MetadataBackend) Store(md *sbom.Metadata, opts *storage.StoreOptions) error {
	if backend.client == nil {
		return fmt.Errorf("failed creating Metadata, Setup Client required")
	}
	id, err := backend.client.Metadata.Create().
		SetID(md.Id).
		SetVersion(md.Version).
		SetName(md.Name).
		SetDate(md.Date.AsTime()).
		SetComment(md.Comment).
		OnConflict().
		Ignore().
		ID(backend.ctx)

	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("failed creating ent.Metadata: %w", err)
	}

	personBackend := &PersonBackend{NewBackend[*sbom.Person](), PersonBackendOptions{MetadataID: id}}
	for i := range md.Authors {
		if err := personBackend.Store(md.Authors[i], opts); err != nil {
			return fmt.Errorf("failed to store metadata authors: %w", err)
		}
	}

	docTypeBackend := &DocumentTypeBackend{NewBackend[*sbom.DocumentType](), DocumentTypeBackendOptions{MetadataID: id}}
	for i := range md.DocumentTypes {
		if err := docTypeBackend.Store(md.DocumentTypes[i], opts); err != nil {
			return fmt.Errorf("failed to store metadata document types: %w", err)
		}
	}

	toolBackend := &ToolBackend{NewBackend[*sbom.Tool](), ToolBackendOptions{MetadataID: id}}
	for i := range md.Tools {
		if err := toolBackend.Store(md.Tools[i], opts); err != nil {
			return fmt.Errorf("failed to store metadata tools: %w", err)
		}
	}

	return nil
}

func (backend *MetadataBackend) Retrieve(_id string, _opts *storage.RetrieveOptions) (*sbom.Metadata, error) {
	return nil, nil
}
