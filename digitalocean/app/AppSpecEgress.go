// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecEgress struct {
	// The app egress type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/app#type App#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

