// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasekafkatopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v10/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v10/databasekafkatopic/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseKafkaTopicConfigAOutputReference interface {
	cdktf.ComplexObject
	CleanupPolicy() *string
	SetCleanupPolicy(val *string)
	CleanupPolicyInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	CompressionType() *string
	SetCompressionType(val *string)
	CompressionTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeleteRetentionMs() *string
	SetDeleteRetentionMs(val *string)
	DeleteRetentionMsInput() *string
	FileDeleteDelayMs() *string
	SetFileDeleteDelayMs(val *string)
	FileDeleteDelayMsInput() *string
	FlushMessages() *string
	SetFlushMessages(val *string)
	FlushMessagesInput() *string
	FlushMs() *string
	SetFlushMs(val *string)
	FlushMsInput() *string
	// Experimental.
	Fqn() *string
	IndexIntervalBytes() *string
	SetIndexIntervalBytes(val *string)
	IndexIntervalBytesInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxCompactionLagMs() *string
	SetMaxCompactionLagMs(val *string)
	MaxCompactionLagMsInput() *string
	MaxMessageBytes() *string
	SetMaxMessageBytes(val *string)
	MaxMessageBytesInput() *string
	MessageDownConversionEnable() interface{}
	SetMessageDownConversionEnable(val interface{})
	MessageDownConversionEnableInput() interface{}
	MessageFormatVersion() *string
	SetMessageFormatVersion(val *string)
	MessageFormatVersionInput() *string
	MessageTimestampDifferenceMaxMs() *string
	SetMessageTimestampDifferenceMaxMs(val *string)
	MessageTimestampDifferenceMaxMsInput() *string
	MessageTimestampType() *string
	SetMessageTimestampType(val *string)
	MessageTimestampTypeInput() *string
	MinCleanableDirtyRatio() *float64
	SetMinCleanableDirtyRatio(val *float64)
	MinCleanableDirtyRatioInput() *float64
	MinCompactionLagMs() *string
	SetMinCompactionLagMs(val *string)
	MinCompactionLagMsInput() *string
	MinInsyncReplicas() *float64
	SetMinInsyncReplicas(val *float64)
	MinInsyncReplicasInput() *float64
	Preallocate() interface{}
	SetPreallocate(val interface{})
	PreallocateInput() interface{}
	RetentionBytes() *string
	SetRetentionBytes(val *string)
	RetentionBytesInput() *string
	RetentionMs() *string
	SetRetentionMs(val *string)
	RetentionMsInput() *string
	SegmentBytes() *string
	SetSegmentBytes(val *string)
	SegmentBytesInput() *string
	SegmentIndexBytes() *string
	SetSegmentIndexBytes(val *string)
	SegmentIndexBytesInput() *string
	SegmentJitterMs() *string
	SetSegmentJitterMs(val *string)
	SegmentJitterMsInput() *string
	SegmentMs() *string
	SetSegmentMs(val *string)
	SegmentMsInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UncleanLeaderElectionEnable() interface{}
	SetUncleanLeaderElectionEnable(val interface{})
	UncleanLeaderElectionEnableInput() interface{}
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCleanupPolicy()
	ResetCompressionType()
	ResetDeleteRetentionMs()
	ResetFileDeleteDelayMs()
	ResetFlushMessages()
	ResetFlushMs()
	ResetIndexIntervalBytes()
	ResetMaxCompactionLagMs()
	ResetMaxMessageBytes()
	ResetMessageDownConversionEnable()
	ResetMessageFormatVersion()
	ResetMessageTimestampDifferenceMaxMs()
	ResetMessageTimestampType()
	ResetMinCleanableDirtyRatio()
	ResetMinCompactionLagMs()
	ResetMinInsyncReplicas()
	ResetPreallocate()
	ResetRetentionBytes()
	ResetRetentionMs()
	ResetSegmentBytes()
	ResetSegmentIndexBytes()
	ResetSegmentJitterMs()
	ResetSegmentMs()
	ResetUncleanLeaderElectionEnable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabaseKafkaTopicConfigAOutputReference
type jsiiProxy_DatabaseKafkaTopicConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) CleanupPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanupPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) CleanupPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cleanupPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) DeleteRetentionMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteRetentionMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) DeleteRetentionMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteRetentionMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) FileDeleteDelayMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileDeleteDelayMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) FileDeleteDelayMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileDeleteDelayMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) FlushMessages() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMessages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) FlushMessagesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMessagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) FlushMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) FlushMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flushMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) IndexIntervalBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexIntervalBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) IndexIntervalBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexIntervalBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MaxCompactionLagMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxCompactionLagMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MaxCompactionLagMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxCompactionLagMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MaxMessageBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxMessageBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MaxMessageBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxMessageBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MessageDownConversionEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageDownConversionEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MessageDownConversionEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageDownConversionEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MessageFormatVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormatVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MessageFormatVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormatVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MessageTimestampDifferenceMaxMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageTimestampDifferenceMaxMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MessageTimestampDifferenceMaxMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageTimestampDifferenceMaxMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MessageTimestampType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageTimestampType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MessageTimestampTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageTimestampTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MinCleanableDirtyRatio() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCleanableDirtyRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MinCleanableDirtyRatioInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCleanableDirtyRatioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MinCompactionLagMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCompactionLagMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MinCompactionLagMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCompactionLagMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MinInsyncReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInsyncReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) MinInsyncReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInsyncReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) Preallocate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preallocate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) PreallocateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preallocateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) RetentionBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) RetentionBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) RetentionMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) RetentionMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) SegmentBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) SegmentBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) SegmentIndexBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentIndexBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) SegmentIndexBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentIndexBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) SegmentJitterMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentJitterMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) SegmentJitterMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentJitterMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) SegmentMs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) SegmentMsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) UncleanLeaderElectionEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uncleanLeaderElectionEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) UncleanLeaderElectionEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uncleanLeaderElectionEnableInput",
		&returns,
	)
	return returns
}


func NewDatabaseKafkaTopicConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DatabaseKafkaTopicConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseKafkaTopicConfigAOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseKafkaTopicConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.databaseKafkaTopic.DatabaseKafkaTopicConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDatabaseKafkaTopicConfigAOutputReference_Override(d DatabaseKafkaTopicConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.databaseKafkaTopic.DatabaseKafkaTopicConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetCleanupPolicy(val *string) {
	if err := j.validateSetCleanupPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanupPolicy",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetCompressionType(val *string) {
	if err := j.validateSetCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetDeleteRetentionMs(val *string) {
	if err := j.validateSetDeleteRetentionMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteRetentionMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetFileDeleteDelayMs(val *string) {
	if err := j.validateSetFileDeleteDelayMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileDeleteDelayMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetFlushMessages(val *string) {
	if err := j.validateSetFlushMessagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flushMessages",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetFlushMs(val *string) {
	if err := j.validateSetFlushMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flushMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetIndexIntervalBytes(val *string) {
	if err := j.validateSetIndexIntervalBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexIntervalBytes",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetMaxCompactionLagMs(val *string) {
	if err := j.validateSetMaxCompactionLagMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCompactionLagMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetMaxMessageBytes(val *string) {
	if err := j.validateSetMaxMessageBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxMessageBytes",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetMessageDownConversionEnable(val interface{}) {
	if err := j.validateSetMessageDownConversionEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageDownConversionEnable",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetMessageFormatVersion(val *string) {
	if err := j.validateSetMessageFormatVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageFormatVersion",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetMessageTimestampDifferenceMaxMs(val *string) {
	if err := j.validateSetMessageTimestampDifferenceMaxMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageTimestampDifferenceMaxMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetMessageTimestampType(val *string) {
	if err := j.validateSetMessageTimestampTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageTimestampType",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetMinCleanableDirtyRatio(val *float64) {
	if err := j.validateSetMinCleanableDirtyRatioParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCleanableDirtyRatio",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetMinCompactionLagMs(val *string) {
	if err := j.validateSetMinCompactionLagMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCompactionLagMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetMinInsyncReplicas(val *float64) {
	if err := j.validateSetMinInsyncReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minInsyncReplicas",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetPreallocate(val interface{}) {
	if err := j.validateSetPreallocateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preallocate",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetRetentionBytes(val *string) {
	if err := j.validateSetRetentionBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionBytes",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetRetentionMs(val *string) {
	if err := j.validateSetRetentionMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetSegmentBytes(val *string) {
	if err := j.validateSetSegmentBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentBytes",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetSegmentIndexBytes(val *string) {
	if err := j.validateSetSegmentIndexBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentIndexBytes",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetSegmentJitterMs(val *string) {
	if err := j.validateSetSegmentJitterMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentJitterMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetSegmentMs(val *string) {
	if err := j.validateSetSegmentMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentMs",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference)SetUncleanLeaderElectionEnable(val interface{}) {
	if err := j.validateSetUncleanLeaderElectionEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uncleanLeaderElectionEnable",
		val,
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetCleanupPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetCleanupPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetCompressionType() {
	_jsii_.InvokeVoid(
		d,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetDeleteRetentionMs() {
	_jsii_.InvokeVoid(
		d,
		"resetDeleteRetentionMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetFileDeleteDelayMs() {
	_jsii_.InvokeVoid(
		d,
		"resetFileDeleteDelayMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetFlushMessages() {
	_jsii_.InvokeVoid(
		d,
		"resetFlushMessages",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetFlushMs() {
	_jsii_.InvokeVoid(
		d,
		"resetFlushMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetIndexIntervalBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetIndexIntervalBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetMaxCompactionLagMs() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxCompactionLagMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetMaxMessageBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxMessageBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetMessageDownConversionEnable() {
	_jsii_.InvokeVoid(
		d,
		"resetMessageDownConversionEnable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetMessageFormatVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetMessageFormatVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetMessageTimestampDifferenceMaxMs() {
	_jsii_.InvokeVoid(
		d,
		"resetMessageTimestampDifferenceMaxMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetMessageTimestampType() {
	_jsii_.InvokeVoid(
		d,
		"resetMessageTimestampType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetMinCleanableDirtyRatio() {
	_jsii_.InvokeVoid(
		d,
		"resetMinCleanableDirtyRatio",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetMinCompactionLagMs() {
	_jsii_.InvokeVoid(
		d,
		"resetMinCompactionLagMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetMinInsyncReplicas() {
	_jsii_.InvokeVoid(
		d,
		"resetMinInsyncReplicas",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetPreallocate() {
	_jsii_.InvokeVoid(
		d,
		"resetPreallocate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetRetentionBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetRetentionMs() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetSegmentBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetSegmentBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetSegmentIndexBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetSegmentIndexBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetSegmentJitterMs() {
	_jsii_.InvokeVoid(
		d,
		"resetSegmentJitterMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetSegmentMs() {
	_jsii_.InvokeVoid(
		d,
		"resetSegmentMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ResetUncleanLeaderElectionEnable() {
	_jsii_.InvokeVoid(
		d,
		"resetUncleanLeaderElectionEnable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseKafkaTopicConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

