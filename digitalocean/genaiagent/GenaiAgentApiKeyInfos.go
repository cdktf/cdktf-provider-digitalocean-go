// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent


type GenaiAgentApiKeyInfos struct {
	// Created By user ID for the API Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/genai_agent#created_by GenaiAgent#created_by}
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// Name of the API Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/genai_agent#name GenaiAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Updated At timestamp for the API Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/genai_agent#secret_key GenaiAgent#secret_key}
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// API Key value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/genai_agent#uuid GenaiAgent#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}

