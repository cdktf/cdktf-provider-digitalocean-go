// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceangenaiagent


type DataDigitaloceanGenaiAgentApiKeyInfos struct {
	// Created By user ID for the API Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/data-sources/genai_agent#created_by DataDigitaloceanGenaiAgent#created_by}
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// Name of the API Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/data-sources/genai_agent#name DataDigitaloceanGenaiAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Updated At timestamp for the API Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/data-sources/genai_agent#secret_key DataDigitaloceanGenaiAgent#secret_key}
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// API Key value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/data-sources/genai_agent#uuid DataDigitaloceanGenaiAgent#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}

