// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/datadigitaloceangenaiagent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDigitaloceanGenaiAgentModelOutputReference interface {
	cdktf.ComplexObject
	Agreement() DataDigitaloceanGenaiAgentModelAgreementList
	AgreementInput() interface{}
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
	// Experimental.
	Fqn() *string
	InferenceName() *string
	SetInferenceName(val *string)
	InferenceNameInput() *string
	InferenceVersion() *string
	SetInferenceVersion(val *string)
	InferenceVersionInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsFoundational() interface{}
	SetIsFoundational(val interface{})
	IsFoundationalInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	ParentUuid() *string
	SetParentUuid(val *string)
	ParentUuidInput() *string
	Provider() *string
	SetProvider(val *string)
	ProviderInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdatedAt() *string
	UploadComplete() interface{}
	SetUploadComplete(val interface{})
	UploadCompleteInput() interface{}
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	Usecases() *[]*string
	SetUsecases(val *[]*string)
	UsecasesInput() *[]*string
	Versions() DataDigitaloceanGenaiAgentModelVersionsList
	VersionsInput() interface{}
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
	PutAgreement(value interface{})
	PutVersions(value interface{})
	ResetAgreement()
	ResetInferenceName()
	ResetInferenceVersion()
	ResetIsFoundational()
	ResetName()
	ResetParentUuid()
	ResetProvider()
	ResetUploadComplete()
	ResetUrl()
	ResetUsecases()
	ResetVersions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDigitaloceanGenaiAgentModelOutputReference
type jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) Agreement() DataDigitaloceanGenaiAgentModelAgreementList {
	var returns DataDigitaloceanGenaiAgentModelAgreementList
	_jsii_.Get(
		j,
		"agreement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) AgreementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agreementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) InferenceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) InferenceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) InferenceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) InferenceVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) IsFoundational() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isFoundational",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) IsFoundationalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isFoundationalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ParentUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ParentUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) Provider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) UploadComplete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uploadComplete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) UploadCompleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uploadCompleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) Usecases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usecases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) UsecasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usecasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) Versions() DataDigitaloceanGenaiAgentModelVersionsList {
	var returns DataDigitaloceanGenaiAgentModelVersionsList
	_jsii_.Get(
		j,
		"versions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) VersionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionsInput",
		&returns,
	)
	return returns
}


func NewDataDigitaloceanGenaiAgentModelOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDigitaloceanGenaiAgentModelOutputReference {
	_init_.Initialize()

	if err := validateNewDataDigitaloceanGenaiAgentModelOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanGenaiAgent.DataDigitaloceanGenaiAgentModelOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDigitaloceanGenaiAgentModelOutputReference_Override(d DataDigitaloceanGenaiAgentModelOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanGenaiAgent.DataDigitaloceanGenaiAgentModelOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference)SetInferenceName(val *string) {
	if err := j.validateSetInferenceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceName",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference)SetInferenceVersion(val *string) {
	if err := j.validateSetInferenceVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceVersion",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference)SetIsFoundational(val interface{}) {
	if err := j.validateSetIsFoundationalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isFoundational",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference)SetParentUuid(val *string) {
	if err := j.validateSetParentUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentUuid",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference)SetProvider(val *string) {
	if err := j.validateSetProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference)SetUploadComplete(val interface{}) {
	if err := j.validateSetUploadCompleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadComplete",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference)SetUsecases(val *[]*string) {
	if err := j.validateSetUsecasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usecases",
		val,
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) PutAgreement(value interface{}) {
	if err := d.validatePutAgreementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAgreement",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) PutVersions(value interface{}) {
	if err := d.validatePutVersionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVersions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ResetAgreement() {
	_jsii_.InvokeVoid(
		d,
		"resetAgreement",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ResetInferenceName() {
	_jsii_.InvokeVoid(
		d,
		"resetInferenceName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ResetInferenceVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetInferenceVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ResetIsFoundational() {
	_jsii_.InvokeVoid(
		d,
		"resetIsFoundational",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ResetParentUuid() {
	_jsii_.InvokeVoid(
		d,
		"resetParentUuid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ResetProvider() {
	_jsii_.InvokeVoid(
		d,
		"resetProvider",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ResetUploadComplete() {
	_jsii_.InvokeVoid(
		d,
		"resetUploadComplete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ResetUsecases() {
	_jsii_.InvokeVoid(
		d,
		"resetUsecases",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ResetVersions() {
	_jsii_.InvokeVoid(
		d,
		"resetVersions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentModelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

