// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent


type GenaiAgentOpenAiApiKey struct {
	// Created By user ID for the API Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_agent#created_by GenaiAgent#created_by}
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// Name of the API Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_agent#name GenaiAgent#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// API Key value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_agent#uuid GenaiAgent#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}

