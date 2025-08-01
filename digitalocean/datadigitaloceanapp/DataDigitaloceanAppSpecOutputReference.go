// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceanapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/datadigitaloceanapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDigitaloceanAppSpecOutputReference interface {
	cdktf.ComplexObject
	Alert() DataDigitaloceanAppSpecAlertList
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
	Database() DataDigitaloceanAppSpecDatabaseList
	DisableEdgeCache() cdktf.IResolvable
	DisableEmailObfuscation() cdktf.IResolvable
	Domain() DataDigitaloceanAppSpecDomainList
	Domains() *[]*string
	Egress() DataDigitaloceanAppSpecEgressList
	EnhancedThreatControlEnabled() cdktf.IResolvable
	Env() DataDigitaloceanAppSpecEnvList
	Features() *[]*string
	// Experimental.
	Fqn() *string
	Function() DataDigitaloceanAppSpecFunctionList
	Ingress() DataDigitaloceanAppSpecIngressList
	InternalValue() *DataDigitaloceanAppSpec
	SetInternalValue(val *DataDigitaloceanAppSpec)
	Job() DataDigitaloceanAppSpecJobList
	Name() *string
	Region() *string
	Service() DataDigitaloceanAppSpecServiceList
	StaticSite() DataDigitaloceanAppSpecStaticSiteList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Vpc() DataDigitaloceanAppSpecVpcList
	Worker() DataDigitaloceanAppSpecWorkerList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDigitaloceanAppSpecOutputReference
type jsiiProxy_DataDigitaloceanAppSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) Alert() DataDigitaloceanAppSpecAlertList {
	var returns DataDigitaloceanAppSpecAlertList
	_jsii_.Get(
		j,
		"alert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) Database() DataDigitaloceanAppSpecDatabaseList {
	var returns DataDigitaloceanAppSpecDatabaseList
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) DisableEdgeCache() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableEdgeCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) DisableEmailObfuscation() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableEmailObfuscation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) Domain() DataDigitaloceanAppSpecDomainList {
	var returns DataDigitaloceanAppSpecDomainList
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) Domains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) Egress() DataDigitaloceanAppSpecEgressList {
	var returns DataDigitaloceanAppSpecEgressList
	_jsii_.Get(
		j,
		"egress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) EnhancedThreatControlEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enhancedThreatControlEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) Env() DataDigitaloceanAppSpecEnvList {
	var returns DataDigitaloceanAppSpecEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) Features() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"features",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) Function() DataDigitaloceanAppSpecFunctionList {
	var returns DataDigitaloceanAppSpecFunctionList
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) Ingress() DataDigitaloceanAppSpecIngressList {
	var returns DataDigitaloceanAppSpecIngressList
	_jsii_.Get(
		j,
		"ingress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) InternalValue() *DataDigitaloceanAppSpec {
	var returns *DataDigitaloceanAppSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) Job() DataDigitaloceanAppSpecJobList {
	var returns DataDigitaloceanAppSpecJobList
	_jsii_.Get(
		j,
		"job",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) Service() DataDigitaloceanAppSpecServiceList {
	var returns DataDigitaloceanAppSpecServiceList
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) StaticSite() DataDigitaloceanAppSpecStaticSiteList {
	var returns DataDigitaloceanAppSpecStaticSiteList
	_jsii_.Get(
		j,
		"staticSite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) Vpc() DataDigitaloceanAppSpecVpcList {
	var returns DataDigitaloceanAppSpecVpcList
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference) Worker() DataDigitaloceanAppSpecWorkerList {
	var returns DataDigitaloceanAppSpecWorkerList
	_jsii_.Get(
		j,
		"worker",
		&returns,
	)
	return returns
}


func NewDataDigitaloceanAppSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDigitaloceanAppSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataDigitaloceanAppSpecOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDigitaloceanAppSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanApp.DataDigitaloceanAppSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDigitaloceanAppSpecOutputReference_Override(d DataDigitaloceanAppSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanApp.DataDigitaloceanAppSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference)SetInternalValue(val *DataDigitaloceanAppSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanAppSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDigitaloceanAppSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanAppSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDigitaloceanAppSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDigitaloceanAppSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDigitaloceanAppSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDigitaloceanAppSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDigitaloceanAppSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanAppSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanAppSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDigitaloceanAppSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDigitaloceanAppSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanAppSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDigitaloceanAppSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDigitaloceanAppSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

