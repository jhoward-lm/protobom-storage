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

type documentTypeSuite struct {
	backend *backend.DocumentTypeBackend
	backendSuite[*sbom.DocumentType]
}

func (dts *documentTypeSuite) BeforeTest(_suiteName, _testName string) {
	dts.backend = &backend.DocumentTypeBackend{
		Backend: backend.NewBackend[*sbom.DocumentType]().WithDatabaseFile(dts.dbFile),
	}

	if err := dts.backend.ClientSetup(); err != nil {
		dts.T().Fatalf("%v", err)
	}
}

func (*documentTypeSuite) TestDocumentTypeBackend_Store() {}

func (*documentTypeSuite) TestDocumentTypeBackend_Retrieve() {}

func TestDocumentTypeBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(documentTypeSuite))
}
