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
	backend *backend.Backend[sbom.Purpose]
	backendSuite[sbom.Purpose]
}

func (ps *purposeSuite) BeforeTest(_suiteName, _testName string) {
	ps.backend = backend.NewBackend[sbom.Purpose]().WithDatabaseFile(ps.dbFile)

	if err := ps.backend.ClientSetup(); err != nil {
		ps.T().Fatalf("%v", err)
	}
}

func (*purposeSuite) TestPurposeBackend_Store() {}

func (*purposeSuite) TestPurposeBackend_Retrieve() {}

func TestPurposeBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(purposeSuite))
}
