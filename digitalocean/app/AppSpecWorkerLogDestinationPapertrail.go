// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecWorkerLogDestinationPapertrail struct {
	// Papertrail syslog endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.34.1/docs/resources/app#endpoint App#endpoint}
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
}

