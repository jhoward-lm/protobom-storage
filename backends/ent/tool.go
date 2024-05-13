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

type (
	ToolBackend struct {
		*Backend[*sbom.Tool]
		Options ToolBackendOptions
	}

	ToolBackendOptions struct {
		*BackendOptions[*sbom.Tool]
		MetadataID string
	}
)

var _ storage.StoreRetriever[*sbom.Tool] = (*ToolBackend)(nil)

func (backend *ToolBackend) Store(tool *sbom.Tool, opts *storage.StoreOptions) error {
	entOpts, ok := opts.BackendOptions.(ToolBackendOptions)
	if !ok || entOpts.MetadataID == "" {
		return fmt.Errorf("%w: %v", errInvalidEntOptions, opts.BackendOptions)
	}

	err := backend.client.Tool.Create().
		SetName(tool.Name).
		SetVendor(tool.Vendor).
		SetVersion(tool.Version).
		SetMetadataID(entOpts.MetadataID).
		OnConflict().
		Ignore().
		Exec(backend.ctx)

	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("failed creating ent.Tool: %w", err)
	}

	return nil
}

func (backend *ToolBackend) Retrieve(_id string, _opts *storage.RetrieveOptions) (*sbom.Tool, error) {
	return nil, nil
}

func (backend *ToolBackend) WithMetadataID(id string) *ToolBackend {
	backend.Options.MetadataID = id

	return backend
}
