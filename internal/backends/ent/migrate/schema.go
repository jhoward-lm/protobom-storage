// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AnnotationsColumns holds the columns for the "annotations" table.
	AnnotationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "is_unique", Type: field.TypeBool, Default: false},
		{Name: "document_id", Type: field.TypeUUID, Nullable: true},
		{Name: "node_id", Type: field.TypeUUID, Nullable: true},
	}
	// AnnotationsTable holds the schema information for the "annotations" table.
	AnnotationsTable = &schema.Table{
		Name:       "annotations",
		Columns:    AnnotationsColumns,
		PrimaryKey: []*schema.Column{AnnotationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "annotations_documents_annotations",
				Columns:    []*schema.Column{AnnotationsColumns[4]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "annotations_nodes_annotations",
				Columns:    []*schema.Column{AnnotationsColumns[5]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_annotations_node_id",
				Unique:  false,
				Columns: []*schema.Column{AnnotationsColumns[5]},
			},
			{
				Name:    "idx_annotations_document_id",
				Unique:  false,
				Columns: []*schema.Column{AnnotationsColumns[4]},
			},
			{
				Name:    "idx_node_annotations",
				Unique:  true,
				Columns: []*schema.Column{AnnotationsColumns[5], AnnotationsColumns[1], AnnotationsColumns[2]},
				Annotation: &entsql.IndexAnnotation{
					Where: "node_id IS NOT NULL AND TRIM(node_id) != ''",
				},
			},
			{
				Name:    "idx_node_unique_annotations",
				Unique:  true,
				Columns: []*schema.Column{AnnotationsColumns[5], AnnotationsColumns[1]},
				Annotation: &entsql.IndexAnnotation{
					Where: "node_id IS NOT NULL AND TRIM(node_id) != '' AND is_unique",
				},
			},
			{
				Name:    "idx_document_annotations",
				Unique:  true,
				Columns: []*schema.Column{AnnotationsColumns[4], AnnotationsColumns[1], AnnotationsColumns[2]},
				Annotation: &entsql.IndexAnnotation{
					Where: "document_id IS NOT NULL AND TRIM(document_id) != ''",
				},
			},
			{
				Name:    "idx_document_unique_annotations",
				Unique:  true,
				Columns: []*schema.Column{AnnotationsColumns[4], AnnotationsColumns[1]},
				Annotation: &entsql.IndexAnnotation{
					Where: "document_id IS NOT NULL AND TRIM(document_id) != '' AND is_unique",
				},
			},
		},
	}
	// DocumentsColumns holds the columns for the "documents" table.
	DocumentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "metadata_id", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "node_list_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// DocumentsTable holds the schema information for the "documents" table.
	DocumentsTable = &schema.Table{
		Name:       "documents",
		Columns:    DocumentsColumns,
		PrimaryKey: []*schema.Column{DocumentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "documents_metadata_document",
				Columns:    []*schema.Column{DocumentsColumns[1]},
				RefColumns: []*schema.Column{MetadataColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "documents_node_lists_document",
				Columns:    []*schema.Column{DocumentsColumns[2]},
				RefColumns: []*schema.Column{NodeListsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DocumentTypesColumns holds the columns for the "document_types" table.
	DocumentTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "proto_message", Type: field.TypeBytes, Unique: true},
		{Name: "type", Type: field.TypeEnum, Nullable: true, Enums: []string{"OTHER", "DESIGN", "SOURCE", "BUILD", "ANALYZED", "DEPLOYED", "RUNTIME", "DISCOVERY", "DECOMISSION"}},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "document_id", Type: field.TypeUUID, Nullable: true},
		{Name: "metadata_id", Type: field.TypeUUID, Nullable: true},
	}
	// DocumentTypesTable holds the schema information for the "document_types" table.
	DocumentTypesTable = &schema.Table{
		Name:       "document_types",
		Columns:    DocumentTypesColumns,
		PrimaryKey: []*schema.Column{DocumentTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "document_types_documents_document",
				Columns:    []*schema.Column{DocumentTypesColumns[5]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "document_types_metadata_document_types",
				Columns:    []*schema.Column{DocumentTypesColumns[6]},
				RefColumns: []*schema.Column{MetadataColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_document_types",
				Unique:  true,
				Columns: []*schema.Column{DocumentTypesColumns[6], DocumentTypesColumns[2], DocumentTypesColumns[3], DocumentTypesColumns[4]},
			},
		},
	}
	// EdgeTypesColumns holds the columns for the "edge_types" table.
	EdgeTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "proto_message", Type: field.TypeBytes, Unique: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"UNKNOWN", "amends", "ancestor", "buildDependency", "buildTool", "contains", "contained_by", "copy", "dataFile", "dependencyManifest", "dependsOn", "dependencyOf", "descendant", "describes", "describedBy", "devDependency", "devTool", "distributionArtifact", "documentation", "dynamicLink", "example", "expandedFromArchive", "fileAdded", "fileDeleted", "fileModified", "generates", "generatedFrom", "metafile", "optionalComponent", "optionalDependency", "other", "packages", "patch", "prerequisite", "prerequisiteFor", "providedDependency", "requirementFor", "runtimeDependency", "specificationFor", "staticLink", "test", "testCase", "testDependency", "testTool", "variant"}},
		{Name: "document_id", Type: field.TypeUUID, Nullable: true},
		{Name: "node_id", Type: field.TypeUUID},
		{Name: "to_node_id", Type: field.TypeUUID},
	}
	// EdgeTypesTable holds the schema information for the "edge_types" table.
	EdgeTypesTable = &schema.Table{
		Name:       "edge_types",
		Columns:    EdgeTypesColumns,
		PrimaryKey: []*schema.Column{EdgeTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "edge_types_documents_document",
				Columns:    []*schema.Column{EdgeTypesColumns[3]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "edge_types_nodes_from",
				Columns:    []*schema.Column{EdgeTypesColumns[4]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "edge_types_nodes_to",
				Columns:    []*schema.Column{EdgeTypesColumns[5]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_edge_types",
				Unique:  true,
				Columns: []*schema.Column{EdgeTypesColumns[2], EdgeTypesColumns[4], EdgeTypesColumns[5]},
			},
			{
				Name:    "edgetype_node_id_to_node_id",
				Unique:  true,
				Columns: []*schema.Column{EdgeTypesColumns[4], EdgeTypesColumns[5]},
			},
		},
	}
	// ExternalReferencesColumns holds the columns for the "external_references" table.
	ExternalReferencesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "proto_message", Type: field.TypeBytes, Unique: true},
		{Name: "url", Type: field.TypeString},
		{Name: "comment", Type: field.TypeString},
		{Name: "authority", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"UNKNOWN", "ATTESTATION", "BINARY", "BOM", "BOWER", "BUILD_META", "BUILD_SYSTEM", "CERTIFICATION_REPORT", "CHAT", "CODIFIED_INFRASTRUCTURE", "COMPONENT_ANALYSIS_REPORT", "CONFIGURATION", "DISTRIBUTION_INTAKE", "DOCUMENTATION", "DOWNLOAD", "DYNAMIC_ANALYSIS_REPORT", "EOL_NOTICE", "EVIDENCE", "EXPORT_CONTROL_ASSESSMENT", "FORMULATION", "FUNDING", "ISSUE_TRACKER", "LICENSE", "LOG", "MAILING_LIST", "MATURITY_REPORT", "MAVEN_CENTRAL", "METRICS", "MODEL_CARD", "NPM", "NUGET", "OTHER", "POAM", "PRIVACY_ASSESSMENT", "PRODUCT_METADATA", "PURCHASE_ORDER", "QUALITY_ASSESSMENT_REPORT", "QUALITY_METRICS", "RELEASE_HISTORY", "RELEASE_NOTES", "RISK_ASSESSMENT", "RUNTIME_ANALYSIS_REPORT", "SECURE_SOFTWARE_ATTESTATION", "SECURITY_ADVERSARY_MODEL", "SECURITY_ADVISORY", "SECURITY_CONTACT", "SECURITY_FIX", "SECURITY_OTHER", "SECURITY_PENTEST_REPORT", "SECURITY_POLICY", "SECURITY_SWID", "SECURITY_THREAT_MODEL", "SOCIAL", "SOURCE_ARTIFACT", "STATIC_ANALYSIS_REPORT", "SUPPORT", "VCS", "VULNERABILITY_ASSERTION", "VULNERABILITY_DISCLOSURE_REPORT", "VULNERABILITY_EXPLOITABILITY_ASSESSMENT", "WEBSITE"}},
		{Name: "document_id", Type: field.TypeUUID, Nullable: true},
	}
	// ExternalReferencesTable holds the schema information for the "external_references" table.
	ExternalReferencesTable = &schema.Table{
		Name:       "external_references",
		Columns:    ExternalReferencesColumns,
		PrimaryKey: []*schema.Column{ExternalReferencesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "external_references_documents_document",
				Columns:    []*schema.Column{ExternalReferencesColumns[6]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// HashesEntriesColumns holds the columns for the "hashes_entries" table.
	HashesEntriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "hash_algorithm", Type: field.TypeEnum, Enums: []string{"UNKNOWN", "MD5", "SHA1", "SHA256", "SHA384", "SHA512", "SHA3_256", "SHA3_384", "SHA3_512", "BLAKE2B_256", "BLAKE2B_384", "BLAKE2B_512", "BLAKE3", "MD2", "ADLER32", "MD4", "MD6", "SHA224"}},
		{Name: "hash_data", Type: field.TypeString},
		{Name: "document_id", Type: field.TypeUUID, Nullable: true},
	}
	// HashesEntriesTable holds the schema information for the "hashes_entries" table.
	HashesEntriesTable = &schema.Table{
		Name:       "hashes_entries",
		Columns:    HashesEntriesColumns,
		PrimaryKey: []*schema.Column{HashesEntriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "hashes_entries_documents_document",
				Columns:    []*schema.Column{HashesEntriesColumns[3]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_hashes",
				Unique:  true,
				Columns: []*schema.Column{HashesEntriesColumns[1], HashesEntriesColumns[2]},
			},
		},
	}
	// IdentifiersEntriesColumns holds the columns for the "identifiers_entries" table.
	IdentifiersEntriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"UNKNOWN_IDENTIFIER_TYPE", "PURL", "CPE22", "CPE23", "GITOID"}},
		{Name: "value", Type: field.TypeString},
		{Name: "document_id", Type: field.TypeUUID, Nullable: true},
	}
	// IdentifiersEntriesTable holds the schema information for the "identifiers_entries" table.
	IdentifiersEntriesTable = &schema.Table{
		Name:       "identifiers_entries",
		Columns:    IdentifiersEntriesColumns,
		PrimaryKey: []*schema.Column{IdentifiersEntriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "identifiers_entries_documents_document",
				Columns:    []*schema.Column{IdentifiersEntriesColumns[3]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_identifiers",
				Unique:  true,
				Columns: []*schema.Column{IdentifiersEntriesColumns[1], IdentifiersEntriesColumns[2]},
			},
		},
	}
	// MetadataColumns holds the columns for the "metadata" table.
	MetadataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "proto_message", Type: field.TypeBytes, Unique: true},
		{Name: "native_id", Type: field.TypeString},
		{Name: "version", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "date", Type: field.TypeTime},
		{Name: "comment", Type: field.TypeString},
	}
	// MetadataTable holds the schema information for the "metadata" table.
	MetadataTable = &schema.Table{
		Name:       "metadata",
		Columns:    MetadataColumns,
		PrimaryKey: []*schema.Column{MetadataColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "idx_metadata",
				Unique:  true,
				Columns: []*schema.Column{MetadataColumns[2], MetadataColumns[3], MetadataColumns[4]},
			},
		},
	}
	// NodesColumns holds the columns for the "nodes" table.
	NodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "proto_message", Type: field.TypeBytes, Unique: true},
		{Name: "native_id", Type: field.TypeString},
		{Name: "node_list_id", Type: field.TypeUUID, Nullable: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"PACKAGE", "FILE"}},
		{Name: "name", Type: field.TypeString},
		{Name: "version", Type: field.TypeString},
		{Name: "file_name", Type: field.TypeString},
		{Name: "url_home", Type: field.TypeString},
		{Name: "url_download", Type: field.TypeString},
		{Name: "licenses", Type: field.TypeJSON},
		{Name: "license_concluded", Type: field.TypeString},
		{Name: "license_comments", Type: field.TypeString},
		{Name: "copyright", Type: field.TypeString},
		{Name: "source_info", Type: field.TypeString},
		{Name: "comment", Type: field.TypeString},
		{Name: "summary", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "release_date", Type: field.TypeTime},
		{Name: "build_date", Type: field.TypeTime},
		{Name: "valid_until_date", Type: field.TypeTime},
		{Name: "attribution", Type: field.TypeJSON},
		{Name: "file_types", Type: field.TypeJSON},
		{Name: "document_id", Type: field.TypeUUID, Nullable: true},
	}
	// NodesTable holds the schema information for the "nodes" table.
	NodesTable = &schema.Table{
		Name:       "nodes",
		Columns:    NodesColumns,
		PrimaryKey: []*schema.Column{NodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "nodes_documents_document",
				Columns:    []*schema.Column{NodesColumns[23]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_nodes",
				Unique:  true,
				Columns: []*schema.Column{NodesColumns[2], NodesColumns[3]},
			},
		},
	}
	// NodeListsColumns holds the columns for the "node_lists" table.
	NodeListsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "proto_message", Type: field.TypeBytes, Unique: true},
		{Name: "root_elements", Type: field.TypeJSON},
	}
	// NodeListsTable holds the schema information for the "node_lists" table.
	NodeListsTable = &schema.Table{
		Name:       "node_lists",
		Columns:    NodeListsColumns,
		PrimaryKey: []*schema.Column{NodeListsColumns[0]},
	}
	// PersonsColumns holds the columns for the "persons" table.
	PersonsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "proto_message", Type: field.TypeBytes, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "is_org", Type: field.TypeBool},
		{Name: "email", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "metadata_id", Type: field.TypeUUID, Nullable: true},
		{Name: "node_suppliers", Type: field.TypeUUID, Nullable: true},
		{Name: "node_id", Type: field.TypeUUID, Nullable: true},
		{Name: "document_id", Type: field.TypeUUID, Nullable: true},
		{Name: "person_contacts", Type: field.TypeUUID, Nullable: true},
	}
	// PersonsTable holds the schema information for the "persons" table.
	PersonsTable = &schema.Table{
		Name:       "persons",
		Columns:    PersonsColumns,
		PrimaryKey: []*schema.Column{PersonsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "persons_metadata_authors",
				Columns:    []*schema.Column{PersonsColumns[7]},
				RefColumns: []*schema.Column{MetadataColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "persons_nodes_suppliers",
				Columns:    []*schema.Column{PersonsColumns[8]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "persons_nodes_originators",
				Columns:    []*schema.Column{PersonsColumns[9]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "persons_documents_document",
				Columns:    []*schema.Column{PersonsColumns[10]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "persons_persons_contacts",
				Columns:    []*schema.Column{PersonsColumns[11]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_person_metadata_id",
				Unique:  true,
				Columns: []*schema.Column{PersonsColumns[7], PersonsColumns[2], PersonsColumns[3], PersonsColumns[4], PersonsColumns[5], PersonsColumns[6]},
				Annotation: &entsql.IndexAnnotation{
					Where: "metadata_id IS NOT NULL AND node_id IS NULL",
				},
			},
			{
				Name:    "idx_person_node_id",
				Unique:  true,
				Columns: []*schema.Column{PersonsColumns[9], PersonsColumns[2], PersonsColumns[3], PersonsColumns[4], PersonsColumns[5], PersonsColumns[6]},
				Annotation: &entsql.IndexAnnotation{
					Where: "metadata_id IS NULL AND node_id IS NOT NULL",
				},
			},
		},
	}
	// PropertiesColumns holds the columns for the "properties" table.
	PropertiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "proto_message", Type: field.TypeBytes, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "data", Type: field.TypeString},
		{Name: "node_id", Type: field.TypeUUID, Nullable: true},
		{Name: "document_id", Type: field.TypeUUID, Nullable: true},
	}
	// PropertiesTable holds the schema information for the "properties" table.
	PropertiesTable = &schema.Table{
		Name:       "properties",
		Columns:    PropertiesColumns,
		PrimaryKey: []*schema.Column{PropertiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "properties_nodes_properties",
				Columns:    []*schema.Column{PropertiesColumns[4]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "properties_documents_document",
				Columns:    []*schema.Column{PropertiesColumns[5]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_property",
				Unique:  true,
				Columns: []*schema.Column{PropertiesColumns[2], PropertiesColumns[3]},
			},
		},
	}
	// PurposesColumns holds the columns for the "purposes" table.
	PurposesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "primary_purpose", Type: field.TypeEnum, Enums: []string{"UNKNOWN_PURPOSE", "APPLICATION", "ARCHIVE", "BOM", "CONFIGURATION", "CONTAINER", "DATA", "DEVICE", "DEVICE_DRIVER", "DOCUMENTATION", "EVIDENCE", "EXECUTABLE", "FILE", "FIRMWARE", "FRAMEWORK", "INSTALL", "LIBRARY", "MACHINE_LEARNING_MODEL", "MANIFEST", "MODEL", "MODULE", "OPERATING_SYSTEM", "OTHER", "PATCH", "PLATFORM", "REQUIREMENT", "SOURCE", "SPECIFICATION", "TEST"}},
		{Name: "node_id", Type: field.TypeUUID, Nullable: true},
		{Name: "document_id", Type: field.TypeUUID, Nullable: true},
	}
	// PurposesTable holds the schema information for the "purposes" table.
	PurposesTable = &schema.Table{
		Name:       "purposes",
		Columns:    PurposesColumns,
		PrimaryKey: []*schema.Column{PurposesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "purposes_nodes_primary_purpose",
				Columns:    []*schema.Column{PurposesColumns[2]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "purposes_documents_document",
				Columns:    []*schema.Column{PurposesColumns[3]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_purposes",
				Unique:  true,
				Columns: []*schema.Column{PurposesColumns[2], PurposesColumns[1]},
			},
		},
	}
	// SourceDataColumns holds the columns for the "source_data" table.
	SourceDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "proto_message", Type: field.TypeBytes, Unique: true},
		{Name: "format", Type: field.TypeString},
		{Name: "size", Type: field.TypeInt64},
		{Name: "uri", Type: field.TypeString, Nullable: true},
		{Name: "hashes", Type: field.TypeJSON, Nullable: true},
		{Name: "metadata_id", Type: field.TypeUUID},
		{Name: "document_id", Type: field.TypeUUID, Nullable: true},
	}
	// SourceDataTable holds the schema information for the "source_data" table.
	SourceDataTable = &schema.Table{
		Name:       "source_data",
		Columns:    SourceDataColumns,
		PrimaryKey: []*schema.Column{SourceDataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "source_data_metadata_source_data",
				Columns:    []*schema.Column{SourceDataColumns[6]},
				RefColumns: []*schema.Column{MetadataColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "source_data_documents_document",
				Columns:    []*schema.Column{SourceDataColumns[7]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_source_data",
				Unique:  true,
				Columns: []*schema.Column{SourceDataColumns[2], SourceDataColumns[3], SourceDataColumns[4]},
			},
		},
	}
	// ToolsColumns holds the columns for the "tools" table.
	ToolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "proto_message", Type: field.TypeBytes, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "version", Type: field.TypeString},
		{Name: "vendor", Type: field.TypeString},
		{Name: "metadata_id", Type: field.TypeUUID, Nullable: true},
		{Name: "document_id", Type: field.TypeUUID, Nullable: true},
	}
	// ToolsTable holds the schema information for the "tools" table.
	ToolsTable = &schema.Table{
		Name:       "tools",
		Columns:    ToolsColumns,
		PrimaryKey: []*schema.Column{ToolsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tools_metadata_tools",
				Columns:    []*schema.Column{ToolsColumns[5]},
				RefColumns: []*schema.Column{MetadataColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tools_documents_document",
				Columns:    []*schema.Column{ToolsColumns[6]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "idx_tools",
				Unique:  true,
				Columns: []*schema.Column{ToolsColumns[5], ToolsColumns[2], ToolsColumns[3], ToolsColumns[4]},
			},
		},
	}
	// ExtRefHashesColumns holds the columns for the "ext_ref_hashes" table.
	ExtRefHashesColumns = []*schema.Column{
		{Name: "ext_ref_id", Type: field.TypeUUID},
		{Name: "hash_entry_id", Type: field.TypeUUID},
	}
	// ExtRefHashesTable holds the schema information for the "ext_ref_hashes" table.
	ExtRefHashesTable = &schema.Table{
		Name:       "ext_ref_hashes",
		Columns:    ExtRefHashesColumns,
		PrimaryKey: []*schema.Column{ExtRefHashesColumns[0], ExtRefHashesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "ext_ref_hashes_ext_ref_id",
				Columns:    []*schema.Column{ExtRefHashesColumns[0]},
				RefColumns: []*schema.Column{ExternalReferencesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "ext_ref_hashes_hash_entry_id",
				Columns:    []*schema.Column{ExtRefHashesColumns[1]},
				RefColumns: []*schema.Column{HashesEntriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// NodeExternalReferencesColumns holds the columns for the "node_external_references" table.
	NodeExternalReferencesColumns = []*schema.Column{
		{Name: "node_id", Type: field.TypeUUID},
		{Name: "external_reference_id", Type: field.TypeUUID},
	}
	// NodeExternalReferencesTable holds the schema information for the "node_external_references" table.
	NodeExternalReferencesTable = &schema.Table{
		Name:       "node_external_references",
		Columns:    NodeExternalReferencesColumns,
		PrimaryKey: []*schema.Column{NodeExternalReferencesColumns[0], NodeExternalReferencesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "node_external_references_node_id",
				Columns:    []*schema.Column{NodeExternalReferencesColumns[0]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "node_external_references_external_reference_id",
				Columns:    []*schema.Column{NodeExternalReferencesColumns[1]},
				RefColumns: []*schema.Column{ExternalReferencesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// NodeHashesColumns holds the columns for the "node_hashes" table.
	NodeHashesColumns = []*schema.Column{
		{Name: "node_id", Type: field.TypeUUID},
		{Name: "hash_entry_id", Type: field.TypeUUID},
	}
	// NodeHashesTable holds the schema information for the "node_hashes" table.
	NodeHashesTable = &schema.Table{
		Name:       "node_hashes",
		Columns:    NodeHashesColumns,
		PrimaryKey: []*schema.Column{NodeHashesColumns[0], NodeHashesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "node_hashes_node_id",
				Columns:    []*schema.Column{NodeHashesColumns[0]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "node_hashes_hash_entry_id",
				Columns:    []*schema.Column{NodeHashesColumns[1]},
				RefColumns: []*schema.Column{HashesEntriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// NodeIdentifiersColumns holds the columns for the "node_identifiers" table.
	NodeIdentifiersColumns = []*schema.Column{
		{Name: "node_id", Type: field.TypeUUID},
		{Name: "identifier_entry_id", Type: field.TypeUUID},
	}
	// NodeIdentifiersTable holds the schema information for the "node_identifiers" table.
	NodeIdentifiersTable = &schema.Table{
		Name:       "node_identifiers",
		Columns:    NodeIdentifiersColumns,
		PrimaryKey: []*schema.Column{NodeIdentifiersColumns[0], NodeIdentifiersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "node_identifiers_node_id",
				Columns:    []*schema.Column{NodeIdentifiersColumns[0]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "node_identifiers_identifier_entry_id",
				Columns:    []*schema.Column{NodeIdentifiersColumns[1]},
				RefColumns: []*schema.Column{IdentifiersEntriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// NodeListEdgesColumns holds the columns for the "node_list_edges" table.
	NodeListEdgesColumns = []*schema.Column{
		{Name: "node_list_id", Type: field.TypeUUID},
		{Name: "edge_type_id", Type: field.TypeUUID},
	}
	// NodeListEdgesTable holds the schema information for the "node_list_edges" table.
	NodeListEdgesTable = &schema.Table{
		Name:       "node_list_edges",
		Columns:    NodeListEdgesColumns,
		PrimaryKey: []*schema.Column{NodeListEdgesColumns[0], NodeListEdgesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "node_list_edges_node_list_id",
				Columns:    []*schema.Column{NodeListEdgesColumns[0]},
				RefColumns: []*schema.Column{NodeListsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "node_list_edges_edge_type_id",
				Columns:    []*schema.Column{NodeListEdgesColumns[1]},
				RefColumns: []*schema.Column{EdgeTypesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// NodeListNodesColumns holds the columns for the "node_list_nodes" table.
	NodeListNodesColumns = []*schema.Column{
		{Name: "node_list_id", Type: field.TypeUUID},
		{Name: "node_id", Type: field.TypeUUID},
	}
	// NodeListNodesTable holds the schema information for the "node_list_nodes" table.
	NodeListNodesTable = &schema.Table{
		Name:       "node_list_nodes",
		Columns:    NodeListNodesColumns,
		PrimaryKey: []*schema.Column{NodeListNodesColumns[0], NodeListNodesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "node_list_nodes_node_list_id",
				Columns:    []*schema.Column{NodeListNodesColumns[0]},
				RefColumns: []*schema.Column{NodeListsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "node_list_nodes_node_id",
				Columns:    []*schema.Column{NodeListNodesColumns[1]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AnnotationsTable,
		DocumentsTable,
		DocumentTypesTable,
		EdgeTypesTable,
		ExternalReferencesTable,
		HashesEntriesTable,
		IdentifiersEntriesTable,
		MetadataTable,
		NodesTable,
		NodeListsTable,
		PersonsTable,
		PropertiesTable,
		PurposesTable,
		SourceDataTable,
		ToolsTable,
		ExtRefHashesTable,
		NodeExternalReferencesTable,
		NodeHashesTable,
		NodeIdentifiersTable,
		NodeListEdgesTable,
		NodeListNodesTable,
	}
)

func init() {
	AnnotationsTable.ForeignKeys[0].RefTable = DocumentsTable
	AnnotationsTable.ForeignKeys[1].RefTable = NodesTable
	DocumentsTable.ForeignKeys[0].RefTable = MetadataTable
	DocumentsTable.ForeignKeys[1].RefTable = NodeListsTable
	DocumentTypesTable.ForeignKeys[0].RefTable = DocumentsTable
	DocumentTypesTable.ForeignKeys[1].RefTable = MetadataTable
	DocumentTypesTable.Annotation = &entsql.Annotation{}
	EdgeTypesTable.ForeignKeys[0].RefTable = DocumentsTable
	EdgeTypesTable.ForeignKeys[1].RefTable = NodesTable
	EdgeTypesTable.ForeignKeys[2].RefTable = NodesTable
	EdgeTypesTable.Annotation = &entsql.Annotation{}
	ExternalReferencesTable.ForeignKeys[0].RefTable = DocumentsTable
	ExternalReferencesTable.Annotation = &entsql.Annotation{}
	HashesEntriesTable.ForeignKeys[0].RefTable = DocumentsTable
	HashesEntriesTable.Annotation = &entsql.Annotation{}
	IdentifiersEntriesTable.ForeignKeys[0].RefTable = DocumentsTable
	IdentifiersEntriesTable.Annotation = &entsql.Annotation{}
	NodesTable.ForeignKeys[0].RefTable = DocumentsTable
	NodesTable.Annotation = &entsql.Annotation{}
	PersonsTable.ForeignKeys[0].RefTable = MetadataTable
	PersonsTable.ForeignKeys[1].RefTable = NodesTable
	PersonsTable.ForeignKeys[2].RefTable = NodesTable
	PersonsTable.ForeignKeys[3].RefTable = DocumentsTable
	PersonsTable.ForeignKeys[4].RefTable = PersonsTable
	PersonsTable.Annotation = &entsql.Annotation{}
	PropertiesTable.ForeignKeys[0].RefTable = NodesTable
	PropertiesTable.ForeignKeys[1].RefTable = DocumentsTable
	PropertiesTable.Annotation = &entsql.Annotation{}
	PurposesTable.ForeignKeys[0].RefTable = NodesTable
	PurposesTable.ForeignKeys[1].RefTable = DocumentsTable
	PurposesTable.Annotation = &entsql.Annotation{}
	SourceDataTable.ForeignKeys[0].RefTable = MetadataTable
	SourceDataTable.ForeignKeys[1].RefTable = DocumentsTable
	SourceDataTable.Annotation = &entsql.Annotation{}
	ToolsTable.ForeignKeys[0].RefTable = MetadataTable
	ToolsTable.ForeignKeys[1].RefTable = DocumentsTable
	ToolsTable.Annotation = &entsql.Annotation{}
	ExtRefHashesTable.ForeignKeys[0].RefTable = ExternalReferencesTable
	ExtRefHashesTable.ForeignKeys[1].RefTable = HashesEntriesTable
	ExtRefHashesTable.Annotation = &entsql.Annotation{}
	NodeExternalReferencesTable.ForeignKeys[0].RefTable = NodesTable
	NodeExternalReferencesTable.ForeignKeys[1].RefTable = ExternalReferencesTable
	NodeExternalReferencesTable.Annotation = &entsql.Annotation{}
	NodeHashesTable.ForeignKeys[0].RefTable = NodesTable
	NodeHashesTable.ForeignKeys[1].RefTable = HashesEntriesTable
	NodeHashesTable.Annotation = &entsql.Annotation{}
	NodeIdentifiersTable.ForeignKeys[0].RefTable = NodesTable
	NodeIdentifiersTable.ForeignKeys[1].RefTable = IdentifiersEntriesTable
	NodeIdentifiersTable.Annotation = &entsql.Annotation{}
	NodeListEdgesTable.ForeignKeys[0].RefTable = NodeListsTable
	NodeListEdgesTable.ForeignKeys[1].RefTable = EdgeTypesTable
	NodeListEdgesTable.Annotation = &entsql.Annotation{}
	NodeListNodesTable.ForeignKeys[0].RefTable = NodeListsTable
	NodeListNodesTable.ForeignKeys[1].RefTable = NodesTable
	NodeListNodesTable.Annotation = &entsql.Annotation{}
}
