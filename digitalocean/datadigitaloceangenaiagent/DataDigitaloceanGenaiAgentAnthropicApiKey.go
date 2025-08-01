// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagent


type DataDigitaloceanGenaiAgentAnthropicApiKey struct {
	// Created By user ID for the API Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/data-sources/genai_agent#created_by DataDigitaloceanGenaiAgent#created_by}
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// Name of the API Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/data-sources/genai_agent#name DataDigitaloceanGenaiAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// API Key value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/data-sources/genai_agent#uuid DataDigitaloceanGenaiAgent#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}

