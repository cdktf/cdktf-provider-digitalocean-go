// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/genaiagent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GenaiAgentChildAgentsChatbotOutputReference interface {
	cdktf.ComplexObject
	ButtonBackgroundColor() *string
	SetButtonBackgroundColor(val *string)
	ButtonBackgroundColorInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Logo() *string
	SetLogo(val *string)
	LogoInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	PrimaryColor() *string
	SetPrimaryColor(val *string)
	PrimaryColorInput() *string
	SecondaryColor() *string
	SetSecondaryColor(val *string)
	SecondaryColorInput() *string
	StartingMessage() *string
	SetStartingMessage(val *string)
	StartingMessageInput() *string
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
	ResetButtonBackgroundColor()
	ResetLogo()
	ResetName()
	ResetPrimaryColor()
	ResetSecondaryColor()
	ResetStartingMessage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GenaiAgentChildAgentsChatbotOutputReference
type jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) ButtonBackgroundColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buttonBackgroundColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) ButtonBackgroundColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buttonBackgroundColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) Logo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) LogoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) PrimaryColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) PrimaryColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) SecondaryColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) SecondaryColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) StartingMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) StartingMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGenaiAgentChildAgentsChatbotOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GenaiAgentChildAgentsChatbotOutputReference {
	_init_.Initialize()

	if err := validateNewGenaiAgentChildAgentsChatbotOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiAgent.GenaiAgentChildAgentsChatbotOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGenaiAgentChildAgentsChatbotOutputReference_Override(g GenaiAgentChildAgentsChatbotOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiAgent.GenaiAgentChildAgentsChatbotOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference)SetButtonBackgroundColor(val *string) {
	if err := j.validateSetButtonBackgroundColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buttonBackgroundColor",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference)SetLogo(val *string) {
	if err := j.validateSetLogoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logo",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference)SetPrimaryColor(val *string) {
	if err := j.validateSetPrimaryColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryColor",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference)SetSecondaryColor(val *string) {
	if err := j.validateSetSecondaryColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryColor",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference)SetStartingMessage(val *string) {
	if err := j.validateSetStartingMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingMessage",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) ResetButtonBackgroundColor() {
	_jsii_.InvokeVoid(
		g,
		"resetButtonBackgroundColor",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) ResetLogo() {
	_jsii_.InvokeVoid(
		g,
		"resetLogo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) ResetPrimaryColor() {
	_jsii_.InvokeVoid(
		g,
		"resetPrimaryColor",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) ResetSecondaryColor() {
	_jsii_.InvokeVoid(
		g,
		"resetSecondaryColor",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) ResetStartingMessage() {
	_jsii_.InvokeVoid(
		g,
		"resetStartingMessage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentChildAgentsChatbotOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

