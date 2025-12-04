// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasemysqlconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/databasemysqlconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/database_mysql_config digitalocean_database_mysql_config}.
type DatabaseMysqlConfig interface {
	cdktf.TerraformResource
	BackupHour() *float64
	SetBackupHour(val *float64)
	BackupHourInput() *float64
	BackupMinute() *float64
	SetBackupMinute(val *float64)
	BackupMinuteInput() *float64
	BinlogRetentionPeriod() *float64
	SetBinlogRetentionPeriod(val *float64)
	BinlogRetentionPeriodInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectTimeout() *float64
	SetConnectTimeout(val *float64)
	ConnectTimeoutInput() *float64
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DefaultTimeZone() *string
	SetDefaultTimeZone(val *string)
	DefaultTimeZoneInput() *string
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
	GroupConcatMaxLen() *float64
	SetGroupConcatMaxLen(val *float64)
	GroupConcatMaxLenInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	InformationSchemaStatsExpiry() *float64
	SetInformationSchemaStatsExpiry(val *float64)
	InformationSchemaStatsExpiryInput() *float64
	InnodbFtMinTokenSize() *float64
	SetInnodbFtMinTokenSize(val *float64)
	InnodbFtMinTokenSizeInput() *float64
	InnodbFtServerStopwordTable() *string
	SetInnodbFtServerStopwordTable(val *string)
	InnodbFtServerStopwordTableInput() *string
	InnodbLockWaitTimeout() *float64
	SetInnodbLockWaitTimeout(val *float64)
	InnodbLockWaitTimeoutInput() *float64
	InnodbLogBufferSize() *float64
	SetInnodbLogBufferSize(val *float64)
	InnodbLogBufferSizeInput() *float64
	InnodbOnlineAlterLogMaxSize() *float64
	SetInnodbOnlineAlterLogMaxSize(val *float64)
	InnodbOnlineAlterLogMaxSizeInput() *float64
	InnodbPrintAllDeadlocks() interface{}
	SetInnodbPrintAllDeadlocks(val interface{})
	InnodbPrintAllDeadlocksInput() interface{}
	InnodbRollbackOnTimeout() interface{}
	SetInnodbRollbackOnTimeout(val interface{})
	InnodbRollbackOnTimeoutInput() interface{}
	InteractiveTimeout() *float64
	SetInteractiveTimeout(val *float64)
	InteractiveTimeoutInput() *float64
	InternalTmpMemStorageEngine() *string
	SetInternalTmpMemStorageEngine(val *string)
	InternalTmpMemStorageEngineInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LongQueryTime() *float64
	SetLongQueryTime(val *float64)
	LongQueryTimeInput() *float64
	MaxAllowedPacket() *float64
	SetMaxAllowedPacket(val *float64)
	MaxAllowedPacketInput() *float64
	MaxHeapTableSize() *float64
	SetMaxHeapTableSize(val *float64)
	MaxHeapTableSizeInput() *float64
	NetReadTimeout() *float64
	SetNetReadTimeout(val *float64)
	NetReadTimeoutInput() *float64
	NetWriteTimeout() *float64
	SetNetWriteTimeout(val *float64)
	NetWriteTimeoutInput() *float64
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
	SlowQueryLog() interface{}
	SetSlowQueryLog(val interface{})
	SlowQueryLogInput() interface{}
	SortBufferSize() *float64
	SetSortBufferSize(val *float64)
	SortBufferSizeInput() *float64
	SqlMode() *string
	SetSqlMode(val *string)
	SqlModeInput() *string
	SqlRequirePrimaryKey() interface{}
	SetSqlRequirePrimaryKey(val interface{})
	SqlRequirePrimaryKeyInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TmpTableSize() *float64
	SetTmpTableSize(val *float64)
	TmpTableSizeInput() *float64
	WaitTimeout() *float64
	SetWaitTimeout(val *float64)
	WaitTimeoutInput() *float64
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
	ResetBackupHour()
	ResetBackupMinute()
	ResetBinlogRetentionPeriod()
	ResetConnectTimeout()
	ResetDefaultTimeZone()
	ResetGroupConcatMaxLen()
	ResetId()
	ResetInformationSchemaStatsExpiry()
	ResetInnodbFtMinTokenSize()
	ResetInnodbFtServerStopwordTable()
	ResetInnodbLockWaitTimeout()
	ResetInnodbLogBufferSize()
	ResetInnodbOnlineAlterLogMaxSize()
	ResetInnodbPrintAllDeadlocks()
	ResetInnodbRollbackOnTimeout()
	ResetInteractiveTimeout()
	ResetInternalTmpMemStorageEngine()
	ResetLongQueryTime()
	ResetMaxAllowedPacket()
	ResetMaxHeapTableSize()
	ResetNetReadTimeout()
	ResetNetWriteTimeout()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSlowQueryLog()
	ResetSortBufferSize()
	ResetSqlMode()
	ResetSqlRequirePrimaryKey()
	ResetTmpTableSize()
	ResetWaitTimeout()
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

// The jsii proxy struct for DatabaseMysqlConfig
type jsiiProxy_DatabaseMysqlConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatabaseMysqlConfig) BackupHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) BackupHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) BackupMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) BackupMinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupMinuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) BinlogRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"binlogRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) BinlogRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"binlogRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) ConnectTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) ConnectTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) DefaultTimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTimeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) DefaultTimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTimeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) GroupConcatMaxLen() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupConcatMaxLen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) GroupConcatMaxLenInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupConcatMaxLenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InformationSchemaStatsExpiry() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"informationSchemaStatsExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InformationSchemaStatsExpiryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"informationSchemaStatsExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InnodbFtMinTokenSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbFtMinTokenSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InnodbFtMinTokenSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbFtMinTokenSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InnodbFtServerStopwordTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"innodbFtServerStopwordTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InnodbFtServerStopwordTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"innodbFtServerStopwordTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InnodbLockWaitTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbLockWaitTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InnodbLockWaitTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbLockWaitTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InnodbLogBufferSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbLogBufferSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InnodbLogBufferSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbLogBufferSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InnodbOnlineAlterLogMaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbOnlineAlterLogMaxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InnodbOnlineAlterLogMaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbOnlineAlterLogMaxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InnodbPrintAllDeadlocks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"innodbPrintAllDeadlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InnodbPrintAllDeadlocksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"innodbPrintAllDeadlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InnodbRollbackOnTimeout() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"innodbRollbackOnTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InnodbRollbackOnTimeoutInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"innodbRollbackOnTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InteractiveTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interactiveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InteractiveTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interactiveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InternalTmpMemStorageEngine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalTmpMemStorageEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) InternalTmpMemStorageEngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalTmpMemStorageEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) LongQueryTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longQueryTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) LongQueryTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longQueryTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) MaxAllowedPacket() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllowedPacket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) MaxAllowedPacketInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllowedPacketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) MaxHeapTableSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHeapTableSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) MaxHeapTableSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHeapTableSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) NetReadTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netReadTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) NetReadTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netReadTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) NetWriteTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netWriteTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) NetWriteTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netWriteTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) SlowQueryLog() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slowQueryLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) SlowQueryLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slowQueryLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) SortBufferSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sortBufferSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) SortBufferSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sortBufferSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) SqlMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) SqlModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) SqlRequirePrimaryKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlRequirePrimaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) SqlRequirePrimaryKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlRequirePrimaryKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) TmpTableSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tmpTableSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) TmpTableSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tmpTableSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) WaitTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMysqlConfig) WaitTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitTimeoutInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/database_mysql_config digitalocean_database_mysql_config} Resource.
func NewDatabaseMysqlConfig(scope constructs.Construct, id *string, config *DatabaseMysqlConfigConfig) DatabaseMysqlConfig {
	_init_.Initialize()

	if err := validateNewDatabaseMysqlConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseMysqlConfig{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.databaseMysqlConfig.DatabaseMysqlConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/database_mysql_config digitalocean_database_mysql_config} Resource.
func NewDatabaseMysqlConfig_Override(d DatabaseMysqlConfig, scope constructs.Construct, id *string, config *DatabaseMysqlConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.databaseMysqlConfig.DatabaseMysqlConfig",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetBackupHour(val *float64) {
	if err := j.validateSetBackupHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupHour",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetBackupMinute(val *float64) {
	if err := j.validateSetBackupMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupMinute",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetBinlogRetentionPeriod(val *float64) {
	if err := j.validateSetBinlogRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binlogRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetConnectTimeout(val *float64) {
	if err := j.validateSetConnectTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectTimeout",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetDefaultTimeZone(val *string) {
	if err := j.validateSetDefaultTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTimeZone",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetGroupConcatMaxLen(val *float64) {
	if err := j.validateSetGroupConcatMaxLenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupConcatMaxLen",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetInformationSchemaStatsExpiry(val *float64) {
	if err := j.validateSetInformationSchemaStatsExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"informationSchemaStatsExpiry",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetInnodbFtMinTokenSize(val *float64) {
	if err := j.validateSetInnodbFtMinTokenSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbFtMinTokenSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetInnodbFtServerStopwordTable(val *string) {
	if err := j.validateSetInnodbFtServerStopwordTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbFtServerStopwordTable",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetInnodbLockWaitTimeout(val *float64) {
	if err := j.validateSetInnodbLockWaitTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbLockWaitTimeout",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetInnodbLogBufferSize(val *float64) {
	if err := j.validateSetInnodbLogBufferSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbLogBufferSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetInnodbOnlineAlterLogMaxSize(val *float64) {
	if err := j.validateSetInnodbOnlineAlterLogMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbOnlineAlterLogMaxSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetInnodbPrintAllDeadlocks(val interface{}) {
	if err := j.validateSetInnodbPrintAllDeadlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbPrintAllDeadlocks",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetInnodbRollbackOnTimeout(val interface{}) {
	if err := j.validateSetInnodbRollbackOnTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbRollbackOnTimeout",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetInteractiveTimeout(val *float64) {
	if err := j.validateSetInteractiveTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interactiveTimeout",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetInternalTmpMemStorageEngine(val *string) {
	if err := j.validateSetInternalTmpMemStorageEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalTmpMemStorageEngine",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetLongQueryTime(val *float64) {
	if err := j.validateSetLongQueryTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longQueryTime",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetMaxAllowedPacket(val *float64) {
	if err := j.validateSetMaxAllowedPacketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAllowedPacket",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetMaxHeapTableSize(val *float64) {
	if err := j.validateSetMaxHeapTableSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxHeapTableSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetNetReadTimeout(val *float64) {
	if err := j.validateSetNetReadTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netReadTimeout",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetNetWriteTimeout(val *float64) {
	if err := j.validateSetNetWriteTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netWriteTimeout",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetSlowQueryLog(val interface{}) {
	if err := j.validateSetSlowQueryLogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slowQueryLog",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetSortBufferSize(val *float64) {
	if err := j.validateSetSortBufferSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortBufferSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetSqlMode(val *string) {
	if err := j.validateSetSqlModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlMode",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetSqlRequirePrimaryKey(val interface{}) {
	if err := j.validateSetSqlRequirePrimaryKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlRequirePrimaryKey",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetTmpTableSize(val *float64) {
	if err := j.validateSetTmpTableSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tmpTableSize",
		val,
	)
}

func (j *jsiiProxy_DatabaseMysqlConfig)SetWaitTimeout(val *float64) {
	if err := j.validateSetWaitTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitTimeout",
		val,
	)
}

// Generates CDKTF code for importing a DatabaseMysqlConfig resource upon running "cdktf plan <stack-name>".
func DatabaseMysqlConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatabaseMysqlConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseMysqlConfig.DatabaseMysqlConfig",
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
func DatabaseMysqlConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseMysqlConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseMysqlConfig.DatabaseMysqlConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseMysqlConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseMysqlConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseMysqlConfig.DatabaseMysqlConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseMysqlConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseMysqlConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databaseMysqlConfig.DatabaseMysqlConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatabaseMysqlConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-digitalocean.databaseMysqlConfig.DatabaseMysqlConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabaseMysqlConfig) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseMysqlConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseMysqlConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseMysqlConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseMysqlConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseMysqlConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseMysqlConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseMysqlConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseMysqlConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseMysqlConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMysqlConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseMysqlConfig) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetBackupHour() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupHour",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetBackupMinute() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupMinute",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetBinlogRetentionPeriod() {
	_jsii_.InvokeVoid(
		d,
		"resetBinlogRetentionPeriod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetConnectTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetDefaultTimeZone() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultTimeZone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetGroupConcatMaxLen() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupConcatMaxLen",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetInformationSchemaStatsExpiry() {
	_jsii_.InvokeVoid(
		d,
		"resetInformationSchemaStatsExpiry",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetInnodbFtMinTokenSize() {
	_jsii_.InvokeVoid(
		d,
		"resetInnodbFtMinTokenSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetInnodbFtServerStopwordTable() {
	_jsii_.InvokeVoid(
		d,
		"resetInnodbFtServerStopwordTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetInnodbLockWaitTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetInnodbLockWaitTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetInnodbLogBufferSize() {
	_jsii_.InvokeVoid(
		d,
		"resetInnodbLogBufferSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetInnodbOnlineAlterLogMaxSize() {
	_jsii_.InvokeVoid(
		d,
		"resetInnodbOnlineAlterLogMaxSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetInnodbPrintAllDeadlocks() {
	_jsii_.InvokeVoid(
		d,
		"resetInnodbPrintAllDeadlocks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetInnodbRollbackOnTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetInnodbRollbackOnTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetInteractiveTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetInteractiveTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetInternalTmpMemStorageEngine() {
	_jsii_.InvokeVoid(
		d,
		"resetInternalTmpMemStorageEngine",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetLongQueryTime() {
	_jsii_.InvokeVoid(
		d,
		"resetLongQueryTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetMaxAllowedPacket() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxAllowedPacket",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetMaxHeapTableSize() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxHeapTableSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetNetReadTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetNetReadTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetNetWriteTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetNetWriteTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetSlowQueryLog() {
	_jsii_.InvokeVoid(
		d,
		"resetSlowQueryLog",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetSortBufferSize() {
	_jsii_.InvokeVoid(
		d,
		"resetSortBufferSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetSqlMode() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetSqlRequirePrimaryKey() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlRequirePrimaryKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetTmpTableSize() {
	_jsii_.InvokeVoid(
		d,
		"resetTmpTableSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) ResetWaitTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetWaitTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMysqlConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMysqlConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMysqlConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMysqlConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMysqlConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMysqlConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

