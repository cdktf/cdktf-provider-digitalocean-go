// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecIngressRuleMatchAuthority struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/app#exact App#exact}.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
}

