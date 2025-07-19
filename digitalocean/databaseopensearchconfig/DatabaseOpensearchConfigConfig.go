// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseopensearchconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseOpensearchConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#cluster_id DatabaseOpensearchConfig#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#action_auto_create_index_enabled DatabaseOpensearchConfig#action_auto_create_index_enabled}.
	ActionAutoCreateIndexEnabled interface{} `field:"optional" json:"actionAutoCreateIndexEnabled" yaml:"actionAutoCreateIndexEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#action_destructive_requires_name DatabaseOpensearchConfig#action_destructive_requires_name}.
	ActionDestructiveRequiresName interface{} `field:"optional" json:"actionDestructiveRequiresName" yaml:"actionDestructiveRequiresName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#cluster_max_shards_per_node DatabaseOpensearchConfig#cluster_max_shards_per_node}.
	ClusterMaxShardsPerNode *float64 `field:"optional" json:"clusterMaxShardsPerNode" yaml:"clusterMaxShardsPerNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#cluster_routing_allocation_node_concurrent_recoveries DatabaseOpensearchConfig#cluster_routing_allocation_node_concurrent_recoveries}.
	ClusterRoutingAllocationNodeConcurrentRecoveries *float64 `field:"optional" json:"clusterRoutingAllocationNodeConcurrentRecoveries" yaml:"clusterRoutingAllocationNodeConcurrentRecoveries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#enable_security_audit DatabaseOpensearchConfig#enable_security_audit}.
	EnableSecurityAudit interface{} `field:"optional" json:"enableSecurityAudit" yaml:"enableSecurityAudit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#http_max_content_length_bytes DatabaseOpensearchConfig#http_max_content_length_bytes}.
	HttpMaxContentLengthBytes *float64 `field:"optional" json:"httpMaxContentLengthBytes" yaml:"httpMaxContentLengthBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#http_max_header_size_bytes DatabaseOpensearchConfig#http_max_header_size_bytes}.
	HttpMaxHeaderSizeBytes *float64 `field:"optional" json:"httpMaxHeaderSizeBytes" yaml:"httpMaxHeaderSizeBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#http_max_initial_line_length_bytes DatabaseOpensearchConfig#http_max_initial_line_length_bytes}.
	HttpMaxInitialLineLengthBytes *float64 `field:"optional" json:"httpMaxInitialLineLengthBytes" yaml:"httpMaxInitialLineLengthBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#id DatabaseOpensearchConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#indices_fielddata_cache_size_percentage DatabaseOpensearchConfig#indices_fielddata_cache_size_percentage}.
	IndicesFielddataCacheSizePercentage *float64 `field:"optional" json:"indicesFielddataCacheSizePercentage" yaml:"indicesFielddataCacheSizePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#indices_memory_index_buffer_size_percentage DatabaseOpensearchConfig#indices_memory_index_buffer_size_percentage}.
	IndicesMemoryIndexBufferSizePercentage *float64 `field:"optional" json:"indicesMemoryIndexBufferSizePercentage" yaml:"indicesMemoryIndexBufferSizePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#indices_memory_max_index_buffer_size_mb DatabaseOpensearchConfig#indices_memory_max_index_buffer_size_mb}.
	IndicesMemoryMaxIndexBufferSizeMb *float64 `field:"optional" json:"indicesMemoryMaxIndexBufferSizeMb" yaml:"indicesMemoryMaxIndexBufferSizeMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#indices_memory_min_index_buffer_size_mb DatabaseOpensearchConfig#indices_memory_min_index_buffer_size_mb}.
	IndicesMemoryMinIndexBufferSizeMb *float64 `field:"optional" json:"indicesMemoryMinIndexBufferSizeMb" yaml:"indicesMemoryMinIndexBufferSizeMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#indices_queries_cache_size_percentage DatabaseOpensearchConfig#indices_queries_cache_size_percentage}.
	IndicesQueriesCacheSizePercentage *float64 `field:"optional" json:"indicesQueriesCacheSizePercentage" yaml:"indicesQueriesCacheSizePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#indices_query_bool_max_clause_count DatabaseOpensearchConfig#indices_query_bool_max_clause_count}.
	IndicesQueryBoolMaxClauseCount *float64 `field:"optional" json:"indicesQueryBoolMaxClauseCount" yaml:"indicesQueryBoolMaxClauseCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#indices_recovery_max_concurrent_file_chunks DatabaseOpensearchConfig#indices_recovery_max_concurrent_file_chunks}.
	IndicesRecoveryMaxConcurrentFileChunks *float64 `field:"optional" json:"indicesRecoveryMaxConcurrentFileChunks" yaml:"indicesRecoveryMaxConcurrentFileChunks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#indices_recovery_max_mb_per_sec DatabaseOpensearchConfig#indices_recovery_max_mb_per_sec}.
	IndicesRecoveryMaxMbPerSec *float64 `field:"optional" json:"indicesRecoveryMaxMbPerSec" yaml:"indicesRecoveryMaxMbPerSec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#ism_enabled DatabaseOpensearchConfig#ism_enabled}.
	IsmEnabled interface{} `field:"optional" json:"ismEnabled" yaml:"ismEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#ism_history_enabled DatabaseOpensearchConfig#ism_history_enabled}.
	IsmHistoryEnabled interface{} `field:"optional" json:"ismHistoryEnabled" yaml:"ismHistoryEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#ism_history_max_age_hours DatabaseOpensearchConfig#ism_history_max_age_hours}.
	IsmHistoryMaxAgeHours *float64 `field:"optional" json:"ismHistoryMaxAgeHours" yaml:"ismHistoryMaxAgeHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#ism_history_max_docs DatabaseOpensearchConfig#ism_history_max_docs}.
	IsmHistoryMaxDocs *float64 `field:"optional" json:"ismHistoryMaxDocs" yaml:"ismHistoryMaxDocs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#ism_history_rollover_check_period_hours DatabaseOpensearchConfig#ism_history_rollover_check_period_hours}.
	IsmHistoryRolloverCheckPeriodHours *float64 `field:"optional" json:"ismHistoryRolloverCheckPeriodHours" yaml:"ismHistoryRolloverCheckPeriodHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#ism_history_rollover_retention_period_days DatabaseOpensearchConfig#ism_history_rollover_retention_period_days}.
	IsmHistoryRolloverRetentionPeriodDays *float64 `field:"optional" json:"ismHistoryRolloverRetentionPeriodDays" yaml:"ismHistoryRolloverRetentionPeriodDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#override_main_response_version DatabaseOpensearchConfig#override_main_response_version}.
	OverrideMainResponseVersion interface{} `field:"optional" json:"overrideMainResponseVersion" yaml:"overrideMainResponseVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#plugins_alerting_filter_by_backend_roles_enabled DatabaseOpensearchConfig#plugins_alerting_filter_by_backend_roles_enabled}.
	PluginsAlertingFilterByBackendRolesEnabled interface{} `field:"optional" json:"pluginsAlertingFilterByBackendRolesEnabled" yaml:"pluginsAlertingFilterByBackendRolesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#reindex_remote_whitelist DatabaseOpensearchConfig#reindex_remote_whitelist}.
	ReindexRemoteWhitelist *[]*string `field:"optional" json:"reindexRemoteWhitelist" yaml:"reindexRemoteWhitelist"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#script_max_compilations_rate DatabaseOpensearchConfig#script_max_compilations_rate}.
	ScriptMaxCompilationsRate *string `field:"optional" json:"scriptMaxCompilationsRate" yaml:"scriptMaxCompilationsRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#search_max_buckets DatabaseOpensearchConfig#search_max_buckets}.
	SearchMaxBuckets *float64 `field:"optional" json:"searchMaxBuckets" yaml:"searchMaxBuckets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#thread_pool_analyze_queue_size DatabaseOpensearchConfig#thread_pool_analyze_queue_size}.
	ThreadPoolAnalyzeQueueSize *float64 `field:"optional" json:"threadPoolAnalyzeQueueSize" yaml:"threadPoolAnalyzeQueueSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#thread_pool_analyze_size DatabaseOpensearchConfig#thread_pool_analyze_size}.
	ThreadPoolAnalyzeSize *float64 `field:"optional" json:"threadPoolAnalyzeSize" yaml:"threadPoolAnalyzeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#thread_pool_force_merge_size DatabaseOpensearchConfig#thread_pool_force_merge_size}.
	ThreadPoolForceMergeSize *float64 `field:"optional" json:"threadPoolForceMergeSize" yaml:"threadPoolForceMergeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#thread_pool_get_queue_size DatabaseOpensearchConfig#thread_pool_get_queue_size}.
	ThreadPoolGetQueueSize *float64 `field:"optional" json:"threadPoolGetQueueSize" yaml:"threadPoolGetQueueSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#thread_pool_get_size DatabaseOpensearchConfig#thread_pool_get_size}.
	ThreadPoolGetSize *float64 `field:"optional" json:"threadPoolGetSize" yaml:"threadPoolGetSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#thread_pool_search_queue_size DatabaseOpensearchConfig#thread_pool_search_queue_size}.
	ThreadPoolSearchQueueSize *float64 `field:"optional" json:"threadPoolSearchQueueSize" yaml:"threadPoolSearchQueueSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#thread_pool_search_size DatabaseOpensearchConfig#thread_pool_search_size}.
	ThreadPoolSearchSize *float64 `field:"optional" json:"threadPoolSearchSize" yaml:"threadPoolSearchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#thread_pool_search_throttled_queue_size DatabaseOpensearchConfig#thread_pool_search_throttled_queue_size}.
	ThreadPoolSearchThrottledQueueSize *float64 `field:"optional" json:"threadPoolSearchThrottledQueueSize" yaml:"threadPoolSearchThrottledQueueSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#thread_pool_search_throttled_size DatabaseOpensearchConfig#thread_pool_search_throttled_size}.
	ThreadPoolSearchThrottledSize *float64 `field:"optional" json:"threadPoolSearchThrottledSize" yaml:"threadPoolSearchThrottledSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#thread_pool_write_queue_size DatabaseOpensearchConfig#thread_pool_write_queue_size}.
	ThreadPoolWriteQueueSize *float64 `field:"optional" json:"threadPoolWriteQueueSize" yaml:"threadPoolWriteQueueSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_opensearch_config#thread_pool_write_size DatabaseOpensearchConfig#thread_pool_write_size}.
	ThreadPoolWriteSize *float64 `field:"optional" json:"threadPoolWriteSize" yaml:"threadPoolWriteSize"`
}

