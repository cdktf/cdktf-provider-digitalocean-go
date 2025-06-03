// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v12/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v12/app/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppSpecFunctionLogDestinationOutputReference interface {
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
	Datadog() AppSpecFunctionLogDestinationDatadogOutputReference
	DatadogInput() *AppSpecFunctionLogDestinationDatadog
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Logtail() AppSpecFunctionLogDestinationLogtailOutputReference
	LogtailInput() *AppSpecFunctionLogDestinationLogtail
	Name() *string
	SetName(val *string)
	NameInput() *string
	OpenSearch() AppSpecFunctionLogDestinationOpenSearchOutputReference
	OpenSearchInput() *AppSpecFunctionLogDestinationOpenSearch
	Papertrail() AppSpecFunctionLogDestinationPapertrailOutputReference
	PapertrailInput() *AppSpecFunctionLogDestinationPapertrail
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutDatadog(value *AppSpecFunctionLogDestinationDatadog)
	PutLogtail(value *AppSpecFunctionLogDestinationLogtail)
	PutOpenSearch(value *AppSpecFunctionLogDestinationOpenSearch)
	PutPapertrail(value *AppSpecFunctionLogDestinationPapertrail)
	ResetDatadog()
	ResetLogtail()
	ResetOpenSearch()
	ResetPapertrail()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppSpecFunctionLogDestinationOutputReference
type jsiiProxy_AppSpecFunctionLogDestinationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) Datadog() AppSpecFunctionLogDestinationDatadogOutputReference {
	var returns AppSpecFunctionLogDestinationDatadogOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) DatadogInput() *AppSpecFunctionLogDestinationDatadog {
	var returns *AppSpecFunctionLogDestinationDatadog
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) Logtail() AppSpecFunctionLogDestinationLogtailOutputReference {
	var returns AppSpecFunctionLogDestinationLogtailOutputReference
	_jsii_.Get(
		j,
		"logtail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) LogtailInput() *AppSpecFunctionLogDestinationLogtail {
	var returns *AppSpecFunctionLogDestinationLogtail
	_jsii_.Get(
		j,
		"logtailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) OpenSearch() AppSpecFunctionLogDestinationOpenSearchOutputReference {
	var returns AppSpecFunctionLogDestinationOpenSearchOutputReference
	_jsii_.Get(
		j,
		"openSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) OpenSearchInput() *AppSpecFunctionLogDestinationOpenSearch {
	var returns *AppSpecFunctionLogDestinationOpenSearch
	_jsii_.Get(
		j,
		"openSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) Papertrail() AppSpecFunctionLogDestinationPapertrailOutputReference {
	var returns AppSpecFunctionLogDestinationPapertrailOutputReference
	_jsii_.Get(
		j,
		"papertrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) PapertrailInput() *AppSpecFunctionLogDestinationPapertrail {
	var returns *AppSpecFunctionLogDestinationPapertrail
	_jsii_.Get(
		j,
		"papertrailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppSpecFunctionLogDestinationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppSpecFunctionLogDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewAppSpecFunctionLogDestinationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSpecFunctionLogDestinationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.app.AppSpecFunctionLogDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppSpecFunctionLogDestinationOutputReference_Override(a AppSpecFunctionLogDestinationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.app.AppSpecFunctionLogDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppSpecFunctionLogDestinationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) PutDatadog(value *AppSpecFunctionLogDestinationDatadog) {
	if err := a.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatadog",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) PutLogtail(value *AppSpecFunctionLogDestinationLogtail) {
	if err := a.validatePutLogtailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogtail",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) PutOpenSearch(value *AppSpecFunctionLogDestinationOpenSearch) {
	if err := a.validatePutOpenSearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpenSearch",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) PutPapertrail(value *AppSpecFunctionLogDestinationPapertrail) {
	if err := a.validatePutPapertrailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPapertrail",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		a,
		"resetDatadog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) ResetLogtail() {
	_jsii_.InvokeVoid(
		a,
		"resetLogtail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) ResetOpenSearch() {
	_jsii_.InvokeVoid(
		a,
		"resetOpenSearch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) ResetPapertrail() {
	_jsii_.InvokeVoid(
		a,
		"resetPapertrail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSpecFunctionLogDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

