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
	PersonBackend struct {
		*Backend[sbom.Person]
		Options PersonBackendOptions
	}

	PersonBackendOptions struct {
		BackendOptions
		MetadataID string
	}
)

var _ storage.Backend[sbom.Person] = (*PersonBackend)(nil)

func (backend *PersonBackend) Store(person *sbom.Person, opts *storage.StoreOptions) error {
	for i := range person.Contacts {
		if err := backend.Store(person.Contacts[i], opts); err != nil {
			return fmt.Errorf("failed to store person contact: %w", err)
		}
	}

	entOpts, ok := opts.BackendOptions.(PersonBackendOptions)
	if !ok || entOpts.MetadataID == "" {
		return fmt.Errorf("%w: %v", errInvalidEntOptions, opts.BackendOptions)
	}

	err := backend.client.Person.Create().
		SetName(person.Name).
		SetEmail(person.Email).
		SetIsOrg(person.IsOrg).
		SetPhone(person.Phone).
		SetURL(person.Url).
		SetMetadataID(entOpts.MetadataID).
		OnConflict().
		Ignore().
		Exec(backend.ctx)

	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("failed creating ent.Person: %w", err)
	}

	return nil
}

func (backend *PersonBackend) Retrieve(_id string, _opts *storage.RetrieveOptions) (*sbom.Person, error) {
	return nil, nil
}

func (backend *PersonBackend) WithMetadataID(id string) *PersonBackend {
	backend.Options.MetadataID = id

	return backend
}
