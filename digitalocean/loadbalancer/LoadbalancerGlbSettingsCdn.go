// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadbalancerGlbSettingsCdn struct {
	// cache enable flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.48.2/docs/resources/loadbalancer#is_enabled Loadbalancer#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
}

