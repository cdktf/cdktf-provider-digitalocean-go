// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dropletautoscale

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/dropletautoscale/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DropletAutoscaleConfigAOutputReference interface {
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
	CooldownMinutes() *float64
	SetCooldownMinutes(val *float64)
	CooldownMinutesInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DropletAutoscaleConfigA
	SetInternalValue(val *DropletAutoscaleConfigA)
	MaxInstances() *float64
	SetMaxInstances(val *float64)
	MaxInstancesInput() *float64
	MinInstances() *float64
	SetMinInstances(val *float64)
	MinInstancesInput() *float64
	TargetCpuUtilization() *float64
	SetTargetCpuUtilization(val *float64)
	TargetCpuUtilizationInput() *float64
	TargetMemoryUtilization() *float64
	SetTargetMemoryUtilization(val *float64)
	TargetMemoryUtilizationInput() *float64
	TargetNumberInstances() *float64
	SetTargetNumberInstances(val *float64)
	TargetNumberInstancesInput() *float64
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
	ResetCooldownMinutes()
	ResetMaxInstances()
	ResetMinInstances()
	ResetTargetCpuUtilization()
	ResetTargetMemoryUtilization()
	ResetTargetNumberInstances()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DropletAutoscaleConfigAOutputReference
type jsiiProxy_DropletAutoscaleConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) CooldownMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldownMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) CooldownMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldownMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) InternalValue() *DropletAutoscaleConfigA {
	var returns *DropletAutoscaleConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) MaxInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) MaxInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) MinInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) MinInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) TargetCpuUtilization() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetCpuUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) TargetCpuUtilizationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetCpuUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) TargetMemoryUtilization() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetMemoryUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) TargetMemoryUtilizationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetMemoryUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) TargetNumberInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetNumberInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) TargetNumberInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetNumberInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDropletAutoscaleConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DropletAutoscaleConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewDropletAutoscaleConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DropletAutoscaleConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dropletAutoscale.DropletAutoscaleConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDropletAutoscaleConfigAOutputReference_Override(d DropletAutoscaleConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dropletAutoscale.DropletAutoscaleConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference)SetCooldownMinutes(val *float64) {
	if err := j.validateSetCooldownMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cooldownMinutes",
		val,
	)
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference)SetInternalValue(val *DropletAutoscaleConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference)SetMaxInstances(val *float64) {
	if err := j.validateSetMaxInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstances",
		val,
	)
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference)SetMinInstances(val *float64) {
	if err := j.validateSetMinInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minInstances",
		val,
	)
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference)SetTargetCpuUtilization(val *float64) {
	if err := j.validateSetTargetCpuUtilizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCpuUtilization",
		val,
	)
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference)SetTargetMemoryUtilization(val *float64) {
	if err := j.validateSetTargetMemoryUtilizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetMemoryUtilization",
		val,
	)
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference)SetTargetNumberInstances(val *float64) {
	if err := j.validateSetTargetNumberInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetNumberInstances",
		val,
	)
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DropletAutoscaleConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) ResetCooldownMinutes() {
	_jsii_.InvokeVoid(
		d,
		"resetCooldownMinutes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) ResetMaxInstances() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxInstances",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) ResetMinInstances() {
	_jsii_.InvokeVoid(
		d,
		"resetMinInstances",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) ResetTargetCpuUtilization() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetCpuUtilization",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) ResetTargetMemoryUtilization() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetMemoryUtilization",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) ResetTargetNumberInstances() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetNumberInstances",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DropletAutoscaleConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

