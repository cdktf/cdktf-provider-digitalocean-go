// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/datadigitaloceangenaiagent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDigitaloceanGenaiAgentTemplateModelOutputReference interface {
	cdktf.ComplexObject
	Agreement() DataDigitaloceanGenaiAgentTemplateModelAgreementList
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
	Versions() DataDigitaloceanGenaiAgentTemplateModelVersionsList
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
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
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDigitaloceanGenaiAgentTemplateModelOutputReference
type jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) Agreement() DataDigitaloceanGenaiAgentTemplateModelAgreementList {
	var returns DataDigitaloceanGenaiAgentTemplateModelAgreementList
	_jsii_.Get(
		j,
		"agreement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) AgreementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agreementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) InferenceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) InferenceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) InferenceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) InferenceVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) IsFoundational() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isFoundational",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) IsFoundationalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isFoundationalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ParentUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ParentUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) Provider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) UploadComplete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uploadComplete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) UploadCompleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uploadCompleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) Usecases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usecases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) UsecasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usecasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) Versions() DataDigitaloceanGenaiAgentTemplateModelVersionsList {
	var returns DataDigitaloceanGenaiAgentTemplateModelVersionsList
	_jsii_.Get(
		j,
		"versions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) VersionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionsInput",
		&returns,
	)
	return returns
}


func NewDataDigitaloceanGenaiAgentTemplateModelOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDigitaloceanGenaiAgentTemplateModelOutputReference {
	_init_.Initialize()

	if err := validateNewDataDigitaloceanGenaiAgentTemplateModelOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanGenaiAgent.DataDigitaloceanGenaiAgentTemplateModelOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDigitaloceanGenaiAgentTemplateModelOutputReference_Override(d DataDigitaloceanGenaiAgentTemplateModelOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanGenaiAgent.DataDigitaloceanGenaiAgentTemplateModelOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference)SetInferenceName(val *string) {
	if err := j.validateSetInferenceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceName",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference)SetInferenceVersion(val *string) {
	if err := j.validateSetInferenceVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceVersion",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference)SetIsFoundational(val interface{}) {
	if err := j.validateSetIsFoundationalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isFoundational",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference)SetParentUuid(val *string) {
	if err := j.validateSetParentUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentUuid",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference)SetProvider(val *string) {
	if err := j.validateSetProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference)SetUploadComplete(val interface{}) {
	if err := j.validateSetUploadCompleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadComplete",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference)SetUsecases(val *[]*string) {
	if err := j.validateSetUsecasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usecases",
		val,
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) PutAgreement(value interface{}) {
	if err := d.validatePutAgreementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAgreement",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) PutVersions(value interface{}) {
	if err := d.validatePutVersionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVersions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ResetAgreement() {
	_jsii_.InvokeVoid(
		d,
		"resetAgreement",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ResetInferenceName() {
	_jsii_.InvokeVoid(
		d,
		"resetInferenceName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ResetInferenceVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetInferenceVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ResetIsFoundational() {
	_jsii_.InvokeVoid(
		d,
		"resetIsFoundational",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ResetParentUuid() {
	_jsii_.InvokeVoid(
		d,
		"resetParentUuid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ResetProvider() {
	_jsii_.InvokeVoid(
		d,
		"resetProvider",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ResetUploadComplete() {
	_jsii_.InvokeVoid(
		d,
		"resetUploadComplete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ResetUsecases() {
	_jsii_.InvokeVoid(
		d,
		"resetUsecases",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ResetVersions() {
	_jsii_.InvokeVoid(
		d,
		"resetVersions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentTemplateModelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

