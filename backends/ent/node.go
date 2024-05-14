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
	"github.com/protobom/storage/internal/backends/ent/node"
)

type (
	NodeBackend struct {
		*Backend[*sbom.Node]
		Options NodeBackendOptions
	}

	NodeBackendOptions struct {
		*BackendOptions[*sbom.Node]
		NodeListID int
	}
)

var _ storage.StoreRetriever[*sbom.Node] = (*NodeBackend)(nil)

func (backend *NodeBackend) SetClient(client *ent.Client) {
	backend.client = client
}

func (backend *NodeBackend) Store(n *sbom.Node, _opts *storage.StoreOptions) error {
	if backend.client == nil {
		return fmt.Errorf("failed creating NodeList, Setup Client required")
	}
	err := backend.client.Node.Create().
		SetID(n.Id).
		SetAttribution(n.Attribution).
		SetBuildDate(n.BuildDate.AsTime()).
		SetComment(n.Comment).
		SetCopyright(n.Copyright).
		SetDescription(n.Description).
		SetFileName(n.FileName).
		SetFileTypes(n.FileTypes).
		SetLicenseComments(n.LicenseComments).
		SetLicenseConcluded(n.LicenseConcluded).
		SetLicenses(n.Licenses).
		SetName(n.Name).
		SetReleaseDate(n.ReleaseDate.AsTime()).
		SetSourceInfo(n.SourceInfo).
		SetSummary(n.Summary).
		SetType(node.Type(n.Type.String())).
		SetURLDownload(n.UrlDownload).
		SetURLHome(n.UrlHome).
		SetValidUntilDate(n.ValidUntilDate.AsTime()).
		SetVersion(n.Version).
		// TODO
		// AddExternalReferences(refsToEnt(ctx, client, n.ExternalReferences, n.Id)...).
		// AddHashes(hashesToEnt(ctx, client, n.Hashes)...).
		// AddIdentifiers(identifiersToEnt(ctx, client, n.Identifiers)...).
		// AddOriginators(personsToEnt(ctx, client, n.Originators)...).
		// AddSuppliers(personsToEnt(ctx, client, n.Suppliers)...).
		// AddPrimaryPurpose(purposesToEnt(ctx, client, n.PrimaryPurpose)...).
		// AddEdgeTypes().
		// AddNodes().
		OnConflict().
		Ignore().
		Exec(backend.ctx)

	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("failed creating ent.NodeList: %w", err)
	}

	return nil
}

func (backend *NodeBackend) Retrieve(_id string, _opts *storage.RetrieveOptions) (*sbom.Node, error) {
	return nil, nil
}

func WithNodeListID(id int) Option[*sbom.Node] {
	return func(backend *Backend[*sbom.Node]) {
		nodeBackend := &NodeBackend{backend, NodeBackendOptions{NodeListID: id}}
		nodeBackend.WithNodeListID(id)
	}
}

func (backend *NodeBackend) WithNodeListID(id int) *NodeBackend {
	backend.Options.NodeListID = id

	return backend
}
