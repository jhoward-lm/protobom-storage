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

type NodeListBackend Backend[sbom.NodeList]

var _ storage.Backend[sbom.NodeList] = (*NodeListBackend)(nil)

func (backend *NodeListBackend) Store(nodeList *sbom.NodeList, _opts *storage.StoreOptions) error {
	err := backend.client.NodeList.Create().
		SetRootElements(nodeList.RootElements).
		OnConflict().
		Ignore().
		Exec(backend.ctx)

	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("failed creating ent.NodeList: %w", err)
	}

	return nil
}

func (backend *NodeListBackend) Retrieve(_id string, _opts *storage.RetrieveOptions) (*sbom.NodeList, error) {
	return nil, nil
}
