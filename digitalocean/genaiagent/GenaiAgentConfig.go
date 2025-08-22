// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GenaiAgentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Instruction for the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#instruction GenaiAgent#instruction}
	Instruction *string `field:"required" json:"instruction" yaml:"instruction"`
	// Model UUID of the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#model_uuid GenaiAgent#model_uuid}
	ModelUuid *string `field:"required" json:"modelUuid" yaml:"modelUuid"`
	// Name of the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#name GenaiAgent#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Project ID of the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#project_id GenaiAgent#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Region where the Agent is deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#region GenaiAgent#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// agent_guardrail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#agent_guardrail GenaiAgent#agent_guardrail}
	AgentGuardrail interface{} `field:"optional" json:"agentGuardrail" yaml:"agentGuardrail"`
	// anthropic_api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#anthropic_api_key GenaiAgent#anthropic_api_key}
	AnthropicApiKey interface{} `field:"optional" json:"anthropicApiKey" yaml:"anthropicApiKey"`
	// Optional Anthropic API key ID to use with Anthropic models.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#anthropic_key_uuid GenaiAgent#anthropic_key_uuid}
	AnthropicKeyUuid *string `field:"optional" json:"anthropicKeyUuid" yaml:"anthropicKeyUuid"`
	// api_key_infos block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#api_key_infos GenaiAgent#api_key_infos}
	ApiKeyInfos interface{} `field:"optional" json:"apiKeyInfos" yaml:"apiKeyInfos"`
	// api_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#api_keys GenaiAgent#api_keys}
	ApiKeys interface{} `field:"optional" json:"apiKeys" yaml:"apiKeys"`
	// chatbot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#chatbot GenaiAgent#chatbot}
	Chatbot interface{} `field:"optional" json:"chatbot" yaml:"chatbot"`
	// chatbot_identifiers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#chatbot_identifiers GenaiAgent#chatbot_identifiers}
	ChatbotIdentifiers interface{} `field:"optional" json:"chatbotIdentifiers" yaml:"chatbotIdentifiers"`
	// child_agents block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#child_agents GenaiAgent#child_agents}
	ChildAgents interface{} `field:"optional" json:"childAgents" yaml:"childAgents"`
	// Timestamp when the Agent was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#created_at GenaiAgent#created_at}
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// deployment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#deployment GenaiAgent#deployment}
	Deployment interface{} `field:"optional" json:"deployment" yaml:"deployment"`
	// Description for the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#description GenaiAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// functions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#functions GenaiAgent#functions}
	Functions interface{} `field:"optional" json:"functions" yaml:"functions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#id GenaiAgent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If case condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#if_case GenaiAgent#if_case}
	IfCase *string `field:"optional" json:"ifCase" yaml:"ifCase"`
	// K value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#k GenaiAgent#k}
	K *float64 `field:"optional" json:"k" yaml:"k"`
	// knowledge_bases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#knowledge_bases GenaiAgent#knowledge_bases}
	KnowledgeBases interface{} `field:"optional" json:"knowledgeBases" yaml:"knowledgeBases"`
	// Ids of the knowledge base(s) to attach to the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#knowledge_base_uuid GenaiAgent#knowledge_base_uuid}
	KnowledgeBaseUuid *[]*string `field:"optional" json:"knowledgeBaseUuid" yaml:"knowledgeBaseUuid"`
	// Maximum tokens allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#max_tokens GenaiAgent#max_tokens}
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#model GenaiAgent#model}
	Model interface{} `field:"optional" json:"model" yaml:"model"`
	// open_ai_api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#open_ai_api_key GenaiAgent#open_ai_api_key}
	OpenAiApiKey interface{} `field:"optional" json:"openAiApiKey" yaml:"openAiApiKey"`
	// Optional OpenAI API key ID to use with OpenAI models.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#open_ai_key_uuid GenaiAgent#open_ai_key_uuid}
	OpenAiKeyUuid *string `field:"optional" json:"openAiKeyUuid" yaml:"openAiKeyUuid"`
	// parent_agents block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#parent_agents GenaiAgent#parent_agents}
	ParentAgents interface{} `field:"optional" json:"parentAgents" yaml:"parentAgents"`
	// Indicates if the agent should provide citations in responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#provide_citations GenaiAgent#provide_citations}
	ProvideCitations interface{} `field:"optional" json:"provideCitations" yaml:"provideCitations"`
	// Retrieval method used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#retrieval_method GenaiAgent#retrieval_method}
	RetrievalMethod *string `field:"optional" json:"retrievalMethod" yaml:"retrievalMethod"`
	// User who created the route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#route_created_by GenaiAgent#route_created_by}
	RouteCreatedBy *string `field:"optional" json:"routeCreatedBy" yaml:"routeCreatedBy"`
	// Route name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#route_name GenaiAgent#route_name}
	RouteName *string `field:"optional" json:"routeName" yaml:"routeName"`
	// Route UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#route_uuid GenaiAgent#route_uuid}
	RouteUuid *string `field:"optional" json:"routeUuid" yaml:"routeUuid"`
	// List of Tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#tags GenaiAgent#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Agent temperature setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#temperature GenaiAgent#temperature}
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#template GenaiAgent#template}
	Template interface{} `field:"optional" json:"template" yaml:"template"`
	// Top P sampling parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#top_p GenaiAgent#top_p}
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
	// URL for the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#url GenaiAgent#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// User ID linked with the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#user_id GenaiAgent#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

