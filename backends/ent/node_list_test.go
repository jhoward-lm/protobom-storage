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

type nodeListSuite struct {
	backend *backend.Backend[sbom.NodeList]
	backendSuite[sbom.NodeList]
}

func (nls *nodeListSuite) BeforeTest(_suiteName, _testName string) {
	nls.backend = backend.NewBackend[sbom.NodeList]().WithDatabaseFile(nls.dbFile)

	if err := nls.backend.ClientSetup(); err != nil {
		nls.T().Fatalf("%v", err)
	}
}

func (*nodeListSuite) TestNodeListBackend_Store() {}

func (*nodeListSuite) TestNodeListBackend_Retrieve() {}

func TestNodeListBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(nodeListSuite))
}
