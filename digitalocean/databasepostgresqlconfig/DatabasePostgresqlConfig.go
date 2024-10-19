// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasepostgresqlconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v11/databasepostgresqlconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.43.0/docs/resources/database_postgresql_config digitalocean_database_postgresql_config}.
type DatabasePostgresqlConfig interface {
	cdktf.TerraformResource
	AutovacuumAnalyzeScaleFactor() *float64
	SetAutovacuumAnalyzeScaleFactor(val *float64)
	AutovacuumAnalyzeScaleFactorInput() *float64
	AutovacuumAnalyzeThreshold() *float64
	SetAutovacuumAnalyzeThreshold(val *float64)
	AutovacuumAnalyzeThresholdInput() *float64
	AutovacuumFreezeMaxAge() *float64
	SetAutovacuumFreezeMaxAge(val *float64)
	AutovacuumFreezeMaxAgeInput() *float64
	AutovacuumMaxWorkers() *float64
	SetAutovacuumMaxWorkers(val *float64)
	AutovacuumMaxWorkersInput() *float64
	AutovacuumNaptime() *float64
	SetAutovacuumNaptime(val *float64)
	AutovacuumNaptimeInput() *float64
	AutovacuumVacuumCostDelay() *float64
	SetAutovacuumVacuumCostDelay(val *float64)
	AutovacuumVacuumCostDelayInput() *float64
	AutovacuumVacuumCostLimit() *float64
	SetAutovacuumVacuumCostLimit(val *float64)
	AutovacuumVacuumCostLimitInput() *float64
	AutovacuumVacuumScaleFactor() *float64
	SetAutovacuumVacuumScaleFactor(val *float64)
	AutovacuumVacuumScaleFactorInput() *float64
	AutovacuumVacuumThreshold() *float64
	SetAutovacuumVacuumThreshold(val *float64)
	AutovacuumVacuumThresholdInput() *float64
	BackupHour() *float64
	SetBackupHour(val *float64)
	BackupHourInput() *float64
	BackupMinute() *float64
	SetBackupMinute(val *float64)
	BackupMinuteInput() *float64
	BgwriterDelay() *float64
	SetBgwriterDelay(val *float64)
	BgwriterDelayInput() *float64
	BgwriterFlushAfter() *float64
	SetBgwriterFlushAfter(val *float64)
	BgwriterFlushAfterInput() *float64
	BgwriterLruMaxpages() *float64
	SetBgwriterLruMaxpages(val *float64)
	BgwriterLruMaxpagesInput() *float64
	BgwriterLruMultiplier() *float64
	SetBgwriterLruMultiplier(val *float64)
	BgwriterLruMultiplierInput() *float64
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
	DeadlockTimeout() *float64
	SetDeadlockTimeout(val *float64)
	DeadlockTimeoutInput() *float64
	DefaultToastCompression() *string
	SetDefaultToastCompression(val *string)
	DefaultToastCompressionInput() *string
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
	IdleInTransactionSessionTimeout() *float64
	SetIdleInTransactionSessionTimeout(val *float64)
	IdleInTransactionSessionTimeoutInput() *float64
	Jit() interface{}
	SetJit(val interface{})
	JitInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogAutovacuumMinDuration() *float64
	SetLogAutovacuumMinDuration(val *float64)
	LogAutovacuumMinDurationInput() *float64
	LogErrorVerbosity() *string
	SetLogErrorVerbosity(val *string)
	LogErrorVerbosityInput() *string
	LogLinePrefix() *string
	SetLogLinePrefix(val *string)
	LogLinePrefixInput() *string
	LogMinDurationStatement() *float64
	SetLogMinDurationStatement(val *float64)
	LogMinDurationStatementInput() *float64
	MaxFilesPerProcess() *float64
	SetMaxFilesPerProcess(val *float64)
	MaxFilesPerProcessInput() *float64
	MaxLocksPerTransaction() *float64
	SetMaxLocksPerTransaction(val *float64)
	MaxLocksPerTransactionInput() *float64
	MaxLogicalReplicationWorkers() *float64
	SetMaxLogicalReplicationWorkers(val *float64)
	MaxLogicalReplicationWorkersInput() *float64
	MaxParallelWorkers() *float64
	SetMaxParallelWorkers(val *float64)
	MaxParallelWorkersInput() *float64
	MaxParallelWorkersPerGather() *float64
	SetMaxParallelWorkersPerGather(val *float64)
	MaxParallelWorkersPerGatherInput() *float64
	MaxPredLocksPerTransaction() *float64
	SetMaxPredLocksPerTransaction(val *float64)
	MaxPredLocksPerTransactionInput() *float64
	MaxPreparedTransactions() *float64
	SetMaxPreparedTransactions(val *float64)
	MaxPreparedTransactionsInput() *float64
	MaxReplicationSlots() *float64
	SetMaxReplicationSlots(val *float64)
	MaxReplicationSlotsInput() *float64
	MaxStackDepth() *float64
	SetMaxStackDepth(val *float64)
	MaxStackDepthInput() *float64
	MaxStandbyArchiveDelay() *float64
	SetMaxStandbyArchiveDelay(val *float64)
	MaxStandbyArchiveDelayInput() *float64
	MaxStandbyStreamingDelay() *float64
	SetMaxStandbyStreamingDelay(val *float64)
	MaxStandbyStreamingDelayInput() *float64
	MaxWalSenders() *float64
	SetMaxWalSenders(val *float64)
	MaxWalSendersInput() *float64
	MaxWorkerProcesses() *float64
	SetMaxWorkerProcesses(val *float64)
	MaxWorkerProcessesInput() *float64
	// The tree node.
	Node() constructs.Node
	Pgbouncer() DatabasePostgresqlConfigPgbouncerList
	PgbouncerInput() interface{}
	PgPartmanBgwInterval() *float64
	SetPgPartmanBgwInterval(val *float64)
	PgPartmanBgwIntervalInput() *float64
	PgPartmanBgwRole() *string
	SetPgPartmanBgwRole(val *string)
	PgPartmanBgwRoleInput() *string
	PgStatStatementsTrack() *string
	SetPgStatStatementsTrack(val *string)
	PgStatStatementsTrackInput() *string
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
	SharedBuffersPercentage() *float64
	SetSharedBuffersPercentage(val *float64)
	SharedBuffersPercentageInput() *float64
	TempFileLimit() *float64
	SetTempFileLimit(val *float64)
	TempFileLimitInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timescaledb() DatabasePostgresqlConfigTimescaledbOutputReference
	TimescaledbInput() *DatabasePostgresqlConfigTimescaledb
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	TrackActivityQuerySize() *float64
	SetTrackActivityQuerySize(val *float64)
	TrackActivityQuerySizeInput() *float64
	TrackCommitTimestamp() *string
	SetTrackCommitTimestamp(val *string)
	TrackCommitTimestampInput() *string
	TrackFunctions() *string
	SetTrackFunctions(val *string)
	TrackFunctionsInput() *string
	TrackIoTiming() *string
	SetTrackIoTiming(val *string)
	TrackIoTimingInput() *string
	WalSenderTimeout() *float64
	SetWalSenderTimeout(val *float64)
	WalSenderTimeoutInput() *float64
	WalWriterDelay() *float64
	SetWalWriterDelay(val *float64)
	WalWriterDelayInput() *float64
	WorkMem() *float64
	SetWorkMem(val *float64)
	WorkMemInput() *float64
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
	PutPgbouncer(value interface{})
	PutTimescaledb(value *DatabasePostgresqlConfigTimescaledb)
	ResetAutovacuumAnalyzeScaleFactor()
	ResetAutovacuumAnalyzeThreshold()
	ResetAutovacuumFreezeMaxAge()
	ResetAutovacuumMaxWorkers()
	ResetAutovacuumNaptime()
	ResetAutovacuumVacuumCostDelay()
	ResetAutovacuumVacuumCostLimit()
	ResetAutovacuumVacuumScaleFactor()
	ResetAutovacuumVacuumThreshold()
	ResetBackupHour()
	ResetBackupMinute()
	ResetBgwriterDelay()
	ResetBgwriterFlushAfter()
	ResetBgwriterLruMaxpages()
	ResetBgwriterLruMultiplier()
	ResetDeadlockTimeout()
	ResetDefaultToastCompression()
	ResetId()
	ResetIdleInTransactionSessionTimeout()
	ResetJit()
	ResetLogAutovacuumMinDuration()
	ResetLogErrorVerbosity()
	ResetLogLinePrefix()
	ResetLogMinDurationStatement()
	ResetMaxFilesPerProcess()
	ResetMaxLocksPerTransaction()
	ResetMaxLogicalReplicationWorkers()
	ResetMaxParallelWorkers()
	ResetMaxParallelWorkersPerGather()
	ResetMaxPredLocksPerTransaction()
	ResetMaxPreparedTransactions()
	ResetMaxReplicationSlots()
	ResetMaxStackDepth()
	ResetMaxStandbyArchiveDelay()
	ResetMaxStandbyStreamingDelay()
	ResetMaxWalSenders()
	ResetMaxWorkerProcesses()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPgbouncer()
	ResetPgPartmanBgwInterval()
	ResetPgPartmanBgwRole()
	ResetPgStatStatementsTrack()
	ResetSharedBuffersPercentage()
	ResetTempFileLimit()
	ResetTimescaledb()
	ResetTimezone()
	ResetTrackActivityQuerySize()
	ResetTrackCommitTimestamp()
	ResetTrackFunctions()
	ResetTrackIoTiming()
	ResetWalSenderTimeout()
	ResetWalWriterDelay()
	ResetWorkMem()
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

// The jsii proxy struct for DatabasePostgresqlConfig
type jsiiProxy_DatabasePostgresqlConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumAnalyzeScaleFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumAnalyzeScaleFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumAnalyzeScaleFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumAnalyzeScaleFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumAnalyzeThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumAnalyzeThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumAnalyzeThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumAnalyzeThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumFreezeMaxAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumFreezeMaxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumFreezeMaxAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumFreezeMaxAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumMaxWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumMaxWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumMaxWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumMaxWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumNaptime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumNaptime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumNaptimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumNaptimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumVacuumCostDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumVacuumCostDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumVacuumCostDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumVacuumCostDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumVacuumCostLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumVacuumCostLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumVacuumCostLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumVacuumCostLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumVacuumScaleFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumVacuumScaleFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumVacuumScaleFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumVacuumScaleFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumVacuumThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumVacuumThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) AutovacuumVacuumThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumVacuumThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) BackupHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) BackupHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) BackupMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) BackupMinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupMinuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) BgwriterDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgwriterDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) BgwriterDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgwriterDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) BgwriterFlushAfter() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgwriterFlushAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) BgwriterFlushAfterInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgwriterFlushAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) BgwriterLruMaxpages() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgwriterLruMaxpages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) BgwriterLruMaxpagesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgwriterLruMaxpagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) BgwriterLruMultiplier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgwriterLruMultiplier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) BgwriterLruMultiplierInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgwriterLruMultiplierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) DeadlockTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deadlockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) DeadlockTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deadlockTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) DefaultToastCompression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultToastCompression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) DefaultToastCompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultToastCompressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) IdleInTransactionSessionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleInTransactionSessionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) IdleInTransactionSessionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleInTransactionSessionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) Jit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) JitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) LogAutovacuumMinDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logAutovacuumMinDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) LogAutovacuumMinDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logAutovacuumMinDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) LogErrorVerbosity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logErrorVerbosity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) LogErrorVerbosityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logErrorVerbosityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) LogLinePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLinePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) LogLinePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLinePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) LogMinDurationStatement() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logMinDurationStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) LogMinDurationStatementInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logMinDurationStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxFilesPerProcess() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFilesPerProcess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxFilesPerProcessInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFilesPerProcessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxLocksPerTransaction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLocksPerTransaction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxLocksPerTransactionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLocksPerTransactionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxLogicalReplicationWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLogicalReplicationWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxLogicalReplicationWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLogicalReplicationWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxParallelWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxParallelWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxParallelWorkersPerGather() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelWorkersPerGather",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxParallelWorkersPerGatherInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelWorkersPerGatherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxPredLocksPerTransaction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPredLocksPerTransaction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxPredLocksPerTransactionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPredLocksPerTransactionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxPreparedTransactions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPreparedTransactions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxPreparedTransactionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPreparedTransactionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxReplicationSlots() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicationSlots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxReplicationSlotsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicationSlotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxStackDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStackDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxStackDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStackDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxStandbyArchiveDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStandbyArchiveDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxStandbyArchiveDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStandbyArchiveDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxStandbyStreamingDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStandbyStreamingDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxStandbyStreamingDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStandbyStreamingDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxWalSenders() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWalSenders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxWalSendersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWalSendersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxWorkerProcesses() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkerProcesses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) MaxWorkerProcessesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkerProcessesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) Pgbouncer() DatabasePostgresqlConfigPgbouncerList {
	var returns DatabasePostgresqlConfigPgbouncerList
	_jsii_.Get(
		j,
		"pgbouncer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) PgbouncerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pgbouncerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) PgPartmanBgwInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pgPartmanBgwInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) PgPartmanBgwIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pgPartmanBgwIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) PgPartmanBgwRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pgPartmanBgwRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) PgPartmanBgwRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pgPartmanBgwRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) PgStatStatementsTrack() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pgStatStatementsTrack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) PgStatStatementsTrackInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pgStatStatementsTrackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) SharedBuffersPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharedBuffersPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) SharedBuffersPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharedBuffersPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) TempFileLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tempFileLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) TempFileLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tempFileLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) Timescaledb() DatabasePostgresqlConfigTimescaledbOutputReference {
	var returns DatabasePostgresqlConfigTimescaledbOutputReference
	_jsii_.Get(
		j,
		"timescaledb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) TimescaledbInput() *DatabasePostgresqlConfigTimescaledb {
	var returns *DatabasePostgresqlConfigTimescaledb
	_jsii_.Get(
		j,
		"timescaledbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) TrackActivityQuerySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trackActivityQuerySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) TrackActivityQuerySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trackActivityQuerySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) TrackCommitTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackCommitTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) TrackCommitTimestampInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackCommitTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) TrackFunctions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) TrackFunctionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) TrackIoTiming() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackIoTiming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) TrackIoTimingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackIoTimingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) WalSenderTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"walSenderTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) WalSenderTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"walSenderTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) WalWriterDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"walWriterDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) WalWriterDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"walWriterDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) WorkMem() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workMem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfig) WorkMemInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workMemInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.43.0/docs/resources/database_postgresql_config digitalocean_database_postgresql_config} Resource.
func NewDatabasePostgresqlConfig(scope constructs.Construct, id *string, config *DatabasePostgresqlConfigConfig) DatabasePostgresqlConfig {
	_init_.Initialize()

	if err := validateNewDatabasePostgresqlConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabasePostgresqlConfig{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.databasePostgresqlConfig.DatabasePostgresqlConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.43.0/docs/resources/database_postgresql_config digitalocean_database_postgresql_config} Resource.
func NewDatabasePostgresqlConfig_Override(d DatabasePostgresqlConfig, scope constructs.Construct, id *string, config *DatabasePostgresqlConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.databasePostgresqlConfig.DatabasePostgresqlConfig",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetAutovacuumAnalyzeScaleFactor(val *float64) {
	if err := j.validateSetAutovacuumAnalyzeScaleFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumAnalyzeScaleFactor",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetAutovacuumAnalyzeThreshold(val *float64) {
	if err := j.validateSetAutovacuumAnalyzeThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumAnalyzeThreshold",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetAutovacuumFreezeMaxAge(val *float64) {
	if err := j.validateSetAutovacuumFreezeMaxAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumFreezeMaxAge",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetAutovacuumMaxWorkers(val *float64) {
	if err := j.validateSetAutovacuumMaxWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumMaxWorkers",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetAutovacuumNaptime(val *float64) {
	if err := j.validateSetAutovacuumNaptimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumNaptime",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetAutovacuumVacuumCostDelay(val *float64) {
	if err := j.validateSetAutovacuumVacuumCostDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumVacuumCostDelay",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetAutovacuumVacuumCostLimit(val *float64) {
	if err := j.validateSetAutovacuumVacuumCostLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumVacuumCostLimit",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetAutovacuumVacuumScaleFactor(val *float64) {
	if err := j.validateSetAutovacuumVacuumScaleFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumVacuumScaleFactor",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetAutovacuumVacuumThreshold(val *float64) {
	if err := j.validateSetAutovacuumVacuumThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumVacuumThreshold",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetBackupHour(val *float64) {
	if err := j.validateSetBackupHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupHour",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetBackupMinute(val *float64) {
	if err := j.validateSetBackupMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupMinute",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetBgwriterDelay(val *float64) {
	if err := j.validateSetBgwriterDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgwriterDelay",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetBgwriterFlushAfter(val *float64) {
	if err := j.validateSetBgwriterFlushAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgwriterFlushAfter",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetBgwriterLruMaxpages(val *float64) {
	if err := j.validateSetBgwriterLruMaxpagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgwriterLruMaxpages",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetBgwriterLruMultiplier(val *float64) {
	if err := j.validateSetBgwriterLruMultiplierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgwriterLruMultiplier",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetDeadlockTimeout(val *float64) {
	if err := j.validateSetDeadlockTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deadlockTimeout",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetDefaultToastCompression(val *string) {
	if err := j.validateSetDefaultToastCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultToastCompression",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetIdleInTransactionSessionTimeout(val *float64) {
	if err := j.validateSetIdleInTransactionSessionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleInTransactionSessionTimeout",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetJit(val interface{}) {
	if err := j.validateSetJitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jit",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetLogAutovacuumMinDuration(val *float64) {
	if err := j.validateSetLogAutovacuumMinDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logAutovacuumMinDuration",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetLogErrorVerbosity(val *string) {
	if err := j.validateSetLogErrorVerbosityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logErrorVerbosity",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetLogLinePrefix(val *string) {
	if err := j.validateSetLogLinePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLinePrefix",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetLogMinDurationStatement(val *float64) {
	if err := j.validateSetLogMinDurationStatementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logMinDurationStatement",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetMaxFilesPerProcess(val *float64) {
	if err := j.validateSetMaxFilesPerProcessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFilesPerProcess",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetMaxLocksPerTransaction(val *float64) {
	if err := j.validateSetMaxLocksPerTransactionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLocksPerTransaction",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetMaxLogicalReplicationWorkers(val *float64) {
	if err := j.validateSetMaxLogicalReplicationWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLogicalReplicationWorkers",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetMaxParallelWorkers(val *float64) {
	if err := j.validateSetMaxParallelWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParallelWorkers",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetMaxParallelWorkersPerGather(val *float64) {
	if err := j.validateSetMaxParallelWorkersPerGatherParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParallelWorkersPerGather",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetMaxPredLocksPerTransaction(val *float64) {
	if err := j.validateSetMaxPredLocksPerTransactionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPredLocksPerTransaction",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetMaxPreparedTransactions(val *float64) {
	if err := j.validateSetMaxPreparedTransactionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPreparedTransactions",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetMaxReplicationSlots(val *float64) {
	if err := j.validateSetMaxReplicationSlotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxReplicationSlots",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetMaxStackDepth(val *float64) {
	if err := j.validateSetMaxStackDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxStackDepth",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetMaxStandbyArchiveDelay(val *float64) {
	if err := j.validateSetMaxStandbyArchiveDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxStandbyArchiveDelay",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetMaxStandbyStreamingDelay(val *float64) {
	if err := j.validateSetMaxStandbyStreamingDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxStandbyStreamingDelay",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetMaxWalSenders(val *float64) {
	if err := j.validateSetMaxWalSendersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWalSenders",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetMaxWorkerProcesses(val *float64) {
	if err := j.validateSetMaxWorkerProcessesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWorkerProcesses",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetPgPartmanBgwInterval(val *float64) {
	if err := j.validateSetPgPartmanBgwIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pgPartmanBgwInterval",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetPgPartmanBgwRole(val *string) {
	if err := j.validateSetPgPartmanBgwRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pgPartmanBgwRole",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetPgStatStatementsTrack(val *string) {
	if err := j.validateSetPgStatStatementsTrackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pgStatStatementsTrack",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetSharedBuffersPercentage(val *float64) {
	if err := j.validateSetSharedBuffersPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedBuffersPercentage",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetTempFileLimit(val *float64) {
	if err := j.validateSetTempFileLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tempFileLimit",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetTrackActivityQuerySize(val *float64) {
	if err := j.validateSetTrackActivityQuerySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackActivityQuerySize",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetTrackCommitTimestamp(val *string) {
	if err := j.validateSetTrackCommitTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackCommitTimestamp",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetTrackFunctions(val *string) {
	if err := j.validateSetTrackFunctionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackFunctions",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetTrackIoTiming(val *string) {
	if err := j.validateSetTrackIoTimingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackIoTiming",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetWalSenderTimeout(val *float64) {
	if err := j.validateSetWalSenderTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"walSenderTimeout",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetWalWriterDelay(val *float64) {
	if err := j.validateSetWalWriterDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"walWriterDelay",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfig)SetWorkMem(val *float64) {
	if err := j.validateSetWorkMemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workMem",
		val,
	)
}

// Generates CDKTF code for importing a DatabasePostgresqlConfig resource upon running "cdktf plan <stack-name>".
func DatabasePostgresqlConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatabasePostgresqlConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databasePostgresqlConfig.DatabasePostgresqlConfig",
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
func DatabasePostgresqlConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabasePostgresqlConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databasePostgresqlConfig.DatabasePostgresqlConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabasePostgresqlConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabasePostgresqlConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databasePostgresqlConfig.DatabasePostgresqlConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabasePostgresqlConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabasePostgresqlConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.databasePostgresqlConfig.DatabasePostgresqlConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatabasePostgresqlConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-digitalocean.databasePostgresqlConfig.DatabasePostgresqlConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabasePostgresqlConfig) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabasePostgresqlConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabasePostgresqlConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabasePostgresqlConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabasePostgresqlConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabasePostgresqlConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabasePostgresqlConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabasePostgresqlConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabasePostgresqlConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabasePostgresqlConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabasePostgresqlConfig) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) PutPgbouncer(value interface{}) {
	if err := d.validatePutPgbouncerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPgbouncer",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) PutTimescaledb(value *DatabasePostgresqlConfigTimescaledb) {
	if err := d.validatePutTimescaledbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimescaledb",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetAutovacuumAnalyzeScaleFactor() {
	_jsii_.InvokeVoid(
		d,
		"resetAutovacuumAnalyzeScaleFactor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetAutovacuumAnalyzeThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetAutovacuumAnalyzeThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetAutovacuumFreezeMaxAge() {
	_jsii_.InvokeVoid(
		d,
		"resetAutovacuumFreezeMaxAge",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetAutovacuumMaxWorkers() {
	_jsii_.InvokeVoid(
		d,
		"resetAutovacuumMaxWorkers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetAutovacuumNaptime() {
	_jsii_.InvokeVoid(
		d,
		"resetAutovacuumNaptime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetAutovacuumVacuumCostDelay() {
	_jsii_.InvokeVoid(
		d,
		"resetAutovacuumVacuumCostDelay",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetAutovacuumVacuumCostLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetAutovacuumVacuumCostLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetAutovacuumVacuumScaleFactor() {
	_jsii_.InvokeVoid(
		d,
		"resetAutovacuumVacuumScaleFactor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetAutovacuumVacuumThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetAutovacuumVacuumThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetBackupHour() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupHour",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetBackupMinute() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupMinute",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetBgwriterDelay() {
	_jsii_.InvokeVoid(
		d,
		"resetBgwriterDelay",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetBgwriterFlushAfter() {
	_jsii_.InvokeVoid(
		d,
		"resetBgwriterFlushAfter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetBgwriterLruMaxpages() {
	_jsii_.InvokeVoid(
		d,
		"resetBgwriterLruMaxpages",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetBgwriterLruMultiplier() {
	_jsii_.InvokeVoid(
		d,
		"resetBgwriterLruMultiplier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetDeadlockTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetDeadlockTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetDefaultToastCompression() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultToastCompression",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetIdleInTransactionSessionTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetIdleInTransactionSessionTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetJit() {
	_jsii_.InvokeVoid(
		d,
		"resetJit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetLogAutovacuumMinDuration() {
	_jsii_.InvokeVoid(
		d,
		"resetLogAutovacuumMinDuration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetLogErrorVerbosity() {
	_jsii_.InvokeVoid(
		d,
		"resetLogErrorVerbosity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetLogLinePrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetLogLinePrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetLogMinDurationStatement() {
	_jsii_.InvokeVoid(
		d,
		"resetLogMinDurationStatement",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetMaxFilesPerProcess() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxFilesPerProcess",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetMaxLocksPerTransaction() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxLocksPerTransaction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetMaxLogicalReplicationWorkers() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxLogicalReplicationWorkers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetMaxParallelWorkers() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxParallelWorkers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetMaxParallelWorkersPerGather() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxParallelWorkersPerGather",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetMaxPredLocksPerTransaction() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxPredLocksPerTransaction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetMaxPreparedTransactions() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxPreparedTransactions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetMaxReplicationSlots() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxReplicationSlots",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetMaxStackDepth() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxStackDepth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetMaxStandbyArchiveDelay() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxStandbyArchiveDelay",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetMaxStandbyStreamingDelay() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxStandbyStreamingDelay",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetMaxWalSenders() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxWalSenders",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetMaxWorkerProcesses() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxWorkerProcesses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetPgbouncer() {
	_jsii_.InvokeVoid(
		d,
		"resetPgbouncer",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetPgPartmanBgwInterval() {
	_jsii_.InvokeVoid(
		d,
		"resetPgPartmanBgwInterval",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetPgPartmanBgwRole() {
	_jsii_.InvokeVoid(
		d,
		"resetPgPartmanBgwRole",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetPgStatStatementsTrack() {
	_jsii_.InvokeVoid(
		d,
		"resetPgStatStatementsTrack",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetSharedBuffersPercentage() {
	_jsii_.InvokeVoid(
		d,
		"resetSharedBuffersPercentage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetTempFileLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetTempFileLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetTimescaledb() {
	_jsii_.InvokeVoid(
		d,
		"resetTimescaledb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetTimezone() {
	_jsii_.InvokeVoid(
		d,
		"resetTimezone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetTrackActivityQuerySize() {
	_jsii_.InvokeVoid(
		d,
		"resetTrackActivityQuerySize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetTrackCommitTimestamp() {
	_jsii_.InvokeVoid(
		d,
		"resetTrackCommitTimestamp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetTrackFunctions() {
	_jsii_.InvokeVoid(
		d,
		"resetTrackFunctions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetTrackIoTiming() {
	_jsii_.InvokeVoid(
		d,
		"resetTrackIoTiming",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetWalSenderTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetWalSenderTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetWalWriterDelay() {
	_jsii_.InvokeVoid(
		d,
		"resetWalWriterDelay",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ResetWorkMem() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkMem",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabasePostgresqlConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabasePostgresqlConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

