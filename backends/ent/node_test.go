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

type nodeSuite struct {
	backend *backend.Backend[sbom.Node]
	backendSuite[sbom.Node]
}

func (ns *nodeSuite) BeforeTest(_suiteName, _testName string) {
	ns.backend = backend.NewBackend[sbom.Node]().WithDatabaseFile(ns.dbFile)

	if err := ns.backend.ClientSetup(); err != nil {
		ns.T().Fatalf("%v", err)
	}
}

func (*nodeSuite) TestNodeBackend_Store() {}

func (*nodeSuite) TestNodeBackend_Retrieve() {}

func TestNodeBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(nodeSuite))
}
