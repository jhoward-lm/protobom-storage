// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"github.com/bom-squad/protobom/pkg/sbom"

	"github.com/protobom/storage/pkg/options"
)

// Retrieve implements the model.v1.storage.Backend interface
func (EntBackend) Retrieve(string, *options.RetrieveOptions) (*sbom.Document, error) {
	return nil, nil
}
