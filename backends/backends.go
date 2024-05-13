// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package backends

import (
	"github.com/protobom/protobom/pkg/storage"

	"github.com/protobom/storage/backends/ent"
)

type EntBackend[T storage.ProtobomType] ent.Backend[T]
