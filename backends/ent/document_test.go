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

type documentSuite struct {
	backendSuite[sbom.Document]
	backend *backend.Backend[sbom.Document]
}

func (suite *documentSuite) BeforeTest(_suiteName, _testName string) {
	suite.backend = backend.NewBackend[sbom.Document]().WithDatabaseFile(suite.dbFile)

	if err := suite.backend.ClientSetup(); err != nil {
		suite.T().Fatalf("%v", err)
	}
}

func (suite *documentSuite) TestDocumentBackend_Store() {}

func (suite *documentSuite) TestDocumentBackend_Retrieve() {}

func TestDocumentBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(documentSuite))
}
