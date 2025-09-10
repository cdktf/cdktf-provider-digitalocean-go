// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagent


type GenaiAgentApiKeys struct {
	// API Key value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/resources/genai_agent#api_key GenaiAgent#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
}

