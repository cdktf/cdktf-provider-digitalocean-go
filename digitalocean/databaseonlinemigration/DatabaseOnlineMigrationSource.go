// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseonlinemigration


type DatabaseOnlineMigrationSource struct {
	// The name of the default database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/database_online_migration#db_name DatabaseOnlineMigration#db_name}
	DbName *string `field:"required" json:"dbName" yaml:"dbName"`
	// The FQDN pointing to the database cluster's current primary node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/database_online_migration#host DatabaseOnlineMigration#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// The password of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/database_online_migration#password DatabaseOnlineMigration#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The port on which the database cluster is listening.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/database_online_migration#port DatabaseOnlineMigration#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The default user of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/database_online_migration#username DatabaseOnlineMigration#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

