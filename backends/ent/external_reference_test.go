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

type externalReferenceSuite struct {
	backendSuite[sbom.ExternalReference]
	backend *backend.Backend[sbom.ExternalReference]
}

func (suite *externalReferenceSuite) BeforeTest(_suiteName, _testName string) {
	suite.backend = backend.NewBackend[sbom.ExternalReference]().WithDatabaseFile(suite.dbFile)

	if err := suite.backend.ClientSetup(); err != nil {
		suite.T().Fatalf("%v", err)
	}
}

func (suite *externalReferenceSuite) TestExternalReferenceBackend_Store() {}

func (suite *externalReferenceSuite) TestExternalReferenceBackend_Retrieve() {}

func TestExternalReferenceBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(externalReferenceSuite))
}
