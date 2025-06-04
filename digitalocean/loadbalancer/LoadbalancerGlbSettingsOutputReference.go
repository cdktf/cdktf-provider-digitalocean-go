// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/loadbalancer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerGlbSettingsOutputReference interface {
	cdktf.ComplexObject
	Cdn() LoadbalancerGlbSettingsCdnOutputReference
	CdnInput() *LoadbalancerGlbSettingsCdn
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
	FailoverThreshold() *float64
	SetFailoverThreshold(val *float64)
	FailoverThresholdInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *LoadbalancerGlbSettings
	SetInternalValue(val *LoadbalancerGlbSettings)
	RegionPriorities() *map[string]*float64
	SetRegionPriorities(val *map[string]*float64)
	RegionPrioritiesInput() *map[string]*float64
	TargetPort() *float64
	SetTargetPort(val *float64)
	TargetPortInput() *float64
	TargetProtocol() *string
	SetTargetProtocol(val *string)
	TargetProtocolInput() *string
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
	PutCdn(value *LoadbalancerGlbSettingsCdn)
	ResetCdn()
	ResetFailoverThreshold()
	ResetRegionPriorities()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LoadbalancerGlbSettingsOutputReference
type jsiiProxy_LoadbalancerGlbSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference) Cdn() LoadbalancerGlbSettingsCdnOutputReference {
	var returns LoadbalancerGlbSettingsCdnOutputReference
	_jsii_.Get(
		j,
		"cdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference) CdnInput() *LoadbalancerGlbSettingsCdn {
	var returns *LoadbalancerGlbSettingsCdn
	_jsii_.Get(
		j,
		"cdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference) FailoverThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failoverThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference) FailoverThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failoverThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference) InternalValue() *LoadbalancerGlbSettings {
	var returns *LoadbalancerGlbSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference) RegionPriorities() *map[string]*float64 {
	var returns *map[string]*float64
	_jsii_.Get(
		j,
		"regionPriorities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference) RegionPrioritiesInput() *map[string]*float64 {
	var returns *map[string]*float64
	_jsii_.Get(
		j,
		"regionPrioritiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference) TargetPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference) TargetPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference) TargetProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference) TargetProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLoadbalancerGlbSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LoadbalancerGlbSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewLoadbalancerGlbSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LoadbalancerGlbSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.loadbalancer.LoadbalancerGlbSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLoadbalancerGlbSettingsOutputReference_Override(l LoadbalancerGlbSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.loadbalancer.LoadbalancerGlbSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference)SetFailoverThreshold(val *float64) {
	if err := j.validateSetFailoverThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failoverThreshold",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference)SetInternalValue(val *LoadbalancerGlbSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference)SetRegionPriorities(val *map[string]*float64) {
	if err := j.validateSetRegionPrioritiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionPriorities",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference)SetTargetPort(val *float64) {
	if err := j.validateSetTargetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetPort",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference)SetTargetProtocol(val *string) {
	if err := j.validateSetTargetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetProtocol",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerGlbSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) PutCdn(value *LoadbalancerGlbSettingsCdn) {
	if err := l.validatePutCdnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCdn",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) ResetCdn() {
	_jsii_.InvokeVoid(
		l,
		"resetCdn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) ResetFailoverThreshold() {
	_jsii_.InvokeVoid(
		l,
		"resetFailoverThreshold",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) ResetRegionPriorities() {
	_jsii_.InvokeVoid(
		l,
		"resetRegionPriorities",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerGlbSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

