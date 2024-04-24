// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasepostgresqlconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v11/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v11/databasepostgresqlconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabasePostgresqlConfigPgbouncerOutputReference interface {
	cdktf.ComplexObject
	AutodbIdleTimeout() *float64
	SetAutodbIdleTimeout(val *float64)
	AutodbIdleTimeoutInput() *float64
	AutodbMaxDbConnections() *float64
	SetAutodbMaxDbConnections(val *float64)
	AutodbMaxDbConnectionsInput() *float64
	AutodbPoolMode() *string
	SetAutodbPoolMode(val *string)
	AutodbPoolModeInput() *string
	AutodbPoolSize() *float64
	SetAutodbPoolSize(val *float64)
	AutodbPoolSizeInput() *float64
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
	// Experimental.
	Fqn() *string
	IgnoreStartupParameters() *[]*string
	SetIgnoreStartupParameters(val *[]*string)
	IgnoreStartupParametersInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MinPoolSize() *float64
	SetMinPoolSize(val *float64)
	MinPoolSizeInput() *float64
	ServerIdleTimeout() *float64
	SetServerIdleTimeout(val *float64)
	ServerIdleTimeoutInput() *float64
	ServerLifetime() *float64
	SetServerLifetime(val *float64)
	ServerLifetimeInput() *float64
	ServerResetQueryAlways() interface{}
	SetServerResetQueryAlways(val interface{})
	ServerResetQueryAlwaysInput() interface{}
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
	ResetAutodbIdleTimeout()
	ResetAutodbMaxDbConnections()
	ResetAutodbPoolMode()
	ResetAutodbPoolSize()
	ResetIgnoreStartupParameters()
	ResetMinPoolSize()
	ResetServerIdleTimeout()
	ResetServerLifetime()
	ResetServerResetQueryAlways()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabasePostgresqlConfigPgbouncerOutputReference
type jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) AutodbIdleTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autodbIdleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) AutodbIdleTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autodbIdleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) AutodbMaxDbConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autodbMaxDbConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) AutodbMaxDbConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autodbMaxDbConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) AutodbPoolMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autodbPoolMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) AutodbPoolModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autodbPoolModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) AutodbPoolSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autodbPoolSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) AutodbPoolSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autodbPoolSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) IgnoreStartupParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreStartupParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) IgnoreStartupParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreStartupParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) MinPoolSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minPoolSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) MinPoolSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minPoolSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ServerIdleTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serverIdleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ServerIdleTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serverIdleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ServerLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serverLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ServerLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serverLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ServerResetQueryAlways() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverResetQueryAlways",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ServerResetQueryAlwaysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverResetQueryAlwaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatabasePostgresqlConfigPgbouncerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DatabasePostgresqlConfigPgbouncerOutputReference {
	_init_.Initialize()

	if err := validateNewDatabasePostgresqlConfigPgbouncerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.databasePostgresqlConfig.DatabasePostgresqlConfigPgbouncerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDatabasePostgresqlConfigPgbouncerOutputReference_Override(d DatabasePostgresqlConfigPgbouncerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.databasePostgresqlConfig.DatabasePostgresqlConfigPgbouncerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference)SetAutodbIdleTimeout(val *float64) {
	if err := j.validateSetAutodbIdleTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autodbIdleTimeout",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference)SetAutodbMaxDbConnections(val *float64) {
	if err := j.validateSetAutodbMaxDbConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autodbMaxDbConnections",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference)SetAutodbPoolMode(val *string) {
	if err := j.validateSetAutodbPoolModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autodbPoolMode",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference)SetAutodbPoolSize(val *float64) {
	if err := j.validateSetAutodbPoolSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autodbPoolSize",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference)SetIgnoreStartupParameters(val *[]*string) {
	if err := j.validateSetIgnoreStartupParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreStartupParameters",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference)SetMinPoolSize(val *float64) {
	if err := j.validateSetMinPoolSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minPoolSize",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference)SetServerIdleTimeout(val *float64) {
	if err := j.validateSetServerIdleTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverIdleTimeout",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference)SetServerLifetime(val *float64) {
	if err := j.validateSetServerLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverLifetime",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference)SetServerResetQueryAlways(val interface{}) {
	if err := j.validateSetServerResetQueryAlwaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverResetQueryAlways",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ResetAutodbIdleTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetAutodbIdleTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ResetAutodbMaxDbConnections() {
	_jsii_.InvokeVoid(
		d,
		"resetAutodbMaxDbConnections",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ResetAutodbPoolMode() {
	_jsii_.InvokeVoid(
		d,
		"resetAutodbPoolMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ResetAutodbPoolSize() {
	_jsii_.InvokeVoid(
		d,
		"resetAutodbPoolSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ResetIgnoreStartupParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetIgnoreStartupParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ResetMinPoolSize() {
	_jsii_.InvokeVoid(
		d,
		"resetMinPoolSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ResetServerIdleTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetServerIdleTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ResetServerLifetime() {
	_jsii_.InvokeVoid(
		d,
		"resetServerLifetime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ResetServerResetQueryAlways() {
	_jsii_.InvokeVoid(
		d,
		"resetServerResetQueryAlways",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabasePostgresqlConfigPgbouncerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

