// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagent


type DataDigitaloceanGenaiAgentAgentGuardrail struct {
	// Agent UUID for the Guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/data-sources/genai_agent#agent_uuid DataDigitaloceanGenaiAgent#agent_uuid}
	AgentUuid *string `field:"optional" json:"agentUuid" yaml:"agentUuid"`
	// Default response for the Guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/data-sources/genai_agent#default_response DataDigitaloceanGenaiAgent#default_response}
	DefaultResponse *string `field:"optional" json:"defaultResponse" yaml:"defaultResponse"`
	// Description of the Guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/data-sources/genai_agent#description DataDigitaloceanGenaiAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Guardrail UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/data-sources/genai_agent#guardrail_uuid DataDigitaloceanGenaiAgent#guardrail_uuid}
	GuardrailUuid *string `field:"optional" json:"guardrailUuid" yaml:"guardrailUuid"`
	// Indicates if the Guardrail is default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/data-sources/genai_agent#is_default DataDigitaloceanGenaiAgent#is_default}
	IsDefault interface{} `field:"optional" json:"isDefault" yaml:"isDefault"`
	// Name of Guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/data-sources/genai_agent#name DataDigitaloceanGenaiAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Priority of the Guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/data-sources/genai_agent#priority DataDigitaloceanGenaiAgent#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Type of the Guardrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/data-sources/genai_agent#type DataDigitaloceanGenaiAgent#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Guardrail UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/data-sources/genai_agent#uuid DataDigitaloceanGenaiAgent#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}

