// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecWorkerLogDestinationOpenSearchBasicAuth struct {
	// Password for basic authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.2/docs/resources/app#password App#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// user for basic authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.2/docs/resources/app#user App#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

