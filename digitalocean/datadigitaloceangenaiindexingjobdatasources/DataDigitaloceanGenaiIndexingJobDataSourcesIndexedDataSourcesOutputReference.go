// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiindexingjobdatasources

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/datadigitaloceangenaiindexingjobdatasources/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference interface {
	cdktf.ComplexObject
	CompletedAt() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataSourceUuid() *string
	ErrorDetails() *string
	ErrorMsg() *string
	FailedItemCount() *string
	// Experimental.
	Fqn() *string
	IndexedFileCount() *string
	IndexedItemCount() *string
	InternalValue() *DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSources
	SetInternalValue(val *DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSources)
	RemovedItemCount() *string
	SkippedItemCount() *string
	StartedAt() *string
	Status() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TotalBytes() *string
	TotalBytesIndexed() *string
	TotalFileCount() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference
type jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) CompletedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"completedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) DataSourceUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) ErrorDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) ErrorMsg() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorMsg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) FailedItemCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failedItemCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) IndexedFileCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexedFileCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) IndexedItemCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexedItemCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) InternalValue() *DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSources {
	var returns *DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSources
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) RemovedItemCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"removedItemCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) SkippedItemCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skippedItemCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) StartedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) TotalBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) TotalBytesIndexed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalBytesIndexed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) TotalFileCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalFileCount",
		&returns,
	)
	return returns
}


func NewDataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference {
	_init_.Initialize()

	if err := validateNewDataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanGenaiIndexingJobDataSources.DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference_Override(d DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanGenaiIndexingJobDataSources.DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference)SetInternalValue(val *DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSources) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanGenaiIndexingJobDataSourcesIndexedDataSourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

