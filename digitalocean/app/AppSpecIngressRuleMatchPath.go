// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecIngressRuleMatchPath struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/resources/app#prefix App#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

