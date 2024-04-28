// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"slices"

	"github.com/bom-squad/protobom/pkg/sbom"
	sqlite "github.com/glebarez/go-sqlite"

	"github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/documenttype"
	"github.com/protobom/storage/internal/backends/ent/externalreference"
	"github.com/protobom/storage/internal/backends/ent/hashesentry"
	"github.com/protobom/storage/internal/backends/ent/identifiersentry"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/purpose"
	"github.com/protobom/storage/pkg/options"
)

// Enable SQLite foreign key support.
const dsnParams string = "?_pragma=foreign_keys(1)"

var errInvalidEntOptions = errors.New("invalid ent backend options")

// Store implements the model.v1.storage.Backend interface.
func (backend *Backend) Store(bom *sbom.Document, opts *options.StoreOptions) error {
	var err error

	if opts == nil {
		opts = &options.StoreOptions{
			BackendOptions: backend.Options,
		}
	}

	entOpts, ok := opts.BackendOptions.(BackendOptions)
	if !ok {
		return fmt.Errorf("%w: %v", errInvalidEntOptions, opts.BackendOptions)
	}

	tx, ctx, err := clientSetup(entOpts.DatabaseFile)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	defer tx.Client().Close()

	metadata := metadataToEnt(ctx, tx, bom.Metadata)
	nodeList := nodeListToEnt(ctx, tx, bom.NodeList)

	if metadata == nil || nodeList == nil {
		// Document already present in database.
		return nil
	}

	_, err = tx.Document.Create().SetMetadata(metadata).SetNodeList(nodeList).Save(ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("failed to save document: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit database transaction: %w", err)
	}

	return nil
}

func clientSetup(dbFile string) (*ent.Tx, context.Context, error) {
	// Register the SQLite driver as "sqlite3".
	if !slices.Contains(sql.Drivers(), "sqlite3") {
		sqlite.RegisterAsSQLITE3()
	}

	client, err := ent.Open("sqlite3", fmt.Sprintf("%s%s", dbFile, dsnParams))
	if err != nil {
		return nil, nil, fmt.Errorf("failed opening connection to sqlite: %w", err)
	}

	tx, err := client.Tx(context.Background())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create transactional client: %w", err)
	}

	ctx := ent.NewTxContext(context.Background(), tx)

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		return nil, nil, fmt.Errorf("failed creating schema resources: %w", err)
	}

	return tx, ctx, nil
}

func documentTypesToEnt(ctx context.Context, tx *ent.Tx, docTypes []*sbom.DocumentType) ent.DocumentTypes {
	entDocTypes := ent.DocumentTypes{}

	for _, docType := range docTypes {
		typeName := documenttype.Type(docType.Type.String())
		docTypeCreate := tx.DocumentType.Create().
			SetNillableName(docType.Name).
			SetNillableDescription(docType.Description).
			SetNillableType(&typeName)

		entDocType, err := docTypeCreate.Save(ctx)
		if err != nil && !ent.IsConstraintError(err) {
			panic(err)
		}

		entDocTypes = append(entDocTypes, entDocType)
	}

	return entDocTypes
}

func hashesToEnt(ctx context.Context, tx *ent.Tx, hashes map[int32]string) ent.HashesEntries {
	var entHashes ent.HashesEntries

	for alg, content := range hashes {
		algString := sbom.HashAlgorithm(alg).String()
		hashCreate := tx.HashesEntry.Create().
			SetHashAlgorithmType(hashesentry.HashAlgorithmType(algString)).
			SetHashData(content)

		entHash, err := hashCreate.Save(ctx)
		if err != nil && !ent.IsConstraintError(err) {
			panic(err)
		}

		entHashes = append(entHashes, entHash)
	}

	return entHashes
}

func identifiersToEnt(ctx context.Context, tx *ent.Tx, identifiers map[int32]string) ent.IdentifiersEntries {
	var entIdentifiers ent.IdentifiersEntries

	for typ, value := range identifiers {
		typeString := sbom.SoftwareIdentifierType(typ).String()
		identifierCreate := tx.IdentifiersEntry.Create().
			SetSoftwareIdentifierType(identifiersentry.SoftwareIdentifierType(typeString)).
			SetSoftwareIdentifierValue(value)

		entIdentifier, err := identifierCreate.Save(ctx)
		if err != nil && !ent.IsConstraintError(err) {
			panic(err)
		}

		entIdentifiers = append(entIdentifiers, entIdentifier)
	}

	return entIdentifiers
}

func metadataToEnt(ctx context.Context, tx *ent.Tx, md *sbom.Metadata) *ent.Metadata {
	metadataCreate := tx.Metadata.Create().
		SetID(md.Id).
		SetVersion(md.Version).
		SetName(md.Name).
		SetDate(md.Date.AsTime()).
		SetComment(md.Comment).
		AddAuthors(personsToEnt(ctx, tx, md.Authors)...).
		AddDocumentTypes(documentTypesToEnt(ctx, tx, md.DocumentTypes)...).
		AddTools(toolsToEnt(ctx, tx, md.Tools)...)

	metadata, err := metadataCreate.Save(ctx)
	if err != nil && !ent.IsConstraintError(err) {
		panic(err)
	}

	return metadata
}

func nodeListToEnt(ctx context.Context, tx *ent.Tx, nodeList *sbom.NodeList) *ent.NodeList {
	nodeListCreate := tx.NodeList.Create().SetRootElements(nodeList.RootElements)

	entNodeList, err := nodeListCreate.Save(ctx)
	if err != nil && !ent.IsConstraintError(err) {
		panic(err)
	}

	if entNodeList != nil {
		nodesToEnt(ctx, tx, nodeList.Nodes, entNodeList.ID)
	}

	return entNodeList
}

func nodesToEnt(ctx context.Context, tx *ent.Tx, nodes []*sbom.Node, nodeListID int) ent.Nodes {
	var entNodes ent.Nodes

	for _, n := range nodes {
		nodeCreate := tx.Node.Create().
			SetID(n.Id).
			SetNodeListID(nodeListID).
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
			AddExternalReferences(refsToEnt(ctx, tx, n.ExternalReferences, n.Id)...).
			AddHashes(hashesToEnt(ctx, tx, n.Hashes)...).
			AddIdentifiers(identifiersToEnt(ctx, tx, n.Identifiers)...).
			AddOriginators(personsToEnt(ctx, tx, n.Originators)...).
			AddSuppliers(personsToEnt(ctx, tx, n.Suppliers)...).
			AddPrimaryPurpose(purposesToEnt(ctx, tx, n.PrimaryPurpose)...).
			// TODO
			AddEdgeTypes().
			AddNodes()

		entNode, err := nodeCreate.Save(ctx)
		if err != nil {
			if !ent.IsConstraintError(err) {
				panic(err)
			}

			continue
		}

		entNodes = append(entNodes, entNode)
	}

	return entNodes
}

func personsToEnt(ctx context.Context, tx *ent.Tx, persons []*sbom.Person) ent.Persons {
	var entPersons ent.Persons

	for _, person := range persons {
		entPerson, err := tx.Person.Create().
			SetName(person.Name).
			SetEmail(person.Email).
			SetIsOrg(person.IsOrg).
			SetPhone(person.Phone).
			SetURL(person.Url).
			AddContacts(personsToEnt(ctx, tx, person.Contacts)...).
			Save(ctx)

		if err != nil && !ent.IsConstraintError(err) {
			panic(err)
		}

		entPersons = append(entPersons, entPerson)
	}

	return entPersons
}

func purposesToEnt(ctx context.Context, tx *ent.Tx, purposes []sbom.Purpose) ent.Purposes {
	var entPurposes ent.Purposes

	for _, p := range purposes {
		purposeCreate := tx.Purpose.Create().
			SetPrimaryPurpose(purpose.PrimaryPurpose(p.String()))

		entPurpose, err := purposeCreate.Save(ctx)
		if err != nil && !ent.IsConstraintError(err) {
			panic(err)
		}

		entPurposes = append(entPurposes, entPurpose)
	}

	return entPurposes
}

func refsToEnt(ctx context.Context, tx *ent.Tx, refs []*sbom.ExternalReference, nodeID string) ent.ExternalReferences {
	var entRefs ent.ExternalReferences

	for _, r := range refs {
		extRefCreate := tx.ExternalReference.Create().
			SetNodeID(nodeID).
			SetAuthority(r.Authority).
			SetComment(r.Comment).
			SetType(externalreference.Type(r.Type.String())).
			SetURL(r.Url).
			AddHashes(hashesToEnt(ctx, tx, r.Hashes)...)

		entRef, err := extRefCreate.Save(ctx)
		if err != nil {
			if !ent.IsConstraintError(err) {
				panic(err)
			}

			continue
		}

		entRefs = append(entRefs, entRef)
	}

	return entRefs
}

func toolsToEnt(ctx context.Context, tx *ent.Tx, tools []*sbom.Tool) ent.Tools {
	var entTools ent.Tools

	for _, tool := range tools {
		entTool, err := tx.Tool.Create().
			SetName(tool.Name).
			SetVendor(tool.Vendor).
			SetVersion(tool.Version).
			Save(ctx)

		if err != nil && !ent.IsConstraintError(err) {
			panic(err)
		}

		entTools = append(entTools, entTool)
	}

	return entTools
}
