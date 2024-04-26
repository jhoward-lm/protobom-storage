// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package externalreference

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the externalreference type in the database.
	Label = "external_reference"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// FieldAuthority holds the string denoting the authority field in the database.
	FieldAuthority = "authority"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeHashes holds the string denoting the hashes edge name in mutations.
	EdgeHashes = "hashes"
	// EdgeNode holds the string denoting the node edge name in mutations.
	EdgeNode = "node"
	// Table holds the table name of the externalreference in the database.
	Table = "external_references"
	// HashesTable is the table that holds the hashes relation/edge.
	HashesTable = "hashes_entries"
	// HashesInverseTable is the table name for the HashesEntry entity.
	// It exists in this package in order to avoid circular dependency with the "hashesentry" package.
	HashesInverseTable = "hashes_entries"
	// HashesColumn is the table column denoting the hashes relation/edge.
	HashesColumn = "external_reference_hashes"
	// NodeTable is the table that holds the node relation/edge.
	NodeTable = "external_references"
	// NodeInverseTable is the table name for the Node entity.
	// It exists in this package in order to avoid circular dependency with the "node" package.
	NodeInverseTable = "nodes"
	// NodeColumn is the table column denoting the node relation/edge.
	NodeColumn = "node_external_references"
)

// Columns holds all SQL columns for externalreference fields.
var Columns = []string{
	FieldID,
	FieldURL,
	FieldComment,
	FieldAuthority,
	FieldType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "external_references"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"node_external_references",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeUNKNOWN                                 Type = "UNKNOWN"
	TypeATTESTATION                             Type = "ATTESTATION"
	TypeBINARY                                  Type = "BINARY"
	TypeBOM                                     Type = "BOM"
	TypeBOWER                                   Type = "BOWER"
	TypeBUILD_META                              Type = "BUILD_META"
	TypeBUILD_SYSTEM                            Type = "BUILD_SYSTEM"
	TypeCERTIFICATION_REPORT                    Type = "CERTIFICATION_REPORT"
	TypeCHAT                                    Type = "CHAT"
	TypeCODIFIED_INFRASTRUCTURE                 Type = "CODIFIED_INFRASTRUCTURE"
	TypeCOMPONENT_ANALYSIS_REPORT               Type = "COMPONENT_ANALYSIS_REPORT"
	TypeCONFIGURATION                           Type = "CONFIGURATION"
	TypeDISTRIBUTION_INTAKE                     Type = "DISTRIBUTION_INTAKE"
	TypeDOCUMENTATION                           Type = "DOCUMENTATION"
	TypeDOWNLOAD                                Type = "DOWNLOAD"
	TypeDYNAMIC_ANALYSIS_REPORT                 Type = "DYNAMIC_ANALYSIS_REPORT"
	TypeEOL_NOTICE                              Type = "EOL_NOTICE"
	TypeEVIDENCE                                Type = "EVIDENCE"
	TypeEXPORT_CONTROL_ASSESSMENT               Type = "EXPORT_CONTROL_ASSESSMENT"
	TypeFORMULATION                             Type = "FORMULATION"
	TypeFUNDING                                 Type = "FUNDING"
	TypeISSUE_TRACKER                           Type = "ISSUE_TRACKER"
	TypeLICENSE                                 Type = "LICENSE"
	TypeLOG                                     Type = "LOG"
	TypeMAILING_LIST                            Type = "MAILING_LIST"
	TypeMATURITY_REPORT                         Type = "MATURITY_REPORT"
	TypeMAVEN_CENTRAL                           Type = "MAVEN_CENTRAL"
	TypeMETRICS                                 Type = "METRICS"
	TypeMODEL_CARD                              Type = "MODEL_CARD"
	TypeNPM                                     Type = "NPM"
	TypeNUGET                                   Type = "NUGET"
	TypeOTHER                                   Type = "OTHER"
	TypePOAM                                    Type = "POAM"
	TypePRIVACY_ASSESSMENT                      Type = "PRIVACY_ASSESSMENT"
	TypePRODUCT_METADATA                        Type = "PRODUCT_METADATA"
	TypePURCHASE_ORDER                          Type = "PURCHASE_ORDER"
	TypeQUALITY_ASSESSMENT_REPORT               Type = "QUALITY_ASSESSMENT_REPORT"
	TypeQUALITY_METRICS                         Type = "QUALITY_METRICS"
	TypeRELEASE_HISTORY                         Type = "RELEASE_HISTORY"
	TypeRELEASE_NOTES                           Type = "RELEASE_NOTES"
	TypeRISK_ASSESSMENT                         Type = "RISK_ASSESSMENT"
	TypeRUNTIME_ANALYSIS_REPORT                 Type = "RUNTIME_ANALYSIS_REPORT"
	TypeSECURE_SOFTWARE_ATTESTATION             Type = "SECURE_SOFTWARE_ATTESTATION"
	TypeSECURITY_ADVERSARY_MODEL                Type = "SECURITY_ADVERSARY_MODEL"
	TypeSECURITY_ADVISORY                       Type = "SECURITY_ADVISORY"
	TypeSECURITY_CONTACT                        Type = "SECURITY_CONTACT"
	TypeSECURITY_FIX                            Type = "SECURITY_FIX"
	TypeSECURITY_OTHER                          Type = "SECURITY_OTHER"
	TypeSECURITY_PENTEST_REPORT                 Type = "SECURITY_PENTEST_REPORT"
	TypeSECURITY_POLICY                         Type = "SECURITY_POLICY"
	TypeSECURITY_SWID                           Type = "SECURITY_SWID"
	TypeSECURITY_THREAT_MODEL                   Type = "SECURITY_THREAT_MODEL"
	TypeSOCIAL                                  Type = "SOCIAL"
	TypeSOURCE_ARTIFACT                         Type = "SOURCE_ARTIFACT"
	TypeSTATIC_ANALYSIS_REPORT                  Type = "STATIC_ANALYSIS_REPORT"
	TypeSUPPORT                                 Type = "SUPPORT"
	TypeVCS                                     Type = "VCS"
	TypeVULNERABILITY_ASSERTION                 Type = "VULNERABILITY_ASSERTION"
	TypeVULNERABILITY_DISCLOSURE_REPORT         Type = "VULNERABILITY_DISCLOSURE_REPORT"
	TypeVULNERABILITY_EXPLOITABILITY_ASSESSMENT Type = "VULNERABILITY_EXPLOITABILITY_ASSESSMENT"
	TypeWEBSITE                                 Type = "WEBSITE"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeUNKNOWN, TypeATTESTATION, TypeBINARY, TypeBOM, TypeBOWER, TypeBUILD_META, TypeBUILD_SYSTEM, TypeCERTIFICATION_REPORT, TypeCHAT, TypeCODIFIED_INFRASTRUCTURE, TypeCOMPONENT_ANALYSIS_REPORT, TypeCONFIGURATION, TypeDISTRIBUTION_INTAKE, TypeDOCUMENTATION, TypeDOWNLOAD, TypeDYNAMIC_ANALYSIS_REPORT, TypeEOL_NOTICE, TypeEVIDENCE, TypeEXPORT_CONTROL_ASSESSMENT, TypeFORMULATION, TypeFUNDING, TypeISSUE_TRACKER, TypeLICENSE, TypeLOG, TypeMAILING_LIST, TypeMATURITY_REPORT, TypeMAVEN_CENTRAL, TypeMETRICS, TypeMODEL_CARD, TypeNPM, TypeNUGET, TypeOTHER, TypePOAM, TypePRIVACY_ASSESSMENT, TypePRODUCT_METADATA, TypePURCHASE_ORDER, TypeQUALITY_ASSESSMENT_REPORT, TypeQUALITY_METRICS, TypeRELEASE_HISTORY, TypeRELEASE_NOTES, TypeRISK_ASSESSMENT, TypeRUNTIME_ANALYSIS_REPORT, TypeSECURE_SOFTWARE_ATTESTATION, TypeSECURITY_ADVERSARY_MODEL, TypeSECURITY_ADVISORY, TypeSECURITY_CONTACT, TypeSECURITY_FIX, TypeSECURITY_OTHER, TypeSECURITY_PENTEST_REPORT, TypeSECURITY_POLICY, TypeSECURITY_SWID, TypeSECURITY_THREAT_MODEL, TypeSOCIAL, TypeSOURCE_ARTIFACT, TypeSTATIC_ANALYSIS_REPORT, TypeSUPPORT, TypeVCS, TypeVULNERABILITY_ASSERTION, TypeVULNERABILITY_DISCLOSURE_REPORT, TypeVULNERABILITY_EXPLOITABILITY_ASSESSMENT, TypeWEBSITE:
		return nil
	default:
		return fmt.Errorf("externalreference: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the ExternalReference queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByComment orders the results by the comment field.
func ByComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComment, opts...).ToFunc()
}

// ByAuthority orders the results by the authority field.
func ByAuthority(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthority, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByHashesCount orders the results by hashes count.
func ByHashesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHashesStep(), opts...)
	}
}

// ByHashes orders the results by hashes terms.
func ByHashes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHashesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNodeField orders the results by node field.
func ByNodeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNodeStep(), sql.OrderByField(field, opts...))
	}
}
func newHashesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HashesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HashesTable, HashesColumn),
	)
}
func newNodeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NodeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, NodeTable, NodeColumn),
	)
}