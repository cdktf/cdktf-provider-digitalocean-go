// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecIngressRuleComponent struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.53.0/docs/resources/app#name App#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.53.0/docs/resources/app#preserve_path_prefix App#preserve_path_prefix}.
	PreservePathPrefix interface{} `field:"optional" json:"preservePathPrefix" yaml:"preservePathPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.53.0/docs/resources/app#rewrite App#rewrite}.
	Rewrite *string `field:"optional" json:"rewrite" yaml:"rewrite"`
}

