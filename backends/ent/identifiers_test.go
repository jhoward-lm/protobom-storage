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

type identifiersSuite struct {
	backendSuite[map[sbom.SoftwareIdentifierType]string]
	backend *backend.Backend[map[sbom.SoftwareIdentifierType]string]
}

func (suite *identifiersSuite) BeforeTest(_suiteName, _testName string) {
	suite.backend = backend.NewBackend[map[sbom.SoftwareIdentifierType]string]().WithDatabaseFile(suite.dbFile)

	if err := suite.backend.ClientSetup(); err != nil {
		suite.T().Fatalf("%v", err)
	}
}

func (suite *identifiersSuite) TestIdentifiersBackend_Store() {}

func (suite *identifiersSuite) TestIdentifiersBackend_Retrieve() {}

func TestIdentifiersBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(identifiersSuite))
}
