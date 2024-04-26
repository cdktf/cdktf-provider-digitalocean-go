// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasepostgresqlconfig


type DatabasePostgresqlConfigPgbouncer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.1/docs/resources/database_postgresql_config#autodb_idle_timeout DatabasePostgresqlConfig#autodb_idle_timeout}.
	AutodbIdleTimeout *float64 `field:"optional" json:"autodbIdleTimeout" yaml:"autodbIdleTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.1/docs/resources/database_postgresql_config#autodb_max_db_connections DatabasePostgresqlConfig#autodb_max_db_connections}.
	AutodbMaxDbConnections *float64 `field:"optional" json:"autodbMaxDbConnections" yaml:"autodbMaxDbConnections"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.1/docs/resources/database_postgresql_config#autodb_pool_mode DatabasePostgresqlConfig#autodb_pool_mode}.
	AutodbPoolMode *string `field:"optional" json:"autodbPoolMode" yaml:"autodbPoolMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.1/docs/resources/database_postgresql_config#autodb_pool_size DatabasePostgresqlConfig#autodb_pool_size}.
	AutodbPoolSize *float64 `field:"optional" json:"autodbPoolSize" yaml:"autodbPoolSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.1/docs/resources/database_postgresql_config#ignore_startup_parameters DatabasePostgresqlConfig#ignore_startup_parameters}.
	IgnoreStartupParameters *[]*string `field:"optional" json:"ignoreStartupParameters" yaml:"ignoreStartupParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.1/docs/resources/database_postgresql_config#min_pool_size DatabasePostgresqlConfig#min_pool_size}.
	MinPoolSize *float64 `field:"optional" json:"minPoolSize" yaml:"minPoolSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.1/docs/resources/database_postgresql_config#server_idle_timeout DatabasePostgresqlConfig#server_idle_timeout}.
	ServerIdleTimeout *float64 `field:"optional" json:"serverIdleTimeout" yaml:"serverIdleTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.1/docs/resources/database_postgresql_config#server_lifetime DatabasePostgresqlConfig#server_lifetime}.
	ServerLifetime *float64 `field:"optional" json:"serverLifetime" yaml:"serverLifetime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.1/docs/resources/database_postgresql_config#server_reset_query_always DatabasePostgresqlConfig#server_reset_query_always}.
	ServerResetQueryAlways interface{} `field:"optional" json:"serverResetQueryAlways" yaml:"serverResetQueryAlways"`
}

