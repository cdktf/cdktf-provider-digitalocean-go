// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecIngressRuleRedirect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/app#authority App#authority}.
	Authority *string `field:"optional" json:"authority" yaml:"authority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/app#port App#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/app#redirect_code App#redirect_code}.
	RedirectCode *float64 `field:"optional" json:"redirectCode" yaml:"redirectCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/app#scheme App#scheme}.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.38.0/docs/resources/app#uri App#uri}.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

