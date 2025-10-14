// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewall


type FirewallOutboundRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/firewall#protocol Firewall#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/firewall#destination_addresses Firewall#destination_addresses}.
	DestinationAddresses *[]*string `field:"optional" json:"destinationAddresses" yaml:"destinationAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/firewall#destination_droplet_ids Firewall#destination_droplet_ids}.
	DestinationDropletIds *[]*float64 `field:"optional" json:"destinationDropletIds" yaml:"destinationDropletIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/firewall#destination_kubernetes_ids Firewall#destination_kubernetes_ids}.
	DestinationKubernetesIds *[]*string `field:"optional" json:"destinationKubernetesIds" yaml:"destinationKubernetesIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/firewall#destination_load_balancer_uids Firewall#destination_load_balancer_uids}.
	DestinationLoadBalancerUids *[]*string `field:"optional" json:"destinationLoadBalancerUids" yaml:"destinationLoadBalancerUids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/firewall#destination_tags Firewall#destination_tags}.
	DestinationTags *[]*string `field:"optional" json:"destinationTags" yaml:"destinationTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/firewall#port_range Firewall#port_range}.
	PortRange *string `field:"optional" json:"portRange" yaml:"portRange"`
}

