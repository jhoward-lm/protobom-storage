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

type purposeSuite struct {
	backendSuite[sbom.Purpose]
	backend *backend.Backend[sbom.Purpose]
}

func (suite *purposeSuite) BeforeTest(_suiteName, _testName string) {
	suite.backend = backend.NewBackend[sbom.Purpose]().WithDatabaseFile(suite.dbFile)

	if err := suite.backend.ClientSetup(); err != nil {
		suite.T().Fatalf("%v", err)
	}
}

func (suite *purposeSuite) TestPurposeBackend_Store() {}

func (suite *purposeSuite) TestPurposeBackend_Retrieve() {}

func TestPurposeBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(purposeSuite))
}
