// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasepostgresqlconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabasePostgresqlConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#cluster_id DatabasePostgresqlConfig#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#autovacuum_analyze_scale_factor DatabasePostgresqlConfig#autovacuum_analyze_scale_factor}.
	AutovacuumAnalyzeScaleFactor *float64 `field:"optional" json:"autovacuumAnalyzeScaleFactor" yaml:"autovacuumAnalyzeScaleFactor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#autovacuum_analyze_threshold DatabasePostgresqlConfig#autovacuum_analyze_threshold}.
	AutovacuumAnalyzeThreshold *float64 `field:"optional" json:"autovacuumAnalyzeThreshold" yaml:"autovacuumAnalyzeThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#autovacuum_freeze_max_age DatabasePostgresqlConfig#autovacuum_freeze_max_age}.
	AutovacuumFreezeMaxAge *float64 `field:"optional" json:"autovacuumFreezeMaxAge" yaml:"autovacuumFreezeMaxAge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#autovacuum_max_workers DatabasePostgresqlConfig#autovacuum_max_workers}.
	AutovacuumMaxWorkers *float64 `field:"optional" json:"autovacuumMaxWorkers" yaml:"autovacuumMaxWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#autovacuum_naptime DatabasePostgresqlConfig#autovacuum_naptime}.
	AutovacuumNaptime *float64 `field:"optional" json:"autovacuumNaptime" yaml:"autovacuumNaptime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#autovacuum_vacuum_cost_delay DatabasePostgresqlConfig#autovacuum_vacuum_cost_delay}.
	AutovacuumVacuumCostDelay *float64 `field:"optional" json:"autovacuumVacuumCostDelay" yaml:"autovacuumVacuumCostDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#autovacuum_vacuum_cost_limit DatabasePostgresqlConfig#autovacuum_vacuum_cost_limit}.
	AutovacuumVacuumCostLimit *float64 `field:"optional" json:"autovacuumVacuumCostLimit" yaml:"autovacuumVacuumCostLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#autovacuum_vacuum_scale_factor DatabasePostgresqlConfig#autovacuum_vacuum_scale_factor}.
	AutovacuumVacuumScaleFactor *float64 `field:"optional" json:"autovacuumVacuumScaleFactor" yaml:"autovacuumVacuumScaleFactor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#autovacuum_vacuum_threshold DatabasePostgresqlConfig#autovacuum_vacuum_threshold}.
	AutovacuumVacuumThreshold *float64 `field:"optional" json:"autovacuumVacuumThreshold" yaml:"autovacuumVacuumThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#backup_hour DatabasePostgresqlConfig#backup_hour}.
	BackupHour *float64 `field:"optional" json:"backupHour" yaml:"backupHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#backup_minute DatabasePostgresqlConfig#backup_minute}.
	BackupMinute *float64 `field:"optional" json:"backupMinute" yaml:"backupMinute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#bgwriter_delay DatabasePostgresqlConfig#bgwriter_delay}.
	BgwriterDelay *float64 `field:"optional" json:"bgwriterDelay" yaml:"bgwriterDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#bgwriter_flush_after DatabasePostgresqlConfig#bgwriter_flush_after}.
	BgwriterFlushAfter *float64 `field:"optional" json:"bgwriterFlushAfter" yaml:"bgwriterFlushAfter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#bgwriter_lru_maxpages DatabasePostgresqlConfig#bgwriter_lru_maxpages}.
	BgwriterLruMaxpages *float64 `field:"optional" json:"bgwriterLruMaxpages" yaml:"bgwriterLruMaxpages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#bgwriter_lru_multiplier DatabasePostgresqlConfig#bgwriter_lru_multiplier}.
	BgwriterLruMultiplier *float64 `field:"optional" json:"bgwriterLruMultiplier" yaml:"bgwriterLruMultiplier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#deadlock_timeout DatabasePostgresqlConfig#deadlock_timeout}.
	DeadlockTimeout *float64 `field:"optional" json:"deadlockTimeout" yaml:"deadlockTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#default_toast_compression DatabasePostgresqlConfig#default_toast_compression}.
	DefaultToastCompression *string `field:"optional" json:"defaultToastCompression" yaml:"defaultToastCompression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#id DatabasePostgresqlConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#idle_in_transaction_session_timeout DatabasePostgresqlConfig#idle_in_transaction_session_timeout}.
	IdleInTransactionSessionTimeout *float64 `field:"optional" json:"idleInTransactionSessionTimeout" yaml:"idleInTransactionSessionTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#jit DatabasePostgresqlConfig#jit}.
	Jit interface{} `field:"optional" json:"jit" yaml:"jit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#log_autovacuum_min_duration DatabasePostgresqlConfig#log_autovacuum_min_duration}.
	LogAutovacuumMinDuration *float64 `field:"optional" json:"logAutovacuumMinDuration" yaml:"logAutovacuumMinDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#log_error_verbosity DatabasePostgresqlConfig#log_error_verbosity}.
	LogErrorVerbosity *string `field:"optional" json:"logErrorVerbosity" yaml:"logErrorVerbosity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#log_line_prefix DatabasePostgresqlConfig#log_line_prefix}.
	LogLinePrefix *string `field:"optional" json:"logLinePrefix" yaml:"logLinePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#log_min_duration_statement DatabasePostgresqlConfig#log_min_duration_statement}.
	LogMinDurationStatement *float64 `field:"optional" json:"logMinDurationStatement" yaml:"logMinDurationStatement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#max_files_per_process DatabasePostgresqlConfig#max_files_per_process}.
	MaxFilesPerProcess *float64 `field:"optional" json:"maxFilesPerProcess" yaml:"maxFilesPerProcess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#max_locks_per_transaction DatabasePostgresqlConfig#max_locks_per_transaction}.
	MaxLocksPerTransaction *float64 `field:"optional" json:"maxLocksPerTransaction" yaml:"maxLocksPerTransaction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#max_logical_replication_workers DatabasePostgresqlConfig#max_logical_replication_workers}.
	MaxLogicalReplicationWorkers *float64 `field:"optional" json:"maxLogicalReplicationWorkers" yaml:"maxLogicalReplicationWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#max_parallel_workers DatabasePostgresqlConfig#max_parallel_workers}.
	MaxParallelWorkers *float64 `field:"optional" json:"maxParallelWorkers" yaml:"maxParallelWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#max_parallel_workers_per_gather DatabasePostgresqlConfig#max_parallel_workers_per_gather}.
	MaxParallelWorkersPerGather *float64 `field:"optional" json:"maxParallelWorkersPerGather" yaml:"maxParallelWorkersPerGather"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#max_pred_locks_per_transaction DatabasePostgresqlConfig#max_pred_locks_per_transaction}.
	MaxPredLocksPerTransaction *float64 `field:"optional" json:"maxPredLocksPerTransaction" yaml:"maxPredLocksPerTransaction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#max_prepared_transactions DatabasePostgresqlConfig#max_prepared_transactions}.
	MaxPreparedTransactions *float64 `field:"optional" json:"maxPreparedTransactions" yaml:"maxPreparedTransactions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#max_replication_slots DatabasePostgresqlConfig#max_replication_slots}.
	MaxReplicationSlots *float64 `field:"optional" json:"maxReplicationSlots" yaml:"maxReplicationSlots"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#max_stack_depth DatabasePostgresqlConfig#max_stack_depth}.
	MaxStackDepth *float64 `field:"optional" json:"maxStackDepth" yaml:"maxStackDepth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#max_standby_archive_delay DatabasePostgresqlConfig#max_standby_archive_delay}.
	MaxStandbyArchiveDelay *float64 `field:"optional" json:"maxStandbyArchiveDelay" yaml:"maxStandbyArchiveDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#max_standby_streaming_delay DatabasePostgresqlConfig#max_standby_streaming_delay}.
	MaxStandbyStreamingDelay *float64 `field:"optional" json:"maxStandbyStreamingDelay" yaml:"maxStandbyStreamingDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#max_wal_senders DatabasePostgresqlConfig#max_wal_senders}.
	MaxWalSenders *float64 `field:"optional" json:"maxWalSenders" yaml:"maxWalSenders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#max_worker_processes DatabasePostgresqlConfig#max_worker_processes}.
	MaxWorkerProcesses *float64 `field:"optional" json:"maxWorkerProcesses" yaml:"maxWorkerProcesses"`
	// pgbouncer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#pgbouncer DatabasePostgresqlConfig#pgbouncer}
	Pgbouncer interface{} `field:"optional" json:"pgbouncer" yaml:"pgbouncer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#pg_partman_bgw_interval DatabasePostgresqlConfig#pg_partman_bgw_interval}.
	PgPartmanBgwInterval *float64 `field:"optional" json:"pgPartmanBgwInterval" yaml:"pgPartmanBgwInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#pg_partman_bgw_role DatabasePostgresqlConfig#pg_partman_bgw_role}.
	PgPartmanBgwRole *string `field:"optional" json:"pgPartmanBgwRole" yaml:"pgPartmanBgwRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#pg_stat_statements_track DatabasePostgresqlConfig#pg_stat_statements_track}.
	PgStatStatementsTrack *string `field:"optional" json:"pgStatStatementsTrack" yaml:"pgStatStatementsTrack"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#shared_buffers_percentage DatabasePostgresqlConfig#shared_buffers_percentage}.
	SharedBuffersPercentage *float64 `field:"optional" json:"sharedBuffersPercentage" yaml:"sharedBuffersPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#temp_file_limit DatabasePostgresqlConfig#temp_file_limit}.
	TempFileLimit *float64 `field:"optional" json:"tempFileLimit" yaml:"tempFileLimit"`
	// timescaledb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#timescaledb DatabasePostgresqlConfig#timescaledb}
	Timescaledb *DatabasePostgresqlConfigTimescaledb `field:"optional" json:"timescaledb" yaml:"timescaledb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#timezone DatabasePostgresqlConfig#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#track_activity_query_size DatabasePostgresqlConfig#track_activity_query_size}.
	TrackActivityQuerySize *float64 `field:"optional" json:"trackActivityQuerySize" yaml:"trackActivityQuerySize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#track_commit_timestamp DatabasePostgresqlConfig#track_commit_timestamp}.
	TrackCommitTimestamp *string `field:"optional" json:"trackCommitTimestamp" yaml:"trackCommitTimestamp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#track_functions DatabasePostgresqlConfig#track_functions}.
	TrackFunctions *string `field:"optional" json:"trackFunctions" yaml:"trackFunctions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#track_io_timing DatabasePostgresqlConfig#track_io_timing}.
	TrackIoTiming *string `field:"optional" json:"trackIoTiming" yaml:"trackIoTiming"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#wal_sender_timeout DatabasePostgresqlConfig#wal_sender_timeout}.
	WalSenderTimeout *float64 `field:"optional" json:"walSenderTimeout" yaml:"walSenderTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#wal_writer_delay DatabasePostgresqlConfig#wal_writer_delay}.
	WalWriterDelay *float64 `field:"optional" json:"walWriterDelay" yaml:"walWriterDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/database_postgresql_config#work_mem DatabasePostgresqlConfig#work_mem}.
	WorkMem *float64 `field:"optional" json:"workMem" yaml:"workMem"`
}

