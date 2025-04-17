// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasecluster


type DatabaseClusterBackupRestore struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.51.0/docs/resources/database_cluster#database_name DatabaseCluster#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.51.0/docs/resources/database_cluster#backup_created_at DatabaseCluster#backup_created_at}.
	BackupCreatedAt *string `field:"optional" json:"backupCreatedAt" yaml:"backupCreatedAt"`
}

