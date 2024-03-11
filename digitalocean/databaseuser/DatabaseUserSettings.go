// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseuser


type DatabaseUserSettings struct {
	// acl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.35.0/docs/resources/database_user#acl DatabaseUser#acl}
	Acl interface{} `field:"optional" json:"acl" yaml:"acl"`
}

