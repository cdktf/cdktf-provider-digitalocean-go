// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecIngressRuleMatchPath struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.1/docs/resources/app#prefix App#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

