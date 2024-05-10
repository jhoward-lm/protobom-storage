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

type toolSuite struct {
	backend *backend.Backend[sbom.Tool]
	backendSuite[sbom.Tool]
}

func (ts *toolSuite) BeforeTest(_suiteName, _testName string) {
	ts.backend = backend.NewBackend[sbom.Tool]().WithDatabaseFile(ts.dbFile)

	if err := ts.backend.ClientSetup(); err != nil {
		ts.T().Fatalf("%v", err)
	}
}

func (*toolSuite) TestToolBackend_Store() {}

func (*toolSuite) TestToolBackend_Retrieve() {}

func TestToolBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(toolSuite))
}
