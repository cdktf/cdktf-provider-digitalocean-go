// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagent


type DataDigitaloceanGenaiAgentFunctions struct {
	// API Key value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/data-sources/genai_agent#api_key DataDigitaloceanGenaiAgent#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Description of the Function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/data-sources/genai_agent#description DataDigitaloceanGenaiAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Name of function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/data-sources/genai_agent#faasname DataDigitaloceanGenaiAgent#faasname}
	Faasname *string `field:"optional" json:"faasname" yaml:"faasname"`
	// Namespace of function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/data-sources/genai_agent#faasnamespace DataDigitaloceanGenaiAgent#faasnamespace}
	Faasnamespace *string `field:"optional" json:"faasnamespace" yaml:"faasnamespace"`
	// Guardrail UUID for the Function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/data-sources/genai_agent#guardrail_uuid DataDigitaloceanGenaiAgent#guardrail_uuid}
	GuardrailUuid *string `field:"optional" json:"guardrailUuid" yaml:"guardrailUuid"`
	// Name of function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/data-sources/genai_agent#name DataDigitaloceanGenaiAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Url of the Deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/data-sources/genai_agent#url DataDigitaloceanGenaiAgent#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// API Key value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/data-sources/genai_agent#uuid DataDigitaloceanGenaiAgent#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}

