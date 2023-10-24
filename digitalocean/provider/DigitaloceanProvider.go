// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v10/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.31.0/docs digitalocean}.
type DigitaloceanProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	ApiEndpoint() *string
	SetApiEndpoint(val *string)
	ApiEndpointInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpRetryMax() *float64
	SetHttpRetryMax(val *float64)
	HttpRetryMaxInput() *float64
	HttpRetryWaitMax() *float64
	SetHttpRetryWaitMax(val *float64)
	HttpRetryWaitMaxInput() *float64
	HttpRetryWaitMin() *float64
	SetHttpRetryWaitMin(val *float64)
	HttpRetryWaitMinInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	RequestsPerSecond() *float64
	SetRequestsPerSecond(val *float64)
	RequestsPerSecondInput() *float64
	SpacesAccessId() *string
	SetSpacesAccessId(val *string)
	SpacesAccessIdInput() *string
	SpacesEndpoint() *string
	SetSpacesEndpoint(val *string)
	SpacesEndpointInput() *string
	SpacesSecretKey() *string
	SetSpacesSecretKey(val *string)
	SpacesSecretKeyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetApiEndpoint()
	ResetHttpRetryMax()
	ResetHttpRetryWaitMax()
	ResetHttpRetryWaitMin()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequestsPerSecond()
	ResetSpacesAccessId()
	ResetSpacesEndpoint()
	ResetSpacesSecretKey()
	ResetToken()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DigitaloceanProvider
type jsiiProxy_DigitaloceanProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_DigitaloceanProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) ApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) ApiEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) HttpRetryMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpRetryMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) HttpRetryMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpRetryMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) HttpRetryWaitMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpRetryWaitMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) HttpRetryWaitMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpRetryWaitMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) HttpRetryWaitMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpRetryWaitMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) HttpRetryWaitMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpRetryWaitMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) RequestsPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) RequestsPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) SpacesAccessId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spacesAccessId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) SpacesAccessIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spacesAccessIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) SpacesEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spacesEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) SpacesEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spacesEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) SpacesSecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spacesSecretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) SpacesSecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spacesSecretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DigitaloceanProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.31.0/docs digitalocean} Resource.
func NewDigitaloceanProvider(scope constructs.Construct, id *string, config *DigitaloceanProviderConfig) DigitaloceanProvider {
	_init_.Initialize()

	if err := validateNewDigitaloceanProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DigitaloceanProvider{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.provider.DigitaloceanProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.31.0/docs digitalocean} Resource.
func NewDigitaloceanProvider_Override(d DigitaloceanProvider, scope constructs.Construct, id *string, config *DigitaloceanProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.provider.DigitaloceanProvider",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DigitaloceanProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_DigitaloceanProvider)SetApiEndpoint(val *string) {
	_jsii_.Set(
		j,
		"apiEndpoint",
		val,
	)
}

func (j *jsiiProxy_DigitaloceanProvider)SetHttpRetryMax(val *float64) {
	_jsii_.Set(
		j,
		"httpRetryMax",
		val,
	)
}

func (j *jsiiProxy_DigitaloceanProvider)SetHttpRetryWaitMax(val *float64) {
	_jsii_.Set(
		j,
		"httpRetryWaitMax",
		val,
	)
}

func (j *jsiiProxy_DigitaloceanProvider)SetHttpRetryWaitMin(val *float64) {
	_jsii_.Set(
		j,
		"httpRetryWaitMin",
		val,
	)
}

func (j *jsiiProxy_DigitaloceanProvider)SetRequestsPerSecond(val *float64) {
	_jsii_.Set(
		j,
		"requestsPerSecond",
		val,
	)
}

func (j *jsiiProxy_DigitaloceanProvider)SetSpacesAccessId(val *string) {
	_jsii_.Set(
		j,
		"spacesAccessId",
		val,
	)
}

func (j *jsiiProxy_DigitaloceanProvider)SetSpacesEndpoint(val *string) {
	_jsii_.Set(
		j,
		"spacesEndpoint",
		val,
	)
}

func (j *jsiiProxy_DigitaloceanProvider)SetSpacesSecretKey(val *string) {
	_jsii_.Set(
		j,
		"spacesSecretKey",
		val,
	)
}

func (j *jsiiProxy_DigitaloceanProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

// Generates CDKTF code for importing a DigitaloceanProvider resource upon running "cdktf plan <stack-name>".
func DigitaloceanProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDigitaloceanProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.provider.DigitaloceanProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DigitaloceanProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDigitaloceanProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.provider.DigitaloceanProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DigitaloceanProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDigitaloceanProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.provider.DigitaloceanProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DigitaloceanProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDigitaloceanProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.provider.DigitaloceanProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DigitaloceanProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-digitalocean.provider.DigitaloceanProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DigitaloceanProvider) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DigitaloceanProvider) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DigitaloceanProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		d,
		"resetAlias",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DigitaloceanProvider) ResetApiEndpoint() {
	_jsii_.InvokeVoid(
		d,
		"resetApiEndpoint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DigitaloceanProvider) ResetHttpRetryMax() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpRetryMax",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DigitaloceanProvider) ResetHttpRetryWaitMax() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpRetryWaitMax",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DigitaloceanProvider) ResetHttpRetryWaitMin() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpRetryWaitMin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DigitaloceanProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DigitaloceanProvider) ResetRequestsPerSecond() {
	_jsii_.InvokeVoid(
		d,
		"resetRequestsPerSecond",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DigitaloceanProvider) ResetSpacesAccessId() {
	_jsii_.InvokeVoid(
		d,
		"resetSpacesAccessId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DigitaloceanProvider) ResetSpacesEndpoint() {
	_jsii_.InvokeVoid(
		d,
		"resetSpacesEndpoint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DigitaloceanProvider) ResetSpacesSecretKey() {
	_jsii_.InvokeVoid(
		d,
		"resetSpacesSecretKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DigitaloceanProvider) ResetToken() {
	_jsii_.InvokeVoid(
		d,
		"resetToken",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DigitaloceanProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DigitaloceanProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DigitaloceanProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DigitaloceanProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

