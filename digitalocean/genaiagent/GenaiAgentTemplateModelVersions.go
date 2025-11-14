// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent


type GenaiAgentTemplateModelVersions struct {
	// Major version of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/resources/genai_agent#major GenaiAgent#major}
	Major *float64 `field:"optional" json:"major" yaml:"major"`
	// Minor version of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/resources/genai_agent#minor GenaiAgent#minor}
	Minor *float64 `field:"optional" json:"minor" yaml:"minor"`
	// Patch version of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/resources/genai_agent#patch GenaiAgent#patch}
	Patch *float64 `field:"optional" json:"patch" yaml:"patch"`
}

