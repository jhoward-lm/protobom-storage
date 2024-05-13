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
	"github.com/protobom/storage/internal/backends/ent/externalreference"
)

type (
	ExternalReferenceBackend struct {
		*Backend[*sbom.ExternalReference]
		Options ExternalReferenceBackendOptions
	}

	ExternalReferenceBackendOptions struct {
		BackendOptions[*sbom.ExternalReference]
		NodeID string
	}
)

var _ storage.StoreRetriever[*sbom.ExternalReference] = (*ExternalReferenceBackend)(nil)

func (backend *ExternalReferenceBackend) Store(ref *sbom.ExternalReference, opts *storage.StoreOptions) error {
	hashes := make(map[sbom.HashAlgorithm]string)
	for alg, content := range ref.Hashes {
		hashes[sbom.HashAlgorithm(alg)] = content
	}

	hashesBackend := HashesBackend{}
	if err := hashesBackend.Store(hashes, opts); err != nil {
		return fmt.Errorf("failed to store external reference hashes: %w", err)
	}

	err := backend.client.ExternalReference.Create().
		// SetNodeID(nodeID).
		SetAuthority(ref.Authority).
		SetComment(ref.Comment).
		SetType(externalreference.Type(ref.Type.String())).
		SetURL(ref.Url).
		// AddHashes(hashesToEnt(ctx, client, ref.Hashes)...).
		OnConflict().
		Ignore().
		Exec(backend.ctx)

	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("failed creating ent.ExternalReference: %w", err)
	}

	return nil
}

func (backend *ExternalReferenceBackend) Retrieve(
	_id string,
	_opts *storage.RetrieveOptions,
) (*sbom.ExternalReference, error) {
	return nil, nil
}

func WithNodeID(id string) Option[*sbom.ExternalReference] {
	return func(backend *Backend[*sbom.ExternalReference]) {
		refBackend := &ExternalReferenceBackend{backend, ExternalReferenceBackendOptions{NodeID: id}}
		refBackend.WithNodeID(id)
	}
}

func (backend *ExternalReferenceBackend) WithNodeID(id string) *ExternalReferenceBackend {
	backend.Options.NodeID = id

	return backend
}
