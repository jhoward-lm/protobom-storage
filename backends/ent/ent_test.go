// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent_test

import (
	"context"
	"os"
	"path/filepath"

	"github.com/protobom/protobom/pkg/reader"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"
	"github.com/stretchr/testify/suite"

	backend "github.com/protobom/storage/backends/ent"
	"github.com/protobom/storage/internal/backends/ent"
)

type backendSuite[T storage.ProtobomType] struct {
	suite.Suite
	backend   backend.Backend[T]
	dbFile    string
	ctx       context.Context
	client    *ent.Client
	documents []*sbom.Document
}

func (bs *backendSuite[any]) SetupSuite() {
	sbomReader := reader.New()

	sboms, err := os.ReadDir("testdata")
	if err != nil {
		bs.T().Fatalf("%v", err)
	}

	for sbomIdx := range sboms {
		doc, err := sbomReader.ParseFile(filepath.Join("testdata", sboms[sbomIdx].Name()))
		if err != nil {
			bs.T().Fatalf("%v", err)
		}

		bs.documents = append(bs.documents, doc)
	}

	bs.dbFile = filepath.Join(bs.T().TempDir(), "test.db")
}

func (bs *backendSuite[any]) TearDownSuite() {
	bs.client.Close()
}
