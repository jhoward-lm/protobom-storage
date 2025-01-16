// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/protobom/protobom/pkg/sbom"
)

type NodeList struct {
	ent.Schema
}

func (NodeList) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ProtoMessageMixin[*sbom.NodeList]{},
		UUIDMixin{},
	}
}

func (NodeList) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("root_elements"),
	}
}

func (NodeList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("document", Document.Type).
			Required().
			Unique().
			Immutable(),
		edge.To("edge_types", EdgeType.Type).
			StorageKey(edge.Table("node_list_edges"), edge.Columns("node_list_id", "edge_type_id")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("nodes", Node.Type).
			StorageKey(edge.Table("node_list_nodes"), edge.Columns("node_list_id", "node_id")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
