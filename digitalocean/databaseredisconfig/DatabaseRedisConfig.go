// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseredisconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v11/databaseredisconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/database_redis_config digitalocean_database_redis_config}.
type DatabaseRedisConfig interface {
	cdktf.TerraformResource
	AclChannelsDefault() *string
	SetAclChannelsDefault(val *string)
	AclChannelsDefaultInput() *string
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	IoThreads() *float64
	SetIoThreads(val *float64)
	IoThreadsInput() *float64
	LfuDecayTime() *float64
	SetLfuDecayTime(val *float64)
	LfuDecayTimeInput() *float64
	LfuLogFactor() *float64
	SetLfuLogFactor(val *float64)
	LfuLogFactorInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxmemoryPolicy() *string
	SetMaxmemoryPolicy(val *string)
	MaxmemoryPolicyInput() *string
	// The tree node.
	Node() constructs.Node
	NotifyKeyspaceEvents() *string
	SetNotifyKeyspaceEvents(val *string)
	NotifyKeyspaceEventsInput() *string
	NumberOfDatabases() *float64
	SetNumberOfDatabases(val *float64)
	NumberOfDatabasesInput() *float64
	Persistence() *string
	SetPersistence(val *string)
	PersistenceInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PubsubClientOutputBufferLimit() *float64
	SetPubsubClientOutputBufferLimit(val *float64)
	PubsubClientOutputBufferLimitInput() *float64
	// Experimental.
	RawOverrides() interface{}
	Ssl() interface{}
	SetSsl(val interface{})
	SslInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
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
	ResetAclChannelsDefault()
	ResetId()
	ResetIoThreads()
	ResetLfuDecayTime()
	ResetLfuLogFactor()
	ResetMaxmemoryPolicy()
	ResetNotifyKeyspaceEvents()
	ResetNumberOfDatabases()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPersistence()
	ResetPubsubClientOutputBufferLimit()
	ResetSsl()
	ResetTimeout()
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

// The jsii proxy struct for DatabaseRedisConfig
type jsiiProxy_DatabaseRedisConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatabaseRedisConfig) AclChannelsDefault() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclChannelsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) AclChannelsDefaultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclChannelsDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) IoThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ioThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) IoThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ioThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) LfuDecayTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lfuDecayTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) LfuDecayTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lfuDecayTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) LfuLogFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lfuLogFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) LfuLogFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lfuLogFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) MaxmemoryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxmemoryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) MaxmemoryPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxmemoryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) NotifyKeyspaceEvents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyKeyspaceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) NotifyKeyspaceEventsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyKeyspaceEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) NumberOfDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) NumberOfDatabasesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfDatabasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) Persistence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) PersistenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) PubsubClientOutputBufferLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pubsubClientOutputBufferLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) PubsubClientOutputBufferLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pubsubClientOutputBufferLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) Ssl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) SslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseRedisConfig) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/database_redis_config digitalocean_database_redis_config} Resource.
func NewDatabaseRedisConfig(scope constructs.Construct, id *string, config *DatabaseRedisConfigConfig) DatabaseRedisConfig {
	_init_.Initialize()

	if err := validateNewDatabaseRedisConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseRedisConfig{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.databaseRedisConfig.DatabaseRedisConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/database_redis_config digitalocean_database_redis_config} Resource.
func NewDatabaseRedisConfig_Override(d DatabaseRedisConfig, scope constructs.Construct, id *string, config *DatabaseRedisConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.databaseRedisConfig.DatabaseRedisConfig",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetAclChannelsDefault(val *string) {
	if err := j.validateSetAclChannelsDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aclChannelsDefault",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetIoThreads(val *float64) {
	if err := j.validateSetIoThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ioThreads",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetLfuDecayTime(val *float64) {
	if err := j.validateSetLfuDecayTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lfuDecayTime",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetLfuLogFactor(val *float64) {
	if err := j.validateSetLfuLogFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lfuLogFactor",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetMaxmemoryPolicy(val *string) {
	if err := j.validateSetMaxmemoryPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxmemoryPolicy",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetNotifyKeyspaceEvents(val *string) {
	if err := j.validateSetNotifyKeyspaceEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyKeyspaceEvents",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetNumberOfDatabases(val *float64) {
	if err := j.validateSetNumberOfDatabasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfDatabases",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetPersistence(val *string) {
	if err := j.validateSetPersistenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistence",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetPubsubClientOutputBufferLimit(val *float64) {
	if err := j.validateSetPubsubClientOutputBufferLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pubsubClientOutputBufferLimit",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetSsl(val interface{}) {
	if err := j.validateSetSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssl",
		val,
	)
}

func (j *jsiiProxy_DatabaseRedisConfig)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

// Generates CDKTF code for importing a DatabaseRedisConfig resource upon running "cdktf plan <stack-name>".
func DatabaseRedisConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatabaseRedisConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseRedisConfig.DatabaseRedisConfig",
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
func DatabaseRedisConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseRedisConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseRedisConfig.DatabaseRedisConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseRedisConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseRedisConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseRedisConfig.DatabaseRedisConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseRedisConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseRedisConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseRedisConfig.DatabaseRedisConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatabaseRedisConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-digitalocean.databaseRedisConfig.DatabaseRedisConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabaseRedisConfig) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseRedisConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseRedisConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseRedisConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseRedisConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseRedisConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseRedisConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseRedisConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseRedisConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseRedisConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseRedisConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseRedisConfig) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) ResetAclChannelsDefault() {
	_jsii_.InvokeVoid(
		d,
		"resetAclChannelsDefault",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) ResetIoThreads() {
	_jsii_.InvokeVoid(
		d,
		"resetIoThreads",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) ResetLfuDecayTime() {
	_jsii_.InvokeVoid(
		d,
		"resetLfuDecayTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) ResetLfuLogFactor() {
	_jsii_.InvokeVoid(
		d,
		"resetLfuLogFactor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) ResetMaxmemoryPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxmemoryPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) ResetNotifyKeyspaceEvents() {
	_jsii_.InvokeVoid(
		d,
		"resetNotifyKeyspaceEvents",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) ResetNumberOfDatabases() {
	_jsii_.InvokeVoid(
		d,
		"resetNumberOfDatabases",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) ResetPersistence() {
	_jsii_.InvokeVoid(
		d,
		"resetPersistence",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) ResetPubsubClientOutputBufferLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetPubsubClientOutputBufferLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) ResetSsl() {
	_jsii_.InvokeVoid(
		d,
		"resetSsl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) ResetTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseRedisConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseRedisConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseRedisConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseRedisConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseRedisConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseRedisConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

