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
	backend *backend.Backend[sbom.Person]
	backendSuite[sbom.Person]
}

func (ps *personSuite) BeforeTest(_suiteName, _testName string) {
	ps.backend = backend.NewBackend[sbom.Person]().WithDatabaseFile(ps.dbFile)

	if err := ps.backend.ClientSetup(); err != nil {
		ps.T().Fatalf("%v", err)
	}
}

func (*personSuite) TestPersonBackend_Store() {}

func (*personSuite) TestPersonBackend_Retrieve() {}

func TestPersonBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(personSuite))
}
