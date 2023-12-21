// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecFunctionLogDestinationDatadog struct {
	// Datadog API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.34.1/docs/resources/app#api_key App#api_key}
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// Datadog HTTP log intake endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.34.1/docs/resources/app#endpoint App#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

