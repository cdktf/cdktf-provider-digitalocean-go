package firewall


type FirewallInboundRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/firewall#protocol Firewall#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/firewall#port_range Firewall#port_range}.
	PortRange *string `field:"optional" json:"portRange" yaml:"portRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/firewall#source_addresses Firewall#source_addresses}.
	SourceAddresses *[]*string `field:"optional" json:"sourceAddresses" yaml:"sourceAddresses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/firewall#source_droplet_ids Firewall#source_droplet_ids}.
	SourceDropletIds *[]*float64 `field:"optional" json:"sourceDropletIds" yaml:"sourceDropletIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/firewall#source_kubernetes_ids Firewall#source_kubernetes_ids}.
	SourceKubernetesIds *[]*string `field:"optional" json:"sourceKubernetesIds" yaml:"sourceKubernetesIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/firewall#source_load_balancer_uids Firewall#source_load_balancer_uids}.
	SourceLoadBalancerUids *[]*string `field:"optional" json:"sourceLoadBalancerUids" yaml:"sourceLoadBalancerUids"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/firewall#source_tags Firewall#source_tags}.
	SourceTags *[]*string `field:"optional" json:"sourceTags" yaml:"sourceTags"`
}

