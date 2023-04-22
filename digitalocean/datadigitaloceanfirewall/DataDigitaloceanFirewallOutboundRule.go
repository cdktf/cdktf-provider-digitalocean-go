package datadigitaloceanfirewall


type DataDigitaloceanFirewallOutboundRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.0/docs/data-sources/firewall#protocol DataDigitaloceanFirewall#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.0/docs/data-sources/firewall#destination_addresses DataDigitaloceanFirewall#destination_addresses}.
	DestinationAddresses *[]*string `field:"optional" json:"destinationAddresses" yaml:"destinationAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.0/docs/data-sources/firewall#destination_droplet_ids DataDigitaloceanFirewall#destination_droplet_ids}.
	DestinationDropletIds *[]*float64 `field:"optional" json:"destinationDropletIds" yaml:"destinationDropletIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.0/docs/data-sources/firewall#destination_kubernetes_ids DataDigitaloceanFirewall#destination_kubernetes_ids}.
	DestinationKubernetesIds *[]*string `field:"optional" json:"destinationKubernetesIds" yaml:"destinationKubernetesIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.0/docs/data-sources/firewall#destination_load_balancer_uids DataDigitaloceanFirewall#destination_load_balancer_uids}.
	DestinationLoadBalancerUids *[]*string `field:"optional" json:"destinationLoadBalancerUids" yaml:"destinationLoadBalancerUids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.0/docs/data-sources/firewall#destination_tags DataDigitaloceanFirewall#destination_tags}.
	DestinationTags *[]*string `field:"optional" json:"destinationTags" yaml:"destinationTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.0/docs/data-sources/firewall#port_range DataDigitaloceanFirewall#port_range}.
	PortRange *string `field:"optional" json:"portRange" yaml:"portRange"`
}

