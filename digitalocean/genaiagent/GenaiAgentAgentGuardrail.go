// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent


type GenaiAgentAgentGuardrail struct {
	// Agent UUID for the Guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/genai_agent#agent_uuid GenaiAgent#agent_uuid}
	AgentUuid *string `field:"optional" json:"agentUuid" yaml:"agentUuid"`
	// Default response for the Guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/genai_agent#default_response GenaiAgent#default_response}
	DefaultResponse *string `field:"optional" json:"defaultResponse" yaml:"defaultResponse"`
	// Description of the Guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/genai_agent#description GenaiAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Guardrail UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/genai_agent#guardrail_uuid GenaiAgent#guardrail_uuid}
	GuardrailUuid *string `field:"optional" json:"guardrailUuid" yaml:"guardrailUuid"`
	// Indicates if the Guardrail is default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/genai_agent#is_default GenaiAgent#is_default}
	IsDefault interface{} `field:"optional" json:"isDefault" yaml:"isDefault"`
	// Name of Guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/genai_agent#name GenaiAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Priority of the Guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/genai_agent#priority GenaiAgent#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Type of the Guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/genai_agent#type GenaiAgent#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Guardrail UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/genai_agent#uuid GenaiAgent#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}

