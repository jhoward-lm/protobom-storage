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

type nodeSuite struct {
	backendSuite[sbom.Node]
	backend *backend.Backend[sbom.Node]
}

func (suite *nodeSuite) BeforeTest(_suiteName, _testName string) {
	suite.backend = backend.NewBackend[sbom.Node]().WithDatabaseFile(suite.dbFile)

	if err := suite.backend.ClientSetup(); err != nil {
		suite.T().Fatalf("%v", err)
	}
}

func (suite *nodeSuite) TestNodeBackend_Store() {}

func (suite *nodeSuite) TestNodeBackend_Retrieve() {}

func TestNodeBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(nodeSuite))
}
