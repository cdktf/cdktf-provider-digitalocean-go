// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecWorkerLogDestinationLogtail struct {
	// Logtail token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/app#token App#token}
	Token *string `field:"required" json:"token" yaml:"token"`
}

