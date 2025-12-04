// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecAlert struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/app#rule App#rule}.
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// destinations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/app#destinations App#destinations}
	Destinations *AppSpecAlertDestinations `field:"optional" json:"destinations" yaml:"destinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/app#disabled App#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

