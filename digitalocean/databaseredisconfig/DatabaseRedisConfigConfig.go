// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseredisconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseRedisConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.0/docs/resources/database_redis_config#cluster_id DatabaseRedisConfig#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.0/docs/resources/database_redis_config#acl_channels_default DatabaseRedisConfig#acl_channels_default}.
	AclChannelsDefault *string `field:"optional" json:"aclChannelsDefault" yaml:"aclChannelsDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.0/docs/resources/database_redis_config#id DatabaseRedisConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.0/docs/resources/database_redis_config#io_threads DatabaseRedisConfig#io_threads}.
	IoThreads *float64 `field:"optional" json:"ioThreads" yaml:"ioThreads"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.0/docs/resources/database_redis_config#lfu_decay_time DatabaseRedisConfig#lfu_decay_time}.
	LfuDecayTime *float64 `field:"optional" json:"lfuDecayTime" yaml:"lfuDecayTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.0/docs/resources/database_redis_config#lfu_log_factor DatabaseRedisConfig#lfu_log_factor}.
	LfuLogFactor *float64 `field:"optional" json:"lfuLogFactor" yaml:"lfuLogFactor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.0/docs/resources/database_redis_config#maxmemory_policy DatabaseRedisConfig#maxmemory_policy}.
	MaxmemoryPolicy *string `field:"optional" json:"maxmemoryPolicy" yaml:"maxmemoryPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.0/docs/resources/database_redis_config#notify_keyspace_events DatabaseRedisConfig#notify_keyspace_events}.
	NotifyKeyspaceEvents *string `field:"optional" json:"notifyKeyspaceEvents" yaml:"notifyKeyspaceEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.0/docs/resources/database_redis_config#number_of_databases DatabaseRedisConfig#number_of_databases}.
	NumberOfDatabases *float64 `field:"optional" json:"numberOfDatabases" yaml:"numberOfDatabases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.0/docs/resources/database_redis_config#persistence DatabaseRedisConfig#persistence}.
	Persistence *string `field:"optional" json:"persistence" yaml:"persistence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.0/docs/resources/database_redis_config#pubsub_client_output_buffer_limit DatabaseRedisConfig#pubsub_client_output_buffer_limit}.
	PubsubClientOutputBufferLimit *float64 `field:"optional" json:"pubsubClientOutputBufferLimit" yaml:"pubsubClientOutputBufferLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.0/docs/resources/database_redis_config#ssl DatabaseRedisConfig#ssl}.
	Ssl interface{} `field:"optional" json:"ssl" yaml:"ssl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.0/docs/resources/database_redis_config#timeout DatabaseRedisConfig#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

