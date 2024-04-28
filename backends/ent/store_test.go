// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	backend "github.com/protobom/storage/backends/ent"
	"github.com/protobom/storage/internal/backends/ent"
)

type storeSuite struct {
	suite.Suite
	tx  *ent.Tx
	ctx context.Context
}

func (ss *storeSuite) BeforeTest(_suiteName, _testName string) {
	tx, ctx, err := backend.ClientSetup(":memory:")
	if err != nil {
		ss.T().Fatalf("%v", err)
	}

	ss.ctx = ctx
	ss.tx = tx
}

func (ss *storeSuite) AfterTest(_suiteName, _testName string) {
	ss.tx.Client().Close()
}

func (ss *storeSuite) TestBackendStore() {}

func (ss *storeSuite) TestDocumentTypesToEnt() {}

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
