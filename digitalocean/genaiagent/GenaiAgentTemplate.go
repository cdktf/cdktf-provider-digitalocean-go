// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent


type GenaiAgentTemplate struct {
	// Description of the Agent Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_agent#description GenaiAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Instruction for the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_agent#instruction GenaiAgent#instruction}
	Instruction *string `field:"optional" json:"instruction" yaml:"instruction"`
	// K value for the Agent Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_agent#k GenaiAgent#k}
	K *float64 `field:"optional" json:"k" yaml:"k"`
	// knowledge_bases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_agent#knowledge_bases GenaiAgent#knowledge_bases}
	KnowledgeBases interface{} `field:"optional" json:"knowledgeBases" yaml:"knowledgeBases"`
	// Maximum tokens allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_agent#max_tokens GenaiAgent#max_tokens}
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_agent#model GenaiAgent#model}
	Model interface{} `field:"optional" json:"model" yaml:"model"`
	// Name of the Agent Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_agent#name GenaiAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Agent temperature setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_agent#temperature GenaiAgent#temperature}
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// Top P sampling parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_agent#top_p GenaiAgent#top_p}
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
	// uuid of the Agent Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/genai_agent#uuid GenaiAgent#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}

