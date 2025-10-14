// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent


type GenaiAgentChildAgents struct {
	// Instruction for the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_agent#instruction GenaiAgent#instruction}
	Instruction *string `field:"required" json:"instruction" yaml:"instruction"`
	// Model UUID of the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_agent#model_uuid GenaiAgent#model_uuid}
	ModelUuid *string `field:"required" json:"modelUuid" yaml:"modelUuid"`
	// Name of the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_agent#name GenaiAgent#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Project ID of the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_agent#project_id GenaiAgent#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Region where the Agent is deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_agent#region GenaiAgent#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// anthropic_api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_agent#anthropic_api_key GenaiAgent#anthropic_api_key}
	AnthropicApiKey interface{} `field:"optional" json:"anthropicApiKey" yaml:"anthropicApiKey"`
	// api_key_infos block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_agent#api_key_infos GenaiAgent#api_key_infos}
	ApiKeyInfos interface{} `field:"optional" json:"apiKeyInfos" yaml:"apiKeyInfos"`
	// api_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_agent#api_keys GenaiAgent#api_keys}
	ApiKeys interface{} `field:"optional" json:"apiKeys" yaml:"apiKeys"`
	// chatbot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_agent#chatbot GenaiAgent#chatbot}
	Chatbot interface{} `field:"optional" json:"chatbot" yaml:"chatbot"`
	// chatbot_identifiers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_agent#chatbot_identifiers GenaiAgent#chatbot_identifiers}
	ChatbotIdentifiers interface{} `field:"optional" json:"chatbotIdentifiers" yaml:"chatbotIdentifiers"`
	// deployment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_agent#deployment GenaiAgent#deployment}
	Deployment interface{} `field:"optional" json:"deployment" yaml:"deployment"`
	// Description for the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/genai_agent#description GenaiAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

