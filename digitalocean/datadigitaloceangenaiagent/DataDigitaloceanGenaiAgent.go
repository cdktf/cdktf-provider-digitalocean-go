// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/datadigitaloceangenaiagent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/data-sources/genai_agent digitalocean_genai_agent}.
type DataDigitaloceanGenaiAgent interface {
	cdktf.TerraformDataSource
	AgentGuardrail() DataDigitaloceanGenaiAgentAgentGuardrailList
	AgentGuardrailInput() interface{}
	AgentId() *string
	SetAgentId(val *string)
	AgentIdInput() *string
	AnthropicApiKey() DataDigitaloceanGenaiAgentAnthropicApiKeyList
	AnthropicApiKeyInput() interface{}
	ApiKeyInfos() DataDigitaloceanGenaiAgentApiKeyInfosList
	ApiKeyInfosInput() interface{}
	ApiKeys() DataDigitaloceanGenaiAgentApiKeysList
	ApiKeysInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Chatbot() DataDigitaloceanGenaiAgentChatbotList
	ChatbotIdentifiers() DataDigitaloceanGenaiAgentChatbotIdentifiersList
	ChatbotIdentifiersInput() interface{}
	ChatbotInput() interface{}
	ChildAgents() DataDigitaloceanGenaiAgentChildAgentsList
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Deployment() DataDigitaloceanGenaiAgentDeploymentList
	DeploymentInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Functions() DataDigitaloceanGenaiAgentFunctionsList
	FunctionsInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	IfCase() *string
	SetIfCase(val *string)
	IfCaseInput() *string
	Instruction() *string
	K() *float64
	SetK(val *float64)
	KInput() *float64
	KnowledgeBases() DataDigitaloceanGenaiAgentKnowledgeBasesList
	KnowledgeBasesInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxTokens() *float64
	SetMaxTokens(val *float64)
	MaxTokensInput() *float64
	Model() DataDigitaloceanGenaiAgentModelList
	ModelInput() interface{}
	ModelUuid() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	OpenAiApiKey() DataDigitaloceanGenaiAgentOpenAiApiKeyList
	OpenAiApiKeyInput() interface{}
	ParentAgents() DataDigitaloceanGenaiAgentParentAgentsList
	ProjectId() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	RetrievalMethod() *string
	SetRetrievalMethod(val *string)
	RetrievalMethodInput() *string
	RouteCreatedAt() *string
	RouteCreatedBy() *string
	SetRouteCreatedBy(val *string)
	RouteCreatedByInput() *string
	RouteName() *string
	SetRouteName(val *string)
	RouteNameInput() *string
	RouteUuid() *string
	SetRouteUuid(val *string)
	RouteUuidInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	Temperature() *float64
	SetTemperature(val *float64)
	TemperatureInput() *float64
	Template() DataDigitaloceanGenaiAgentTemplateList
	TemplateInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TopP() *float64
	SetTopP(val *float64)
	TopPInput() *float64
	UpdatedAt() *string
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	UserId() *string
	SetUserId(val *string)
	UserIdInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAgentGuardrail(value interface{})
	PutAnthropicApiKey(value interface{})
	PutApiKeyInfos(value interface{})
	PutApiKeys(value interface{})
	PutChatbot(value interface{})
	PutChatbotIdentifiers(value interface{})
	PutDeployment(value interface{})
	PutFunctions(value interface{})
	PutKnowledgeBases(value interface{})
	PutModel(value interface{})
	PutOpenAiApiKey(value interface{})
	PutTemplate(value interface{})
	ResetAgentGuardrail()
	ResetAnthropicApiKey()
	ResetApiKeyInfos()
	ResetApiKeys()
	ResetChatbot()
	ResetChatbotIdentifiers()
	ResetDeployment()
	ResetDescription()
	ResetFunctions()
	ResetId()
	ResetIfCase()
	ResetK()
	ResetKnowledgeBases()
	ResetMaxTokens()
	ResetModel()
	ResetOpenAiApiKey()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRetrievalMethod()
	ResetRouteCreatedBy()
	ResetRouteName()
	ResetRouteUuid()
	ResetTags()
	ResetTemperature()
	ResetTemplate()
	ResetTopP()
	ResetUrl()
	ResetUserId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataDigitaloceanGenaiAgent
type jsiiProxy_DataDigitaloceanGenaiAgent struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) AgentGuardrail() DataDigitaloceanGenaiAgentAgentGuardrailList {
	var returns DataDigitaloceanGenaiAgentAgentGuardrailList
	_jsii_.Get(
		j,
		"agentGuardrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) AgentGuardrailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentGuardrailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) AgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) AgentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) AnthropicApiKey() DataDigitaloceanGenaiAgentAnthropicApiKeyList {
	var returns DataDigitaloceanGenaiAgentAnthropicApiKeyList
	_jsii_.Get(
		j,
		"anthropicApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) AnthropicApiKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anthropicApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) ApiKeyInfos() DataDigitaloceanGenaiAgentApiKeyInfosList {
	var returns DataDigitaloceanGenaiAgentApiKeyInfosList
	_jsii_.Get(
		j,
		"apiKeyInfos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) ApiKeyInfosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyInfosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) ApiKeys() DataDigitaloceanGenaiAgentApiKeysList {
	var returns DataDigitaloceanGenaiAgentApiKeysList
	_jsii_.Get(
		j,
		"apiKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) ApiKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Chatbot() DataDigitaloceanGenaiAgentChatbotList {
	var returns DataDigitaloceanGenaiAgentChatbotList
	_jsii_.Get(
		j,
		"chatbot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) ChatbotIdentifiers() DataDigitaloceanGenaiAgentChatbotIdentifiersList {
	var returns DataDigitaloceanGenaiAgentChatbotIdentifiersList
	_jsii_.Get(
		j,
		"chatbotIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) ChatbotIdentifiersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chatbotIdentifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) ChatbotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chatbotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) ChildAgents() DataDigitaloceanGenaiAgentChildAgentsList {
	var returns DataDigitaloceanGenaiAgentChildAgentsList
	_jsii_.Get(
		j,
		"childAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Deployment() DataDigitaloceanGenaiAgentDeploymentList {
	var returns DataDigitaloceanGenaiAgentDeploymentList
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) DeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Functions() DataDigitaloceanGenaiAgentFunctionsList {
	var returns DataDigitaloceanGenaiAgentFunctionsList
	_jsii_.Get(
		j,
		"functions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) FunctionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) IfCase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ifCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) IfCaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ifCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Instruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) K() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"k",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) KInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"kInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) KnowledgeBases() DataDigitaloceanGenaiAgentKnowledgeBasesList {
	var returns DataDigitaloceanGenaiAgentKnowledgeBasesList
	_jsii_.Get(
		j,
		"knowledgeBases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) KnowledgeBasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"knowledgeBasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) MaxTokens() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) MaxTokensInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Model() DataDigitaloceanGenaiAgentModelList {
	var returns DataDigitaloceanGenaiAgentModelList
	_jsii_.Get(
		j,
		"model",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) ModelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) ModelUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) OpenAiApiKey() DataDigitaloceanGenaiAgentOpenAiApiKeyList {
	var returns DataDigitaloceanGenaiAgentOpenAiApiKeyList
	_jsii_.Get(
		j,
		"openAiApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) OpenAiApiKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openAiApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) ParentAgents() DataDigitaloceanGenaiAgentParentAgentsList {
	var returns DataDigitaloceanGenaiAgentParentAgentsList
	_jsii_.Get(
		j,
		"parentAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) RetrievalMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retrievalMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) RetrievalMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retrievalMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) RouteCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) RouteCreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeCreatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) RouteCreatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeCreatedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) RouteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) RouteNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) RouteUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) RouteUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Temperature() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temperature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) TemperatureInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temperatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Template() DataDigitaloceanGenaiAgentTemplateList {
	var returns DataDigitaloceanGenaiAgentTemplateList
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) TemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"templateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) TopP() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topP",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) TopPInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topPInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent) UserIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/data-sources/genai_agent digitalocean_genai_agent} Data Source.
func NewDataDigitaloceanGenaiAgent(scope constructs.Construct, id *string, config *DataDigitaloceanGenaiAgentConfig) DataDigitaloceanGenaiAgent {
	_init_.Initialize()

	if err := validateNewDataDigitaloceanGenaiAgentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDigitaloceanGenaiAgent{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanGenaiAgent.DataDigitaloceanGenaiAgent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/data-sources/genai_agent digitalocean_genai_agent} Data Source.
func NewDataDigitaloceanGenaiAgent_Override(d DataDigitaloceanGenaiAgent, scope constructs.Construct, id *string, config *DataDigitaloceanGenaiAgentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.dataDigitaloceanGenaiAgent.DataDigitaloceanGenaiAgent",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetAgentId(val *string) {
	if err := j.validateSetAgentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentId",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetIfCase(val *string) {
	if err := j.validateSetIfCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ifCase",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetK(val *float64) {
	if err := j.validateSetKParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"k",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetMaxTokens(val *float64) {
	if err := j.validateSetMaxTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTokens",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetRetrievalMethod(val *string) {
	if err := j.validateSetRetrievalMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retrievalMethod",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetRouteCreatedBy(val *string) {
	if err := j.validateSetRouteCreatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeCreatedBy",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetRouteName(val *string) {
	if err := j.validateSetRouteNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeName",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetRouteUuid(val *string) {
	if err := j.validateSetRouteUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeUuid",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetTemperature(val *float64) {
	if err := j.validateSetTemperatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"temperature",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetTopP(val *float64) {
	if err := j.validateSetTopPParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topP",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_DataDigitaloceanGenaiAgent)SetUserId(val *string) {
	if err := j.validateSetUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userId",
		val,
	)
}

// Generates CDKTF code for importing a DataDigitaloceanGenaiAgent resource upon running "cdktf plan <stack-name>".
func DataDigitaloceanGenaiAgent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataDigitaloceanGenaiAgent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.dataDigitaloceanGenaiAgent.DataDigitaloceanGenaiAgent",
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
func DataDigitaloceanGenaiAgent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDigitaloceanGenaiAgent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.dataDigitaloceanGenaiAgent.DataDigitaloceanGenaiAgent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDigitaloceanGenaiAgent_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDigitaloceanGenaiAgent_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.dataDigitaloceanGenaiAgent.DataDigitaloceanGenaiAgent",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDigitaloceanGenaiAgent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDigitaloceanGenaiAgent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.dataDigitaloceanGenaiAgent.DataDigitaloceanGenaiAgent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDigitaloceanGenaiAgent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-digitalocean.dataDigitaloceanGenaiAgent.DataDigitaloceanGenaiAgent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) PutAgentGuardrail(value interface{}) {
	if err := d.validatePutAgentGuardrailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAgentGuardrail",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) PutAnthropicApiKey(value interface{}) {
	if err := d.validatePutAnthropicApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAnthropicApiKey",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) PutApiKeyInfos(value interface{}) {
	if err := d.validatePutApiKeyInfosParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApiKeyInfos",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) PutApiKeys(value interface{}) {
	if err := d.validatePutApiKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApiKeys",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) PutChatbot(value interface{}) {
	if err := d.validatePutChatbotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putChatbot",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) PutChatbotIdentifiers(value interface{}) {
	if err := d.validatePutChatbotIdentifiersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putChatbotIdentifiers",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) PutDeployment(value interface{}) {
	if err := d.validatePutDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDeployment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) PutFunctions(value interface{}) {
	if err := d.validatePutFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFunctions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) PutKnowledgeBases(value interface{}) {
	if err := d.validatePutKnowledgeBasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKnowledgeBases",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) PutModel(value interface{}) {
	if err := d.validatePutModelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putModel",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) PutOpenAiApiKey(value interface{}) {
	if err := d.validatePutOpenAiApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOpenAiApiKey",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) PutTemplate(value interface{}) {
	if err := d.validatePutTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTemplate",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetAgentGuardrail() {
	_jsii_.InvokeVoid(
		d,
		"resetAgentGuardrail",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetAnthropicApiKey() {
	_jsii_.InvokeVoid(
		d,
		"resetAnthropicApiKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetApiKeyInfos() {
	_jsii_.InvokeVoid(
		d,
		"resetApiKeyInfos",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetApiKeys() {
	_jsii_.InvokeVoid(
		d,
		"resetApiKeys",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetChatbot() {
	_jsii_.InvokeVoid(
		d,
		"resetChatbot",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetChatbotIdentifiers() {
	_jsii_.InvokeVoid(
		d,
		"resetChatbotIdentifiers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetDeployment() {
	_jsii_.InvokeVoid(
		d,
		"resetDeployment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetFunctions() {
	_jsii_.InvokeVoid(
		d,
		"resetFunctions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetIfCase() {
	_jsii_.InvokeVoid(
		d,
		"resetIfCase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetK() {
	_jsii_.InvokeVoid(
		d,
		"resetK",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetKnowledgeBases() {
	_jsii_.InvokeVoid(
		d,
		"resetKnowledgeBases",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetMaxTokens() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxTokens",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetModel() {
	_jsii_.InvokeVoid(
		d,
		"resetModel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetOpenAiApiKey() {
	_jsii_.InvokeVoid(
		d,
		"resetOpenAiApiKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetRetrievalMethod() {
	_jsii_.InvokeVoid(
		d,
		"resetRetrievalMethod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetRouteCreatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetRouteCreatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetRouteName() {
	_jsii_.InvokeVoid(
		d,
		"resetRouteName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetRouteUuid() {
	_jsii_.InvokeVoid(
		d,
		"resetRouteUuid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetTemperature() {
	_jsii_.InvokeVoid(
		d,
		"resetTemperature",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetTemplate() {
	_jsii_.InvokeVoid(
		d,
		"resetTemplate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetTopP() {
	_jsii_.InvokeVoid(
		d,
		"resetTopP",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ResetUserId() {
	_jsii_.InvokeVoid(
		d,
		"resetUserId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDigitaloceanGenaiAgent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

