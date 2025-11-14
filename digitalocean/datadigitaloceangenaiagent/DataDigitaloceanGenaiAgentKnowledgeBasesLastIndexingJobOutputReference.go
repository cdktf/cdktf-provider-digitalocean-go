// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/datadigitaloceangenaiagent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference interface {
	cdktf.ComplexObject
	CompletedDatasources() *float64
	SetCompletedDatasources(val *float64)
	CompletedDatasourcesInput() *float64
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
	CreatedAt() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataSourceUuids() *[]*string
	SetDataSourceUuids(val *[]*string)
	DataSourceUuidsInput() *[]*string
	FinishedAt() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJob
	SetInternalValue(val *DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJob)
	KnowledgeBaseUuid() *string
	Phase() *string
	SetPhase(val *string)
	PhaseInput() *string
	StartedAt() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tokens() *float64
	SetTokens(val *float64)
	TokensInput() *float64
	TotalDatasources() *float64
	SetTotalDatasources(val *float64)
	TotalDatasourcesInput() *float64
	UpdatedAt() *string
	Uuid() *string
	SetUuid(val *string)
	UuidInput() *string
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
	ResetCompletedDatasources()
	ResetDataSourceUuids()
	ResetPhase()
	ResetTokens()
	ResetTotalDatasources()
	ResetUuid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference
type jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) CompletedDatasources() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"completedDatasources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) CompletedDatasourcesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"completedDatasourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) DataSourceUuids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataSourceUuids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) DataSourceUuidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataSourceUuidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) FinishedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finishedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) InternalValue() *DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJob {
	var returns *DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) KnowledgeBaseUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) Phase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) PhaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) StartedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) Tokens() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) TokensInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) TotalDatasources() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalDatasources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) TotalDatasourcesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalDatasourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) UuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuidInput",
		&returns,
	)
	return returns
}


func NewDataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference {
	_init_.Initialize()

	if err := validateNewDataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanGenaiAgent.DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference_Override(d DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanGenaiAgent.DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference)SetCompletedDatasources(val *float64) {
	if err := j.validateSetCompletedDatasourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"completedDatasources",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference)SetDataSourceUuids(val *[]*string) {
	if err := j.validateSetDataSourceUuidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceUuids",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference)SetInternalValue(val *DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference)SetPhase(val *string) {
	if err := j.validateSetPhaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phase",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference)SetTokens(val *float64) {
	if err := j.validateSetTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokens",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference)SetTotalDatasources(val *float64) {
	if err := j.validateSetTotalDatasourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalDatasources",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference)SetUuid(val *string) {
	if err := j.validateSetUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uuid",
		val,
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) ResetCompletedDatasources() {
	_jsii_.InvokeVoid(
		d,
		"resetCompletedDatasources",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) ResetDataSourceUuids() {
	_jsii_.InvokeVoid(
		d,
		"resetDataSourceUuids",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) ResetPhase() {
	_jsii_.InvokeVoid(
		d,
		"resetPhase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) ResetTokens() {
	_jsii_.InvokeVoid(
		d,
		"resetTokens",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) ResetTotalDatasources() {
	_jsii_.InvokeVoid(
		d,
		"resetTotalDatasources",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) ResetUuid() {
	_jsii_.InvokeVoid(
		d,
		"resetUuid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentKnowledgeBasesLastIndexingJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

