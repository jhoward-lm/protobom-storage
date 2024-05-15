// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Document struct {
	ent.Schema
}

func (Document) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique().Immutable(),
	}
}

func (Document) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("metadata", Metadata.Type).
			Required().
			Unique().
			Immutable().
			StorageKey(edge.Column("id")),
		edge.From("node_list", NodeList.Type).
			Ref("document").
			Required().
			Unique(),
	}
}

func (Document) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("node_list").Fields("id").Unique(),
	}
}

func (Document) Annotations() []schema.Annotation { return nil }
