// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/app/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppSpecJobImageOutputReference interface {
	cdktf.ComplexObject
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
	DeployOnPush() AppSpecJobImageDeployOnPushList
	DeployOnPushInput() interface{}
	Digest() *string
	SetDigest(val *string)
	DigestInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *AppSpecJobImage
	SetInternalValue(val *AppSpecJobImage)
	Registry() *string
	SetRegistry(val *string)
	RegistryCredentials() *string
	SetRegistryCredentials(val *string)
	RegistryCredentialsInput() *string
	RegistryInput() *string
	RegistryType() *string
	SetRegistryType(val *string)
	RegistryTypeInput() *string
	Repository() *string
	SetRepository(val *string)
	RepositoryInput() *string
	Tag() *string
	SetTag(val *string)
	TagInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutDeployOnPush(value interface{})
	ResetDeployOnPush()
	ResetDigest()
	ResetRegistry()
	ResetRegistryCredentials()
	ResetTag()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppSpecJobImageOutputReference
type jsiiProxy_AppSpecJobImageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) DeployOnPush() AppSpecJobImageDeployOnPushList {
	var returns AppSpecJobImageDeployOnPushList
	_jsii_.Get(
		j,
		"deployOnPush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) DeployOnPushInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deployOnPushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) Digest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) DigestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) InternalValue() *AppSpecJobImage {
	var returns *AppSpecJobImage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) Registry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) RegistryCredentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) RegistryCredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) RegistryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) RegistryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) RegistryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) Repository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) RepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) Tag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) TagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecJobImageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppSpecJobImageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppSpecJobImageOutputReference {
	_init_.Initialize()

	if err := validateNewAppSpecJobImageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSpecJobImageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.app.AppSpecJobImageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppSpecJobImageOutputReference_Override(a AppSpecJobImageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.app.AppSpecJobImageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppSpecJobImageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobImageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobImageOutputReference)SetDigest(val *string) {
	if err := j.validateSetDigestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"digest",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobImageOutputReference)SetInternalValue(val *AppSpecJobImage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobImageOutputReference)SetRegistry(val *string) {
	if err := j.validateSetRegistryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registry",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobImageOutputReference)SetRegistryCredentials(val *string) {
	if err := j.validateSetRegistryCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registryCredentials",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobImageOutputReference)SetRegistryType(val *string) {
	if err := j.validateSetRegistryTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registryType",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobImageOutputReference)SetRepository(val *string) {
	if err := j.validateSetRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repository",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobImageOutputReference)SetTag(val *string) {
	if err := j.validateSetTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tag",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobImageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppSpecJobImageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) PutDeployOnPush(value interface{}) {
	if err := a.validatePutDeployOnPushParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeployOnPush",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) ResetDeployOnPush() {
	_jsii_.InvokeVoid(
		a,
		"resetDeployOnPush",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) ResetDigest() {
	_jsii_.InvokeVoid(
		a,
		"resetDigest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) ResetRegistry() {
	_jsii_.InvokeVoid(
		a,
		"resetRegistry",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) ResetRegistryCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetRegistryCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) ResetTag() {
	_jsii_.InvokeVoid(
		a,
		"resetTag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecJobImageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

