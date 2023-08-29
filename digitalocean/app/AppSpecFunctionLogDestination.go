// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecFunctionLogDestination struct {
	// Name of the log destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/app#name App#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/app#datadog App#datadog}
	Datadog *AppSpecFunctionLogDestinationDatadog `field:"optional" json:"datadog" yaml:"datadog"`
	// logtail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/app#logtail App#logtail}
	Logtail *AppSpecFunctionLogDestinationLogtail `field:"optional" json:"logtail" yaml:"logtail"`
	// papertrail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/app#papertrail App#papertrail}
	Papertrail *AppSpecFunctionLogDestinationPapertrail `field:"optional" json:"papertrail" yaml:"papertrail"`
}

