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
	backend *backend.DocumentBackend
	backendSuite[*sbom.Document]
}

func (ds *documentSuite) BeforeTest(_suiteName, _testName string) {
	ds.backend = &backend.DocumentBackend{
		Backend: backend.NewBackend[*sbom.Document]().WithDatabaseFile(ds.dbFile),
	}

	if err := ds.backend.ClientSetup(); err != nil {
		ds.T().Fatalf("%v", err)
	}
}

func (*documentSuite) TestDocumentBackend_Store() {}

func (*documentSuite) TestDocumentBackend_Retrieve() {}

func TestDocumentBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(documentSuite))
}
