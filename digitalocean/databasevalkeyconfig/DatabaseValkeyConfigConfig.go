// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasevalkeyconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseValkeyConfigConfig struct {
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
	// A unique identifier for the database cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/database_valkey_config#cluster_id DatabaseValkeyConfig#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Determines default pub/sub channels' ACL for new users if ACL is not supplied.
	//
	// When this option is not defined, all_channels is assumed to keep backward compatibility. This option doesn't affect Valkey configuration acl-pubsub-default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/database_valkey_config#acl_channels_default DatabaseValkeyConfig#acl_channels_default}
	AclChannelsDefault *string `field:"optional" json:"aclChannelsDefault" yaml:"aclChannelsDefault"`
	// Frequent RDB snapshots.
	//
	// When enabled, Valkey will create frequent local RDB snapshots. When disabled, Valkey will only take RDB snapshots when a backup is created, based on the backup schedule. This setting is ignored when valkey_persistence is set to off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/database_valkey_config#frequent_snapshots DatabaseValkeyConfig#frequent_snapshots}
	FrequentSnapshots interface{} `field:"optional" json:"frequentSnapshots" yaml:"frequentSnapshots"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/database_valkey_config#id DatabaseValkeyConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The number of IO threads used by Valkey. Must be between 1 and 32.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/database_valkey_config#io_threads DatabaseValkeyConfig#io_threads}
	IoThreads *float64 `field:"optional" json:"ioThreads" yaml:"ioThreads"`
	// The decay time for Valkey's LFU cache eviction. Must be between 1 and 120.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/database_valkey_config#lfu_decay_time DatabaseValkeyConfig#lfu_decay_time}
	LfuDecayTime *float64 `field:"optional" json:"lfuDecayTime" yaml:"lfuDecayTime"`
	// The log factor for Valkey's LFU (Least Frequently Used) cache eviction. Must be between 1 and 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/database_valkey_config#lfu_log_factor DatabaseValkeyConfig#lfu_log_factor}
	LfuLogFactor *float64 `field:"optional" json:"lfuLogFactor" yaml:"lfuLogFactor"`
	// Set notify-keyspace-events option.
	//
	// Requires at least K or E and accepts any combination of the following options. Setting the parameter to "" disables notifications.
	//
	// K — Keyspace events
	// E — Keyevent events
	// g — Generic commands (e.g. DEL, EXPIRE, RENAME, ...)
	// $ — String commands
	// l — List commands
	// s — Set commands
	// h — Hash commands
	// z — Sorted set commands
	// t — Stream commands
	// d — Module key type events
	// x — Expired events
	// e — Evicted events
	// m — Key miss events
	// n — New key events
	// A — Alias for "g$lshztxed"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/database_valkey_config#notify_keyspace_events DatabaseValkeyConfig#notify_keyspace_events}
	NotifyKeyspaceEvents *string `field:"optional" json:"notifyKeyspaceEvents" yaml:"notifyKeyspaceEvents"`
	// The number of logical databases in the Valkey cluster. Must be between 1 and 128.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/database_valkey_config#number_of_databases DatabaseValkeyConfig#number_of_databases}
	NumberOfDatabases *float64 `field:"optional" json:"numberOfDatabases" yaml:"numberOfDatabases"`
	// When persistence is 'rdb', Valkey does RDB dumps each 10 minutes if any key is changed.
	//
	// Also RDB dumps are done according to backup schedule for backup purposes. When persistence is 'off', no RDB dumps and backups are done, so data can be lost at any moment if service is restarted for any reason, or if service is powered off. Also service can't be forked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/database_valkey_config#persistence DatabaseValkeyConfig#persistence}
	Persistence *string `field:"optional" json:"persistence" yaml:"persistence"`
	// Set output buffer limit for pub / sub clients in MB.
	//
	// The value is the hard limit, the soft limit is 1/4 of the hard limit. When setting the limit, be mindful of the available memory in the selected service plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/database_valkey_config#pubsub_client_output_buffer_limit DatabaseValkeyConfig#pubsub_client_output_buffer_limit}
	PubsubClientOutputBufferLimit *float64 `field:"optional" json:"pubsubClientOutputBufferLimit" yaml:"pubsubClientOutputBufferLimit"`
	// Whether to enable SSL/TLS for connections to the Valkey cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/database_valkey_config#ssl DatabaseValkeyConfig#ssl}
	Ssl interface{} `field:"optional" json:"ssl" yaml:"ssl"`
	// The timeout (in seconds) for Valkey client connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/database_valkey_config#timeout DatabaseValkeyConfig#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Active expire effort.
	//
	// Valkey reclaims expired keys both when accessed and in the background. The background process scans for expired keys to free memory. Increasing the active-expire-effort setting (default 1, max 10) uses more CPU to reclaim expired keys faster, reducing memory usage but potentially increasing latency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/database_valkey_config#valkey_active_expire_effort DatabaseValkeyConfig#valkey_active_expire_effort}
	ValkeyActiveExpireEffort *float64 `field:"optional" json:"valkeyActiveExpireEffort" yaml:"valkeyActiveExpireEffort"`
}

