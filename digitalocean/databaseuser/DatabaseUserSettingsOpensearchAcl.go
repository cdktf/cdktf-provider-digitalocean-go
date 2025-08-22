// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseuser


type DatabaseUserSettingsOpensearchAcl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/database_user#index DatabaseUser#index}.
	Index *string `field:"required" json:"index" yaml:"index"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/database_user#permission DatabaseUser#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

