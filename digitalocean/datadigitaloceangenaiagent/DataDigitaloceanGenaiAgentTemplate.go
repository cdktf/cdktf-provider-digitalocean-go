// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagent


type DataDigitaloceanGenaiAgentTemplate struct {
	// Description of the Agent Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/data-sources/genai_agent#description DataDigitaloceanGenaiAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Instruction for the Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/data-sources/genai_agent#instruction DataDigitaloceanGenaiAgent#instruction}
	Instruction *string `field:"optional" json:"instruction" yaml:"instruction"`
	// K value for the Agent Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/data-sources/genai_agent#k DataDigitaloceanGenaiAgent#k}
	K *float64 `field:"optional" json:"k" yaml:"k"`
	// knowledge_bases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/data-sources/genai_agent#knowledge_bases DataDigitaloceanGenaiAgent#knowledge_bases}
	KnowledgeBases interface{} `field:"optional" json:"knowledgeBases" yaml:"knowledgeBases"`
	// Maximum tokens allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/data-sources/genai_agent#max_tokens DataDigitaloceanGenaiAgent#max_tokens}
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/data-sources/genai_agent#model DataDigitaloceanGenaiAgent#model}
	Model interface{} `field:"optional" json:"model" yaml:"model"`
	// Name of the Agent Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/data-sources/genai_agent#name DataDigitaloceanGenaiAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Agent temperature setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/data-sources/genai_agent#temperature DataDigitaloceanGenaiAgent#temperature}
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// Top P sampling parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/data-sources/genai_agent#top_p DataDigitaloceanGenaiAgent#top_p}
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
	// uuid of the Agent Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/data-sources/genai_agent#uuid DataDigitaloceanGenaiAgent#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}

