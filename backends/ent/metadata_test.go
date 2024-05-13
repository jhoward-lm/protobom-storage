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
	backend *backend.MetadataBackend
	backendSuite[*sbom.Metadata]
}

func (ms *metadataSuite) BeforeTest(_suiteName, _testName string) {
	ms.backend = &backend.MetadataBackend{
		Backend: backend.NewBackend[*sbom.Metadata]().WithDatabaseFile(ms.dbFile),
	}

	if err := ms.backend.ClientSetup(); err != nil {
		ms.T().Fatalf("%v", err)
	}
}

func (*metadataSuite) TestMetadataBackend_Store() {}

func (*metadataSuite) TestMetadataBackend_Retrieve() {}

func TestMetadataBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(metadataSuite))
}
