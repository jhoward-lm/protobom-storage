// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package person

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldID, id))
}

// DocumentID applies equality check predicate on the "document_id" field. It's identical to DocumentIDEQ.
func DocumentID(v uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldDocumentID, v))
}

// ProtoMessage applies equality check predicate on the "proto_message" field. It's identical to ProtoMessageEQ.
func ProtoMessage(v *sbom.Person) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldProtoMessage, v))
}

// MetadataID applies equality check predicate on the "metadata_id" field. It's identical to MetadataIDEQ.
func MetadataID(v uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldMetadataID, v))
}

// NodeID applies equality check predicate on the "node_id" field. It's identical to NodeIDEQ.
func NodeID(v uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldNodeID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldName, v))
}

// IsOrg applies equality check predicate on the "is_org" field. It's identical to IsOrgEQ.
func IsOrg(v bool) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldIsOrg, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldEmail, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldURL, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldPhone, v))
}

// DocumentIDEQ applies the EQ predicate on the "document_id" field.
func DocumentIDEQ(v uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldDocumentID, v))
}

// DocumentIDNEQ applies the NEQ predicate on the "document_id" field.
func DocumentIDNEQ(v uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldDocumentID, v))
}

// DocumentIDIn applies the In predicate on the "document_id" field.
func DocumentIDIn(vs ...uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldDocumentID, vs...))
}

// DocumentIDNotIn applies the NotIn predicate on the "document_id" field.
func DocumentIDNotIn(vs ...uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldDocumentID, vs...))
}

// DocumentIDIsNil applies the IsNil predicate on the "document_id" field.
func DocumentIDIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldDocumentID))
}

// DocumentIDNotNil applies the NotNil predicate on the "document_id" field.
func DocumentIDNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldDocumentID))
}

// ProtoMessageEQ applies the EQ predicate on the "proto_message" field.
func ProtoMessageEQ(v *sbom.Person) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldProtoMessage, v))
}

// ProtoMessageNEQ applies the NEQ predicate on the "proto_message" field.
func ProtoMessageNEQ(v *sbom.Person) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldProtoMessage, v))
}

// ProtoMessageIn applies the In predicate on the "proto_message" field.
func ProtoMessageIn(vs ...*sbom.Person) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldProtoMessage, vs...))
}

// ProtoMessageNotIn applies the NotIn predicate on the "proto_message" field.
func ProtoMessageNotIn(vs ...*sbom.Person) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldProtoMessage, vs...))
}

// ProtoMessageGT applies the GT predicate on the "proto_message" field.
func ProtoMessageGT(v *sbom.Person) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldProtoMessage, v))
}

// ProtoMessageGTE applies the GTE predicate on the "proto_message" field.
func ProtoMessageGTE(v *sbom.Person) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldProtoMessage, v))
}

// ProtoMessageLT applies the LT predicate on the "proto_message" field.
func ProtoMessageLT(v *sbom.Person) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldProtoMessage, v))
}

// ProtoMessageLTE applies the LTE predicate on the "proto_message" field.
func ProtoMessageLTE(v *sbom.Person) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldProtoMessage, v))
}

// MetadataIDEQ applies the EQ predicate on the "metadata_id" field.
func MetadataIDEQ(v uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldMetadataID, v))
}

// MetadataIDNEQ applies the NEQ predicate on the "metadata_id" field.
func MetadataIDNEQ(v uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldMetadataID, v))
}

// MetadataIDIn applies the In predicate on the "metadata_id" field.
func MetadataIDIn(vs ...uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldMetadataID, vs...))
}

// MetadataIDNotIn applies the NotIn predicate on the "metadata_id" field.
func MetadataIDNotIn(vs ...uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldMetadataID, vs...))
}

// MetadataIDIsNil applies the IsNil predicate on the "metadata_id" field.
func MetadataIDIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldMetadataID))
}

// MetadataIDNotNil applies the NotNil predicate on the "metadata_id" field.
func MetadataIDNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldMetadataID))
}

// NodeIDEQ applies the EQ predicate on the "node_id" field.
func NodeIDEQ(v uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldNodeID, v))
}

// NodeIDNEQ applies the NEQ predicate on the "node_id" field.
func NodeIDNEQ(v uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldNodeID, v))
}

// NodeIDIn applies the In predicate on the "node_id" field.
func NodeIDIn(vs ...uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldNodeID, vs...))
}

// NodeIDNotIn applies the NotIn predicate on the "node_id" field.
func NodeIDNotIn(vs ...uuid.UUID) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldNodeID, vs...))
}

// NodeIDIsNil applies the IsNil predicate on the "node_id" field.
func NodeIDIsNil() predicate.Person {
	return predicate.Person(sql.FieldIsNull(FieldNodeID))
}

// NodeIDNotNil applies the NotNil predicate on the "node_id" field.
func NodeIDNotNil() predicate.Person {
	return predicate.Person(sql.FieldNotNull(FieldNodeID))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldName, v))
}

// IsOrgEQ applies the EQ predicate on the "is_org" field.
func IsOrgEQ(v bool) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldIsOrg, v))
}

// IsOrgNEQ applies the NEQ predicate on the "is_org" field.
func IsOrgNEQ(v bool) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldIsOrg, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldEmail, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldURL, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldPhone, v))
}

// HasDocument applies the HasEdge predicate on the "document" edge.
func HasDocument() predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, DocumentTable, DocumentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDocumentWith applies the HasEdge predicate on the "document" edge with a given conditions (other predicates).
func HasDocumentWith(preds ...predicate.Document) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := newDocumentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasContactOwner applies the HasEdge predicate on the "contact_owner" edge.
func HasContactOwner() predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ContactOwnerTable, ContactOwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContactOwnerWith applies the HasEdge predicate on the "contact_owner" edge with a given conditions (other predicates).
func HasContactOwnerWith(preds ...predicate.Person) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := newContactOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasContacts applies the HasEdge predicate on the "contacts" edge.
func HasContacts() predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ContactsTable, ContactsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContactsWith applies the HasEdge predicate on the "contacts" edge with a given conditions (other predicates).
func HasContactsWith(preds ...predicate.Person) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := newContactsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMetadata applies the HasEdge predicate on the "metadata" edge.
func HasMetadata() predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MetadataTable, MetadataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetadataWith applies the HasEdge predicate on the "metadata" edge with a given conditions (other predicates).
func HasMetadataWith(preds ...predicate.Metadata) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := newMetadataStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNode applies the HasEdge predicate on the "node" edge.
func HasNode() predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NodeTable, NodeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNodeWith applies the HasEdge predicate on the "node" edge with a given conditions (other predicates).
func HasNodeWith(preds ...predicate.Node) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := newNodeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Person) predicate.Person {
	return predicate.Person(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Person) predicate.Person {
	return predicate.Person(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Person) predicate.Person {
	return predicate.Person(sql.NotPredicates(p))
}
