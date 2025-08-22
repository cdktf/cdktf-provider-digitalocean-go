// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiopenaiapikey


type GenaiOpenaiApiKeyModelVersions struct {
	// Major version of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_openai_api_key#major GenaiOpenaiApiKey#major}
	Major *float64 `field:"optional" json:"major" yaml:"major"`
	// Minor version of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_openai_api_key#minor GenaiOpenaiApiKey#minor}
	Minor *float64 `field:"optional" json:"minor" yaml:"minor"`
	// Patch version of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/resources/genai_openai_api_key#patch GenaiOpenaiApiKey#patch}
	Patch *float64 `field:"optional" json:"patch" yaml:"patch"`
}

