// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecDatabase struct {
	// The name of the underlying DigitalOcean DBaaS cluster.
	//
	// This is required for production databases. For dev databases, if cluster_name is not set, a new cluster will be provisioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.54.0/docs/resources/app#cluster_name App#cluster_name}
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The name of the MySQL or PostgreSQL database to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.54.0/docs/resources/app#db_name App#db_name}
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// The name of the MySQL or PostgreSQL user to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.54.0/docs/resources/app#db_user App#db_user}
	DbUser *string `field:"optional" json:"dbUser" yaml:"dbUser"`
	// The database engine to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.54.0/docs/resources/app#engine App#engine}
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The name of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.54.0/docs/resources/app#name App#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Whether this is a production or dev database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.54.0/docs/resources/app#production App#production}
	Production interface{} `field:"optional" json:"production" yaml:"production"`
	// The version of the database engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.54.0/docs/resources/app#version App#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

