// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type DataDigitaloceanFirewallInboundRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/firewall#protocol DataDigitaloceanFirewall#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/firewall#port_range DataDigitaloceanFirewall#port_range}.
	PortRange *string `field:"optional" json:"portRange" yaml:"portRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/firewall#source_addresses DataDigitaloceanFirewall#source_addresses}.
	SourceAddresses *[]*string `field:"optional" json:"sourceAddresses" yaml:"sourceAddresses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/firewall#source_droplet_ids DataDigitaloceanFirewall#source_droplet_ids}.
	SourceDropletIds *[]*float64 `field:"optional" json:"sourceDropletIds" yaml:"sourceDropletIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/firewall#source_kubernetes_ids DataDigitaloceanFirewall#source_kubernetes_ids}.
	SourceKubernetesIds *[]*string `field:"optional" json:"sourceKubernetesIds" yaml:"sourceKubernetesIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/firewall#source_load_balancer_uids DataDigitaloceanFirewall#source_load_balancer_uids}.
	SourceLoadBalancerUids *[]*string `field:"optional" json:"sourceLoadBalancerUids" yaml:"sourceLoadBalancerUids"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/firewall#source_tags DataDigitaloceanFirewall#source_tags}.
	SourceTags *[]*string `field:"optional" json:"sourceTags" yaml:"sourceTags"`
}
