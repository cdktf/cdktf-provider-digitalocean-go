// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiopenaiapikey


type GenaiOpenaiApiKeyModelAgreement struct {
	// Description of the agreement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_openai_api_key#description GenaiOpenaiApiKey#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Name of the agreement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_openai_api_key#name GenaiOpenaiApiKey#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// URL of the agreement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_openai_api_key#url GenaiOpenaiApiKey#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// UUID of the agreement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_openai_api_key#uuid GenaiOpenaiApiKey#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}

