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

type hashesSuite struct {
	backendSuite[map[sbom.HashAlgorithm]string]
	backend *backend.Backend[map[sbom.HashAlgorithm]string]
}

func (suite *hashesSuite) BeforeTest(_suiteName, _testName string) {
	suite.backend = backend.NewBackend[map[sbom.HashAlgorithm]string]().WithDatabaseFile(suite.dbFile)

	if err := suite.backend.ClientSetup(); err != nil {
		suite.T().Fatalf("%v", err)
	}
}

func (suite *hashesSuite) TestHashesBackend_Store() {}

func (suite *hashesSuite) TestHashesBackend_Retrieve() {}

func TestHashesBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(hashesSuite))
}
