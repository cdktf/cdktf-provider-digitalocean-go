// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecIngressRule struct {
	// component block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.1/docs/resources/app#component App#component}
	Component *AppSpecIngressRuleComponent `field:"optional" json:"component" yaml:"component"`
	// cors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.1/docs/resources/app#cors App#cors}
	Cors *AppSpecIngressRuleCors `field:"optional" json:"cors" yaml:"cors"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.1/docs/resources/app#match App#match}
	Match *AppSpecIngressRuleMatch `field:"optional" json:"match" yaml:"match"`
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.1/docs/resources/app#redirect App#redirect}
	Redirect *AppSpecIngressRuleRedirect `field:"optional" json:"redirect" yaml:"redirect"`
}

