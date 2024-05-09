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

type toolSuite struct {
	backendSuite[sbom.Tool]
	backend *backend.Backend[sbom.Tool]
}

func (suite *toolSuite) BeforeTest(_suiteName, _testName string) {
	suite.backend = backend.NewBackend[sbom.Tool]().WithDatabaseFile(suite.dbFile)

	if err := suite.backend.ClientSetup(); err != nil {
		suite.T().Fatalf("%v", err)
	}
}

func (suite *toolSuite) TestToolBackend_Store() {}

func (suite *toolSuite) TestToolBackend_Retrieve() {}

func TestToolBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(toolSuite))
}
