// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagent


type DataDigitaloceanGenaiAgentDeployment struct {
	// Name of the API Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/data-sources/genai_agent#name DataDigitaloceanGenaiAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Status of the Deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/data-sources/genai_agent#status DataDigitaloceanGenaiAgent#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Url of the Deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/data-sources/genai_agent#url DataDigitaloceanGenaiAgent#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// API Key value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/data-sources/genai_agent#uuid DataDigitaloceanGenaiAgent#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
	// Visibility of the Deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/data-sources/genai_agent#visibility DataDigitaloceanGenaiAgent#visibility}
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

