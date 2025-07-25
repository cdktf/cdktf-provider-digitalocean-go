// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecJobAlert struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/app#operator App#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/app#rule App#rule}.
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/app#value App#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/app#window App#window}.
	Window *string `field:"required" json:"window" yaml:"window"`
	// destinations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/app#destinations App#destinations}
	Destinations *AppSpecJobAlertDestinations `field:"optional" json:"destinations" yaml:"destinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.61.0/docs/resources/app#disabled App#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

