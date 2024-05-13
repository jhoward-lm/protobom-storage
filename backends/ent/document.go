// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"
)

type DocumentBackend struct {
	*Backend[*sbom.Document]
}

var _ storage.StoreRetriever[*sbom.Document] = (*DocumentBackend)(nil)

func (backend *DocumentBackend) Store(_bom *sbom.Document, _opts *storage.StoreOptions) error {
	return nil
}

func (backend *DocumentBackend) Retrieve(_id string, _opts *storage.RetrieveOptions) (*sbom.Document, error) {
	return nil, nil
}
