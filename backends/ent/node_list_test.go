// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent_test

import (
	"testing"

	"github.com/protobom/protobom/pkg/sbom"
	"github.com/stretchr/testify/suite"

	backend "github.com/protobom/storage/backends/ent"
)

type nodeListSuite struct {
	backendSuite[sbom.NodeList]
	backend *backend.Backend[sbom.NodeList]
}

func (suite *nodeListSuite) BeforeTest(_suiteName, _testName string) {
	suite.backend = backend.NewBackend[sbom.NodeList]().WithDatabaseFile(suite.dbFile)

	if err := suite.backend.ClientSetup(); err != nil {
		suite.T().Fatalf("%v", err)
	}
}

func (suite *nodeListSuite) TestNodeListBackend_Store() {}

func (suite *nodeListSuite) TestNodeListBackend_Retrieve() {}

func TestNodeListBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(nodeListSuite))
}
