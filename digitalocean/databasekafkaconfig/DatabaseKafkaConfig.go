// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasekafkaconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/databasekafkaconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/database_kafka_config digitalocean_database_kafka_config}.
type DatabaseKafkaConfig interface {
	cdktf.TerraformResource
	AutoCreateTopicsEnable() interface{}
	SetAutoCreateTopicsEnable(val interface{})
	AutoCreateTopicsEnableInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupInitialRebalanceDelayMs() *float64
	SetGroupInitialRebalanceDelayMs(val *float64)
	GroupInitialRebalanceDelayMsInput() *float64
	GroupMaxSessionTimeoutMs() *float64
	SetGroupMaxSessionTimeoutMs(val *float64)
	GroupMaxSessionTimeoutMsInput() *float64
	GroupMinSessionTimeoutMs() *float64
	SetGroupMinSessionTimeoutMs(val *float64)
	GroupMinSessionTimeoutMsInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogCleanerDeleteRetentionMs() *float64
	SetLogCleanerDeleteRetentionMs(val *float64)
	LogCleanerDeleteRetentionMsInput() *float64
	LogCleanerMinCompactionLagMs() *string
	SetLogCleanerMinCompactionLagMs(val *string)
	LogCleanerMinCompactionLagMsInput() *string
	LogFlushIntervalMs() *string
	SetLogFlushIntervalMs(val *string)
	LogFlushIntervalMsInput() *string
	LogIndexIntervalBytes() *float64
	SetLogIndexIntervalBytes(val *float64)
	LogIndexIntervalBytesInput() *float64
	LogMessageDownconversionEnable() interface{}
	SetLogMessageDownconversionEnable(val interface{})
	LogMessageDownconversionEnableInput() interface{}
	LogMessageTimestampDifferenceMaxMs() *string
	SetLogMessageTimestampDifferenceMaxMs(val *string)
	LogMessageTimestampDifferenceMaxMsInput() *string
	LogPreallocate() interface{}
	SetLogPreallocate(val interface{})
	LogPreallocateInput() interface{}
	LogRetentionBytes() *string
	SetLogRetentionBytes(val *string)
	LogRetentionBytesInput() *string
	LogRetentionHours() *float64
	SetLogRetentionHours(val *float64)
	LogRetentionHoursInput() *float64
	LogRetentionMs() *string
	SetLogRetentionMs(val *string)
	LogRetentionMsInput() *string
	LogRollJitterMs() *string
	SetLogRollJitterMs(val *string)
	LogRollJitterMsInput() *string
	LogSegmentDeleteDelayMs() *float64
	SetLogSegmentDeleteDelayMs(val *float64)
	LogSegmentDeleteDelayMsInput() *float64
	MessageMaxBytes() *float64
	SetMessageMaxBytes(val *float64)
	MessageMaxBytesInput() *float64
	// The tree node.
	Node() constructs.Node
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetAutoCreateTopicsEnable()
	ResetGroupInitialRebalanceDelayMs()
	ResetGroupMaxSessionTimeoutMs()
	ResetGroupMinSessionTimeoutMs()
	ResetId()
	ResetLogCleanerDeleteRetentionMs()
	ResetLogCleanerMinCompactionLagMs()
	ResetLogFlushIntervalMs()
	ResetLogIndexIntervalBytes()
	ResetLogMessageDownconversionEnable()
	ResetLogMessageTimestampDifferenceMaxMs()
	ResetLogPreallocate()
	ResetLogRetentionBytes()
	ResetLogRetentionHours()
	ResetLogRetentionMs()
	ResetLogRollJitterMs()
	ResetLogSegmentDeleteDelayMs()
	ResetMessageMaxBytes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for DatabaseKafkaConfig
type jsiiProxy_DatabaseKafkaConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatabaseKafkaConfig) AutoCreateTopicsEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoCreateTopicsEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) AutoCreateTopicsEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoCreateTopicsEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) GroupInitialRebalanceDelayMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupInitialRebalanceDelayMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) GroupInitialRebalanceDelayMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupInitialRebalanceDelayMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) GroupMaxSessionTimeoutMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupMaxSessionTimeoutMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) GroupMaxSessionTimeoutMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupMaxSessionTimeoutMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) GroupMinSessionTimeoutMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupMinSessionTimeoutMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) GroupMinSessionTimeoutMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupMinSessionTimeoutMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogCleanerDeleteRetentionMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logCleanerDeleteRetentionMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogCleanerDeleteRetentionMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logCleanerDeleteRetentionMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogCleanerMinCompactionLagMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logCleanerMinCompactionLagMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogCleanerMinCompactionLagMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logCleanerMinCompactionLagMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogFlushIntervalMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFlushIntervalMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogFlushIntervalMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFlushIntervalMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogIndexIntervalBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logIndexIntervalBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogIndexIntervalBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logIndexIntervalBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogMessageDownconversionEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logMessageDownconversionEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogMessageDownconversionEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logMessageDownconversionEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogMessageTimestampDifferenceMaxMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logMessageTimestampDifferenceMaxMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogMessageTimestampDifferenceMaxMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logMessageTimestampDifferenceMaxMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogPreallocate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logPreallocate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogPreallocateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logPreallocateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogRetentionBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogRetentionBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogRetentionHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logRetentionHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogRetentionHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logRetentionHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogRetentionMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogRetentionMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRetentionMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogRollJitterMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRollJitterMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogRollJitterMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logRollJitterMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogSegmentDeleteDelayMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logSegmentDeleteDelayMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) LogSegmentDeleteDelayMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logSegmentDeleteDelayMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) MessageMaxBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageMaxBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) MessageMaxBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageMaxBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/database_kafka_config digitalocean_database_kafka_config} Resource.
func NewDatabaseKafkaConfig(scope constructs.Construct, id *string, config *DatabaseKafkaConfigConfig) DatabaseKafkaConfig {
	_init_.Initialize()

	if err := validateNewDatabaseKafkaConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseKafkaConfig{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.databaseKafkaConfig.DatabaseKafkaConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/database_kafka_config digitalocean_database_kafka_config} Resource.
func NewDatabaseKafkaConfig_Override(d DatabaseKafkaConfig, scope constructs.Construct, id *string, config *DatabaseKafkaConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.databaseKafkaConfig.DatabaseKafkaConfig",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetAutoCreateTopicsEnable(val interface{}) {
	if err := j.validateSetAutoCreateTopicsEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoCreateTopicsEnable",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetGroupInitialRebalanceDelayMs(val *float64) {
	if err := j.validateSetGroupInitialRebalanceDelayMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupInitialRebalanceDelayMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetGroupMaxSessionTimeoutMs(val *float64) {
	if err := j.validateSetGroupMaxSessionTimeoutMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupMaxSessionTimeoutMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetGroupMinSessionTimeoutMs(val *float64) {
	if err := j.validateSetGroupMinSessionTimeoutMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupMinSessionTimeoutMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetLogCleanerDeleteRetentionMs(val *float64) {
	if err := j.validateSetLogCleanerDeleteRetentionMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logCleanerDeleteRetentionMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetLogCleanerMinCompactionLagMs(val *string) {
	if err := j.validateSetLogCleanerMinCompactionLagMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logCleanerMinCompactionLagMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetLogFlushIntervalMs(val *string) {
	if err := j.validateSetLogFlushIntervalMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logFlushIntervalMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetLogIndexIntervalBytes(val *float64) {
	if err := j.validateSetLogIndexIntervalBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logIndexIntervalBytes",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetLogMessageDownconversionEnable(val interface{}) {
	if err := j.validateSetLogMessageDownconversionEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logMessageDownconversionEnable",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetLogMessageTimestampDifferenceMaxMs(val *string) {
	if err := j.validateSetLogMessageTimestampDifferenceMaxMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logMessageTimestampDifferenceMaxMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetLogPreallocate(val interface{}) {
	if err := j.validateSetLogPreallocateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logPreallocate",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetLogRetentionBytes(val *string) {
	if err := j.validateSetLogRetentionBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logRetentionBytes",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetLogRetentionHours(val *float64) {
	if err := j.validateSetLogRetentionHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logRetentionHours",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetLogRetentionMs(val *string) {
	if err := j.validateSetLogRetentionMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logRetentionMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetLogRollJitterMs(val *string) {
	if err := j.validateSetLogRollJitterMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logRollJitterMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetLogSegmentDeleteDelayMs(val *float64) {
	if err := j.validateSetLogSegmentDeleteDelayMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logSegmentDeleteDelayMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetMessageMaxBytes(val *float64) {
	if err := j.validateSetMessageMaxBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageMaxBytes",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a DatabaseKafkaConfig resource upon running "cdktf plan <stack-name>".
func DatabaseKafkaConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatabaseKafkaConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseKafkaConfig.DatabaseKafkaConfig",
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
func DatabaseKafkaConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseKafkaConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseKafkaConfig.DatabaseKafkaConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseKafkaConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseKafkaConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseKafkaConfig.DatabaseKafkaConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseKafkaConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseKafkaConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseKafkaConfig.DatabaseKafkaConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatabaseKafkaConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-digitalocean.databaseKafkaConfig.DatabaseKafkaConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabaseKafkaConfig) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseKafkaConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseKafkaConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseKafkaConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseKafkaConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseKafkaConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseKafkaConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseKafkaConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseKafkaConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseKafkaConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseKafkaConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseKafkaConfig) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetAutoCreateTopicsEnable() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoCreateTopicsEnable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetGroupInitialRebalanceDelayMs() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupInitialRebalanceDelayMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetGroupMaxSessionTimeoutMs() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupMaxSessionTimeoutMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetGroupMinSessionTimeoutMs() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupMinSessionTimeoutMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetLogCleanerDeleteRetentionMs() {
	_jsii_.InvokeVoid(
		d,
		"resetLogCleanerDeleteRetentionMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetLogCleanerMinCompactionLagMs() {
	_jsii_.InvokeVoid(
		d,
		"resetLogCleanerMinCompactionLagMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetLogFlushIntervalMs() {
	_jsii_.InvokeVoid(
		d,
		"resetLogFlushIntervalMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetLogIndexIntervalBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetLogIndexIntervalBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetLogMessageDownconversionEnable() {
	_jsii_.InvokeVoid(
		d,
		"resetLogMessageDownconversionEnable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetLogMessageTimestampDifferenceMaxMs() {
	_jsii_.InvokeVoid(
		d,
		"resetLogMessageTimestampDifferenceMaxMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetLogPreallocate() {
	_jsii_.InvokeVoid(
		d,
		"resetLogPreallocate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetLogRetentionBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetLogRetentionBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetLogRetentionHours() {
	_jsii_.InvokeVoid(
		d,
		"resetLogRetentionHours",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetLogRetentionMs() {
	_jsii_.InvokeVoid(
		d,
		"resetLogRetentionMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetLogRollJitterMs() {
	_jsii_.InvokeVoid(
		d,
		"resetLogRollJitterMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetLogSegmentDeleteDelayMs() {
	_jsii_.InvokeVoid(
		d,
		"resetLogSegmentDeleteDelayMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetMessageMaxBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetMessageMaxBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseKafkaConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseKafkaConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseKafkaConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseKafkaConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseKafkaConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

