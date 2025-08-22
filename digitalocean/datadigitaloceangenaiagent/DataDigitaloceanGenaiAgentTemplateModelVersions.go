// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagent


type DataDigitaloceanGenaiAgentTemplateModelVersions struct {
	// Major version of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/genai_agent#major DataDigitaloceanGenaiAgent#major}
	Major *float64 `field:"optional" json:"major" yaml:"major"`
	// Minor version of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/genai_agent#minor DataDigitaloceanGenaiAgent#minor}
	Minor *float64 `field:"optional" json:"minor" yaml:"minor"`
	// Patch version of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/genai_agent#patch DataDigitaloceanGenaiAgent#patch}
	Patch *float64 `field:"optional" json:"patch" yaml:"patch"`
}

