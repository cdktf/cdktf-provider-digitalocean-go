// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasekafkatopic


type DatabaseKafkaTopicConfigA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#cleanup_policy DatabaseKafkaTopic#cleanup_policy}.
	CleanupPolicy *string `field:"optional" json:"cleanupPolicy" yaml:"cleanupPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#compression_type DatabaseKafkaTopic#compression_type}.
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#delete_retention_ms DatabaseKafkaTopic#delete_retention_ms}.
	DeleteRetentionMs *string `field:"optional" json:"deleteRetentionMs" yaml:"deleteRetentionMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#file_delete_delay_ms DatabaseKafkaTopic#file_delete_delay_ms}.
	FileDeleteDelayMs *string `field:"optional" json:"fileDeleteDelayMs" yaml:"fileDeleteDelayMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#flush_messages DatabaseKafkaTopic#flush_messages}.
	FlushMessages *string `field:"optional" json:"flushMessages" yaml:"flushMessages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#flush_ms DatabaseKafkaTopic#flush_ms}.
	FlushMs *string `field:"optional" json:"flushMs" yaml:"flushMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#index_interval_bytes DatabaseKafkaTopic#index_interval_bytes}.
	IndexIntervalBytes *string `field:"optional" json:"indexIntervalBytes" yaml:"indexIntervalBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#max_compaction_lag_ms DatabaseKafkaTopic#max_compaction_lag_ms}.
	MaxCompactionLagMs *string `field:"optional" json:"maxCompactionLagMs" yaml:"maxCompactionLagMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#max_message_bytes DatabaseKafkaTopic#max_message_bytes}.
	MaxMessageBytes *string `field:"optional" json:"maxMessageBytes" yaml:"maxMessageBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#message_down_conversion_enable DatabaseKafkaTopic#message_down_conversion_enable}.
	MessageDownConversionEnable interface{} `field:"optional" json:"messageDownConversionEnable" yaml:"messageDownConversionEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#message_format_version DatabaseKafkaTopic#message_format_version}.
	MessageFormatVersion *string `field:"optional" json:"messageFormatVersion" yaml:"messageFormatVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#message_timestamp_difference_max_ms DatabaseKafkaTopic#message_timestamp_difference_max_ms}.
	MessageTimestampDifferenceMaxMs *string `field:"optional" json:"messageTimestampDifferenceMaxMs" yaml:"messageTimestampDifferenceMaxMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#message_timestamp_type DatabaseKafkaTopic#message_timestamp_type}.
	MessageTimestampType *string `field:"optional" json:"messageTimestampType" yaml:"messageTimestampType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#min_cleanable_dirty_ratio DatabaseKafkaTopic#min_cleanable_dirty_ratio}.
	MinCleanableDirtyRatio *float64 `field:"optional" json:"minCleanableDirtyRatio" yaml:"minCleanableDirtyRatio"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#min_compaction_lag_ms DatabaseKafkaTopic#min_compaction_lag_ms}.
	MinCompactionLagMs *string `field:"optional" json:"minCompactionLagMs" yaml:"minCompactionLagMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#min_insync_replicas DatabaseKafkaTopic#min_insync_replicas}.
	MinInsyncReplicas *float64 `field:"optional" json:"minInsyncReplicas" yaml:"minInsyncReplicas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#preallocate DatabaseKafkaTopic#preallocate}.
	Preallocate interface{} `field:"optional" json:"preallocate" yaml:"preallocate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#retention_bytes DatabaseKafkaTopic#retention_bytes}.
	RetentionBytes *string `field:"optional" json:"retentionBytes" yaml:"retentionBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#retention_ms DatabaseKafkaTopic#retention_ms}.
	RetentionMs *string `field:"optional" json:"retentionMs" yaml:"retentionMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#segment_bytes DatabaseKafkaTopic#segment_bytes}.
	SegmentBytes *string `field:"optional" json:"segmentBytes" yaml:"segmentBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#segment_index_bytes DatabaseKafkaTopic#segment_index_bytes}.
	SegmentIndexBytes *string `field:"optional" json:"segmentIndexBytes" yaml:"segmentIndexBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#segment_jitter_ms DatabaseKafkaTopic#segment_jitter_ms}.
	SegmentJitterMs *string `field:"optional" json:"segmentJitterMs" yaml:"segmentJitterMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/database_kafka_topic#segment_ms DatabaseKafkaTopic#segment_ms}.
	SegmentMs *string `field:"optional" json:"segmentMs" yaml:"segmentMs"`
}

