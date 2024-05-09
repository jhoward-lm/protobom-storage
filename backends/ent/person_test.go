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

type personSuite struct {
	backendSuite[sbom.Person]
	backend *backend.Backend[sbom.Person]
}

func (suite *personSuite) BeforeTest(_suiteName, _testName string) {
	suite.backend = backend.NewBackend[sbom.Person]().WithDatabaseFile(suite.dbFile)

	if err := suite.backend.ClientSetup(); err != nil {
		suite.T().Fatalf("%v", err)
	}
}

func (suite *personSuite) TestPersonBackend_Store() {}

func (suite *personSuite) TestPersonBackend_Retrieve() {}

func TestPersonBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(personSuite))
}
