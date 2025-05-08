// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseopensearchconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v11/databaseopensearchconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.53.0/docs/resources/database_opensearch_config digitalocean_database_opensearch_config}.
type DatabaseOpensearchConfig interface {
	cdktf.TerraformResource
	ActionAutoCreateIndexEnabled() interface{}
	SetActionAutoCreateIndexEnabled(val interface{})
	ActionAutoCreateIndexEnabledInput() interface{}
	ActionDestructiveRequiresName() interface{}
	SetActionDestructiveRequiresName(val interface{})
	ActionDestructiveRequiresNameInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ClusterMaxShardsPerNode() *float64
	SetClusterMaxShardsPerNode(val *float64)
	ClusterMaxShardsPerNodeInput() *float64
	ClusterRoutingAllocationNodeConcurrentRecoveries() *float64
	SetClusterRoutingAllocationNodeConcurrentRecoveries(val *float64)
	ClusterRoutingAllocationNodeConcurrentRecoveriesInput() *float64
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableSecurityAudit() interface{}
	SetEnableSecurityAudit(val interface{})
	EnableSecurityAuditInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpMaxContentLengthBytes() *float64
	SetHttpMaxContentLengthBytes(val *float64)
	HttpMaxContentLengthBytesInput() *float64
	HttpMaxHeaderSizeBytes() *float64
	SetHttpMaxHeaderSizeBytes(val *float64)
	HttpMaxHeaderSizeBytesInput() *float64
	HttpMaxInitialLineLengthBytes() *float64
	SetHttpMaxInitialLineLengthBytes(val *float64)
	HttpMaxInitialLineLengthBytesInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	IndicesFielddataCacheSizePercentage() *float64
	SetIndicesFielddataCacheSizePercentage(val *float64)
	IndicesFielddataCacheSizePercentageInput() *float64
	IndicesMemoryIndexBufferSizePercentage() *float64
	SetIndicesMemoryIndexBufferSizePercentage(val *float64)
	IndicesMemoryIndexBufferSizePercentageInput() *float64
	IndicesMemoryMaxIndexBufferSizeMb() *float64
	SetIndicesMemoryMaxIndexBufferSizeMb(val *float64)
	IndicesMemoryMaxIndexBufferSizeMbInput() *float64
	IndicesMemoryMinIndexBufferSizeMb() *float64
	SetIndicesMemoryMinIndexBufferSizeMb(val *float64)
	IndicesMemoryMinIndexBufferSizeMbInput() *float64
	IndicesQueriesCacheSizePercentage() *float64
	SetIndicesQueriesCacheSizePercentage(val *float64)
	IndicesQueriesCacheSizePercentageInput() *float64
	IndicesQueryBoolMaxClauseCount() *float64
	SetIndicesQueryBoolMaxClauseCount(val *float64)
	IndicesQueryBoolMaxClauseCountInput() *float64
	IndicesRecoveryMaxConcurrentFileChunks() *float64
	SetIndicesRecoveryMaxConcurrentFileChunks(val *float64)
	IndicesRecoveryMaxConcurrentFileChunksInput() *float64
	IndicesRecoveryMaxMbPerSec() *float64
	SetIndicesRecoveryMaxMbPerSec(val *float64)
	IndicesRecoveryMaxMbPerSecInput() *float64
	IsmEnabled() interface{}
	SetIsmEnabled(val interface{})
	IsmEnabledInput() interface{}
	IsmHistoryEnabled() interface{}
	SetIsmHistoryEnabled(val interface{})
	IsmHistoryEnabledInput() interface{}
	IsmHistoryMaxAgeHours() *float64
	SetIsmHistoryMaxAgeHours(val *float64)
	IsmHistoryMaxAgeHoursInput() *float64
	IsmHistoryMaxDocs() *float64
	SetIsmHistoryMaxDocs(val *float64)
	IsmHistoryMaxDocsInput() *float64
	IsmHistoryRolloverCheckPeriodHours() *float64
	SetIsmHistoryRolloverCheckPeriodHours(val *float64)
	IsmHistoryRolloverCheckPeriodHoursInput() *float64
	IsmHistoryRolloverRetentionPeriodDays() *float64
	SetIsmHistoryRolloverRetentionPeriodDays(val *float64)
	IsmHistoryRolloverRetentionPeriodDaysInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OverrideMainResponseVersion() interface{}
	SetOverrideMainResponseVersion(val interface{})
	OverrideMainResponseVersionInput() interface{}
	PluginsAlertingFilterByBackendRolesEnabled() interface{}
	SetPluginsAlertingFilterByBackendRolesEnabled(val interface{})
	PluginsAlertingFilterByBackendRolesEnabledInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReindexRemoteWhitelist() *[]*string
	SetReindexRemoteWhitelist(val *[]*string)
	ReindexRemoteWhitelistInput() *[]*string
	ScriptMaxCompilationsRate() *string
	SetScriptMaxCompilationsRate(val *string)
	ScriptMaxCompilationsRateInput() *string
	SearchMaxBuckets() *float64
	SetSearchMaxBuckets(val *float64)
	SearchMaxBucketsInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThreadPoolAnalyzeQueueSize() *float64
	SetThreadPoolAnalyzeQueueSize(val *float64)
	ThreadPoolAnalyzeQueueSizeInput() *float64
	ThreadPoolAnalyzeSize() *float64
	SetThreadPoolAnalyzeSize(val *float64)
	ThreadPoolAnalyzeSizeInput() *float64
	ThreadPoolForceMergeSize() *float64
	SetThreadPoolForceMergeSize(val *float64)
	ThreadPoolForceMergeSizeInput() *float64
	ThreadPoolGetQueueSize() *float64
	SetThreadPoolGetQueueSize(val *float64)
	ThreadPoolGetQueueSizeInput() *float64
	ThreadPoolGetSize() *float64
	SetThreadPoolGetSize(val *float64)
	ThreadPoolGetSizeInput() *float64
	ThreadPoolSearchQueueSize() *float64
	SetThreadPoolSearchQueueSize(val *float64)
	ThreadPoolSearchQueueSizeInput() *float64
	ThreadPoolSearchSize() *float64
	SetThreadPoolSearchSize(val *float64)
	ThreadPoolSearchSizeInput() *float64
	ThreadPoolSearchThrottledQueueSize() *float64
	SetThreadPoolSearchThrottledQueueSize(val *float64)
	ThreadPoolSearchThrottledQueueSizeInput() *float64
	ThreadPoolSearchThrottledSize() *float64
	SetThreadPoolSearchThrottledSize(val *float64)
	ThreadPoolSearchThrottledSizeInput() *float64
	ThreadPoolWriteQueueSize() *float64
	SetThreadPoolWriteQueueSize(val *float64)
	ThreadPoolWriteQueueSizeInput() *float64
	ThreadPoolWriteSize() *float64
	SetThreadPoolWriteSize(val *float64)
	ThreadPoolWriteSizeInput() *float64
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetActionAutoCreateIndexEnabled()
	ResetActionDestructiveRequiresName()
	ResetClusterMaxShardsPerNode()
	ResetClusterRoutingAllocationNodeConcurrentRecoveries()
	ResetEnableSecurityAudit()
	ResetHttpMaxContentLengthBytes()
	ResetHttpMaxHeaderSizeBytes()
	ResetHttpMaxInitialLineLengthBytes()
	ResetId()
	ResetIndicesFielddataCacheSizePercentage()
	ResetIndicesMemoryIndexBufferSizePercentage()
	ResetIndicesMemoryMaxIndexBufferSizeMb()
	ResetIndicesMemoryMinIndexBufferSizeMb()
	ResetIndicesQueriesCacheSizePercentage()
	ResetIndicesQueryBoolMaxClauseCount()
	ResetIndicesRecoveryMaxConcurrentFileChunks()
	ResetIndicesRecoveryMaxMbPerSec()
	ResetIsmEnabled()
	ResetIsmHistoryEnabled()
	ResetIsmHistoryMaxAgeHours()
	ResetIsmHistoryMaxDocs()
	ResetIsmHistoryRolloverCheckPeriodHours()
	ResetIsmHistoryRolloverRetentionPeriodDays()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOverrideMainResponseVersion()
	ResetPluginsAlertingFilterByBackendRolesEnabled()
	ResetReindexRemoteWhitelist()
	ResetScriptMaxCompilationsRate()
	ResetSearchMaxBuckets()
	ResetThreadPoolAnalyzeQueueSize()
	ResetThreadPoolAnalyzeSize()
	ResetThreadPoolForceMergeSize()
	ResetThreadPoolGetQueueSize()
	ResetThreadPoolGetSize()
	ResetThreadPoolSearchQueueSize()
	ResetThreadPoolSearchSize()
	ResetThreadPoolSearchThrottledQueueSize()
	ResetThreadPoolSearchThrottledSize()
	ResetThreadPoolWriteQueueSize()
	ResetThreadPoolWriteSize()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DatabaseOpensearchConfig
type jsiiProxy_DatabaseOpensearchConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ActionAutoCreateIndexEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionAutoCreateIndexEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ActionAutoCreateIndexEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionAutoCreateIndexEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ActionDestructiveRequiresName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionDestructiveRequiresName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ActionDestructiveRequiresNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionDestructiveRequiresNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ClusterMaxShardsPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterMaxShardsPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ClusterMaxShardsPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterMaxShardsPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ClusterRoutingAllocationNodeConcurrentRecoveries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterRoutingAllocationNodeConcurrentRecoveries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ClusterRoutingAllocationNodeConcurrentRecoveriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterRoutingAllocationNodeConcurrentRecoveriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) EnableSecurityAudit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSecurityAudit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) EnableSecurityAuditInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSecurityAuditInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) HttpMaxContentLengthBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpMaxContentLengthBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) HttpMaxContentLengthBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpMaxContentLengthBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) HttpMaxHeaderSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpMaxHeaderSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) HttpMaxHeaderSizeBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpMaxHeaderSizeBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) HttpMaxInitialLineLengthBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpMaxInitialLineLengthBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) HttpMaxInitialLineLengthBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpMaxInitialLineLengthBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IndicesFielddataCacheSizePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesFielddataCacheSizePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IndicesFielddataCacheSizePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesFielddataCacheSizePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IndicesMemoryIndexBufferSizePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesMemoryIndexBufferSizePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IndicesMemoryIndexBufferSizePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesMemoryIndexBufferSizePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IndicesMemoryMaxIndexBufferSizeMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesMemoryMaxIndexBufferSizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IndicesMemoryMaxIndexBufferSizeMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesMemoryMaxIndexBufferSizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IndicesMemoryMinIndexBufferSizeMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesMemoryMinIndexBufferSizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IndicesMemoryMinIndexBufferSizeMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesMemoryMinIndexBufferSizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IndicesQueriesCacheSizePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesQueriesCacheSizePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IndicesQueriesCacheSizePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesQueriesCacheSizePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IndicesQueryBoolMaxClauseCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesQueryBoolMaxClauseCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IndicesQueryBoolMaxClauseCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesQueryBoolMaxClauseCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IndicesRecoveryMaxConcurrentFileChunks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesRecoveryMaxConcurrentFileChunks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IndicesRecoveryMaxConcurrentFileChunksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesRecoveryMaxConcurrentFileChunksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IndicesRecoveryMaxMbPerSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesRecoveryMaxMbPerSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IndicesRecoveryMaxMbPerSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesRecoveryMaxMbPerSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IsmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ismEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IsmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ismEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IsmHistoryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ismHistoryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IsmHistoryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ismHistoryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IsmHistoryMaxAgeHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismHistoryMaxAgeHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IsmHistoryMaxAgeHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismHistoryMaxAgeHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IsmHistoryMaxDocs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismHistoryMaxDocs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IsmHistoryMaxDocsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismHistoryMaxDocsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IsmHistoryRolloverCheckPeriodHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismHistoryRolloverCheckPeriodHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IsmHistoryRolloverCheckPeriodHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismHistoryRolloverCheckPeriodHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IsmHistoryRolloverRetentionPeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismHistoryRolloverRetentionPeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) IsmHistoryRolloverRetentionPeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismHistoryRolloverRetentionPeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) OverrideMainResponseVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideMainResponseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) OverrideMainResponseVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideMainResponseVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) PluginsAlertingFilterByBackendRolesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pluginsAlertingFilterByBackendRolesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) PluginsAlertingFilterByBackendRolesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pluginsAlertingFilterByBackendRolesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ReindexRemoteWhitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reindexRemoteWhitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ReindexRemoteWhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reindexRemoteWhitelistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ScriptMaxCompilationsRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptMaxCompilationsRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ScriptMaxCompilationsRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptMaxCompilationsRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) SearchMaxBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"searchMaxBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) SearchMaxBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"searchMaxBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolAnalyzeQueueSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolAnalyzeQueueSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolAnalyzeQueueSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolAnalyzeQueueSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolAnalyzeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolAnalyzeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolAnalyzeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolAnalyzeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolForceMergeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolForceMergeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolForceMergeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolForceMergeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolGetQueueSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolGetQueueSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolGetQueueSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolGetQueueSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolGetSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolGetSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolGetSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolGetSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolSearchQueueSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolSearchQueueSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolSearchQueueSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolSearchQueueSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolSearchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolSearchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolSearchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolSearchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolSearchThrottledQueueSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolSearchThrottledQueueSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolSearchThrottledQueueSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolSearchThrottledQueueSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolSearchThrottledSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolSearchThrottledSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolSearchThrottledSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolSearchThrottledSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolWriteQueueSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolWriteQueueSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolWriteQueueSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolWriteQueueSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolWriteSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolWriteSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseOpensearchConfig) ThreadPoolWriteSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolWriteSizeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.53.0/docs/resources/database_opensearch_config digitalocean_database_opensearch_config} Resource.
func NewDatabaseOpensearchConfig(scope constructs.Construct, id *string, config *DatabaseOpensearchConfigConfig) DatabaseOpensearchConfig {
	_init_.Initialize()

	if err := validateNewDatabaseOpensearchConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseOpensearchConfig{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.databaseOpensearchConfig.DatabaseOpensearchConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.53.0/docs/resources/database_opensearch_config digitalocean_database_opensearch_config} Resource.
func NewDatabaseOpensearchConfig_Override(d DatabaseOpensearchConfig, scope constructs.Construct, id *string, config *DatabaseOpensearchConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.databaseOpensearchConfig.DatabaseOpensearchConfig",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetActionAutoCreateIndexEnabled(val interface{}) {
	if err := j.validateSetActionAutoCreateIndexEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionAutoCreateIndexEnabled",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetActionDestructiveRequiresName(val interface{}) {
	if err := j.validateSetActionDestructiveRequiresNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionDestructiveRequiresName",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetClusterMaxShardsPerNode(val *float64) {
	if err := j.validateSetClusterMaxShardsPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterMaxShardsPerNode",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetClusterRoutingAllocationNodeConcurrentRecoveries(val *float64) {
	if err := j.validateSetClusterRoutingAllocationNodeConcurrentRecoveriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterRoutingAllocationNodeConcurrentRecoveries",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetEnableSecurityAudit(val interface{}) {
	if err := j.validateSetEnableSecurityAuditParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSecurityAudit",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetHttpMaxContentLengthBytes(val *float64) {
	if err := j.validateSetHttpMaxContentLengthBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMaxContentLengthBytes",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetHttpMaxHeaderSizeBytes(val *float64) {
	if err := j.validateSetHttpMaxHeaderSizeBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMaxHeaderSizeBytes",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetHttpMaxInitialLineLengthBytes(val *float64) {
	if err := j.validateSetHttpMaxInitialLineLengthBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMaxInitialLineLengthBytes",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetIndicesFielddataCacheSizePercentage(val *float64) {
	if err := j.validateSetIndicesFielddataCacheSizePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indicesFielddataCacheSizePercentage",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetIndicesMemoryIndexBufferSizePercentage(val *float64) {
	if err := j.validateSetIndicesMemoryIndexBufferSizePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indicesMemoryIndexBufferSizePercentage",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetIndicesMemoryMaxIndexBufferSizeMb(val *float64) {
	if err := j.validateSetIndicesMemoryMaxIndexBufferSizeMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indicesMemoryMaxIndexBufferSizeMb",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetIndicesMemoryMinIndexBufferSizeMb(val *float64) {
	if err := j.validateSetIndicesMemoryMinIndexBufferSizeMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indicesMemoryMinIndexBufferSizeMb",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetIndicesQueriesCacheSizePercentage(val *float64) {
	if err := j.validateSetIndicesQueriesCacheSizePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indicesQueriesCacheSizePercentage",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetIndicesQueryBoolMaxClauseCount(val *float64) {
	if err := j.validateSetIndicesQueryBoolMaxClauseCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indicesQueryBoolMaxClauseCount",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetIndicesRecoveryMaxConcurrentFileChunks(val *float64) {
	if err := j.validateSetIndicesRecoveryMaxConcurrentFileChunksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indicesRecoveryMaxConcurrentFileChunks",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetIndicesRecoveryMaxMbPerSec(val *float64) {
	if err := j.validateSetIndicesRecoveryMaxMbPerSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indicesRecoveryMaxMbPerSec",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetIsmEnabled(val interface{}) {
	if err := j.validateSetIsmEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ismEnabled",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetIsmHistoryEnabled(val interface{}) {
	if err := j.validateSetIsmHistoryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ismHistoryEnabled",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetIsmHistoryMaxAgeHours(val *float64) {
	if err := j.validateSetIsmHistoryMaxAgeHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ismHistoryMaxAgeHours",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetIsmHistoryMaxDocs(val *float64) {
	if err := j.validateSetIsmHistoryMaxDocsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ismHistoryMaxDocs",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetIsmHistoryRolloverCheckPeriodHours(val *float64) {
	if err := j.validateSetIsmHistoryRolloverCheckPeriodHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ismHistoryRolloverCheckPeriodHours",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetIsmHistoryRolloverRetentionPeriodDays(val *float64) {
	if err := j.validateSetIsmHistoryRolloverRetentionPeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ismHistoryRolloverRetentionPeriodDays",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetOverrideMainResponseVersion(val interface{}) {
	if err := j.validateSetOverrideMainResponseVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideMainResponseVersion",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetPluginsAlertingFilterByBackendRolesEnabled(val interface{}) {
	if err := j.validateSetPluginsAlertingFilterByBackendRolesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginsAlertingFilterByBackendRolesEnabled",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetReindexRemoteWhitelist(val *[]*string) {
	if err := j.validateSetReindexRemoteWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reindexRemoteWhitelist",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetScriptMaxCompilationsRate(val *string) {
	if err := j.validateSetScriptMaxCompilationsRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptMaxCompilationsRate",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetSearchMaxBuckets(val *float64) {
	if err := j.validateSetSearchMaxBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchMaxBuckets",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetThreadPoolAnalyzeQueueSize(val *float64) {
	if err := j.validateSetThreadPoolAnalyzeQueueSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolAnalyzeQueueSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetThreadPoolAnalyzeSize(val *float64) {
	if err := j.validateSetThreadPoolAnalyzeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolAnalyzeSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetThreadPoolForceMergeSize(val *float64) {
	if err := j.validateSetThreadPoolForceMergeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolForceMergeSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetThreadPoolGetQueueSize(val *float64) {
	if err := j.validateSetThreadPoolGetQueueSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolGetQueueSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetThreadPoolGetSize(val *float64) {
	if err := j.validateSetThreadPoolGetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolGetSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetThreadPoolSearchQueueSize(val *float64) {
	if err := j.validateSetThreadPoolSearchQueueSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolSearchQueueSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetThreadPoolSearchSize(val *float64) {
	if err := j.validateSetThreadPoolSearchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolSearchSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetThreadPoolSearchThrottledQueueSize(val *float64) {
	if err := j.validateSetThreadPoolSearchThrottledQueueSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolSearchThrottledQueueSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetThreadPoolSearchThrottledSize(val *float64) {
	if err := j.validateSetThreadPoolSearchThrottledSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolSearchThrottledSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetThreadPoolWriteQueueSize(val *float64) {
	if err := j.validateSetThreadPoolWriteQueueSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolWriteQueueSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseOpensearchConfig)SetThreadPoolWriteSize(val *float64) {
	if err := j.validateSetThreadPoolWriteSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolWriteSize",
		val,
	)
}

// Generates CDKTF code for importing a DatabaseOpensearchConfig resource upon running "cdktf plan <stack-name>".
func DatabaseOpensearchConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatabaseOpensearchConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseOpensearchConfig.DatabaseOpensearchConfig",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DatabaseOpensearchConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseOpensearchConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseOpensearchConfig.DatabaseOpensearchConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseOpensearchConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseOpensearchConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseOpensearchConfig.DatabaseOpensearchConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseOpensearchConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseOpensearchConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseOpensearchConfig.DatabaseOpensearchConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatabaseOpensearchConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-digitalocean.databaseOpensearchConfig.DatabaseOpensearchConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabaseOpensearchConfig) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseOpensearchConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseOpensearchConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseOpensearchConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseOpensearchConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseOpensearchConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseOpensearchConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseOpensearchConfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseOpensearchConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseOpensearchConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseOpensearchConfig) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetActionAutoCreateIndexEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetActionAutoCreateIndexEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetActionDestructiveRequiresName() {
	_jsii_.InvokeVoid(
		d,
		"resetActionDestructiveRequiresName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetClusterMaxShardsPerNode() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterMaxShardsPerNode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetClusterRoutingAllocationNodeConcurrentRecoveries() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterRoutingAllocationNodeConcurrentRecoveries",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetEnableSecurityAudit() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableSecurityAudit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetHttpMaxContentLengthBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpMaxContentLengthBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetHttpMaxHeaderSizeBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpMaxHeaderSizeBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetHttpMaxInitialLineLengthBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpMaxInitialLineLengthBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetIndicesFielddataCacheSizePercentage() {
	_jsii_.InvokeVoid(
		d,
		"resetIndicesFielddataCacheSizePercentage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetIndicesMemoryIndexBufferSizePercentage() {
	_jsii_.InvokeVoid(
		d,
		"resetIndicesMemoryIndexBufferSizePercentage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetIndicesMemoryMaxIndexBufferSizeMb() {
	_jsii_.InvokeVoid(
		d,
		"resetIndicesMemoryMaxIndexBufferSizeMb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetIndicesMemoryMinIndexBufferSizeMb() {
	_jsii_.InvokeVoid(
		d,
		"resetIndicesMemoryMinIndexBufferSizeMb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetIndicesQueriesCacheSizePercentage() {
	_jsii_.InvokeVoid(
		d,
		"resetIndicesQueriesCacheSizePercentage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetIndicesQueryBoolMaxClauseCount() {
	_jsii_.InvokeVoid(
		d,
		"resetIndicesQueryBoolMaxClauseCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetIndicesRecoveryMaxConcurrentFileChunks() {
	_jsii_.InvokeVoid(
		d,
		"resetIndicesRecoveryMaxConcurrentFileChunks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetIndicesRecoveryMaxMbPerSec() {
	_jsii_.InvokeVoid(
		d,
		"resetIndicesRecoveryMaxMbPerSec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetIsmEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetIsmEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetIsmHistoryEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetIsmHistoryEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetIsmHistoryMaxAgeHours() {
	_jsii_.InvokeVoid(
		d,
		"resetIsmHistoryMaxAgeHours",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetIsmHistoryMaxDocs() {
	_jsii_.InvokeVoid(
		d,
		"resetIsmHistoryMaxDocs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetIsmHistoryRolloverCheckPeriodHours() {
	_jsii_.InvokeVoid(
		d,
		"resetIsmHistoryRolloverCheckPeriodHours",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetIsmHistoryRolloverRetentionPeriodDays() {
	_jsii_.InvokeVoid(
		d,
		"resetIsmHistoryRolloverRetentionPeriodDays",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetOverrideMainResponseVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideMainResponseVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetPluginsAlertingFilterByBackendRolesEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetPluginsAlertingFilterByBackendRolesEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetReindexRemoteWhitelist() {
	_jsii_.InvokeVoid(
		d,
		"resetReindexRemoteWhitelist",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetScriptMaxCompilationsRate() {
	_jsii_.InvokeVoid(
		d,
		"resetScriptMaxCompilationsRate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetSearchMaxBuckets() {
	_jsii_.InvokeVoid(
		d,
		"resetSearchMaxBuckets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetThreadPoolAnalyzeQueueSize() {
	_jsii_.InvokeVoid(
		d,
		"resetThreadPoolAnalyzeQueueSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetThreadPoolAnalyzeSize() {
	_jsii_.InvokeVoid(
		d,
		"resetThreadPoolAnalyzeSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetThreadPoolForceMergeSize() {
	_jsii_.InvokeVoid(
		d,
		"resetThreadPoolForceMergeSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetThreadPoolGetQueueSize() {
	_jsii_.InvokeVoid(
		d,
		"resetThreadPoolGetQueueSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetThreadPoolGetSize() {
	_jsii_.InvokeVoid(
		d,
		"resetThreadPoolGetSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetThreadPoolSearchQueueSize() {
	_jsii_.InvokeVoid(
		d,
		"resetThreadPoolSearchQueueSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetThreadPoolSearchSize() {
	_jsii_.InvokeVoid(
		d,
		"resetThreadPoolSearchSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetThreadPoolSearchThrottledQueueSize() {
	_jsii_.InvokeVoid(
		d,
		"resetThreadPoolSearchThrottledQueueSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetThreadPoolSearchThrottledSize() {
	_jsii_.InvokeVoid(
		d,
		"resetThreadPoolSearchThrottledSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetThreadPoolWriteQueueSize() {
	_jsii_.InvokeVoid(
		d,
		"resetThreadPoolWriteQueueSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ResetThreadPoolWriteSize() {
	_jsii_.InvokeVoid(
		d,
		"resetThreadPoolWriteSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseOpensearchConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseOpensearchConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseOpensearchConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

