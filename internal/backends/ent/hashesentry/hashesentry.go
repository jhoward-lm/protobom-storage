// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package hashesentry

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the hashesentry type in the database.
	Label = "hashes_entry"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDocumentID holds the string denoting the document_id field in the database.
	FieldDocumentID = "document_id"
	// FieldHashAlgorithm holds the string denoting the hash_algorithm field in the database.
	FieldHashAlgorithm = "hash_algorithm"
	// FieldHashData holds the string denoting the hash_data field in the database.
	FieldHashData = "hash_data"
	// EdgeDocument holds the string denoting the document edge name in mutations.
	EdgeDocument = "document"
	// EdgeExternalReferences holds the string denoting the external_references edge name in mutations.
	EdgeExternalReferences = "external_references"
	// EdgeNodes holds the string denoting the nodes edge name in mutations.
	EdgeNodes = "nodes"
	// Table holds the table name of the hashesentry in the database.
	Table = "hashes_entries"
	// DocumentTable is the table that holds the document relation/edge.
	DocumentTable = "hashes_entries"
	// DocumentInverseTable is the table name for the Document entity.
	// It exists in this package in order to avoid circular dependency with the "document" package.
	DocumentInverseTable = "documents"
	// DocumentColumn is the table column denoting the document relation/edge.
	DocumentColumn = "document_id"
	// ExternalReferencesTable is the table that holds the external_references relation/edge. The primary key declared below.
	ExternalReferencesTable = "ext_ref_hashes"
	// ExternalReferencesInverseTable is the table name for the ExternalReference entity.
	// It exists in this package in order to avoid circular dependency with the "externalreference" package.
	ExternalReferencesInverseTable = "external_references"
	// NodesTable is the table that holds the nodes relation/edge. The primary key declared below.
	NodesTable = "node_hashes"
	// NodesInverseTable is the table name for the Node entity.
	// It exists in this package in order to avoid circular dependency with the "node" package.
	NodesInverseTable = "nodes"
)

// Columns holds all SQL columns for hashesentry fields.
var Columns = []string{
	FieldID,
	FieldDocumentID,
	FieldHashAlgorithm,
	FieldHashData,
}

var (
	// ExternalReferencesPrimaryKey and ExternalReferencesColumn2 are the table columns denoting the
	// primary key for the external_references relation (M2M).
	ExternalReferencesPrimaryKey = []string{"ext_ref_id", "hash_entry_id"}
	// NodesPrimaryKey and NodesColumn2 are the table columns denoting the
	// primary key for the nodes relation (M2M).
	NodesPrimaryKey = []string{"node_id", "hash_entry_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultDocumentID holds the default value on creation for the "document_id" field.
	DefaultDocumentID func() uuid.UUID
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// HashAlgorithm defines the type for the "hash_algorithm" enum field.
type HashAlgorithm string

// HashAlgorithm values.
const (
	HashAlgorithmUNKNOWN     HashAlgorithm = "UNKNOWN"
	HashAlgorithmMD5         HashAlgorithm = "MD5"
	HashAlgorithmSHA1        HashAlgorithm = "SHA1"
	HashAlgorithmSHA256      HashAlgorithm = "SHA256"
	HashAlgorithmSHA384      HashAlgorithm = "SHA384"
	HashAlgorithmSHA512      HashAlgorithm = "SHA512"
	HashAlgorithmSHA3_256    HashAlgorithm = "SHA3_256"
	HashAlgorithmSHA3_384    HashAlgorithm = "SHA3_384"
	HashAlgorithmSHA3_512    HashAlgorithm = "SHA3_512"
	HashAlgorithmBLAKE2B_256 HashAlgorithm = "BLAKE2B_256"
	HashAlgorithmBLAKE2B_384 HashAlgorithm = "BLAKE2B_384"
	HashAlgorithmBLAKE2B_512 HashAlgorithm = "BLAKE2B_512"
	HashAlgorithmBLAKE3      HashAlgorithm = "BLAKE3"
	HashAlgorithmMD2         HashAlgorithm = "MD2"
	HashAlgorithmADLER32     HashAlgorithm = "ADLER32"
	HashAlgorithmMD4         HashAlgorithm = "MD4"
	HashAlgorithmMD6         HashAlgorithm = "MD6"
	HashAlgorithmSHA224      HashAlgorithm = "SHA224"
)

func (ha HashAlgorithm) String() string {
	return string(ha)
}

// HashAlgorithmValidator is a validator for the "hash_algorithm" field enum values. It is called by the builders before save.
func HashAlgorithmValidator(ha HashAlgorithm) error {
	switch ha {
	case HashAlgorithmUNKNOWN, HashAlgorithmMD5, HashAlgorithmSHA1, HashAlgorithmSHA256, HashAlgorithmSHA384, HashAlgorithmSHA512, HashAlgorithmSHA3_256, HashAlgorithmSHA3_384, HashAlgorithmSHA3_512, HashAlgorithmBLAKE2B_256, HashAlgorithmBLAKE2B_384, HashAlgorithmBLAKE2B_512, HashAlgorithmBLAKE3, HashAlgorithmMD2, HashAlgorithmADLER32, HashAlgorithmMD4, HashAlgorithmMD6, HashAlgorithmSHA224:
		return nil
	default:
		return fmt.Errorf("hashesentry: invalid enum value for hash_algorithm field: %q", ha)
	}
}

// OrderOption defines the ordering options for the HashesEntry queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDocumentID orders the results by the document_id field.
func ByDocumentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDocumentID, opts...).ToFunc()
}

// ByHashAlgorithm orders the results by the hash_algorithm field.
func ByHashAlgorithm(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHashAlgorithm, opts...).ToFunc()
}

// ByHashData orders the results by the hash_data field.
func ByHashData(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHashData, opts...).ToFunc()
}

// ByDocumentField orders the results by document field.
func ByDocumentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDocumentStep(), sql.OrderByField(field, opts...))
	}
}

// ByExternalReferencesCount orders the results by external_references count.
func ByExternalReferencesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExternalReferencesStep(), opts...)
	}
}

// ByExternalReferences orders the results by external_references terms.
func ByExternalReferences(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExternalReferencesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNodesCount orders the results by nodes count.
func ByNodesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNodesStep(), opts...)
	}
}

// ByNodes orders the results by nodes terms.
func ByNodes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNodesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDocumentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DocumentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DocumentTable, DocumentColumn),
	)
}
func newExternalReferencesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExternalReferencesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ExternalReferencesTable, ExternalReferencesPrimaryKey...),
	)
}
func newNodesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NodesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, NodesTable, NodesPrimaryKey...),
	)
}
