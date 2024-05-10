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

type IdentifiersBackend Backend[map[sbom.SoftwareIdentifierType]string]

var _ storage.Backend[map[sbom.SoftwareIdentifierType]string] = (*IdentifiersBackend)(nil)

func (backend *IdentifiersBackend) Store(
	idents *map[sbom.SoftwareIdentifierType]string, //nolint: gocritic // Ignore ptrToRefParam rule
	_opts *storage.StoreOptions,
) error {
	for typ, value := range *idents {
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
) (*map[sbom.SoftwareIdentifierType]string, error) { //nolint: gocritic // Ignore ptrToRefParam rule
	return nil, nil
}
