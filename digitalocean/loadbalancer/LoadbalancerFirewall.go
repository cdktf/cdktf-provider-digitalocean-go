// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadbalancerFirewall struct {
	// the rules for ALLOWING traffic to the LB (strings in the form: 'ip:1.2.3.4' or 'cidr:1.2.0.0/16').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.54.0/docs/resources/loadbalancer#allow Loadbalancer#allow}
	Allow *[]*string `field:"optional" json:"allow" yaml:"allow"`
	// the rules for DENYING traffic to the LB (strings in the form: 'ip:1.2.3.4' or 'cidr:1.2.0.0/16').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.54.0/docs/resources/loadbalancer#deny Loadbalancer#deny}
	Deny *[]*string `field:"optional" json:"deny" yaml:"deny"`
}

