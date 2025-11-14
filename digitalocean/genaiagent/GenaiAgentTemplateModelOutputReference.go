// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/genaiagent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GenaiAgentTemplateModelOutputReference interface {
	cdktf.ComplexObject
	Agreement() GenaiAgentTemplateModelAgreementList
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
	Versions() GenaiAgentTemplateModelVersionsList
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

// The jsii proxy struct for GenaiAgentTemplateModelOutputReference
type jsiiProxy_GenaiAgentTemplateModelOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) Agreement() GenaiAgentTemplateModelAgreementList {
	var returns GenaiAgentTemplateModelAgreementList
	_jsii_.Get(
		j,
		"agreement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) AgreementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agreementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) InferenceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) InferenceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) InferenceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) InferenceVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) IsFoundational() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isFoundational",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) IsFoundationalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isFoundationalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) ParentUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) ParentUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) Provider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) ProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) UploadComplete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uploadComplete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) UploadCompleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uploadCompleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) Usecases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usecases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) UsecasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usecasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) Versions() GenaiAgentTemplateModelVersionsList {
	var returns GenaiAgentTemplateModelVersionsList
	_jsii_.Get(
		j,
		"versions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference) VersionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionsInput",
		&returns,
	)
	return returns
}


func NewGenaiAgentTemplateModelOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GenaiAgentTemplateModelOutputReference {
	_init_.Initialize()

	if err := validateNewGenaiAgentTemplateModelOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GenaiAgentTemplateModelOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiAgent.GenaiAgentTemplateModelOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGenaiAgentTemplateModelOutputReference_Override(g GenaiAgentTemplateModelOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiAgent.GenaiAgentTemplateModelOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference)SetInferenceName(val *string) {
	if err := j.validateSetInferenceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceName",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference)SetInferenceVersion(val *string) {
	if err := j.validateSetInferenceVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceVersion",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference)SetIsFoundational(val interface{}) {
	if err := j.validateSetIsFoundationalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isFoundational",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference)SetParentUuid(val *string) {
	if err := j.validateSetParentUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentUuid",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference)SetProvider(val *string) {
	if err := j.validateSetProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference)SetUploadComplete(val interface{}) {
	if err := j.validateSetUploadCompleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadComplete",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentTemplateModelOutputReference)SetUsecases(val *[]*string) {
	if err := j.validateSetUsecasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usecases",
		val,
	)
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) PutAgreement(value interface{}) {
	if err := g.validatePutAgreementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAgreement",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) PutVersions(value interface{}) {
	if err := g.validatePutVersionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVersions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) ResetAgreement() {
	_jsii_.InvokeVoid(
		g,
		"resetAgreement",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) ResetInferenceName() {
	_jsii_.InvokeVoid(
		g,
		"resetInferenceName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) ResetInferenceVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetInferenceVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) ResetIsFoundational() {
	_jsii_.InvokeVoid(
		g,
		"resetIsFoundational",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) ResetParentUuid() {
	_jsii_.InvokeVoid(
		g,
		"resetParentUuid",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) ResetProvider() {
	_jsii_.InvokeVoid(
		g,
		"resetProvider",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) ResetUploadComplete() {
	_jsii_.InvokeVoid(
		g,
		"resetUploadComplete",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) ResetUsecases() {
	_jsii_.InvokeVoid(
		g,
		"resetUsecases",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) ResetVersions() {
	_jsii_.InvokeVoid(
		g,
		"resetVersions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentTemplateModelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

