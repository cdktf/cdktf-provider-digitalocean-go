// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseonlinemigration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseOnlineMigrationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/resources/database_online_migration#cluster_id DatabaseOnlineMigration#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/resources/database_online_migration#source DatabaseOnlineMigration#source}
	Source *DatabaseOnlineMigrationSource `field:"required" json:"source" yaml:"source"`
	// Disables SSL encryption when connecting to the source database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/resources/database_online_migration#disable_ssl DatabaseOnlineMigration#disable_ssl}
	DisableSsl interface{} `field:"optional" json:"disableSsl" yaml:"disableSsl"`
	// The list of databases to be ignored during the migration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/resources/database_online_migration#ignore_dbs DatabaseOnlineMigration#ignore_dbs}
	IgnoreDbs *[]*string `field:"optional" json:"ignoreDbs" yaml:"ignoreDbs"`
}

