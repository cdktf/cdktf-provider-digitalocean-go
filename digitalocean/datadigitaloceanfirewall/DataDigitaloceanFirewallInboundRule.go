// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceanfirewall


type DataDigitaloceanFirewallInboundRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/data-sources/firewall#protocol DataDigitaloceanFirewall#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/data-sources/firewall#port_range DataDigitaloceanFirewall#port_range}.
	PortRange *string `field:"optional" json:"portRange" yaml:"portRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/data-sources/firewall#source_addresses DataDigitaloceanFirewall#source_addresses}.
	SourceAddresses *[]*string `field:"optional" json:"sourceAddresses" yaml:"sourceAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/data-sources/firewall#source_droplet_ids DataDigitaloceanFirewall#source_droplet_ids}.
	SourceDropletIds *[]*float64 `field:"optional" json:"sourceDropletIds" yaml:"sourceDropletIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/data-sources/firewall#source_kubernetes_ids DataDigitaloceanFirewall#source_kubernetes_ids}.
	SourceKubernetesIds *[]*string `field:"optional" json:"sourceKubernetesIds" yaml:"sourceKubernetesIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/data-sources/firewall#source_load_balancer_uids DataDigitaloceanFirewall#source_load_balancer_uids}.
	SourceLoadBalancerUids *[]*string `field:"optional" json:"sourceLoadBalancerUids" yaml:"sourceLoadBalancerUids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/data-sources/firewall#source_tags DataDigitaloceanFirewall#source_tags}.
	SourceTags *[]*string `field:"optional" json:"sourceTags" yaml:"sourceTags"`
}

