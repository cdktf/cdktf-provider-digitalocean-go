// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasemongodbconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseMongodbConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/database_mongodb_config#cluster_id DatabaseMongodbConfig#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/database_mongodb_config#default_read_concern DatabaseMongodbConfig#default_read_concern}.
	DefaultReadConcern *string `field:"optional" json:"defaultReadConcern" yaml:"defaultReadConcern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/database_mongodb_config#default_write_concern DatabaseMongodbConfig#default_write_concern}.
	DefaultWriteConcern *string `field:"optional" json:"defaultWriteConcern" yaml:"defaultWriteConcern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/database_mongodb_config#id DatabaseMongodbConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/database_mongodb_config#slow_op_threshold_ms DatabaseMongodbConfig#slow_op_threshold_ms}.
	SlowOpThresholdMs *float64 `field:"optional" json:"slowOpThresholdMs" yaml:"slowOpThresholdMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/database_mongodb_config#transaction_lifetime_limit_seconds DatabaseMongodbConfig#transaction_lifetime_limit_seconds}.
	TransactionLifetimeLimitSeconds *float64 `field:"optional" json:"transactionLifetimeLimitSeconds" yaml:"transactionLifetimeLimitSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/database_mongodb_config#verbosity DatabaseMongodbConfig#verbosity}.
	Verbosity *float64 `field:"optional" json:"verbosity" yaml:"verbosity"`
}

