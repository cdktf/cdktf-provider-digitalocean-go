// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent


type GenaiAgentFunctions struct {
	// API Key value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#api_key GenaiAgent#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Description of the Function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#description GenaiAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Name of function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#faasname GenaiAgent#faasname}
	Faasname *string `field:"optional" json:"faasname" yaml:"faasname"`
	// Namespace of function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#faasnamespace GenaiAgent#faasnamespace}
	Faasnamespace *string `field:"optional" json:"faasnamespace" yaml:"faasnamespace"`
	// Guardrail UUID for the Function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#guardrail_uuid GenaiAgent#guardrail_uuid}
	GuardrailUuid *string `field:"optional" json:"guardrailUuid" yaml:"guardrailUuid"`
	// Name of function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#name GenaiAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Url of the Deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#url GenaiAgent#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// API Key value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_agent#uuid GenaiAgent#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}

