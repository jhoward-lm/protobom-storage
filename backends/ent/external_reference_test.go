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
	backend *backend.Backend[sbom.ExternalReference]
	backendSuite[sbom.ExternalReference]
}

func (ers *externalReferenceSuite) BeforeTest(_suiteName, _testName string) {
	ers.backend = backend.NewBackend[sbom.ExternalReference]().WithDatabaseFile(ers.dbFile)

	if err := ers.backend.ClientSetup(); err != nil {
		ers.T().Fatalf("%v", err)
	}
}

func (*externalReferenceSuite) TestExternalReferenceBackend_Store() {}

func (*externalReferenceSuite) TestExternalReferenceBackend_Retrieve() {}

func TestExternalReferenceBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(externalReferenceSuite))
}
