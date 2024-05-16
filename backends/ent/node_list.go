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

type NodeListBackend struct {
	*Backend[*sbom.NodeList]
}

var _ storage.StoreRetriever[*sbom.NodeList] = (*NodeListBackend)(nil)

func (backend *NodeListBackend) Store(nodeList *sbom.NodeList, _opts *storage.StoreOptions) error {
	id, err := backend.client.NodeList.Create().
		SetRootElements(nodeList.RootElements).
		OnConflict().
		Ignore().
		ID(backend.ctx)

	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("failed creating ent.NodeList: %w", err)
	}

	storeOpts, backendOpt := backend.getNodeBackendOpt(id)
	nodeBackend := NewNodeBackend(backendOpt)

	for _, n := range nodeList.Nodes {
		if err := nodeBackend.Store(n, storeOpts); err != nil {
			return fmt.Errorf("failed creating ent.Nodes: %w", err)
		}
	}

	return nil
}

func (backend *NodeListBackend) Retrieve(_id string, _opts *storage.RetrieveOptions) (*sbom.NodeList, error) {
	return nil, nil
}

func (backend *NodeListBackend) getNodeBackendOpt(id int) (*storage.StoreOptions, NodeBackendOptions) {
	backendOpt := NodeBackendOptions{
		NodeListID: id,
		BackendOptions: &BackendOptions[*sbom.Node]{
			DatabaseFile: backend.Options.DatabaseFile,
		},
	}

	storeOpt := &storage.StoreOptions{
		BackendOptions: backendOpt,
	}

	return storeOpt, backendOpt
}
