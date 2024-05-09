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

type metadataSuite struct {
	backendSuite[sbom.Metadata]
	backend *backend.Backend[sbom.Metadata]
}

func (suite *metadataSuite) BeforeTest(_suiteName, _testName string) {
	suite.backend = backend.NewBackend[sbom.Metadata]().WithDatabaseFile(suite.dbFile)

	if err := suite.backend.ClientSetup(); err != nil {
		suite.T().Fatalf("%v", err)
	}
}

func (suite *metadataSuite) TestMetadataBackend_Store() {}

func (suite *metadataSuite) TestMetadataBackend_Retrieve() {}

func TestMetadataBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(metadataSuite))
}
