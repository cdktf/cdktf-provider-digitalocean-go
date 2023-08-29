// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecFunctionCorsAllowOrigins struct {
	// Exact string match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/app#exact App#exact}
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Prefix-based match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/app#prefix App#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// RE2 style regex-based match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/app#regex App#regex}
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}

