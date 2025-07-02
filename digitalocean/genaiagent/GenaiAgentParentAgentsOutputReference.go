// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/genaiagent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GenaiAgentParentAgentsOutputReference interface {
	cdktf.ComplexObject
	AgentId() *string
	AnthropicApiKey() GenaiAgentParentAgentsAnthropicApiKeyList
	AnthropicApiKeyInput() interface{}
	ApiKeyInfos() GenaiAgentParentAgentsApiKeyInfosList
	ApiKeyInfosInput() interface{}
	ApiKeys() GenaiAgentParentAgentsApiKeysList
	ApiKeysInput() interface{}
	Chatbot() GenaiAgentParentAgentsChatbotList
	ChatbotIdentifiers() GenaiAgentParentAgentsChatbotIdentifiersList
	ChatbotIdentifiersInput() interface{}
	ChatbotInput() interface{}
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
	Deployment() GenaiAgentParentAgentsDeploymentList
	DeploymentInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	Instruction() *string
	SetInstruction(val *string)
	InstructionInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ModelUuid() *string
	SetModelUuid(val *string)
	ModelUuidInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
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
	PutAnthropicApiKey(value interface{})
	PutApiKeyInfos(value interface{})
	PutApiKeys(value interface{})
	PutChatbot(value interface{})
	PutChatbotIdentifiers(value interface{})
	PutDeployment(value interface{})
	ResetAnthropicApiKey()
	ResetApiKeyInfos()
	ResetApiKeys()
	ResetChatbot()
	ResetChatbotIdentifiers()
	ResetDeployment()
	ResetDescription()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GenaiAgentParentAgentsOutputReference
type jsiiProxy_GenaiAgentParentAgentsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) AgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) AnthropicApiKey() GenaiAgentParentAgentsAnthropicApiKeyList {
	var returns GenaiAgentParentAgentsAnthropicApiKeyList
	_jsii_.Get(
		j,
		"anthropicApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) AnthropicApiKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anthropicApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) ApiKeyInfos() GenaiAgentParentAgentsApiKeyInfosList {
	var returns GenaiAgentParentAgentsApiKeyInfosList
	_jsii_.Get(
		j,
		"apiKeyInfos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) ApiKeyInfosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyInfosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) ApiKeys() GenaiAgentParentAgentsApiKeysList {
	var returns GenaiAgentParentAgentsApiKeysList
	_jsii_.Get(
		j,
		"apiKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) ApiKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) Chatbot() GenaiAgentParentAgentsChatbotList {
	var returns GenaiAgentParentAgentsChatbotList
	_jsii_.Get(
		j,
		"chatbot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) ChatbotIdentifiers() GenaiAgentParentAgentsChatbotIdentifiersList {
	var returns GenaiAgentParentAgentsChatbotIdentifiersList
	_jsii_.Get(
		j,
		"chatbotIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) ChatbotIdentifiersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chatbotIdentifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) ChatbotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chatbotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) Deployment() GenaiAgentParentAgentsDeploymentList {
	var returns GenaiAgentParentAgentsDeploymentList
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) DeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) Instruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) InstructionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instructionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) ModelUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) ModelUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGenaiAgentParentAgentsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GenaiAgentParentAgentsOutputReference {
	_init_.Initialize()

	if err := validateNewGenaiAgentParentAgentsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GenaiAgentParentAgentsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiAgent.GenaiAgentParentAgentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGenaiAgentParentAgentsOutputReference_Override(g GenaiAgentParentAgentsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiAgent.GenaiAgentParentAgentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference)SetInstruction(val *string) {
	if err := j.validateSetInstructionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instruction",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference)SetModelUuid(val *string) {
	if err := j.validateSetModelUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelUuid",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GenaiAgentParentAgentsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) PutAnthropicApiKey(value interface{}) {
	if err := g.validatePutAnthropicApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAnthropicApiKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) PutApiKeyInfos(value interface{}) {
	if err := g.validatePutApiKeyInfosParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApiKeyInfos",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) PutApiKeys(value interface{}) {
	if err := g.validatePutApiKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApiKeys",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) PutChatbot(value interface{}) {
	if err := g.validatePutChatbotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putChatbot",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) PutChatbotIdentifiers(value interface{}) {
	if err := g.validatePutChatbotIdentifiersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putChatbotIdentifiers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) PutDeployment(value interface{}) {
	if err := g.validatePutDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeployment",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) ResetAnthropicApiKey() {
	_jsii_.InvokeVoid(
		g,
		"resetAnthropicApiKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) ResetApiKeyInfos() {
	_jsii_.InvokeVoid(
		g,
		"resetApiKeyInfos",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) ResetApiKeys() {
	_jsii_.InvokeVoid(
		g,
		"resetApiKeys",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) ResetChatbot() {
	_jsii_.InvokeVoid(
		g,
		"resetChatbot",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) ResetChatbotIdentifiers() {
	_jsii_.InvokeVoid(
		g,
		"resetChatbotIdentifiers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) ResetDeployment() {
	_jsii_.InvokeVoid(
		g,
		"resetDeployment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GenaiAgentParentAgentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

