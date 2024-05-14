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
	"github.com/protobom/storage/internal/backends/ent/purpose"
)

type PurposeBackend struct {
	*Backend[*sbom.Purpose]
}

var _ storage.StoreRetriever[*sbom.Purpose] = (*PurposeBackend)(nil)

func (backend *PurposeBackend) Store(p *sbom.Purpose, _opts *storage.StoreOptions) error {
	if backend.client == nil {
		return fmt.Errorf("failed creating Purpose, Setup Client required")
	}
	err := backend.client.Purpose.Create().
		SetPrimaryPurpose(purpose.PrimaryPurpose(p.String())).
		OnConflict().
		Ignore().
		Exec(backend.ctx)

	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("failed creating ent.Purpose: %w", err)
	}

	return nil
}

func (backend *PurposeBackend) Retrieve(_id string, _opts *storage.RetrieveOptions) (*sbom.Purpose, error) {
	return nil, nil
}
