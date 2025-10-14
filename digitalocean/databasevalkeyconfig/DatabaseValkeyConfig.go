// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasevalkeyconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/databasevalkeyconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/database_valkey_config digitalocean_database_valkey_config}.
type DatabaseValkeyConfig interface {
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
	FrequentSnapshots() interface{}
	SetFrequentSnapshots(val interface{})
	FrequentSnapshotsInput() interface{}
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
	ValkeyActiveExpireEffort() *float64
	SetValkeyActiveExpireEffort(val *float64)
	ValkeyActiveExpireEffortInput() *float64
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
	ResetFrequentSnapshots()
	ResetId()
	ResetIoThreads()
	ResetLfuDecayTime()
	ResetLfuLogFactor()
	ResetNotifyKeyspaceEvents()
	ResetNumberOfDatabases()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPersistence()
	ResetPubsubClientOutputBufferLimit()
	ResetSsl()
	ResetTimeout()
	ResetValkeyActiveExpireEffort()
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

// The jsii proxy struct for DatabaseValkeyConfig
type jsiiProxy_DatabaseValkeyConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatabaseValkeyConfig) AclChannelsDefault() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclChannelsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) AclChannelsDefaultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclChannelsDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) FrequentSnapshots() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"frequentSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) FrequentSnapshotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"frequentSnapshotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) IoThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ioThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) IoThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ioThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) LfuDecayTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lfuDecayTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) LfuDecayTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lfuDecayTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) LfuLogFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lfuLogFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) LfuLogFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lfuLogFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) NotifyKeyspaceEvents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyKeyspaceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) NotifyKeyspaceEventsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyKeyspaceEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) NumberOfDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) NumberOfDatabasesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfDatabasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) Persistence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) PersistenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) PubsubClientOutputBufferLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pubsubClientOutputBufferLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) PubsubClientOutputBufferLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pubsubClientOutputBufferLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) Ssl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) SslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) ValkeyActiveExpireEffort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valkeyActiveExpireEffort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseValkeyConfig) ValkeyActiveExpireEffortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valkeyActiveExpireEffortInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/database_valkey_config digitalocean_database_valkey_config} Resource.
func NewDatabaseValkeyConfig(scope constructs.Construct, id *string, config *DatabaseValkeyConfigConfig) DatabaseValkeyConfig {
	_init_.Initialize()

	if err := validateNewDatabaseValkeyConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseValkeyConfig{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.databaseValkeyConfig.DatabaseValkeyConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/database_valkey_config digitalocean_database_valkey_config} Resource.
func NewDatabaseValkeyConfig_Override(d DatabaseValkeyConfig, scope constructs.Construct, id *string, config *DatabaseValkeyConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.databaseValkeyConfig.DatabaseValkeyConfig",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetAclChannelsDefault(val *string) {
	if err := j.validateSetAclChannelsDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aclChannelsDefault",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetFrequentSnapshots(val interface{}) {
	if err := j.validateSetFrequentSnapshotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequentSnapshots",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetIoThreads(val *float64) {
	if err := j.validateSetIoThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ioThreads",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetLfuDecayTime(val *float64) {
	if err := j.validateSetLfuDecayTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lfuDecayTime",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetLfuLogFactor(val *float64) {
	if err := j.validateSetLfuLogFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lfuLogFactor",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetNotifyKeyspaceEvents(val *string) {
	if err := j.validateSetNotifyKeyspaceEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyKeyspaceEvents",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetNumberOfDatabases(val *float64) {
	if err := j.validateSetNumberOfDatabasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfDatabases",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetPersistence(val *string) {
	if err := j.validateSetPersistenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistence",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetPubsubClientOutputBufferLimit(val *float64) {
	if err := j.validateSetPubsubClientOutputBufferLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pubsubClientOutputBufferLimit",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetSsl(val interface{}) {
	if err := j.validateSetSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssl",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_DatabaseValkeyConfig)SetValkeyActiveExpireEffort(val *float64) {
	if err := j.validateSetValkeyActiveExpireEffortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valkeyActiveExpireEffort",
		val,
	)
}

// Generates CDKTF code for importing a DatabaseValkeyConfig resource upon running "cdktf plan <stack-name>".
func DatabaseValkeyConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatabaseValkeyConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseValkeyConfig.DatabaseValkeyConfig",
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
func DatabaseValkeyConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseValkeyConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseValkeyConfig.DatabaseValkeyConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseValkeyConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseValkeyConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseValkeyConfig.DatabaseValkeyConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseValkeyConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseValkeyConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseValkeyConfig.DatabaseValkeyConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatabaseValkeyConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-digitalocean.databaseValkeyConfig.DatabaseValkeyConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabaseValkeyConfig) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseValkeyConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseValkeyConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseValkeyConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseValkeyConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseValkeyConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseValkeyConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseValkeyConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseValkeyConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseValkeyConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseValkeyConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseValkeyConfig) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) ResetAclChannelsDefault() {
	_jsii_.InvokeVoid(
		d,
		"resetAclChannelsDefault",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) ResetFrequentSnapshots() {
	_jsii_.InvokeVoid(
		d,
		"resetFrequentSnapshots",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) ResetIoThreads() {
	_jsii_.InvokeVoid(
		d,
		"resetIoThreads",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) ResetLfuDecayTime() {
	_jsii_.InvokeVoid(
		d,
		"resetLfuDecayTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) ResetLfuLogFactor() {
	_jsii_.InvokeVoid(
		d,
		"resetLfuLogFactor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) ResetNotifyKeyspaceEvents() {
	_jsii_.InvokeVoid(
		d,
		"resetNotifyKeyspaceEvents",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) ResetNumberOfDatabases() {
	_jsii_.InvokeVoid(
		d,
		"resetNumberOfDatabases",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) ResetPersistence() {
	_jsii_.InvokeVoid(
		d,
		"resetPersistence",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) ResetPubsubClientOutputBufferLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetPubsubClientOutputBufferLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) ResetSsl() {
	_jsii_.InvokeVoid(
		d,
		"resetSsl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) ResetTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) ResetValkeyActiveExpireEffort() {
	_jsii_.InvokeVoid(
		d,
		"resetValkeyActiveExpireEffort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseValkeyConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseValkeyConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseValkeyConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseValkeyConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseValkeyConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseValkeyConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

