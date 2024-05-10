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
	backend *backend.Backend[map[sbom.HashAlgorithm]string]
	backendSuite[map[sbom.HashAlgorithm]string]
}

func (hs *hashesSuite) BeforeTest(_suiteName, _testName string) {
	hs.backend = backend.NewBackend[map[sbom.HashAlgorithm]string]().WithDatabaseFile(hs.dbFile)

	if err := hs.backend.ClientSetup(); err != nil {
		hs.T().Fatalf("%v", err)
	}
}

func (*hashesSuite) TestHashesBackend_Store() {}

func (*hashesSuite) TestHashesBackend_Retrieve() {}

func TestHashesBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(hashesSuite))
}
