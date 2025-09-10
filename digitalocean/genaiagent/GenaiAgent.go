// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-digitalocean-go/digitalocean/v13/genaiagent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/resources/genai_agent digitalocean_genai_agent}.
type GenaiAgent interface {
	cdktf.TerraformResource
	AgentGuardrail() GenaiAgentAgentGuardrailList
	AgentGuardrailInput() interface{}
	AnthropicApiKey() GenaiAgentAnthropicApiKeyList
	AnthropicApiKeyInput() interface{}
	AnthropicKeyUuid() *string
	SetAnthropicKeyUuid(val *string)
	AnthropicKeyUuidInput() *string
	ApiKeyInfos() GenaiAgentApiKeyInfosList
	ApiKeyInfosInput() interface{}
	ApiKeys() GenaiAgentApiKeysList
	ApiKeysInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Chatbot() GenaiAgentChatbotList
	ChatbotIdentifiers() GenaiAgentChatbotIdentifiersList
	ChatbotIdentifiersInput() interface{}
	ChatbotInput() interface{}
	ChildAgents() GenaiAgentChildAgentsList
	ChildAgentsInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	SetCreatedAt(val *string)
	CreatedAtInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Deployment() GenaiAgentDeploymentList
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
	Functions() GenaiAgentFunctionsList
	FunctionsInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	IfCase() *string
	SetIfCase(val *string)
	IfCaseInput() *string
	Instruction() *string
	SetInstruction(val *string)
	InstructionInput() *string
	K() *float64
	SetK(val *float64)
	KInput() *float64
	KnowledgeBases() GenaiAgentKnowledgeBasesList
	KnowledgeBasesInput() interface{}
	KnowledgeBaseUuid() *[]*string
	SetKnowledgeBaseUuid(val *[]*string)
	KnowledgeBaseUuidInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxTokens() *float64
	SetMaxTokens(val *float64)
	MaxTokensInput() *float64
	Model() GenaiAgentModelList
	ModelInput() interface{}
	ModelUuid() *string
	SetModelUuid(val *string)
	ModelUuidInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OpenAiApiKey() GenaiAgentOpenAiApiKeyList
	OpenAiApiKeyInput() interface{}
	OpenAiKeyUuid() *string
	SetOpenAiKeyUuid(val *string)
	OpenAiKeyUuidInput() *string
	ParentAgents() GenaiAgentParentAgentsList
	ParentAgentsInput() interface{}
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	ProvideCitations() interface{}
	SetProvideCitations(val interface{})
	ProvideCitationsInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
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
	Template() GenaiAgentTemplateList
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
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAgentGuardrail(value interface{})
	PutAnthropicApiKey(value interface{})
	PutApiKeyInfos(value interface{})
	PutApiKeys(value interface{})
	PutChatbot(value interface{})
	PutChatbotIdentifiers(value interface{})
	PutChildAgents(value interface{})
	PutDeployment(value interface{})
	PutFunctions(value interface{})
	PutKnowledgeBases(value interface{})
	PutModel(value interface{})
	PutOpenAiApiKey(value interface{})
	PutParentAgents(value interface{})
	PutTemplate(value interface{})
	ResetAgentGuardrail()
	ResetAnthropicApiKey()
	ResetAnthropicKeyUuid()
	ResetApiKeyInfos()
	ResetApiKeys()
	ResetChatbot()
	ResetChatbotIdentifiers()
	ResetChildAgents()
	ResetCreatedAt()
	ResetDeployment()
	ResetDescription()
	ResetFunctions()
	ResetId()
	ResetIfCase()
	ResetK()
	ResetKnowledgeBases()
	ResetKnowledgeBaseUuid()
	ResetMaxTokens()
	ResetModel()
	ResetOpenAiApiKey()
	ResetOpenAiKeyUuid()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentAgents()
	ResetProvideCitations()
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

// The jsii proxy struct for GenaiAgent
type jsiiProxy_GenaiAgent struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GenaiAgent) AgentGuardrail() GenaiAgentAgentGuardrailList {
	var returns GenaiAgentAgentGuardrailList
	_jsii_.Get(
		j,
		"agentGuardrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) AgentGuardrailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentGuardrailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) AnthropicApiKey() GenaiAgentAnthropicApiKeyList {
	var returns GenaiAgentAnthropicApiKeyList
	_jsii_.Get(
		j,
		"anthropicApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) AnthropicApiKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anthropicApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) AnthropicKeyUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anthropicKeyUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) AnthropicKeyUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anthropicKeyUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ApiKeyInfos() GenaiAgentApiKeyInfosList {
	var returns GenaiAgentApiKeyInfosList
	_jsii_.Get(
		j,
		"apiKeyInfos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ApiKeyInfosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyInfosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ApiKeys() GenaiAgentApiKeysList {
	var returns GenaiAgentApiKeysList
	_jsii_.Get(
		j,
		"apiKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ApiKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Chatbot() GenaiAgentChatbotList {
	var returns GenaiAgentChatbotList
	_jsii_.Get(
		j,
		"chatbot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ChatbotIdentifiers() GenaiAgentChatbotIdentifiersList {
	var returns GenaiAgentChatbotIdentifiersList
	_jsii_.Get(
		j,
		"chatbotIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ChatbotIdentifiersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chatbotIdentifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ChatbotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chatbotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ChildAgents() GenaiAgentChildAgentsList {
	var returns GenaiAgentChildAgentsList
	_jsii_.Get(
		j,
		"childAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ChildAgentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"childAgentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) CreatedAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Deployment() GenaiAgentDeploymentList {
	var returns GenaiAgentDeploymentList
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) DeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Functions() GenaiAgentFunctionsList {
	var returns GenaiAgentFunctionsList
	_jsii_.Get(
		j,
		"functions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) FunctionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) IfCase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ifCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) IfCaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ifCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Instruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) InstructionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instructionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) K() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"k",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) KInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"kInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) KnowledgeBases() GenaiAgentKnowledgeBasesList {
	var returns GenaiAgentKnowledgeBasesList
	_jsii_.Get(
		j,
		"knowledgeBases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) KnowledgeBasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"knowledgeBasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) KnowledgeBaseUuid() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"knowledgeBaseUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) KnowledgeBaseUuidInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"knowledgeBaseUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) MaxTokens() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) MaxTokensInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Model() GenaiAgentModelList {
	var returns GenaiAgentModelList
	_jsii_.Get(
		j,
		"model",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ModelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ModelUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ModelUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) OpenAiApiKey() GenaiAgentOpenAiApiKeyList {
	var returns GenaiAgentOpenAiApiKeyList
	_jsii_.Get(
		j,
		"openAiApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) OpenAiApiKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openAiApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) OpenAiKeyUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openAiKeyUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) OpenAiKeyUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openAiKeyUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ParentAgents() GenaiAgentParentAgentsList {
	var returns GenaiAgentParentAgentsList
	_jsii_.Get(
		j,
		"parentAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ParentAgentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parentAgentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ProvideCitations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provideCitations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) ProvideCitationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provideCitationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) RetrievalMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retrievalMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) RetrievalMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retrievalMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) RouteCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) RouteCreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeCreatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) RouteCreatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeCreatedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) RouteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) RouteNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) RouteUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) RouteUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Temperature() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temperature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) TemperatureInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temperatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Template() GenaiAgentTemplateList {
	var returns GenaiAgentTemplateList
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) TemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"templateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) TopP() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topP",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) TopPInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topPInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenaiAgent) UserIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/resources/genai_agent digitalocean_genai_agent} Resource.
func NewGenaiAgent(scope constructs.Construct, id *string, config *GenaiAgentConfig) GenaiAgent {
	_init_.Initialize()

	if err := validateNewGenaiAgentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GenaiAgent{}

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiAgent.GenaiAgent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/resources/genai_agent digitalocean_genai_agent} Resource.
func NewGenaiAgent_Override(g GenaiAgent, scope constructs.Construct, id *string, config *GenaiAgentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-digitalocean.genaiAgent.GenaiAgent",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GenaiAgent)SetAnthropicKeyUuid(val *string) {
	if err := j.validateSetAnthropicKeyUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anthropicKeyUuid",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetCreatedAt(val *string) {
	if err := j.validateSetCreatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetIfCase(val *string) {
	if err := j.validateSetIfCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ifCase",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetInstruction(val *string) {
	if err := j.validateSetInstructionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instruction",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetK(val *float64) {
	if err := j.validateSetKParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"k",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetKnowledgeBaseUuid(val *[]*string) {
	if err := j.validateSetKnowledgeBaseUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"knowledgeBaseUuid",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetMaxTokens(val *float64) {
	if err := j.validateSetMaxTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTokens",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetModelUuid(val *string) {
	if err := j.validateSetModelUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelUuid",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetOpenAiKeyUuid(val *string) {
	if err := j.validateSetOpenAiKeyUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openAiKeyUuid",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetProvideCitations(val interface{}) {
	if err := j.validateSetProvideCitationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provideCitations",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetRetrievalMethod(val *string) {
	if err := j.validateSetRetrievalMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retrievalMethod",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetRouteCreatedBy(val *string) {
	if err := j.validateSetRouteCreatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeCreatedBy",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetRouteName(val *string) {
	if err := j.validateSetRouteNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeName",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetRouteUuid(val *string) {
	if err := j.validateSetRouteUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeUuid",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetTemperature(val *float64) {
	if err := j.validateSetTemperatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"temperature",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetTopP(val *float64) {
	if err := j.validateSetTopPParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topP",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_GenaiAgent)SetUserId(val *string) {
	if err := j.validateSetUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userId",
		val,
	)
}

// Generates CDKTF code for importing a GenaiAgent resource upon running "cdktf plan <stack-name>".
func GenaiAgent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGenaiAgent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.genaiAgent.GenaiAgent",
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
func GenaiAgent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGenaiAgent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.genaiAgent.GenaiAgent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GenaiAgent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGenaiAgent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.genaiAgent.GenaiAgent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GenaiAgent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGenaiAgent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-digitalocean.genaiAgent.GenaiAgent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GenaiAgent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-digitalocean.genaiAgent.GenaiAgent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GenaiAgent) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GenaiAgent) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GenaiAgent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GenaiAgent) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GenaiAgent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GenaiAgent) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GenaiAgent) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GenaiAgent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GenaiAgent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GenaiAgent) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GenaiAgent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GenaiAgent) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgent) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GenaiAgent) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgent) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GenaiAgent) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GenaiAgent) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GenaiAgent) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GenaiAgent) PutAgentGuardrail(value interface{}) {
	if err := g.validatePutAgentGuardrailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAgentGuardrail",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgent) PutAnthropicApiKey(value interface{}) {
	if err := g.validatePutAnthropicApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAnthropicApiKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgent) PutApiKeyInfos(value interface{}) {
	if err := g.validatePutApiKeyInfosParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApiKeyInfos",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgent) PutApiKeys(value interface{}) {
	if err := g.validatePutApiKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApiKeys",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgent) PutChatbot(value interface{}) {
	if err := g.validatePutChatbotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putChatbot",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgent) PutChatbotIdentifiers(value interface{}) {
	if err := g.validatePutChatbotIdentifiersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putChatbotIdentifiers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgent) PutChildAgents(value interface{}) {
	if err := g.validatePutChildAgentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putChildAgents",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgent) PutDeployment(value interface{}) {
	if err := g.validatePutDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeployment",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgent) PutFunctions(value interface{}) {
	if err := g.validatePutFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFunctions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgent) PutKnowledgeBases(value interface{}) {
	if err := g.validatePutKnowledgeBasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKnowledgeBases",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgent) PutModel(value interface{}) {
	if err := g.validatePutModelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putModel",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgent) PutOpenAiApiKey(value interface{}) {
	if err := g.validatePutOpenAiApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOpenAiApiKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgent) PutParentAgents(value interface{}) {
	if err := g.validatePutParentAgentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putParentAgents",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgent) PutTemplate(value interface{}) {
	if err := g.validatePutTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTemplate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GenaiAgent) ResetAgentGuardrail() {
	_jsii_.InvokeVoid(
		g,
		"resetAgentGuardrail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetAnthropicApiKey() {
	_jsii_.InvokeVoid(
		g,
		"resetAnthropicApiKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetAnthropicKeyUuid() {
	_jsii_.InvokeVoid(
		g,
		"resetAnthropicKeyUuid",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetApiKeyInfos() {
	_jsii_.InvokeVoid(
		g,
		"resetApiKeyInfos",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetApiKeys() {
	_jsii_.InvokeVoid(
		g,
		"resetApiKeys",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetChatbot() {
	_jsii_.InvokeVoid(
		g,
		"resetChatbot",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetChatbotIdentifiers() {
	_jsii_.InvokeVoid(
		g,
		"resetChatbotIdentifiers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetChildAgents() {
	_jsii_.InvokeVoid(
		g,
		"resetChildAgents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		g,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetDeployment() {
	_jsii_.InvokeVoid(
		g,
		"resetDeployment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetFunctions() {
	_jsii_.InvokeVoid(
		g,
		"resetFunctions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetIfCase() {
	_jsii_.InvokeVoid(
		g,
		"resetIfCase",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetK() {
	_jsii_.InvokeVoid(
		g,
		"resetK",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetKnowledgeBases() {
	_jsii_.InvokeVoid(
		g,
		"resetKnowledgeBases",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetKnowledgeBaseUuid() {
	_jsii_.InvokeVoid(
		g,
		"resetKnowledgeBaseUuid",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetMaxTokens() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxTokens",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetModel() {
	_jsii_.InvokeVoid(
		g,
		"resetModel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetOpenAiApiKey() {
	_jsii_.InvokeVoid(
		g,
		"resetOpenAiApiKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetOpenAiKeyUuid() {
	_jsii_.InvokeVoid(
		g,
		"resetOpenAiKeyUuid",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetParentAgents() {
	_jsii_.InvokeVoid(
		g,
		"resetParentAgents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetProvideCitations() {
	_jsii_.InvokeVoid(
		g,
		"resetProvideCitations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetRetrievalMethod() {
	_jsii_.InvokeVoid(
		g,
		"resetRetrievalMethod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetRouteCreatedBy() {
	_jsii_.InvokeVoid(
		g,
		"resetRouteCreatedBy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetRouteName() {
	_jsii_.InvokeVoid(
		g,
		"resetRouteName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetRouteUuid() {
	_jsii_.InvokeVoid(
		g,
		"resetRouteUuid",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetTemperature() {
	_jsii_.InvokeVoid(
		g,
		"resetTemperature",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetTemplate() {
	_jsii_.InvokeVoid(
		g,
		"resetTemplate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetTopP() {
	_jsii_.InvokeVoid(
		g,
		"resetTopP",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) ResetUserId() {
	_jsii_.InvokeVoid(
		g,
		"resetUserId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GenaiAgent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenaiAgent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

