// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent_test

import (
	"cmp"
	"slices"
	"testing"

	"github.com/protobom/protobom/pkg/sbom"
	"github.com/stretchr/testify/suite"

	backend "github.com/protobom/storage/backends/ent"
	"github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/node"
)

type nodeSuite struct {
	backend *backend.NodeBackend
	backendSuite[*sbom.Node]
}

func (ns *nodeSuite) BeforeTest(_suiteName, _testName string) {
	ns.backend = &backend.NodeBackend{
		Backend: backend.NewBackend[*sbom.Node]().WithDatabaseFile(ns.dbFile),
	}

	if err := ns.backend.ClientSetup(); err != nil {
		ns.T().Fatalf("%v", err)
	}
}

func (ns *nodeSuite) TestNodeBackend_Store() {
	nodes := []*sbom.Node{}
	nodeIDs := []string{}

	for _, doc := range ns.documents {
		for _, n := range doc.NodeList.Nodes {
			nodes = append(nodes, n)
			nodeIDs = append(nodeIDs, n.Id)
		}
	}

	for _, n := range nodes {
		ns.Run(n.Id, func() {
			err := ns.backend.Store(n, nil)
			ns.Nil(err)
		})
	}

	queryResult, err := ns.client.Node.Query().Where(node.IDIn(nodeIDs...)).All(ns.ctx)
	ns.Require().Nil(err)
	ns.Require().Len(queryResult, len(nodes))

	slices.SortStableFunc(nodes, func(a, b *sbom.Node) int { return cmp.Compare(a.Id, b.Id) })
	slices.SortStableFunc(queryResult, func(a, b *ent.Node) int { return cmp.Compare(a.ID, b.ID) })

	for i := range queryResult {
		ns.Equal(queryResult[i].ID, nodes[i].Id)
	}
}

func (*nodeSuite) TestNodeBackend_Retrieve() {}

func TestNodeBackend(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(nodeSuite))
}
