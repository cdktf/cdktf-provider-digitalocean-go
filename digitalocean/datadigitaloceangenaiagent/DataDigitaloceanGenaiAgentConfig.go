// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDigitaloceanGenaiAgentConfig struct {
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
	// ID of the Agent to retrieve.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#agent_id DataDigitaloceanGenaiAgent#agent_id}
	AgentId *string `field:"required" json:"agentId" yaml:"agentId"`
	// agent_guardrail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#agent_guardrail DataDigitaloceanGenaiAgent#agent_guardrail}
	AgentGuardrail interface{} `field:"optional" json:"agentGuardrail" yaml:"agentGuardrail"`
	// anthropic_api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#anthropic_api_key DataDigitaloceanGenaiAgent#anthropic_api_key}
	AnthropicApiKey interface{} `field:"optional" json:"anthropicApiKey" yaml:"anthropicApiKey"`
	// api_key_infos block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#api_key_infos DataDigitaloceanGenaiAgent#api_key_infos}
	ApiKeyInfos interface{} `field:"optional" json:"apiKeyInfos" yaml:"apiKeyInfos"`
	// api_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#api_keys DataDigitaloceanGenaiAgent#api_keys}
	ApiKeys interface{} `field:"optional" json:"apiKeys" yaml:"apiKeys"`
	// chatbot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#chatbot DataDigitaloceanGenaiAgent#chatbot}
	Chatbot interface{} `field:"optional" json:"chatbot" yaml:"chatbot"`
	// chatbot_identifiers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#chatbot_identifiers DataDigitaloceanGenaiAgent#chatbot_identifiers}
	ChatbotIdentifiers interface{} `field:"optional" json:"chatbotIdentifiers" yaml:"chatbotIdentifiers"`
	// deployment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#deployment DataDigitaloceanGenaiAgent#deployment}
	Deployment interface{} `field:"optional" json:"deployment" yaml:"deployment"`
	// Description for the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#description DataDigitaloceanGenaiAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// functions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#functions DataDigitaloceanGenaiAgent#functions}
	Functions interface{} `field:"optional" json:"functions" yaml:"functions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#id DataDigitaloceanGenaiAgent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If case condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#if_case DataDigitaloceanGenaiAgent#if_case}
	IfCase *string `field:"optional" json:"ifCase" yaml:"ifCase"`
	// K value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#k DataDigitaloceanGenaiAgent#k}
	K *float64 `field:"optional" json:"k" yaml:"k"`
	// knowledge_bases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#knowledge_bases DataDigitaloceanGenaiAgent#knowledge_bases}
	KnowledgeBases interface{} `field:"optional" json:"knowledgeBases" yaml:"knowledgeBases"`
	// Maximum tokens allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#max_tokens DataDigitaloceanGenaiAgent#max_tokens}
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#model DataDigitaloceanGenaiAgent#model}
	Model interface{} `field:"optional" json:"model" yaml:"model"`
	// open_ai_api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#open_ai_api_key DataDigitaloceanGenaiAgent#open_ai_api_key}
	OpenAiApiKey interface{} `field:"optional" json:"openAiApiKey" yaml:"openAiApiKey"`
	// Retrieval method used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#retrieval_method DataDigitaloceanGenaiAgent#retrieval_method}
	RetrievalMethod *string `field:"optional" json:"retrievalMethod" yaml:"retrievalMethod"`
	// User who created the route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#route_created_by DataDigitaloceanGenaiAgent#route_created_by}
	RouteCreatedBy *string `field:"optional" json:"routeCreatedBy" yaml:"routeCreatedBy"`
	// Route name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#route_name DataDigitaloceanGenaiAgent#route_name}
	RouteName *string `field:"optional" json:"routeName" yaml:"routeName"`
	// Route UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#route_uuid DataDigitaloceanGenaiAgent#route_uuid}
	RouteUuid *string `field:"optional" json:"routeUuid" yaml:"routeUuid"`
	// List of Tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#tags DataDigitaloceanGenaiAgent#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Agent temperature setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#temperature DataDigitaloceanGenaiAgent#temperature}
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#template DataDigitaloceanGenaiAgent#template}
	Template interface{} `field:"optional" json:"template" yaml:"template"`
	// Top P sampling parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#top_p DataDigitaloceanGenaiAgent#top_p}
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
	// URL for the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#url DataDigitaloceanGenaiAgent#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// User ID linked with the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/genai_agent#user_id DataDigitaloceanGenaiAgent#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

