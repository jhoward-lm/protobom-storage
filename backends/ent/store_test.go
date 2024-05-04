// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/bom-squad/protobom/pkg/reader"
	"github.com/bom-squad/protobom/pkg/sbom"
	"github.com/stretchr/testify/suite"

	backend "github.com/protobom/storage/backends/ent"
	"github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/documenttype"
	"github.com/protobom/storage/internal/backends/ent/metadata"
)

type storeSuite struct {
	suite.Suite
	dbFile    string
	ctx       context.Context
	backend   *backend.Backend
	client    *ent.Client
	documents []*sbom.Document
}

func (ss *storeSuite) SetupSuite() {
	sbomReader := reader.New()

	sboms, err := os.ReadDir("testdata")
	if err != nil {
		ss.T().Fatalf("%v", err)
	}

	for sbomIdx := range sboms {
		doc, err := sbomReader.ParseFile(filepath.Join("testdata", sboms[sbomIdx].Name()))
		if err != nil {
			ss.T().Fatalf("%v", err)
		}

		ss.documents = append(ss.documents, doc)
	}

	ss.dbFile = filepath.Join(ss.T().TempDir(), "test.db")

	client, ctx, err := backend.ClientSetup(ss.dbFile)
	if err != nil {
		ss.T().Fatalf("%v", err)
	}

	ss.ctx = ctx
	ss.client = client
	ss.backend = backend.NewBackend().WithDatabaseFile(ss.dbFile)
}

func (ss *storeSuite) TearDownSuite() {
	ss.client.Close()
}

func (ss *storeSuite) TestBackend_Store() {
	for docIdx := range ss.documents {
		ss.Run(ss.documents[docIdx].Metadata.Id, func() {
			ss.T().Parallel()

			err := ss.backend.Store(ss.documents[docIdx], nil)
			if err != nil {
				ss.T().Fatalf("%v", err)
			}

			dbDocument := ss.client.Document.Query().
				Where(document.HasMetadataWith(metadata.ID(ss.documents[docIdx].Metadata.Id))).
				FirstX(ss.ctx)

			ss.Require().NotNil(dbDocument)

			dbNodes := dbDocument.QueryNodeList().QueryNodes().AllX(ss.ctx)

			ss.Require().Len(dbNodes, len(ss.documents[docIdx].NodeList.Nodes))
		})
	}
}

func (ss *storeSuite) TestDocumentTypesToEnt() {
	docTypes := []*sbom.DocumentType{}
	expected := ent.DocumentTypes{}

	for docIdx := range ss.documents {
		docTypes = append(docTypes, ss.documents[docIdx].Metadata.DocumentTypes...)
	}

	testData := []struct {
		name        string
		description string
		docType     sbom.DocumentType_SBOMType
	}{
		{name: "other", description: "DocumentType_OTHER", docType: sbom.DocumentType_OTHER},
		{name: "design", description: "DocumentType_DESIGN", docType: sbom.DocumentType_DESIGN},
		{name: "source", description: "DocumentType_SOURCE", docType: sbom.DocumentType_SOURCE},
		{name: "build", description: "DocumentType_BUILD", docType: sbom.DocumentType_BUILD},
		{name: "analyzed", description: "DocumentType_ANALYZED", docType: sbom.DocumentType_ANALYZED},
		{name: "deployed", description: "DocumentType_DEPLOYED", docType: sbom.DocumentType_DEPLOYED},
		{name: "runtime", description: "DocumentType_RUNTIME", docType: sbom.DocumentType_RUNTIME},
		{name: "discovery", description: "DocumentType_DISCOVERY", docType: sbom.DocumentType_DISCOVERY},
		{name: "decomission", description: "DocumentType_DECOMISSION", docType: sbom.DocumentType_DECOMISSION},
	}

	for idx := range testData {
		docTypes = append(docTypes, &sbom.DocumentType{
			Name:        &testData[idx].name,
			Description: &testData[idx].description,
			Type:        &testData[idx].docType,
		})
	}

	dbData := []struct {
		name        string
		description string
		docType     documenttype.Type
		id          int
	}{
		{name: "other", description: "DocumentType_OTHER", docType: documenttype.TypeOTHER, id: 1},
		{name: "design", description: "DocumentType_DESIGN", docType: documenttype.TypeDESIGN, id: 2},
		{name: "source", description: "DocumentType_SOURCE", docType: documenttype.TypeSOURCE, id: 3},
		{name: "build", description: "DocumentType_BUILD", docType: documenttype.TypeBUILD, id: 4},
		{name: "analyzed", description: "DocumentType_ANALYZED", docType: documenttype.TypeANALYZED, id: 5},
		{name: "deployed", description: "DocumentType_DEPLOYED", docType: documenttype.TypeDEPLOYED, id: 6},
		{name: "runtime", description: "DocumentType_RUNTIME", docType: documenttype.TypeRUNTIME, id: 7},
		{name: "discovery", description: "DocumentType_DISCOVERY", docType: documenttype.TypeDISCOVERY, id: 8},
		{name: "decomission", description: "DocumentType_DECOMISSION", docType: documenttype.TypeDECOMISSION, id: 9},
	}

	for idx := range dbData {
		expected = append(expected, &ent.DocumentType{
			Name:        &dbData[idx].name,
			Description: &dbData[idx].description,
			Type:        &dbData[idx].docType,
			ID:          dbData[idx].id,
		})
	}

	actual := backend.DocumentTypesToEnt(ss.ctx, ss.client, docTypes)

	queryResult := ent.DocumentTypes(ss.client.DocumentType.Query().
		Select("id", "type", "name", "description").
		Where(documenttype.IDLTE(len(expected))).
		AllX(ss.ctx))

	ss.Require().Len(expected, len(actual))
	ss.Require().Len(expected, len(queryResult))

	for i := range expected {
		ss.Assert().Equal(expected[i].ID, actual[i].ID)
		ss.Assert().Equal(expected[i].ID, queryResult[i].ID)

		ss.Assert().Equal(expected[i].Name, actual[i].Name)
		ss.Assert().Equal(expected[i].Name, queryResult[i].Name)

		ss.Assert().Equal(expected[i].Type, actual[i].Type)
		ss.Assert().Equal(expected[i].Type, queryResult[i].Type)

		ss.Assert().Equal(expected[i].Description, actual[i].Description)
		ss.Assert().Equal(expected[i].Description, queryResult[i].Description)
	}
}

func (ss *storeSuite) TestHashesToEnt() {}

func (ss *storeSuite) TestIdentifiersToEnt() {}

func (ss *storeSuite) TestMetadataToEnt() {}

func (ss *storeSuite) TestNodeListToEnt() {}

func (ss *storeSuite) TestNodesToEnt() {}

func (ss *storeSuite) TestPersonsToEnt() {}

func (ss *storeSuite) TestPurposesToEnt() {}

func (ss *storeSuite) TestRefsToEnt() {}

func (ss *storeSuite) TestToolsToEnt() {}

func TestStoreSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(storeSuite))
}
