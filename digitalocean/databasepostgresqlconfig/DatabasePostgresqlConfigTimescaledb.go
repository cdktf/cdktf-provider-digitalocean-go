// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasepostgresqlconfig


type DatabasePostgresqlConfigTimescaledb struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/database_postgresql_config#max_background_workers DatabasePostgresqlConfig#max_background_workers}.
	MaxBackgroundWorkers *float64 `field:"optional" json:"maxBackgroundWorkers" yaml:"maxBackgroundWorkers"`
}

