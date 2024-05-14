// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent //nolint:dupl

import (
	"fmt"

	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"

	"github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/identifiersentry"
)

type IdentifiersBackend struct {
	*Backend[map[sbom.SoftwareIdentifierType]string]
}

var _ storage.StoreRetriever[map[sbom.SoftwareIdentifierType]string] = (*IdentifiersBackend)(nil)

func (backend *IdentifiersBackend) Store(
	idents map[sbom.SoftwareIdentifierType]string,
	_opts *storage.StoreOptions,
) error {
	if backend.client == nil {
		return fmt.Errorf("failed creating IdentifiersEntry, Setup Client required")
	}
	for typ, value := range idents {
		err := backend.client.IdentifiersEntry.Create().
			SetSoftwareIdentifierType(identifiersentry.SoftwareIdentifierType(typ.String())).
			SetSoftwareIdentifierValue(value).
			OnConflict().
			DoNothing().
			Exec(backend.ctx)

		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("failed creating ent.IdentifiersEntry: %w", err)
		}
	}

	return nil
}

func (backend *IdentifiersBackend) Retrieve(
	_id string,
	_opts *storage.RetrieveOptions,
) (map[sbom.SoftwareIdentifierType]string, error) {
	return nil, nil
}
