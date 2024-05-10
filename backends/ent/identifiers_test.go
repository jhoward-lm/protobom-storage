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
	backend *backend.Backend[map[sbom.SoftwareIdentifierType]string]
	backendSuite[map[sbom.SoftwareIdentifierType]string]
}

func (is *identifiersSuite) BeforeTest(_suiteName, _testName string) {
	is.backend = backend.NewBackend[map[sbom.SoftwareIdentifierType]string]().WithDatabaseFile(is.dbFile)

	if err := is.backend.ClientSetup(); err != nil {
		is.T().Fatalf("%v", err)
	}
}

func (*identifiersSuite) TestIdentifiersBackend_Store() {}

func (*identifiersSuite) TestIdentifiersBackend_Retrieve() {}

func TestIdentifiersBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(identifiersSuite))
}
