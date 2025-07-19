// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasemysqlconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseMysqlConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#cluster_id DatabaseMysqlConfig#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#backup_hour DatabaseMysqlConfig#backup_hour}.
	BackupHour *float64 `field:"optional" json:"backupHour" yaml:"backupHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#backup_minute DatabaseMysqlConfig#backup_minute}.
	BackupMinute *float64 `field:"optional" json:"backupMinute" yaml:"backupMinute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#binlog_retention_period DatabaseMysqlConfig#binlog_retention_period}.
	BinlogRetentionPeriod *float64 `field:"optional" json:"binlogRetentionPeriod" yaml:"binlogRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#connect_timeout DatabaseMysqlConfig#connect_timeout}.
	ConnectTimeout *float64 `field:"optional" json:"connectTimeout" yaml:"connectTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#default_time_zone DatabaseMysqlConfig#default_time_zone}.
	DefaultTimeZone *string `field:"optional" json:"defaultTimeZone" yaml:"defaultTimeZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#group_concat_max_len DatabaseMysqlConfig#group_concat_max_len}.
	GroupConcatMaxLen *float64 `field:"optional" json:"groupConcatMaxLen" yaml:"groupConcatMaxLen"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#id DatabaseMysqlConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#information_schema_stats_expiry DatabaseMysqlConfig#information_schema_stats_expiry}.
	InformationSchemaStatsExpiry *float64 `field:"optional" json:"informationSchemaStatsExpiry" yaml:"informationSchemaStatsExpiry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#innodb_ft_min_token_size DatabaseMysqlConfig#innodb_ft_min_token_size}.
	InnodbFtMinTokenSize *float64 `field:"optional" json:"innodbFtMinTokenSize" yaml:"innodbFtMinTokenSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#innodb_ft_server_stopword_table DatabaseMysqlConfig#innodb_ft_server_stopword_table}.
	InnodbFtServerStopwordTable *string `field:"optional" json:"innodbFtServerStopwordTable" yaml:"innodbFtServerStopwordTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#innodb_lock_wait_timeout DatabaseMysqlConfig#innodb_lock_wait_timeout}.
	InnodbLockWaitTimeout *float64 `field:"optional" json:"innodbLockWaitTimeout" yaml:"innodbLockWaitTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#innodb_log_buffer_size DatabaseMysqlConfig#innodb_log_buffer_size}.
	InnodbLogBufferSize *float64 `field:"optional" json:"innodbLogBufferSize" yaml:"innodbLogBufferSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#innodb_online_alter_log_max_size DatabaseMysqlConfig#innodb_online_alter_log_max_size}.
	InnodbOnlineAlterLogMaxSize *float64 `field:"optional" json:"innodbOnlineAlterLogMaxSize" yaml:"innodbOnlineAlterLogMaxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#innodb_print_all_deadlocks DatabaseMysqlConfig#innodb_print_all_deadlocks}.
	InnodbPrintAllDeadlocks interface{} `field:"optional" json:"innodbPrintAllDeadlocks" yaml:"innodbPrintAllDeadlocks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#innodb_rollback_on_timeout DatabaseMysqlConfig#innodb_rollback_on_timeout}.
	InnodbRollbackOnTimeout interface{} `field:"optional" json:"innodbRollbackOnTimeout" yaml:"innodbRollbackOnTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#interactive_timeout DatabaseMysqlConfig#interactive_timeout}.
	InteractiveTimeout *float64 `field:"optional" json:"interactiveTimeout" yaml:"interactiveTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#internal_tmp_mem_storage_engine DatabaseMysqlConfig#internal_tmp_mem_storage_engine}.
	InternalTmpMemStorageEngine *string `field:"optional" json:"internalTmpMemStorageEngine" yaml:"internalTmpMemStorageEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#long_query_time DatabaseMysqlConfig#long_query_time}.
	LongQueryTime *float64 `field:"optional" json:"longQueryTime" yaml:"longQueryTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#max_allowed_packet DatabaseMysqlConfig#max_allowed_packet}.
	MaxAllowedPacket *float64 `field:"optional" json:"maxAllowedPacket" yaml:"maxAllowedPacket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#max_heap_table_size DatabaseMysqlConfig#max_heap_table_size}.
	MaxHeapTableSize *float64 `field:"optional" json:"maxHeapTableSize" yaml:"maxHeapTableSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#net_read_timeout DatabaseMysqlConfig#net_read_timeout}.
	NetReadTimeout *float64 `field:"optional" json:"netReadTimeout" yaml:"netReadTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#net_write_timeout DatabaseMysqlConfig#net_write_timeout}.
	NetWriteTimeout *float64 `field:"optional" json:"netWriteTimeout" yaml:"netWriteTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#slow_query_log DatabaseMysqlConfig#slow_query_log}.
	SlowQueryLog interface{} `field:"optional" json:"slowQueryLog" yaml:"slowQueryLog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#sort_buffer_size DatabaseMysqlConfig#sort_buffer_size}.
	SortBufferSize *float64 `field:"optional" json:"sortBufferSize" yaml:"sortBufferSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#sql_mode DatabaseMysqlConfig#sql_mode}.
	SqlMode *string `field:"optional" json:"sqlMode" yaml:"sqlMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#sql_require_primary_key DatabaseMysqlConfig#sql_require_primary_key}.
	SqlRequirePrimaryKey interface{} `field:"optional" json:"sqlRequirePrimaryKey" yaml:"sqlRequirePrimaryKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#tmp_table_size DatabaseMysqlConfig#tmp_table_size}.
	TmpTableSize *float64 `field:"optional" json:"tmpTableSize" yaml:"tmpTableSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/database_mysql_config#wait_timeout DatabaseMysqlConfig#wait_timeout}.
	WaitTimeout *float64 `field:"optional" json:"waitTimeout" yaml:"waitTimeout"`
}

