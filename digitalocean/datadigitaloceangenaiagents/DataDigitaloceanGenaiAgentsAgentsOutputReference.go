// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/datadigitaloceangenaiagents/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDigitaloceanGenaiAgentsAgentsOutputReference interface {
	cdktf.ComplexObject
	AgentGuardrail() DataDigitaloceanGenaiAgentsAgentsAgentGuardrailList
	AgentId() *string
	AnthropicApiKey() DataDigitaloceanGenaiAgentsAgentsAnthropicApiKeyList
	ApiKeyInfos() DataDigitaloceanGenaiAgentsAgentsApiKeyInfosList
	ApiKeys() DataDigitaloceanGenaiAgentsAgentsApiKeysList
	Chatbot() DataDigitaloceanGenaiAgentsAgentsChatbotList
	ChatbotIdentifiers() DataDigitaloceanGenaiAgentsAgentsChatbotIdentifiersList
	ChildAgents() DataDigitaloceanGenaiAgentsAgentsChildAgentsList
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
	Deployment() DataDigitaloceanGenaiAgentsAgentsDeploymentList
	Description() *string
	// Experimental.
	Fqn() *string
	Functions() DataDigitaloceanGenaiAgentsAgentsFunctionsList
	IfCase() *string
	Instruction() *string
	InternalValue() *DataDigitaloceanGenaiAgentsAgents
	SetInternalValue(val *DataDigitaloceanGenaiAgentsAgents)
	K() *float64
	KnowledgeBases() DataDigitaloceanGenaiAgentsAgentsKnowledgeBasesList
	MaxTokens() *float64
	Model() DataDigitaloceanGenaiAgentsAgentsModelList
	ModelUuid() *string
	Name() *string
	OpenAiApiKey() DataDigitaloceanGenaiAgentsAgentsOpenAiApiKeyList
	ParentAgents() DataDigitaloceanGenaiAgentsAgentsParentAgentsList
	ProjectId() *string
	Region() *string
	RetrievalMethod() *string
	RouteCreatedAt() *string
	RouteCreatedBy() *string
	RouteName() *string
	RouteUuid() *string
	Tags() *[]*string
	Temperature() *float64
	Template() DataDigitaloceanGenaiAgentsAgentsTemplateList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TopP() *float64
	UpdatedAt() *string
	Url() *string
	UserId() *string
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

// The jsii proxy struct for DataDigitaloceanGenaiAgentsAgentsOutputReference
type jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) AgentGuardrail() DataDigitaloceanGenaiAgentsAgentsAgentGuardrailList {
	var returns DataDigitaloceanGenaiAgentsAgentsAgentGuardrailList
	_jsii_.Get(
		j,
		"agentGuardrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) AgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) AnthropicApiKey() DataDigitaloceanGenaiAgentsAgentsAnthropicApiKeyList {
	var returns DataDigitaloceanGenaiAgentsAgentsAnthropicApiKeyList
	_jsii_.Get(
		j,
		"anthropicApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) ApiKeyInfos() DataDigitaloceanGenaiAgentsAgentsApiKeyInfosList {
	var returns DataDigitaloceanGenaiAgentsAgentsApiKeyInfosList
	_jsii_.Get(
		j,
		"apiKeyInfos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) ApiKeys() DataDigitaloceanGenaiAgentsAgentsApiKeysList {
	var returns DataDigitaloceanGenaiAgentsAgentsApiKeysList
	_jsii_.Get(
		j,
		"apiKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) Chatbot() DataDigitaloceanGenaiAgentsAgentsChatbotList {
	var returns DataDigitaloceanGenaiAgentsAgentsChatbotList
	_jsii_.Get(
		j,
		"chatbot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) ChatbotIdentifiers() DataDigitaloceanGenaiAgentsAgentsChatbotIdentifiersList {
	var returns DataDigitaloceanGenaiAgentsAgentsChatbotIdentifiersList
	_jsii_.Get(
		j,
		"chatbotIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) ChildAgents() DataDigitaloceanGenaiAgentsAgentsChildAgentsList {
	var returns DataDigitaloceanGenaiAgentsAgentsChildAgentsList
	_jsii_.Get(
		j,
		"childAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) Deployment() DataDigitaloceanGenaiAgentsAgentsDeploymentList {
	var returns DataDigitaloceanGenaiAgentsAgentsDeploymentList
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) Functions() DataDigitaloceanGenaiAgentsAgentsFunctionsList {
	var returns DataDigitaloceanGenaiAgentsAgentsFunctionsList
	_jsii_.Get(
		j,
		"functions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) IfCase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ifCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) Instruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) InternalValue() *DataDigitaloceanGenaiAgentsAgents {
	var returns *DataDigitaloceanGenaiAgentsAgents
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) K() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"k",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) KnowledgeBases() DataDigitaloceanGenaiAgentsAgentsKnowledgeBasesList {
	var returns DataDigitaloceanGenaiAgentsAgentsKnowledgeBasesList
	_jsii_.Get(
		j,
		"knowledgeBases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) MaxTokens() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) Model() DataDigitaloceanGenaiAgentsAgentsModelList {
	var returns DataDigitaloceanGenaiAgentsAgentsModelList
	_jsii_.Get(
		j,
		"model",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) ModelUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) OpenAiApiKey() DataDigitaloceanGenaiAgentsAgentsOpenAiApiKeyList {
	var returns DataDigitaloceanGenaiAgentsAgentsOpenAiApiKeyList
	_jsii_.Get(
		j,
		"openAiApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) ParentAgents() DataDigitaloceanGenaiAgentsAgentsParentAgentsList {
	var returns DataDigitaloceanGenaiAgentsAgentsParentAgentsList
	_jsii_.Get(
		j,
		"parentAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) RetrievalMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retrievalMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) RouteCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) RouteCreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeCreatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) RouteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) RouteUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) Temperature() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temperature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) Template() DataDigitaloceanGenaiAgentsAgentsTemplateList {
	var returns DataDigitaloceanGenaiAgentsAgentsTemplateList
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) TopP() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topP",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}


func NewDataDigitaloceanGenaiAgentsAgentsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDigitaloceanGenaiAgentsAgentsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDigitaloceanGenaiAgentsAgentsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanGenaiAgents.DataDigitaloceanGenaiAgentsAgentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDigitaloceanGenaiAgentsAgentsOutputReference_Override(d DataDigitaloceanGenaiAgentsAgentsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanGenaiAgents.DataDigitaloceanGenaiAgentsAgentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference)SetInternalValue(val *DataDigitaloceanGenaiAgentsAgents) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgentsAgentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

