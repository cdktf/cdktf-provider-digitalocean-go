// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseuser


type DatabaseUserSettings struct {
	// acl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/database_user#acl DatabaseUser#acl}
	Acl interface{} `field:"optional" json:"acl" yaml:"acl"`
	// opensearch_acl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/database_user#opensearch_acl DatabaseUser#opensearch_acl}
	OpensearchAcl interface{} `field:"optional" json:"opensearchAcl" yaml:"opensearchAcl"`
}

