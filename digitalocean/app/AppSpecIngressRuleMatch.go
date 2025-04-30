// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecIngressRuleMatch struct {
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.52.0/docs/resources/app#path App#path}
	Path *AppSpecIngressRuleMatchPath `field:"optional" json:"path" yaml:"path"`
}

